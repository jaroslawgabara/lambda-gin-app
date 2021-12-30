package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func StartApi(r *gin.Engine) {
	if os.Getenv("HTTPS") == "true" {
		fmt.Println("running https")
		r.RunTLS(":8080", os.Getenv("CERT_FILE"), os.Getenv("KEY_PATH"))
	} else {
		fmt.Println("running http")
		r.Run()
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
