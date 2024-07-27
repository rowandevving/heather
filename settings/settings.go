package settings

import (
	"encoding/json"
	"log"
	"os"
)

var SettingsPath string
var Config Settings

type Settings struct {
	Token       string   `json:"token"`
	DatabaseDir string   `json:"databaseDir"`
	Prefix      string   `json:"prefix"`
	Tags        []Tag    `json:"tags"`
	Colours     []Colour `json:"colors"`
	Moderation  Moderation
}

type Moderation struct {
	TrustedRole      string `json:"trustedRole"`
	TrustedThreshold uint64 `json:"trustedThreshold"`
}

type Colour struct {
	Name string `json:"name"`
	Hex  string `json:"hex"`
}

type Tag struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	SubTags []Tag  `json:"subtags"`
}

func LoadSettings() {

	raw, err := os.ReadFile(SettingsPath)
	if err != nil {
		log.Fatal("Couldn't read settings file: ", err)
	}

	err = json.Unmarshal([]byte(raw), &Config)
	if err != nil {
		log.Fatal("Couldn't parse settings file: ", err)
	}
}
