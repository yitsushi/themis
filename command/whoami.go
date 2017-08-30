package command

import (
	"fmt"

	"github.com/Yitsushi/go-commander"
	"github.com/Yitsushi/themis/digitalocean"
)

type WhoAmI struct {
}

// Execute is the main function. It will be called on version command
func (c *WhoAmI) Execute(opts *commander.CommandHelper) {
	client := digitalocean.NewClient()
	ctx := digitalocean.NewContext()

	account := digitalocean.AccountInfo(ctx, client)

	fmt.Printf(
		`---
UUID:              %s
Email address:     %s
Email Verified:    %t
Status:            %s
Droplet Limit:     %d
Floating IP Limit: %d
---
`,
		account.UUID,
		account.Email,
		account.EmailVerified,
		account.Status,
		account.DropletLimit,
		account.FloatingIPLimit,
	)
}

func NewWhoAmI(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &WhoAmI{},
		Help: &commander.CommandDescriptor{
			Name:             "whoami",
			ShortDescription: "Info about your current DigitalOcean account",
		},
	}
}
