package format

import (
	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/initializers"
)

func ConstructorAuthor() {
	Author = &discordgo.MessageEmbedAuthor{
		URL:          "",
		Name:         "Julio",
		IconURL:      initializers.AvatarURL,
		ProxyIconURL: "",
	}
}
