package diff

import (
	"fmt"
	"sort"
	"strings"
)

// EventType fired when the trigger conditions are met
type EventType struct {
	Description string
	DisplayName string
	Name        string
}

// Trigger fired when the condition is met
type Trigger struct {
	Active      bool
	BaseObject  string
	Condition   string
	Description string
	EventType   EventType
}

// TriggerDiff contains the differences between the template and the remote environment
type TriggerDiff struct {
	Add    []Trigger
	Remove []Trigger
	Update []Trigger
}

const (
	managedTriggerDescription = "trigger managed by znt"
	managedEventDescription   = "event managed by znt"
)

func (t *Template) triggers() []Trigger {
	result := make([]Trigger, 0)

	for _, n := range t.Notifications {
		for op, condition := range n.Triggers {
			name := "znt-" + n.BaseObject + "-on" + strings.Title(op)

			result = append(result, Trigger{
				Active:      true,
				BaseObject:  n.BaseObject,
				Condition:   condition,
				Description: managedTriggerDescription,
				EventType: EventType{
					Description: managedEventDescription,
					DisplayName: name,
					Name:        name,
				},
			})
		}
	}

	// sort the triggers by name, since the template layout use a map
	// of short name / condition (map are not guaranteed order when
	// parsing JSON)
	sort.Slice(result, func(i, j int) bool {
		return result[i].EventType.Name < result[j].EventType.Name
	})

	return result
}

// Equals verify that two triggers base object and condition matches
func (t Trigger) Equals(another Trigger) bool {
	return t.BaseObject == another.BaseObject && t.Condition == another.Condition
}

func (t Trigger) String() string {
	return fmt.Sprintf("{%s on %q}", t.BaseObject, t.Condition)
}

// NewDiff accepts sorted trigger arrays and return the diff
func NewDiff(template, remote []Trigger) TriggerDiff {
	result := TriggerDiff{}

	// guard clause
	if len(remote) == 0 {
		result.Add = template
		return result
	}

	j := 0
	for i := range template {
		for !template[i].Equals(remote[j]) {
			result.Remove = append(result.Remove, remote[i])
			j++
		}
	}

	return result
}
