package router

import (
	docs "github.com/alexandrejuniorc/gopportunities/docs"
	"github.com/alexandrejuniorc/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	BASE_PATH := "/api/v1"

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Gopportunities API"
	docs.SwaggerInfo.Description = "This is a sample server to create job opportunities using Go."
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = BASE_PATH
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Initialize handler
	handler.InitializeHandler()

	v1 := router.Group(BASE_PATH)
	{
		v1.GET("/openings", handler.GetAllOpeningsHandler)
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
