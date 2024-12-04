package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

// Authentcation Middleware
// Returns unauthorised if user is not 
func AuthentionMiddleware()  gin.HandlerFunc{

	return func(ctx *gin.Context) {
		// Get authorisation Header
		authHeader := ctx.GetHeader(authorizationHeader)

		// check if authHeader is missing
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"ERROR": "Authorisation Header Required",
			})
			return
		}
		
	
		
	}
	
}