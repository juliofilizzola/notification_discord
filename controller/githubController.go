package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/bot_discord/discord"
	types "github.com/juliofilizzola/bot_discord/internal/struct"
)

func GithubController(ctx *gin.Context) {
	var body types.Github

	err := ctx.Bind(&body)

	if err != nil {
		fmt.Println("err", err)
		return
	}
	webhookId := ctx.Param("id")
	webhookToken := ctx.Param("token")

	discord.Send(&body, webhookId, webhookToken)

	ctx.JSON(200, gin.H{
		"message": "Send webhook",
	})
}
