package digitalocean

import "golang.org/x/oauth2"

type Token struct {
	AccessToken string
}

func (t *Token) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}
