package discord

import (
	"github.com/juliofilizzola/bot_discord/internal/send"
	types "github.com/juliofilizzola/bot_discord/internal/struct"
)

func Send(data *types.Github, webhookId, webhookToken string) {
	send.MessageDiscord(data, webhookId, webhookToken)
}
