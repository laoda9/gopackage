package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	v1.Use(middleware.JWTAuthMiddleware())
	//注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	//登录业务路由
	v1.POST("/login", controller.LoginHandler)

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler) //模糊匹配

		//创建帖子
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts/", controller.GetPostListHandler)

		v1.POST("/vote", controller.PostVoteController)

	}
	////要求只有登录网站，才能看到ping路由
	//r.GET("/ping", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
	//	//如果是登录的用户，判断请求头中是否有 有效的JWT
	//	c.String(http.StatusOK, "pong")
	//})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
