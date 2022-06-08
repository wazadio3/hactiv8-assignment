package router

import (
	"final_project/controller"
	"final_project/database"
	"final_project/middleware"
	"final_project/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartService(db *gorm.DB) *gin.Engine {
	userQuery := database.NewUserQuery(db)
	userService := service.NewUserService(userQuery)
	userController := controller.NewUserController(userService)

	photoQuery := database.NewPhotoQuery(db)
	photoService := service.NewPhotoService(photoQuery)
	photoController := controller.NewPhotoController(photoService)

	commentQuery := database.NewCommentQuery(db)
	commentService := service.NewCommentService(commentQuery)
	commentController := controller.NewCommentController(commentService)

	sosMedQuery := database.NewSosMedQuery(db)
	sosMedService := service.NewSosMedService(sosMedQuery)
	sosMedController := controller.NewSosMedController(sosMedService)

	route := gin.Default()

	route.POST("/user/register", userController.CreateUser)
	route.POST("/user/login", userController.LoginUser)
	route.PUT("/user", middleware.Auth(), userController.UpdateUser)
	route.DELETE("/user", middleware.Auth(), userController.DeleteUser)

	route.POST("/photo", middleware.Auth(), photoController.CreatePhoto)
	route.GET("/photo", middleware.Auth(), photoController.GetPhotos)
	route.PUT("/photo/:photoId", middleware.Auth(), photoController.UpdatePhotoById)
	route.DELETE("/photo/:photoId", middleware.Auth(), photoController.DeletePhotoById)

	route.POST("/comment", middleware.Auth(), commentController.CreateComment)
	route.GET("/comment", middleware.Auth(), commentController.GetComments)
	route.PUT("/comment/:commentId", middleware.Auth(), commentController.UpdateCommentById)
	route.DELETE("/comment/:commentId", middleware.Auth(), commentController.DeleteCommentById)

	route.POST("/socialmedia", middleware.Auth(), sosMedController.CreateSosMed)
	route.GET("/socialmedia", middleware.Auth(), sosMedController.GetSosMeds)
	route.PUT("/socialmedia/:socialMediaId", middleware.Auth(), sosMedController.UpdateSosMedById)
	route.DELETE("/socialmedia/:socialMediaId", middleware.Auth(), sosMedController.DeleteSosMedById)

	return route
}
