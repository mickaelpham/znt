package diff

import (
	"sort"
	"strings"
)

type eventType struct {
	description string
	displayName string
	name        string
}

// Trigger fired when the condition is met
type Trigger struct {
	Active      bool
	BaseObject  string
	Condition   string
	Description string
	EventType   eventType
}

func (t *Template) triggers() []Trigger {
	result := make([]Trigger, 0)

	for _, n := range t.Notifications {
		for op, condition := range n.Triggers {
			name := "znt-" + n.BaseObject + "-on" + strings.Title(op)

			result = append(result, Trigger{
				Active:      true,
				BaseObject:  n.BaseObject,
				Condition:   condition,
				Description: "trigger managed by znt",
				EventType: eventType{
					description: "event managed by znt",
					displayName: name,
					name:        name,
				},
			})
		}
	}

	// sort the triggers by name, since the template layout use a map
	// of short name / condition (map are not guaranteed order when
	// parsing JSON)
	sort.Slice(result, func(i, j int) bool {
		return result[i].EventType.name < result[j].EventType.name
	})

	return result
}
