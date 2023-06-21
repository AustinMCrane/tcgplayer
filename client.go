package tcgplayer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

var (
	BaseURL        = "https://api.tcgplayer.com/"
	CurrentVersion = "v1.39.0"
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

type APIResponse struct {
	Results interface{} `json:"results"`
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
	return fmt.Sprintf("%s/%s/%s", BaseURL, CurrentVersion, path)
}

func get(client *Client, path string, p APIParams, apiResponse interface{}) error {

	c := &http.Client{}

	url := generateURL(path)
	req, err := http.NewRequest("GET", url, nil)
	/*
		q := req.URL.Query()
		if p != nil {
			p.SetQueryParams(&q)
		}*/
	req.Header.Set("Authorization", "bearer "+client.authToken.AccessToken)
	if err != nil {
		return err
	}
	res, err := c.Do(req)
	if err != nil {
		return errors.Wrap(err, "response error")
	}

	/*
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)
	*/

	err = json.NewDecoder(res.Body).Decode(&apiResponse)
	if err != nil {
		return errors.Wrap(err, "json parsing error")
	}

	if res.StatusCode != 200 {
		log.Println(res.StatusCode)
		return errors.New("not 200")
	}

	return nil
}
