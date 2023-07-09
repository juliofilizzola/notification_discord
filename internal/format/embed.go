package format

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	Embed     []*discordgo.MessageEmbed
	Author    *discordgo.MessageEmbedAuthor
	Footer    *discordgo.MessageEmbedFooter
	Fields    []*discordgo.MessageEmbedField
	Thumbnail *discordgo.MessageEmbedThumbnail
	Image     *discordgo.MessageEmbedImage
)

func ConstructEmbed() {
	ConstructFields()
	ConstructorAuthor()

	ConstructorImg()
	ConstructorThumbnail()
	ConstructorFooter()

	// todo melhorar esse codigo para a v2
	timeSend := time.Now().Format(`2006-01-02 15:04:05`)
	Embed = append(Embed, &discordgo.MessageEmbed{
		URL:         "",
		Type:        discordgo.EmbedTypeLink,
		Title:       "PR 1000",
		Description: "PR NO REPO",
		Timestamp:   timeSend,
		Color:       0,
		Footer:      Footer,
		Image:       Image,
		Thumbnail:   Thumbnail,
		Video:       nil,
		Provider:    nil,
		Author:      Author,
		Fields:      Fields,
	})
}
