package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/bot_discord/controller"
	"github.com/juliofilizzola/bot_discord/initializers"
)

func init() {
	initializers.Env()
}

func main() {
	r := gin.Default()
	r.POST("/:id/:token", controller.GithubController)
	r.Run()
}
