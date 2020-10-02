package diff

import (
	"log"
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

func (t *Template) notifications(profileIDByName map[string]string) []Notification {
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
