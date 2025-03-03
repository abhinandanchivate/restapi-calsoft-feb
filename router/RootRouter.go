package router

import (
	"user-rest-app/config"
	"user-rest-app/repository"
	"user-rest-app/services"

	"github.com/gorilla/mux"
)

// @RequestMapping

func SetupRootRouter() *mux.Router {

	// connectDB
	db:= config.ConnectDB()
	// User Repository 
	userRepo := repository.NewUserRepositoryImpl(db)

	UserService := services.NewUserServiceImpl(userRepo)

	// User Service

	// required services , repos ==> objects I will create it here
	
	r := mux.NewRouter()
	r.PathPrefix("/api").Subrouter()
// /api as a common spec for all routes
SetupUserRouter(r, UserService)
	// handling routing .
	// movements 
	return r 
}