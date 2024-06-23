package main

import (
	"encoding/binary"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/dgraph-io/badger/v4"
)

var db *badger.DB

func connectDatabase(databasePath string) {

	var err error
	db, err = badger.Open(badger.DefaultOptions(databasePath))
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
}

func incrementCount(session *discordgo.Session, message *discordgo.MessageCreate) {

	keys := [][]byte{
		[]byte(message.Author.ID),
		[]byte(message.GuildID),
	}

	err := db.Update(func(txn *badger.Txn) error {

		for _, key := range keys {
			item, err := txn.Get(key)
			var count uint64

			if err == nil {
				err = item.Value(func(val []byte) error {

					count = binary.BigEndian.Uint64(val)
					return nil
				})
				if err != nil {
					return err
				}
			}

			count++

			countBytes := make([]byte, 8)
			binary.BigEndian.PutUint64(countBytes, count)

			if err := txn.Set(key, countBytes); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal("Couldn't update database: ", err)
	}
}
