package config

import (
	"encoding/json"
	"log"
	"os"
)

var ConfigPath string
var Global BotConfig

type BotConfig struct {
	DatabaseDir string `json:"databaseDir"`
	Prefix      string `json:"prefix"`
	Tags        []Tag  `json:"tags"`
	Colour      Colour `json:"color"`
	Stats       Stats  `json:"stats"`
	Moderation  Moderation
}

type Stats struct {
	Enabled bool `json:"enabled"`
}

type Moderation struct {
	TrustedRole      string `json:"trustedRole"`
	TrustedThreshold uint64 `json:"trustedThreshold"`
}

type Colour struct {
	Enabled bool      `json:"enabled"`
	Colours []Colours `json:"colors"`
}

type Colours struct {
	Name string `json:"name"`
	Hex  string `json:"hex"`
}

type Tag struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	SubTags []Tag  `json:"subtags"`
}

func LoadConfig() {

	raw, err := os.ReadFile(ConfigPath)
	if err != nil {
		log.Fatal("Couldn't read config file: ", err)
	}

	err = json.Unmarshal([]byte(raw), &Global)
	if err != nil {
		log.Fatal("Couldn't parse config file: ", err)
	}
}
