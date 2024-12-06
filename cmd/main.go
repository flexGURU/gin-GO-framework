package main

import (
	"net/http"
	"time"

	"girhub.com/flexGURU/go-gin/handlers"
	"girhub.com/flexGURU/go-gin/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init()  {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,

		})
	logrus.SetLevel(logrus.DebugLevel)
	
}

func main() {



	router := gin.New()


	// Printing using logrus
	router.Use(middleware.LoggingMidd())

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
	


	server := http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}

	server.ListenAndServe()
	



}