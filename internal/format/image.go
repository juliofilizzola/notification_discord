package format

import (
	"github.com/bwmarrin/discordgo"
	types "github.com/juliofilizzola/bot_discord/internal/struct"
)

func ConstructorImg(data *types.Github) {
	Image = &discordgo.MessageEmbedImage{
		URL:      data.Organization.AvatarUrl,
		ProxyURL: "",
		Width:    20,
		Height:   20,
	}
}
