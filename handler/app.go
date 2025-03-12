package handler

import (
	"15min-city/db"
	"15min-city/middleware"
	"15min-city/repository/corridor_route_repository/corridor_route_db"
	"15min-city/repository/dataset_repository/dataset_db"
	"15min-city/repository/user_repository/user_db"
	"15min-city/service"

	"github.com/gin-gonic/gin"
)

func App() {
	db.InitializeDB()
	dbInstance := db.GetDBInstance()

	// User endpoint
	userRepository := user_db.NewUserRepository(dbInstance)
	userService := service.NewUserService(userRepository)
	userHandler := NewUserHandler(userService)

	// Dataset endpoint
	datasetRepository := dataset_db.NewDatasetRepository(dbInstance)
	datasetService := service.NewDatasetService(datasetRepository)
	datasetHandler := NewDatasetHandler(datasetService)

	// CorridorRoute endpoint
	corridorRouteRepository := corridor_route_db.NewCorridorRouteRepository(dbInstance)
	corridorRouteService := service.NewCorridorRouteService(corridorRouteRepository)
	corridorRouteHandler := NewCorridorRouteHandler(corridorRouteService)

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

		homeRoute := userRoute.Group("/me")
		{
			homeRoute.GET("", middleware.Authentication(), userHandler.GetUserByID)
			homeRoute.PATCH("/upload", middleware.Authentication(), userHandler.CreateImage)
			homeRoute.GET("/:fileID", middleware.Authentication(), userHandler.GetImageByUser)
		}
	}

	// Dataset routes
	datasetRoute := api.Group("/datasets")
	{
		datasetRoute.POST("", middleware.Authentication(), middleware.Authorization(), datasetHandler.CreateDataset)
		datasetRoute.POST("/upload", middleware.Authentication(), middleware.Authorization(), datasetHandler.UploadDatasets)
		datasetRoute.GET("/:datasetID", middleware.Authentication(), datasetHandler.GetDatasetByID)
		datasetRoute.GET("/name/:name", middleware.Authentication(), datasetHandler.GetDatasetByName)
		datasetRoute.GET("/kecamatan/:kecamatan", middleware.Authentication(), datasetHandler.GetDatasetByKecamatan)
		datasetRoute.GET("/kelurahan/:kelurahan", middleware.Authentication(), datasetHandler.GetDatasetByKelurahan)
		datasetRoute.PUT("/:datasetID", middleware.Authentication(), middleware.Authorization(), datasetHandler.UpdateDataset)
		datasetRoute.GET("/category/:category", middleware.Authentication(), datasetHandler.GetDatasetByCategory)
		datasetRoute.DELETE("/:datasetID", middleware.Authentication(), middleware.Authorization(), datasetHandler.DeleteDataset)
		datasetRoute.GET("", middleware.Authentication(), datasetHandler.GetAllDatasets)
	}

	// CorridorRoute routes
	corridorRoute := api.Group("/corridor-routes")
	{
		corridorRoute.POST("", middleware.Authentication(), middleware.Authorization(), corridorRouteHandler.CreateCorridorRoute)
		corridorRoute.GET("/:id", middleware.Authentication(), corridorRouteHandler.GetCorridorRouteByID)
		corridorRoute.GET("/name/:name", middleware.Authentication(), corridorRouteHandler.GetCorridorRouteByName)
		corridorRoute.GET("/route/:route", middleware.Authentication(), corridorRouteHandler.GetCorridorRouteByRoute)
		corridorRoute.GET("/direction/:direction", middleware.Authentication(), corridorRouteHandler.GetCorridorRouteByDirection)
		corridorRoute.PUT("/:id", middleware.Authentication(), middleware.Authorization(), corridorRouteHandler.UpdateCorridorRoute)
		corridorRoute.DELETE("/:id", middleware.Authentication(), middleware.Authorization(), corridorRouteHandler.DeleteCorridorRoute)
		corridorRoute.GET("", middleware.Authentication(), corridorRouteHandler.GetAllCorridorRoutes)
	}

	r.Run()
}
