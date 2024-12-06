package routes

import (
	"girhub.com/flexGURU/go-gin/handlers"
	"girhub.com/flexGURU/go-gin/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {


	// Subrouting and using Router Gropus
	// admin routes
	admin := router.Group("/admin")
	{
		admin.GET("/getData", middleware.AuthenticationMiddleware(), middleware.AddHeader, handlers.GetData)
	}

	// client routes

	client := router.Group("/client")
	{
		client.POST("/postData", handlers.PostData)
		client.POST("/getQS", handlers.GetQueryString)
		client.GET("/getQueryParams/:name/:age", handlers.GetQueryParams)
	}
	



}