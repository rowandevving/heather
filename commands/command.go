package commands

import (
	"log"
	"strings"

	"github.com/rowandevving/heather/config"
)

func handleCommand(content string, name string, toggleFlag bool) []string {
	var args []string

	prefix := config.Global.Prefix

	if prefix == "" {
		log.Fatal("No prefix defined in settings")
	}

	if !strings.HasPrefix(content, prefix) || !toggleFlag {
		return nil
	}

	command := content[len(prefix):]
	parts := strings.Fields(command)

	if len(parts) == 0 {
		return nil
	}

	if parts[0] == name {
		args = parts[1:]
	} else {
		return nil
	}

	return args

}
