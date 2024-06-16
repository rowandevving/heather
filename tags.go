package main

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

type Tag struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	SubTags []Tag  `json:"subtags"`
}

const tagMatch = `--([a-zA-Z0-9]+)(?:-([a-zA-Z0-9]+))?`

func handleTag(session *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == session.State.User.ID {
		return
	}

	regex := regexp.MustCompile(tagMatch)
	matches := regex.FindAllStringSubmatch(message.Content, -1)

	processedTags := make(map[string]struct{})
	processedSubTags := make(map[string]struct{})

	for _, match := range matches {
		if len(match) == 3 {
			tag := match[1]
			subtag := ""
			if len(match) > 2 {
				subtag = match[2]
			}
			if subtag == "" {

				if _, found := processedTags[tag]; found {
					continue
				}
				processedTags[tag] = struct{}{}

				for _, currentTag := range settings.Tags {
					if currentTag.Name == tag {
						session.ChannelMessageSend(message.ChannelID, currentTag.Message)
						break
					}
				}
			} else {

				fullTag := tag + "-" + subtag
				if _, found := processedSubTags[fullTag]; found {
					continue
				}
				processedSubTags[fullTag] = struct{}{}

				for _, currentTag := range settings.Tags {
					if currentTag.Name == tag {
						for _, currentSubTag := range currentTag.SubTags {
							if currentSubTag.Name == subtag {
								session.ChannelMessageSend(message.ChannelID, currentSubTag.Message)
								break
							}
						}
					}
				}
			}
		}
	}
}
