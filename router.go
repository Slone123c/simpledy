package main

import (
	"simpledy/service"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", service.Feed)
	apiRouter.GET("/user/", service.UserInfo)
	apiRouter.POST("/user/register/", service.Register)
	apiRouter.POST("/user/login/", service.Login)
	apiRouter.POST("/publish/action/", service.Publish)
	apiRouter.GET("/publish/list/", service.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", service.FavoriteAction)
	apiRouter.GET("/favorite/list/", service.FavoriteList)
	apiRouter.POST("/comment/action/", service.CommentAction)
	apiRouter.GET("/comment/list/", service.CommentList)
	//
	//extra apis - II
	apiRouter.POST("/relation/action/", service.RelationAction)
	apiRouter.GET("/relation/follow/list/", service.FollowList)
	apiRouter.GET("/relation/follower/list/", service.FollowerList)
}
