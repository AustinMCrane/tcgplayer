package tcgplayer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	errors "github.com/AustinMCrane/errorutil"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
)

const BaseURL = "https://api.tcgplayer.com/v1.39.0/"

func GetAuthTokenProvider(publicKey string, privateKey string) (RequestEditorFn, error) {
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
		log.Println(response.StatusCode)
		log.Println(response.Body)
		return nil, errors.New("Status code was not 200")
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

	// Example BearerToken
	// See: https://swagger.io/docs/specification/authentication/bearer-authentication/
	bearerTokenProvider, bearerTokenProviderErr := securityprovider.NewSecurityProviderBearerToken(authResponse.AccessToken)
	if bearerTokenProviderErr != nil {
		return nil, errors.New("Error creating bearer token provider")
	}

	return bearerTokenProvider.Intercept, nil
}
