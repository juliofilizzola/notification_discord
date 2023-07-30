package input

import (
	"github.com/bwmarrin/discordgo"
)

type WebhookDomainService interface {
	Send(dataGit *discordgo.WebhookParams, webhookId, webhookToken, action string) string
}
