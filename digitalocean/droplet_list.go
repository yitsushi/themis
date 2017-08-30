package digitalocean

import (
	"log"

	"github.com/digitalocean/godo"
	"golang.org/x/net/context"
)

func DropletList(ctx context.Context, client *godo.Client) []godo.Droplet {
	opt := &godo.ListOptions{}
	list := []godo.Droplet{}

	for {
		droplets, resp, err := client.Droplets.List(ctx, opt)
		if err != nil {
			log.Fatalln(err)
			return list
		}

		for _, d := range droplets {
			list = append(list, d)
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			log.Fatalln(err)
			return list
		}

		opt.Page = page + 1
	}

	return list
}
