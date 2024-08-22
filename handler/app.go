package handler

import (
	"15min-city/db"
	"15min-city/middleware"
	"15min-city/repository/dataset_repository/dataset_db"
	"15min-city/repository/user_repository/user_db"
	"15min-city/service"

	"github.com/gin-gonic/gin"
)

func App() {
	db.InitializeDB()
	db := db.GetDBInstance()

	// User endpoint
	userRepository := user_db.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := NewUserHandler(userService)

	// Dataset endpoint
	datasetRepository := dataset_db.NewDatasetRepository(db)
	datasetService := service.NewDatasetService(datasetRepository)
	datasetHandler := NewDatasetHandler(datasetService)

	r := gin.Default()
	api := r.Group("/api/v1")
	api.Use(middleware.CORSMiddleware())
	api.OPTIONS("/*any", middleware.CORSMiddleware())

	// User routes
	userRoute := api.Group("/home")
	{
		userRoute.POST("/register", userHandler.Register)
		userRoute.POST("/login", userHandler.Login)
		userRoute.PATCH("/reset-password", userHandler.ResetPassword)

		homeRoute := userRoute.Group("/my-account")
		{
			homeRoute.GET("", middleware.Authentication(), userHandler.GetUserByID)
			homeRoute.PATCH("/upload", middleware.Authentication(), userHandler.CreateImage)
			homeRoute.GET("/:fileID", middleware.Authentication(), userHandler.GetImageByUser)
		}
	}

	// Dataset routes
	datasetRoute := api.Group("/datasets")
	{
		datasetRoute.POST("", datasetHandler.CreateDataset)
		datasetRoute.GET("/:datasetID", datasetHandler.GetDatasetByID)
		datasetRoute.GET("/name/:name", datasetHandler.GetDatasetByName)
		datasetRoute.PUT("/:datasetID", datasetHandler.UpdateDataset)
		datasetRoute.GET("/category/:category", datasetHandler.GetDatasetByCategory)
		datasetRoute.DELETE("/:datasetID", datasetHandler.DeleteDataset)
		datasetRoute.GET("", datasetHandler.GetAllDatasets)
	}
	r.Run()
}
