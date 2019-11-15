package tcgplayer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
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
	authToken  AuthToken
	debug      bool
	httpClient http.Client
}

type response struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}

func (client *Client) get(route string, result interface{}) error {
	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "bearer "+client.authToken.AccessToken)

	// log how long request takes
	startTime := time.Now()

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}

	// if debug print some description of what happened
	if client.debug {
		log.Println("URL: ", req.URL.String())
		log.Println("STATUS: ", resp.StatusCode)
		log.Println("TIME ELAPSED: ", time.Since(startTime))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bodyStr := buf.String()

	// marshal the response to check if it was successful
	var resultStatus response
	err = json.Unmarshal([]byte(bodyStr), &resultStatus)
	if err != nil {
		return err
	}

	if !resultStatus.Success {
		return wrapResponseErrors(resultStatus.Errors)
	}

	err = json.Unmarshal([]byte(bodyStr), &result)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) post(route string, body interface{}, result interface{}) error {
	var reader io.Reader
	var err error
	if body != nil {
		data, err := json.MarshalIndent(body, "", "  ")
		if err != nil {
			return err
		}

		//log.Println("REQ\n", string(data))

		reader = bytes.NewReader(data)
	}

	req, err := http.NewRequest("POST", route, reader)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json;charset=UTF-8")
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Authorization", "bearer "+client.authToken.AccessToken)

	// log how long request takes
	startTime := time.Now()

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}

	// if debug print some description of what happened
	if client.debug {
		log.Println("URL: ", req.URL.String())
		log.Println("STATUS: ", resp.StatusCode)
		log.Println("TIME ELAPSED: ", time.Since(startTime))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bodyStr := buf.String()

	// marshal the response to check if it was successful
	var resultStatus response
	err = json.Unmarshal([]byte(bodyStr), &resultStatus)
	if err != nil {
		return err
	}

	if !resultStatus.Success {
		return wrapResponseErrors(resultStatus.Errors)
	}

	err = json.Unmarshal([]byte(bodyStr), &result)
	if err != nil {
		return err
	}

	return nil
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
