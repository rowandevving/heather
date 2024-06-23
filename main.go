package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var settingsPath string
var settings Settings

type Settings struct {
	Token       string `json:"token"`
	DatabaseDir string `json:"databaseDir"`
	Tags        []Tag  `json:"tags"`
}

func init() {
	flag.StringVar(&settingsPath, "settings", "", "Path to settings file")
	flag.Parse()
}

func main() {

	loadSettings()

	Token := settings.Token

	connectDatabase(settings.DatabaseDir)

	bot, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatal("Error starting session: ", err)
		return
	}

	bot.AddHandler(ping)
	bot.AddHandler(handleTag)
	bot.AddHandler(incrementCount)

	bot.Identify.Intents = discordgo.IntentsGuildMessages

	err = bot.Open()
	if err != nil {
		log.Fatal("Error connecting to discord: ", err)
		return
	}

	log.Println("Connection opened successfully.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Close()
	db.Close()
}

func loadSettings() {

	raw, err := os.ReadFile(settingsPath)
	if err != nil {
		log.Fatal("Couldn't read settings file: ", err)
	}

	err = json.Unmarshal([]byte(raw), &settings)
	if err != nil {
		log.Fatal("Couldn't parse settings file: ", err)
	}
}

func ping(session *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == session.State.User.ID {
		return
	}

	for _, user := range message.Mentions {
		if user.ID == session.State.User.ID {

			err := session.MessageReactionAdd(message.ChannelID, message.ID, "🥺")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
