package main

import (
	"net/http"
	"time"

	"girhub.com/flexGURU/go-gin/middleware"
	"girhub.com/flexGURU/go-gin/routes"
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


	routes.Routes(router)

	server := http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}

	server.ListenAndServe()
	



}