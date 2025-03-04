package controllers

import (
	"encoding/json"
	"net/http"
	"user-rest-app/services"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) { 
	// hello message json via http.ResponseWriter
	// {message: "Hello from user controller"}
	w.Header().Set("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello from user controller"})
	//w.Write([]byte("Hello from user controller"))

	// return anything? 
	// NO 
}
	// req, res 


	
