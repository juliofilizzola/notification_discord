package format

import (
	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/initializers"
	types "github.com/juliofilizzola/bot_discord/internal/struct"
)

func ConstructParams(data *types.Github) discordgo.WebhookParams {
	ConstructEmbed(data)
	return discordgo.WebhookParams{
		Content:    "Nova PR no Repositorio: " + data.Repository.Name,
		Username:   data.PullRequest.Repo.Name,
		AvatarURL:  initializers.AvatarURL,
		TTS:        false,
		Files:      nil,
		Components: nil,
		Embeds:     Embed,
		AllowedMentions: &discordgo.MessageAllowedMentions{
			Parse: []discordgo.AllowedMentionType{
				discordgo.AllowedMentionTypeEveryone,
			},
			Roles:       nil,
			Users:       nil,
			RepliedUser: false,
		},
		Flags: 0,
	}
}
