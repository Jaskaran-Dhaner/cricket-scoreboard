package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		// AllowOrigins:     []string{"http://example.com", "http://anotherdomain.com"}, // specify allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},   // specify allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // specify allowed headers
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // specify how long the results of a preflight request can be cached
	})
}
