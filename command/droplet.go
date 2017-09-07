package command

import (
	"github.com/Yitsushi/go-commander"
	subcommand "github.com/Yitsushi/themis/command/droplet"
)

// Droplet structure is the representation of the Droplet command
type Droplet struct {
}

// Execute is the main function. It will be called on version command
func (c *Droplet) Execute(opts *commander.CommandHelper) {
	registry := commander.NewCommandRegistry()
	registry.Depth = 1
	registry.Register(subcommand.NewDropletList)
	registry.Execute()
}

func NewDroplet(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &Droplet{},
		Help: &commander.CommandDescriptor{
			Name:             "droplet",
			ShortDescription: "Manage droplets",
			Arguments:        "subcommand",
			Examples:         []string{"help"},
		},
	}
}
