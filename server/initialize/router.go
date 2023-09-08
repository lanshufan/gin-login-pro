package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pass-crypto/modules/middleware"
	"pass-crypto/routers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())
	publicRouter := routers.PublicRouterApp
	privateRouter := routers.PrivateRouterApp

	// public路由
	publicRouterGroup := r.Group("v1")
	{
		publicRouter.RegisterRouter.RegisterRouter(publicRouterGroup)
		publicRouter.LoginRouter.LoginRouter(publicRouterGroup)
		publicRouter.MyRouter.BlogLists(publicRouterGroup)
	}

	// private路由
	privateRouterGroup := r.Group("v1")
	privateRouterGroup.Use(middleware.VerifyMiddleware())
	{
		privateRouter.PageRouter.PageRouter(privateRouterGroup)
	}

	return r
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, "+
				"Content-Length, X-CSRF-Token, Token, Session, Access-Control-Allow-Credentials")
			c.Set("Content-Type", "application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
		c.Next()
	}
}
