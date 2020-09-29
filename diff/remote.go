package diff

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mickaelpham/znt/auth"
	"github.com/spf13/viper"
)

func fetchTriggers() ([]Trigger, error) {
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

	// dec := json.NewDecoder(response.Body)
	return nil, nil
}
