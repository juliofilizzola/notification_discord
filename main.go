package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/juliofilizzola/bot_discord/adpter/input/controller"
	"github.com/juliofilizzola/bot_discord/adpter/input/controller/routes"
	"github.com/juliofilizzola/bot_discord/application/services"
	discord2 "github.com/juliofilizzola/bot_discord/config/discord"
	"github.com/juliofilizzola/bot_discord/config/env"
)

func init() {
	env.Env()
}

func main() {
	r := gin.Default()

	webController := initDependencies()
	routes.InitRoutes(&r.RouterGroup, webController)

	if err := r.Run(env.Port); err != nil {
		log.Fatal(err)
	}
}

func initDependencies() controller.WebhookControllerInterface {
	discord, err := discord2.StartDiscord()
	if err != nil {
		log.Fatal(err)
	}
	service := services.NewWebhookDomainService(discord)
	return controller.NewWebhookControllerInterface(service)
}
