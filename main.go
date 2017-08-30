package main

import (
	"github.com/Yitsushi/go-commander"
	"github.com/Yitsushi/themis/command"
)

func registerCommands(registry *commander.CommandRegistry) {
	// Register available commands
	registry.Register(command.NewDroplet)
	registry.Register(command.NewVersion)
	registry.Register(command.NewWhoAmI)
}

func main() {
	registry := commander.NewCommandRegistry()
	registerCommands(registry)
	registry.Execute()
}
