package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wfernandez/rest-api/models"
	"github.com/wfernandez/rest-api/utils"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		utils.HandleError(context, http.StatusBadRequest, err)
		return
	}

	err = user.Save()
	if err != nil {
		utils.HandleError(context, http.StatusInternalServerError, err)
		return
	}
	utils.JSONResponse(context, http.StatusCreated, gin.H{"user": user})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		utils.HandleError(context, http.StatusBadRequest, err)
		return
	}

	err = user.Authenticate()
	if err != nil {
		utils.HandleError(context, http.StatusUnauthorized, err)
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID) 

	if err != nil {
		utils.HandleError(context, http.StatusInternalServerError, err)
		return
	}	

	utils.JSONResponse(context, http.StatusOK, gin.H{"message": "Login successful", "token": token})
}