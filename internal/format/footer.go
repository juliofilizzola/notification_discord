package format

import (
	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/initializers"
)

func ConstructorFooter() {
	Footer = &discordgo.MessageEmbedFooter{
		Text:         "This is @360",
		IconURL:      initializers.AvatarURL,
		ProxyIconURL: "",
	}
}
