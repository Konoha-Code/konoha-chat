package controllers

import (
	"fmt"
	"konoha/api-go/database"
	"konoha/api-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUser(ctx *gin.Context) {
	name := ctx.Query("name")

	if name == "" {
		ctx.JSON(http.StatusBadRequest, "Necessário informar o parâmetro [name]")
		return
	}

	if getUser("name", name).ID != "" {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Usuário [%s] já cadastrado, favor informar um nome diferente", name))
		return
	}

	id := uuid.New().String()

	database.GetDatabase().Create(&models.User{ID: id, Name: name})

	ctx.JSON(http.StatusOK, "Usuário cadastrado com sucesso!")
}

func GetUser(ctx *gin.Context) {
	name := ctx.Query("name")

	if name == "" {
		ctx.JSON(http.StatusBadRequest, "Necessário informar o parâmetro [name]")
		return
	}

	ctx.JSON(http.StatusOK, getUser("name", name))
}

func getUser(key, value string) models.User {
	var user = &models.User{}

	database.GetDatabase().Where(fmt.Sprintf("%s = ?", key), value).First(&user)

	return *user
}
