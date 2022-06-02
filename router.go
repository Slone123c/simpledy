package main

import (
	"github.com/gin-gonic/gin"
	"simpledy/service"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	//apiRouter.GET("/feed/", handler.Feed)
	//apiRouter.GET("/user/", handler.UserInfo)
	apiRouter.POST("/user/register/", service.Register)
	//apiRouter.POST("/user/login/", handler.Login)
	//apiRouter.POST("/publish/action/", handler.Publish)
	//apiRouter.GET("/publish/list/", handler.PublishList)

	// extra apis - I
	//apiRouter.POST("/favorite/action/", handler.FavoriteAction)
	//apiRouter.GET("/favorite/list/", handler.FavoriteList)
	//apiRouter.POST("/comment/action/", handler.CommentAction)
	//apiRouter.GET("/comment/list/", handler.CommentList)

	//extra apis - II
	//apiRouter.POST("/relation/action/", handler.RelationAction)
	//apiRouter.GET("/relation/follow/list/", handler.FollowList)
	//apiRouter.GET("/relation/follower/list/", handler.FollowerList)
}
