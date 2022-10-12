package controllers

import "github.com/gin-gonic/gin"

func GetPing(ctx *gin.Context) {
	ctx.JSON(200, "Chat Online")
}
