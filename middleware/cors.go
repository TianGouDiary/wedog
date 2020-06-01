package middleware

import (
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//// Cors 跨域配置
//func Cors() gin.HandlerFunc {
//	config := cors.DefaultConfig()
//	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
//	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
//	config.AllowOrigins = []string{"http://www.example.com"}
//	config.AllowCredentials = true
//	return cors.New(config)
//}

//处理跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, XMLHttpRequest, "+
			"Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.String(200, "ok")
			return
		}
		c.Next()
	}
}