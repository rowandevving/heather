package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/rowandevving/heather/commands"
	"github.com/rowandevving/heather/database"
	"github.com/rowandevving/heather/moderation"
	"github.com/rowandevving/heather/settings"
	"github.com/rowandevving/heather/tags"
)

func init() {
	flag.StringVar(&settings.SettingsPath, "settings", "", "Path to settings file")
	flag.Parse()
}

func main() {

	settings.LoadSettings()

	Token := settings.Config.Token

	database.ConnectDatabase(settings.Config.DatabaseDir)

	bot, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatal("Error starting session: ", err)
		return
	}

	bot.AddHandler(ping)
	bot.AddHandler(tags.HandleTag)
	bot.AddHandler(database.IncrementCount)
	bot.AddHandler(moderation.HandleTrustedRole)

	addCommands(bot)

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
	database.DB.Close()
}

func addCommands(bot *discordgo.Session) {
	bot.AddHandler(commands.StatsCommand)
	bot.AddHandler(commands.ColourCommand)
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
