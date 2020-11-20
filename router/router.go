package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv01/controller"
)

func Router() *gin.Engine {
	router := gin.Default()
	// 路径映射
	/*
		//router.GET("/user", InitPage)
		router.POST("/user/create", CreateUser)
		//router.GET("/user/create", CreateUser)
		router.GET("/user/list", ListUser)
		router.PUT("/user/update/:id", UpdateUser)
		router.DELETE("/user/:id", DeleteUser)
	*/
	router.GET("/article/getone/:id", controller.NewArticleController().GetOne);
	router.GET("/article/list", controller.NewArticleController().GetList);
	return router
}
