package diff

import (
	"fmt"
	"log"
	"strings"
)

const managedNotificationDescription = "notification managed by znt"

// Notification fires a callout when the associated event is triggered
type Notification struct {
	Active                 bool
	Callout                Callout
	CalloutActive          bool
	CommunicationProfileID string
	Description            string
	EventTypeName          string
	ID                     string
	Name                   string
}

// NotificationDefinitions expected from the template
func (t *Template) NotificationDefinitions(profileIDByName map[string]string) []Notification {
	result := make([]Notification, 0)

	baseCallout := t.Callout
	baseCallout.Active = true
	baseCallout.CalloutRetry = true
	baseCallout.Description = managedNotificationDescription
	baseCallout.HTTPMethod = "POST"
	baseCallout.RequiredAuth = true

	profilesIDs := make([]string, 0)
	for _, profileName := range t.Profiles {
		if profileID, ok := profileIDByName[profileName]; ok {
			profilesIDs = append(profilesIDs, profileID)
		} else {
			log.Fatalf("profile %q not found in Zuora environment", profileName)
		}
	}

	for _, n := range t.Notifications {
		for _, t := range n.Triggers {
			for _, pID := range profilesIDs {
				trigger := NewTrigger(n.BaseObject, t.Name, t.Condition)

				callout := baseCallout

				callout.CalloutParams = n.CalloutParams
				callout.EventTypeName = trigger.EventType.Name
				callout.Name = trigger.EventType.Name

				result = append(result, Notification{
					Active:                 true,
					Callout:                callout,
					CalloutActive:          true,
					CommunicationProfileID: pID,
					Description:            managedNotificationDescription,
					EventTypeName:          trigger.EventType.Name,
					Name:                   trigger.EventType.Name,
				})
			}
		}
	}

	return result
}

// Equals verify that two notification have the same com. profile ID and event type name
func (n Notification) Equals(another Notification) bool {
	return n.CommunicationProfileID == another.CommunicationProfileID && n.EventTypeName == another.EventTypeName
}

// LessThan compares two notifications based on their com. profile ID, further filtered by their event type name
func (n Notification) LessThan(another Notification) bool {
	return n.CommunicationProfileID < another.CommunicationProfileID || n.CommunicationProfileID == another.CommunicationProfileID && n.EventTypeName < another.EventTypeName
}

// NotificationDiff contains the differences between the template and the remote environment
type NotificationDiff struct {
	Add    []Notification
	Remove []Notification
	Update []Notification
}

// NewNotificationDiff accepts sorted trigger arrays and return the diff
func NewNotificationDiff(template, remote []Notification) NotificationDiff {
	result := NotificationDiff{}

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

	// TODO filter out the elements in update which already match with the template
	tmp := make([]Notification, 0)
	for _, needUpdate := range result.Update {
		if !needUpdate.Active {
			tmp = append(tmp, needUpdate)
		}

	}
	result.Update = tmp

	return result
}

func (n Notification) String() string {
	return fmt.Sprintf("(%s) %s", n.CommunicationProfileID, n.EventTypeName)
}

func (d NotificationDiff) String() string {
	var sb strings.Builder

	sb.WriteString("\n--- Notification Diff\n\n")

	if len(d.Add) > 0 {
		sb.WriteString("These notifications will be created: \n")
		for _, t := range d.Add {
			sb.WriteString("  * " + t.String() + "\n")
		}
		sb.WriteString("\n")
	}

	if len(d.Remove) > 0 {
		sb.WriteString("These notifications will be deleted: \n")
		for _, t := range d.Remove {
			sb.WriteString("  * " + t.String() + "\n")
		}
		sb.WriteString("\n")
	}

	if len(d.Update) > 0 {
		sb.WriteString("These notifications will be updated: \n")
		for _, t := range d.Update {
			sb.WriteString("  * " + t.String() + " (activated)\n")
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
