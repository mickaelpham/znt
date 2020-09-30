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

// LessThan compares two trigger instances based on their base object, further filtered by their condition
func (t Trigger) LessThan(another Trigger) bool {
	return t.BaseObject < another.BaseObject || t.BaseObject == another.BaseObject && t.Condition < another.Condition
}

// Stringer interface
func (t Trigger) String() string {
	return fmt.Sprintf("{%s on %q}", t.BaseObject, t.Condition)
}

// NewDiff accepts sorted trigger arrays and return the diff
func NewDiff(template, remote []Trigger) TriggerDiff {
	result := TriggerDiff{}

	i := 0
	j := 0

	for i < len(template) && j < len(remote) {
		if template[i].Equals(remote[j]) {
			result.Update = append(result.Update, remote[j])
			i++
			j++
		} else if template[i].LessThan(remote[j]) {
			result.Add = append(result.Add, template[i])
			i++
		} else {
			result.Remove = append(result.Remove, remote[j])
			j++
		}
	}

	// remaining elements of a need to be added
	for i < len(template) {
		result.Add = append(result.Add, template[i])
		i++
	}

	// remaining elements of remote need to be removed
	for j < len(remote) {
		result.Remove = append(result.Remove, remote[j])
		j++
	}

	// filter out the elements in update which are already active
	tmp := make([]Trigger, 0)
	for _, needUpdate := range result.Update {
		if !needUpdate.Active {
			tmp = append(tmp, needUpdate)
		}

	}
	result.Update = tmp

	return result
}
