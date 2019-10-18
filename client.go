package tcgplayer

import (
	"errors"
	"fmt"
)

const (
	baseURL        = "https://api.tcgplayer.com/"
	CurrentVersion = "v1.36.0"
)

var (
	ErrUnauthorized = errors.New("unauthorized access")
)

type AuthToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Issued      string `json:".issued"`
	Expires     string `json:".expires"`
}

type Client struct {
	authToken AuthToken
}

func New(publicKey string, privateKey string) (*Client, error) {

	token, err := getAuthToken(publicKey, privateKey)
	if err != nil {
		return nil, err
	}

	return &Client{
		authToken: AuthToken{
			AccessToken: token.AccessToken,
			TokenType:   token.TokenType,
			Issued:      token.Issued,
			Expires:     token.Expires,
		},
	}, nil
}

func generateURL(path string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, CurrentVersion, path)
}
