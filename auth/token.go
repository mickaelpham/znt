package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/viper"
)

// Token is an OAuth token from Zuora with an expiration time
type Token struct {
	Val     string
	expires time.Time
}

type createTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// NewToken generates a new token from Zuora
func NewToken() Token {
	form := url.Values{}
	form.Set("client_id", viper.GetString("client"))
	form.Set("client_secret", viper.GetString("secret"))
	form.Set("grant_type", "client_credentials")

	response, err := http.PostForm(viper.GetString("baseurl")+"/oauth/token", form)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(string(body))
	}

	dec := json.NewDecoder(response.Body)
	var body createTokenResponse
	if err = dec.Decode(&body); err != nil {
		log.Fatal(err)
	}

	return Token{
		Val:     body.AccessToken,
		expires: time.Now().Add(time.Duration(body.ExpiresIn)*time.Second - 15*time.Minute),
	}
}
