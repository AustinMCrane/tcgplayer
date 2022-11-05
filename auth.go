package tcgplayer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AuthTokenResponse struct {
	AuthToken
}

func getAuthToken(publicKey string, privateKey string) (*AuthTokenResponse, error) {
	u := BaseURL + "token"
	response, err := http.PostForm(
		u,
		url.Values{
			"grant_type":    {"client_credentials"},
			"client_id":     {publicKey},
			"client_secret": {privateKey},
		},
	)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, ErrUnauthorized
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var authResponse AuthTokenResponse
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		return nil, err
	}

	return &authResponse, nil
}
