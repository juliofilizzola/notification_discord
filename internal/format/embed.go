package format

import (
	"time"

	"github.com/bwmarrin/discordgo"
	types "github.com/juliofilizzola/bot_discord/internal/struct"
)

var (
	Embed     []*discordgo.MessageEmbed
	Author    *discordgo.MessageEmbedAuthor
	Footer    *discordgo.MessageEmbedFooter
	Fields    []*discordgo.MessageEmbedField
	Thumbnail *discordgo.MessageEmbedThumbnail
	Image     *discordgo.MessageEmbedImage
)

func ConstructEmbed(data *types.Github) {
	ConstructFields(data)
	ConstructorAuthor(&data.Sender)

	ConstructorImg(data)
	ConstructorThumbnail(data)
	ConstructorFooter(&data.Organization)

	// todo melhorar esse codigo para a v2
	timeSend := time.Now().Format(`2006-01-02 15:04:05`)
	Embed = append(Embed, &discordgo.MessageEmbed{
		URL:         data.PullRequest.HtmlUrl,
		Type:        discordgo.EmbedTypeLink,
		Title:       data.PullRequest.Title,
		Description: data.PullRequest.Body,
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
