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

// NewTriggerDiff accepts sorted trigger arrays and return the diff
func NewTriggerDiff(template, remote []Trigger) TriggerDiff {
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

func (d TriggerDiff) String() string {
	var sb strings.Builder

	sb.WriteString("\n--- Trigger Diff\n\n")

	if len(d.Add) > 0 {
		sb.WriteString("These triggers will be created: \n")
		for _, t := range d.Add {
			sb.WriteString("  * " + t.String() + "\n")
		}
		sb.WriteString("\n")
	}

	if len(d.Remove) > 0 {
		sb.WriteString("These triggers will be deleted: \n")
		for _, t := range d.Remove {
			sb.WriteString("  * " + t.String() + "\n")
		}
		sb.WriteString("\n")
	}

	if len(d.Update) > 0 {
		sb.WriteString("These triggers will be updated: \n")
		for _, t := range d.Update {
			sb.WriteString("  * " + t.String() + " (activated)\n")
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
