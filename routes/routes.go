package routes

import (
	"github.com/gin-gonic/gin"
	"im-server/controllers"
	"im-server/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 公开路由
	public := r.Group("/api")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}

	// 需要认证的路由
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/profile", controllers.GetProfile)
	}

	return r
} 