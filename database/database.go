package database

import (
	"encoding/binary"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/dgraph-io/badger/v4"
)

var DB *badger.DB

func ConnectDatabase(databasePath string) {

	var err error
	DB, err = badger.Open(badger.DefaultOptions(databasePath))
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
}

func IncrementCount(session *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == session.State.User.ID {
		return
	}

	keys := [][]byte{
		[]byte(message.Author.ID),
		[]byte(message.GuildID),
	}

	err := DB.Update(func(txn *badger.Txn) error {

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

func GetCount(key string) uint64 {

	var count uint64

	err := DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))

		if err == badger.ErrKeyNotFound {
			return nil
		} else if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {

			count = binary.BigEndian.Uint64(val)
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

	return count
}
