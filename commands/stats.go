package commands

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/dgraph-io/badger/v4"
	"github.com/rowandevving/heather/database"
)

func StatsCommand(session *discordgo.Session, message *discordgo.MessageCreate) {

	var key []byte
	var name string
	args := handleCommand(message.Content, "stats")
	if args == nil {
		return
	} else if len(args) > 0 && args[0] == "server" {
		key = []byte(message.GuildID)
		name = "Server stats"
	} else {
		key = []byte(message.Author.ID)
		name = message.Author.GlobalName + "'s stats"
	}

	err := database.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)

		if err == badger.ErrKeyNotFound {
			return nil
		} else if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {

			session.ChannelMessageSendEmbedReply(message.ChannelID, &discordgo.MessageEmbed{
				Title: name,
				Color: 13346551,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "Messages sent",
						Value: fmt.Sprintf("%d", binary.BigEndian.Uint64(val)),
					},
				},
			}, message.Reference())

			return nil
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal("Couldn't read database: ", err)
	}
}
