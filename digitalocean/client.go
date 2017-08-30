package digitalocean

import (
	"os"

	"github.com/digitalocean/godo"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func NewClient() *godo.Client {
	oauthClient := oauth2.NewClient(context.Background(), &Token{AccessToken: os.Getenv("DO_TOKEN")})
	return godo.NewClient(oauthClient)
}

func NewContext() context.Context {
	return context.TODO()
}
