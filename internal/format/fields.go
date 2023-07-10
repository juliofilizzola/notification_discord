package format

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	types "github.com/juliofilizzola/bot_discord/internal/struct"
)

func ConstructFields(data *types.Github) {
	status := &discordgo.MessageEmbedField{
		Name:   "Status:",
		Value:  data.PullRequest.State,
		Inline: false,
	}

	assigner := &discordgo.MessageEmbedField{
		Name:   "Assinado:",
		Value:  data.PullRequest.Assignee.Login,
		Inline: false,
	}

	additions := &discordgo.MessageEmbedField{
		Name:   "Codigo adicionado:",
		Value:  strconv.Itoa(data.PullRequest.Additions),
		Inline: true,
	}
	deletions := &discordgo.MessageEmbedField{
		Name:   "Codigo deletado",
		Value:  strconv.Itoa(data.PullRequest.Deletions),
		Inline: true,
	}
	commits := &discordgo.MessageEmbedField{
		Name:   "Commits:",
		Value:  "[commits](" + data.PullRequest.HtmlUrl + "/commits)",
		Inline: false,
	}

	Fields = append(Fields, status, assigner, additions, deletions, commits)

	var reviewers []*discordgo.MessageEmbedField

	indexRequests := len(data.PullRequest.RequestedReviewers)

	if indexRequests > 0 {
		for _, value := range data.PullRequest.RequestedReviewers {
			review := &discordgo.MessageEmbedField{
				Name:   "Review :",
				Value:  value.Login,
				Inline: false,
			}
			reviewers = append(reviewers, review)
		}
		for _, review := range reviewers {
			Fields = append(Fields, review)
		}
	} else {
		review := &discordgo.MessageEmbedField{
			Name:   "Review :",
			Value:  "sem review...",
			Inline: false,
		}
		reviewers = append(reviewers, review)
	}

}
