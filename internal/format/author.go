package format

import (
	"github.com/bwmarrin/discordgo"
	types "github.com/juliofilizzola/bot_discord/internal/struct"
)

func ConstructorAuthor(data *types.User) {
	Author = &discordgo.MessageEmbedAuthor{
		URL:          data.HtmlUrl,
		Name:         data.Login,
		IconURL:      data.AvatarUrl,
		ProxyIconURL: "",
	}
}
