package digitalocean

import (
	"github.com/digitalocean/godo"
	"golang.org/x/net/context"
)

func AccountInfo(ctx context.Context, client *godo.Client) *godo.Account {
	account, _, _ := client.Account.Get(ctx)

	return account
}
