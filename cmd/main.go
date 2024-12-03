package main

import (
	"net/http"
	"time"

	"girhub.com/flexGURU/go-gin/handlers"
	"github.com/gin-gonic/gin"
)



func main() {

	router := gin.Default()

	// Router
	router.GET("/getData", handlers.GetData)
	router.POST("/postData", handlers.PostData)
	router.POST("/getQS", handlers.GetQueryString)
	router.GET("/getQueryParams/:name/:age", handlers.GetQueryParams)


	server := http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}

	server.ListenAndServe()
	



}