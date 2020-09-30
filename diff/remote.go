package diff

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	"github.com/mickaelpham/znt/auth"
	"github.com/spf13/viper"
)

type triggersResponse struct {
	Data []Trigger
	Next string
}

type notificationsResponse struct {
	Data []Notification
	Next string
}

func fetchTriggers() []Trigger {
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

	// sort the triggers by name
	sort.Slice(result, func(i, j int) bool {
		return result[i].EventType.Name < result[j].EventType.Name
	})

	return result
}

// FetchManagedTriggers retrieves all managed triggers from Zuora
func FetchManagedTriggers() []Trigger {
	result := make([]Trigger, 0)

	for _, rmt := range fetchTriggers() {
		if rmt.Description == managedTriggerDescription {
			result = append(result, rmt)
		}
	}

	return result
}

// FetchManagedNotifications retrieves all managed notifications from Zuora
func FetchManagedNotifications() []Notification {
	token := auth.NewToken()
	result := make([]Notification, 0)
	queryPaths := []string{"/notifications/notification-definitions"}

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
		var body notificationsResponse
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

	// sort the triggers by name
	// sort.Slice(result, func(i, j int) bool {
	// 	return result[i].EventType.Name < result[j].EventType.Name
	// })

	return result
}
