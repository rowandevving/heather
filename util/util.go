package util

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func GetIDFromRoleName(message *discordgo.MessageCreate, session *discordgo.Session, roleName string) string {

	roles, err := session.GuildRoles(message.GuildID)

	if err != nil {
		log.Fatal("Error getting guild roles: ", err)
	}

	for _, role := range roles {

		if role.Name == roleName {
			return role.ID
		}
	}

	return ""

}
