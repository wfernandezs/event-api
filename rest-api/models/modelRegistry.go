package models

import (
	"fmt"
	"reflect"
)

var registry = make([]interface{}, 0)

func RegisterModel(model interface{}) {
	registry = append(registry, model)
	fmt.Printf("Model registered for auto-migration: %T\n", model)
}

func GetRegisteredModels() []interface{} {
	return registry
}

func ListRegisteredModels() []string {
	var modelNames []string
	for _, model := range registry {
		t := reflect.TypeOf(model)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		modelNames = append(modelNames, t.Name())
	}
	return modelNames
}