package main

import (
	"strings"

	"github.com/rowandevving/heather/settings"
)

func handleCommand(content string, name string) []string {

	var args []string

	if string(content[0]) != settings.Config.Prefix {
		args = nil
	}

	parts := strings.Fields(content)
	if len(parts) == 0 {
		args = nil
	}

	if parts[0][1:] == name {
		args = parts[1:]
	}

	return args
}
