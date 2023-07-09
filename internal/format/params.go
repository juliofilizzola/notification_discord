package format

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/initializers"
)

func ConstructParams() discordgo.WebhookParams {
	ConstructEmbed()
	fmt.Println(initializers.Username)
	return discordgo.WebhookParams{
		Content:    "Nova PR",
		Username:   initializers.Username,
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
