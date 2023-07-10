package send

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/initializers"
	"github.com/juliofilizzola/bot_discord/internal/format"
	types "github.com/juliofilizzola/bot_discord/internal/struct"
)

func MessageDiscord(data *types.Github, webhookId, webhookToken string) {
	discord, err := startServer()

	if err != nil {
		log.Fatal(err)
	}

	params := format.ConstructParams(data)

	webhook, err := discord.WebhookExecute(webhookId, webhookToken, false, &params)
	fmt.Println(webhook)
}

func startServer() (s *discordgo.Session, err error) {
	discord, err := discordgo.New("Bot " + initializers.TokenDiscord)
	if err != nil {
		return nil, err
	}

	return discord, err
}
