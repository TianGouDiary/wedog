package server

import (
	"github.com/gin-gonic/gin"
	"wedog/api"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.GET("getRandomLetter", api.GetRandomLetter)
		v1.POST("getRandomLetter", api.GetRandomLetter)
	}
	return r
}
