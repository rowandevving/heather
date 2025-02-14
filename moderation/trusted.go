package moderation

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/rowandevving/heather/config"
	"github.com/rowandevving/heather/database"
	"github.com/rowandevving/heather/util"
)

func HandleTrustedRole(session *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == session.State.User.ID {
		return
	}

	trustedRole := util.GetIDFromRoleName(message, session, config.Global.Moderation.TrustedRole)

	key := message.Author.ID
	roles := message.Member.Roles

	for _, role := range roles {
		if role == trustedRole {
			return
		}
	}

	if database.GetCount(key) >= config.Global.Moderation.TrustedThreshold {

		err := session.GuildMemberRoleAdd(message.GuildID, message.Author.ID, trustedRole)

		if err != nil {
			log.Fatal("Error adding trusted role to user: ", err)
		}
	}
}
