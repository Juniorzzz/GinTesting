package api

import "github.com/gin-gonic/gin"

func Item_List(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "item list success"})
}

func Item_Info(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "item info success"})
}
