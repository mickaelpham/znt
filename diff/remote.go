package diff

import (
	"encoding/json"
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
func FetchTriggers() []Trigger {
	token := auth.NewToken()
	result := make([]Trigger, 0)
	queryPaths := []string{"/events/event-triggers"}

	for len(queryPaths) > 0 {
		// pop the path to query
		path := queryPaths[0]
		queryPaths = queryPaths[1:]

		log.Printf("GET %s\n", path)
		req, err := http.NewRequest("GET", viper.GetString("baseurl")+path, nil)
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

		// append the data to the current results and add the
		// next path to query if present
		result = append(result, body.Data...)
		if body.Next != "" {
			queryPaths = append(queryPaths, body.Next)
		}
	}

	return result
}
