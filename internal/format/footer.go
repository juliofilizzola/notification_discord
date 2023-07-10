package format

import (
	"github.com/bwmarrin/discordgo"
	types "github.com/juliofilizzola/bot_discord/internal/struct"
)

func ConstructorFooter(data *types.Organization) {
	Footer = &discordgo.MessageEmbedFooter{
		Text:         data.Login,
		IconURL:      data.AvatarUrl,
		ProxyIconURL: "",
	}
}
