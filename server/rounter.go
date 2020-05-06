package server

import (
	"GinWeb/api"
	"github.com/gin-gonic/gin"
)

func NewRounter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	user := v1.Group("user")
	{
		user.POST("register", api.Register)
		user.POST("login", api.Login)
	}

	item := v1.Group("item")
	{
		item.POST("list", api.Item_List)
		item.POST("info", api.Item_Info)
	}

	return router
}
