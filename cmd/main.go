package main

import (
	"net/http"
	"time"

	"girhub.com/flexGURU/go-gin/handlers"
	"github.com/gin-gonic/gin"
)



func main() {

	router := gin.Default()

	// Subrouting and using Router Gropus
	// admin routes
	admin := router.Group("/admin")
	{
		admin.GET("/getData", handlers.GetData)
	}

	// client routes

	client := router.Group("/client")
	{
		client.POST("/postData", handlers.PostData)
		client.POST("/getQS", handlers.GetQueryString)
		client.GET("/getQueryParams/:name/:age", handlers.GetQueryParams)
	}
	


	server := http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}

	server.ListenAndServe()
	



}