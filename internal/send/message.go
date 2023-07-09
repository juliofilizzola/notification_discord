package send

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/initializers"
	"github.com/juliofilizzola/bot_discord/internal/format"
)

func MessageDiscord() {
	discord, err := startServer()

	if err != nil {
		log.Fatal(err)
	}

	params := format.ConstructParams()

	webhook, err := discord.WebhookExecute(initializers.WebhookId, initializers.WebhookToken, false, &params)
	fmt.Println(webhook)
}

func startServer() (s *discordgo.Session, err error) {
	discord, err := discordgo.New("Bot " + initializers.TokenDiscord)
	if err != nil {
		return nil, err
	}

	return discord, err
}
