package router

import (
	"github.com/gin-gonic/gin"
	"tgin/handle"
	"tgin/middleware"
)

func InitRouter(r *gin.Engine) {

	r.Use(middleware.Cors())
	// uc先关接口
	GroupV1 := r.Group("/api")
	{
		GroupV1.GET("/words", handle.UCLogin)
		GroupV1.GET("/hello", handle.UCLogin)
	}
	// message相关接口
	GroupV2 := r.Group("/message", middleware.RequestIDMiddleware())
	{
		GroupV2.GET("/hi", handle.UCLogin)
		GroupV2.GET("hello", handle.UCLogin)
	}
}
