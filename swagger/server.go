package swagger

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	httpSwagger "github.com/swaggo/http-swagger"
)

func enableCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for all responses including preflight requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Max-Age", "3600")
		
		// Handle preflight OPTIONS requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		handler(w, r)
	}
}

func StartServer(port string) error {
	if port == "" {
		port = "8081"
	}

	// Get current directory
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	yamlPath := filepath.Join(dir, "docs", "openapi.yaml")
	
	// Create a new mux to handle all routes
	mux := http.NewServeMux()
	
	// Handle serving the OpenAPI spec with proper CORS headers
	mux.HandleFunc("/openapi.yaml", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/yaml")
		http.ServeFile(w, r, yamlPath)
	}))

	// Configure Swagger UI with proper CORS handling
	swaggerHandler := httpSwagger.Handler(
		httpSwagger.URL("/openapi.yaml"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("list"),
		httpSwagger.DomID("swagger-ui"),
	)
	mux.HandleFunc("/swagger/", enableCORS(swaggerHandler))

	// Redirect root to Swagger UI
	mux.HandleFunc("/", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/swagger/", http.StatusSeeOther)
			return
		}
		http.NotFound(w, r)
	}))

	fmt.Printf("Swagger documentation server running on http://localhost:%s/swagger/\n", port)
	return http.ListenAndServe(":"+port, mux)
}