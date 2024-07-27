package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/rowandevving/heather/database"
	"github.com/rowandevving/heather/settings"
)

func StatsCommand(session *discordgo.Session, message *discordgo.MessageCreate) {

	var key string
	var name string
	args := handleCommand(message.Content, "stats", settings.Config.Stats.Enabled)
	if args == nil {
		return
	} else if len(args) > 0 && args[0] == "server" {
		key = message.GuildID
		name = "Server stats"
	} else {
		key = message.Author.ID
		name = message.Author.GlobalName + "'s stats"
	}

	session.ChannelMessageSendEmbedReply(message.ChannelID, &discordgo.MessageEmbed{
		Title: name,
		Color: 13346551,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Messages sent",
				Value: fmt.Sprintf("%d", database.GetCount(key)),
			},
		},
	}, message.Reference())
}
