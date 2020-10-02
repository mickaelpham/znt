package diff

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/mickaelpham/znt/auth"
	"github.com/spf13/viper"
)

// EventType fired when the trigger conditions are met
type EventType struct {
	Description string
	DisplayName string
	Name        string
}

// Trigger fired when the condition is met
type Trigger struct {
	ID          string
	Active      bool
	BaseObject  string
	Condition   string
	Description string
	EventType   EventType
}

const (
	managedTriggerDescription = "trigger managed by znt"
	managedEventDescription   = "event managed by znt"
)

// NewTrigger managed by ZNT
func NewTrigger(baseObject, triggerName, condition string) Trigger {
	name := "znt-" + baseObject + "-on" + strings.Title(triggerName)

	return Trigger{
		Active:      true,
		BaseObject:  baseObject,
		Condition:   condition,
		Description: managedTriggerDescription,
		EventType: EventType{
			Description: managedEventDescription,
			DisplayName: name,
			Name:        name,
		},
	}
}

// Triggers expected from the template
func (t *Template) Triggers() []Trigger {
	result := make([]Trigger, 0)

	for _, n := range t.Notifications {
		for _, t := range n.Triggers {
			result = append(result, NewTrigger(n.BaseObject, t.Name, t.Condition))
		}
	}

	// sort the triggers by event type name
	sort.Slice(result, func(i, j int) bool {
		return result[i].EventType.Name < result[j].EventType.Name
	})

	return result
}

// Equals verify that two triggers base object and condition matches
func (t Trigger) Equals(another Trigger) bool {
	return t.BaseObject == another.BaseObject && t.Condition == another.Condition
}

// LessThan compares two trigger instances based on their base object, further filtered by their condition
func (t Trigger) LessThan(another Trigger) bool {
	return t.BaseObject < another.BaseObject || t.BaseObject == another.BaseObject && t.Condition < another.Condition
}

// Stringer interface
func (t Trigger) String() string {
	return fmt.Sprintf("{%s on %q}", t.BaseObject, t.Condition)
}

// Insert the trigger in the target Zuora environment
func (t Trigger) Insert() error {
	token := auth.NewToken()
	payload, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", viper.GetString("baseurl")+"/events/event-triggers", bytes.NewBuffer(payload))
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

	if response.StatusCode != 201 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(string(body))
	}

	return nil
}

// Destroy the trigger in the targeted Zuora environment
func (t Trigger) Destroy() error {
	if t.ID == "" {
		log.Fatalf("trigger %s doesn't have an ID", t)
	}

	token := auth.NewToken()
	req, err := http.NewRequest("DELETE", viper.GetString("baseurl")+"/events/event-triggers/"+t.ID, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+token.Val)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

	return nil
}
