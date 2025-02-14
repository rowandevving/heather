package config

import (
	"log"
	"os"
	"time"
)

func Poll(done chan bool) {
	var lastModified time.Time

	for {
		select {
		case <-done:
			log.Println("Stopping config poller")

		default:
			file, err := os.Stat(ConfigPath)
			if err != nil {
				if os.IsNotExist(err) {
					log.Fatal("Config file does not exist")
				} else {
					log.Fatal("Error getting config file: ", err)
				}
				time.Sleep(1 * time.Second)
				continue
			}

			if file.ModTime().After(lastModified) {
				log.Println("Config has changed, reloading...")
				lastModified = file.ModTime()
				LoadConfig()
			}

			time.Sleep(1 * time.Second)
		}
	}
}
