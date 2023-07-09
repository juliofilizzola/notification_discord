package format

import (
	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/initializers"
)

func ConstructorImg() {
	Image = &discordgo.MessageEmbedImage{
		URL:      initializers.AvatarURL,
		ProxyURL: "",
		Width:    20,
		Height:   20,
	}
}
