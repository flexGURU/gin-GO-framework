package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// const (
// 	authorizationHeader = "Authorization"
// )

// Authentcation Middleware
// Returns unauthorised if user is not
func AuthenticationMiddleware()  gin.HandlerFunc{

	return func(ctx *gin.Context) {
		// Get authorisation Header

		if !(ctx.Request.Header.Get("Token")  == "auth") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Message" : "Token required",
			})
			return
		} 
		ctx.Next()
		
	}
	
}

func AddHeader(ctx *gin.Context)  {

	ctx.Writer.Header().Set("Key", "Value")
	ctx.Next()
	
}

func LoggingMidd() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// Log request details before processing
		startTime := time.Now()

		logrus.WithFields(logrus.Fields{
			"clientIP": ctx.ClientIP(),
			"method":   ctx.Request.Method,
			"endpoint": ctx.Request.URL.Path,
		}).Info("Incoming request")

		// Process the request
		ctx.Next()

		// Log response details after processing
		latency := time.Since(startTime)
		logrus.WithFields(logrus.Fields{
			"clientIP":    ctx.ClientIP(),
			"method":      ctx.Request.Method,
			"endpoint":    ctx.Request.URL.Path,
			"status":      ctx.Writer.Status(),
			"latency(ms)": latency.Milliseconds(),
		}).Info("Request processed")
	}


	
}



