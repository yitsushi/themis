package command_droplet

import (
	"fmt"

	"github.com/Yitsushi/go-commander"
	"github.com/Yitsushi/themis/digitalocean"
	"golang.org/x/net/context"
)

type DropletList struct {
}

func (c *DropletList) Execute(opts *commander.CommandHelper) {
	fmt.Printf("[not fully implemented yet]\n")

	client := digitalocean.NewClient()
	ctx := context.TODO()

	droplets := digitalocean.DropletList(ctx, client)

	if len(droplets) < 1 {
		fmt.Printf("  --- No Droplets? Never say no to Cloud!")
		return
	}

	for _, d := range droplets {
		fmt.Println(d)
	}
}

func NewDropletList(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &DropletList{},
		Help: &commander.CommandDescriptor{
			Name:             "list",
			ShortDescription: "List droplets",
		},
	}
}
