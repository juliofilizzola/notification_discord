package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/config/env"
)

func StartDiscord() (*discordgo.Session, error) {
	discord, err := discordgo.New("Bot " + env.TokenDiscord)
	if err != nil {
		return discord, err
	}
	return discord, nil
}
