package format

import (
	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/initializers"
)

func ConstructorThumbnail() {
	Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL:      initializers.AvatarURL,
		ProxyURL: "",
		Width:    40,
		Height:   60,
	}
}
