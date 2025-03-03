package router

import (
	"user-rest-app/controllers"
	"user-rest-app/services"

	"github.com/gorilla/mux"
)

// import (
// 	"gorm-crud-api/handlers"
// 	"gorm-crud-api/repository"
// 	"gorm-crud-api/services"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gorilla/mux"
// 	"gorm.io/gorm"
// )

// // User Router
// func SetupRoutes(router *gin.Engine, db *gorm.DB) {
// 	userRepo := repository.NewUserRepository(db)
// 	userService := services.NewUserService(userRepo)
// 	userHandler := handlers.NewUserHandler(userService)

// 	userRoutes := router.Group("/users")
// 	{
// 		userRoutes.POST("", userHandler.CreateUser)
// 		userRoutes.GET("", userHandler.GetUsers)
// 		userRoutes.GET("/:id", userHandler.GetUser)
// 		userRoutes.PUT("/:id", userHandler.UpdateUser)
// 		userRoutes.DELETE("/:id", userHandler.DeleteUser)
// 	}
// }

// package routes

// import (
//     "github.com/gin-gonic/gin"
//     "gin-user-roles/controllers"
//     "gin-user-roles/middlewares"
// )

func SetupUserRouter(r *mux.Router,  UserService services.UserService) { 
	userRouter :=r.PathPrefix("/users").Subrouter()

	RegisterUserRoutes(userRouter, UserService)
	// we want to handle the end point called 
	/// /api/users method : get

	// we have to map the endpointt with handler
	// handlers ===> controllers
	// controllers ===> services==> repository==> db

	//userRouter.HandleFunc("", ).Methods("GET")


}
func RegisterUserRoutes(router *mux.Router, UserService services.UserService) {
	controller := controllers.NewUserController(UserService)
	router.HandleFunc("/", controller.GetAllUsers).Methods("GET")
	// router.HandleFunc("/{id}", controller.GetUserByID).Methods("GET")
	// router.HandleFunc("/", controller.CreateUser).Methods("POST")
	// router.HandleFunc("/{id}", controller.UpdateUser).Methods("PUT")
	// router.HandleFunc("/{id}", controller.DeleteUser).Methods("DELETE")
}

// func SetupRouter() *gin.Engine {
//     router := gin.Default()

//     // Public routes
//     router.GET("/users", controllers.GetAllUsers)

//     // Protected routes (Admin-only)
//     admin := router.Group("/admin")
//     admin.Use(middlewares.RoleMiddleware("admin"))
//     {
//         admin.GET("/dashboard", controllers.AdminOnly)
//     }

//     return router
// }
