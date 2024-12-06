package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler function to get serve the GET request /getData
func GetData(ctx *gin.Context) {


	// Logrus 
	ctx.JSON(http.StatusOK, gin.H{
		"data": "here is the data",
	})
}


// 
func PostData(ctx *gin.Context) {

	body := ctx.Request.Body
	value, _ := io.ReadAll(body)

	ctx.JSON(http.StatusOK, gin.H{
		"data": fmt.Sprintf("%s", value),
	})
}

// Query strings
// eg URL: http://localhost:8080/getQS?name=John&age=30
func GetQueryString(ctx *gin.Context) {

	name := ctx.Query("name")
	age := ctx.Query("age")

	ctx.JSON(http.StatusOK, gin.H{
		"data": fmt.Sprintf("values are name: %s and age: %s", name, age),
	})
}
// Query parameters
// eg URL: http://localhost:8080/getQueryParams/John/30
func GetQueryParams(ctx *gin.Context) {

	name := ctx.Param("name")
	age := ctx.Param("age")

	ctx.JSON(http.StatusOK, gin.H{
		"data": fmt.Sprintf("values are name: %s and age: %s", name, age),
	})



}