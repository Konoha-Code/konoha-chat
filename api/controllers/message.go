package controllers

import (
	"fmt"
	"net/http"
	"sysmo/api-go/database"
	"sysmo/api-go/models"

	"github.com/gin-gonic/gin"
)

func SendMessage(ctx *gin.Context) {
	from := ctx.Query("from")
	message := ctx.Query("message")
	to := ctx.Query("to")

	if from == "" {
		ctx.JSON(http.StatusBadRequest, "Necessário informar o parâmetro [from]")
		return
	}

	if message == "" {
		ctx.JSON(http.StatusBadRequest, "Necessário informar o parâmetro [message]")
		return
	}

	if to == "" {
		ctx.JSON(http.StatusBadRequest, "Necessário informar o parâmetro [to]")
		return
	}

	if getUser("ID", from).ID == "" {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Usuário com ID = [%s] não encontrado", from))
		return
	}

	if getUser("ID", to).ID == "" {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Usuário com ID = [%s] não encontrado", to))
		return
	}

	database.GetDatabase().Create(&models.Message{From: from, Message: message, To: to})

	ctx.JSON(http.StatusOK, "Mensagem enviada!")
}

func GetMessages(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")

	if from == "" {
		ctx.JSON(http.StatusBadRequest, "Necessário informar o parâmetro [from]")
		return
	}

	if to == "" {
		ctx.JSON(http.StatusBadRequest, "Necessário informar o parâmetro [to]")
		return
	}

	if getUser("ID", from).ID == "" {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Usuário com ID = [%s] não encontrado", from))
		return
	}

	if getUser("ID", to).ID == "" {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Usuário com ID = [%s] não encontrado", to))
		return
	}

	var messages []models.Message

	database.GetDatabase().Where(&models.Message{From: from, To: to}).Find(&messages)

	ctx.JSON(http.StatusOK, messages)
}
