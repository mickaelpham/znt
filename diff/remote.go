package diff

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mickaelpham/znt/auth"
	"github.com/spf13/viper"
)

type triggersResponse struct {
	Data []Trigger
	Next string
}

// FetchTriggers retrive all triggers from Zuora
func FetchTriggers() ([]Trigger, error) {
	token := auth.NewToken()

	req, err := http.NewRequest("GET", viper.GetString("baseurl")+"/events/event-triggers", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.Val)

	response, err := http.DefaultClient.Do(req)
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
	var body triggersResponse

	if err = dec.Decode(&body); err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)
	return body.Data, nil
}
