package format

import (
	"github.com/bwmarrin/discordgo"
	types "github.com/juliofilizzola/bot_discord/internal/struct"
)

func ConstructorThumbnail(data *types.Github) {
	Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL:      data.Organization.AvatarUrl,
		ProxyURL: "",
		Width:    40,
		Height:   60,
	}
}
