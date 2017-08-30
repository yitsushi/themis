package command

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/Yitsushi/go-commander"
	"github.com/Yitsushi/themis/digitalocean"
)

// Version structure is the representation of the Version command
type Droplet struct {
}

// Execute is the main function. It will be called on version command
func (c *Droplet) Execute(opts *commander.CommandHelper) {
	subcommand := opts.Arg(0)
	if len(subcommand) < 1 {
		panic("Available subcommands: list")
	}

	if subcommand == "list" {
		c.List()
		return
	}

	panic("Available subcommands: list")
}

func (c *Droplet) List() {
	fmt.Printf("[not implemented yet]\n")

	client := digitalocean.NewClient()
	ctx := context.TODO()

	droplets := digitalocean.DropletList(ctx, client)

	for _, d := range droplets {
		fmt.Println(d)
	}
}

func NewDroplet(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &Droplet{},
		Help: &commander.CommandDescriptor{
			Name:             "droplet",
			ShortDescription: "Manage droplets",
			Arguments:        "subcommand",
			Examples:         []string{"list"},
		},
	}
}
