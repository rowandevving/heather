package commands

import (
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rowandevving/heather/settings"
)

func ColourCommand(session *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == session.State.User.ID {
		return
	}

	args := handleCommand(message.Content, "color", settings.Config.Colour.Enabled)
	if args == nil {
		return
	}

	for _, colour := range settings.Config.Colour.Colours {
		if colour.Name == args[0] {

			roles, err := session.GuildRoles(message.GuildID)
			if err != nil {
				log.Fatal("Error getting guild roles: ", err)
			}

			rolePresent := false
			name := args[0]
			var roleToAdd *discordgo.Role

			for _, role := range roles {
				if role.Name == name {
					rolePresent = true
					roleToAdd = role
					break
				}
			}

			roleColour := convertHexToInt(colour.Hex)

			if !rolePresent {

				newRole, err := session.GuildRoleCreate(message.GuildID, &discordgo.RoleParams{
					Name:  name,
					Color: &roleColour,
				})

				if err != nil {
					log.Fatal("Error creating role: ", err)
				}

				err = session.GuildMemberRoleAdd(message.GuildID, message.Author.ID, newRole.ID)
				if err != nil {
					log.Fatal("Error adding new role: ", err)
				}
			} else {

				updatedRole, err := session.GuildRoleEdit(message.GuildID, roleToAdd.ID, &discordgo.RoleParams{
					Color: &roleColour,
				})
				if err != nil {
					log.Fatal("Error updating role: ", err)
				}
				err = session.GuildMemberRoleAdd(message.GuildID, message.Author.ID, updatedRole.ID)
				if err != nil {
					log.Fatal("Error adding role: ", err)
				}

				pruneColourRoles(session, message, updatedRole)
			}
		}
	}
}

func convertHexToInt(hex string) int {

	hexadecimal := strings.Replace(hex, "#", "0x", -1)

	integer, err := strconv.ParseInt(hexadecimal, 0, 64)

	if err != nil {
		panic(err)
	}

	return int(integer)

}

func pruneColourRoles(session *discordgo.Session, message *discordgo.MessageCreate, excludedRole *discordgo.Role) {

	for _, colour := range settings.Config.Colour.Colours {

		if colour.Name != excludedRole.Name {

			roles, err := session.GuildRoles(message.GuildID)
			if err != nil {
				log.Fatal("Error getting guild roles: ", err)
			}

			for _, role := range roles {
				if role.Name == colour.Name {

					session.GuildMemberRoleRemove(message.GuildID, message.Author.ID, role.ID)
				}
			}
		}
	}
}
