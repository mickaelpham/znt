package diff

import "strings"

type eventType struct {
	description string
	displayName string
	name        string
}

type trigger struct {
	active      bool
	baseObject  string
	condition   string
	description string
	eventType   eventType
}

func (t *Template) triggers() []trigger {
	result := make([]trigger, 0)

	for _, n := range t.Notifications {
		for op, condition := range n.Triggers {
			name := "znt-" + n.BaseObject + "-on" + strings.Title(op)

			result = append(result, trigger{
				active:      true,
				baseObject:  n.BaseObject,
				condition:   condition,
				description: "trigger managed by znt",
				eventType: eventType{
					description: "event managed by znt",
					displayName: name,
					name:        name,
				},
			})
		}
	}

	return result
}
