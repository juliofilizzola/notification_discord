package format

import (
	"github.com/bwmarrin/discordgo"
)

func ConstructFields() {
	Fields = append(Fields, &discordgo.MessageEmbedField{
		Name:   "Test",
		Value:  "ESSE È UM FIELD DE TEST",
		Inline: false,
	})
}
