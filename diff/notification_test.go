package diff

import (
	"reflect"
	"strings"
	"testing"
)

func TestNotifications(t *testing.T) {
	profiles := map[string]string{
		"Profile A": "123456789",
		"Profile B": "987654321",
	}

	names := []string{
		"znt-Account-onInsert",
		"znt-Account-onUpdate",
		"znt-Subscription-onInsert",
		"znt-Subscription-onUpdate",
	}

	t.Run("one object, one trigger, one profile", func(t *testing.T) {
		templateJSON := `
{
  "callout": {
    "calloutAuth": {
      "domain": "example.com",
      "password": "verysecret",
      "preemptive": true,
      "username": "janedoe"
    },
    "calloutBaseurl": "https://example.com/callout"
  },
  "profiles": ["Profile A"],
  "notifications": [
    {
      "baseObject": "Account",
      "triggers": [
        {
          "name": "insert",
          "condition": "changeType == 'INSERT'"
        }
      ],
      "calloutParams": {
        "AccountName": "<Account.Name>"
      }
    }
  ]
}
`
		tpl, err := Parse(strings.NewReader(templateJSON))
		if err != nil {
			t.Error(err)
		}

		got := tpl.notifications(profiles)

		want := []Notification{
			{
				Active: true,
				Callout: Callout{
					Active: true,
					CalloutAuth: CalloutAuth{
						Domain:     "example.com",
						Password:   "verysecret",
						Preemptive: true,
						Username:   "janedoe",
					},
					CalloutBaseURL: "https://example.com/callout",
					CalloutParams: map[string]string{
						"AccountName": "<Account.Name>",
					},
					CalloutRetry:  true,
					Description:   managedNotificationDescription,
					EventTypeName: names[0],
					HTTPMethod:    "POST",
					Name:          names[0],
					RequiredAuth:  true,
				},
				CalloutActive:          true,
				CommunicationProfileID: profiles["Profile A"],
				Description:            managedNotificationDescription,
				EventTypeName:          names[0],
				Name:                   names[0],
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("\ngot:\n%v\nwant:\n%v\ngiven:\n%v", got, want, tpl)
		}

	})

	t.Run("one object, one trigger, two profiles", func(t *testing.T) {
		templateJSON := `
{
  "callout": {
    "calloutAuth": {
      "domain": "example.com",
      "password": "verysecret",
      "preemptive": true,
      "username": "janedoe"
    },
    "calloutBaseurl": "https://example.com/callout"
  },
  "profiles": ["Profile A", "Profile B"],
  "notifications": [
    {
      "baseObject": "Account",
      "triggers": [
        {
          "name": "insert",
          "condition": "changeType == 'INSERT'"
        }
      ],
      "calloutParams": {
        "AccountName": "<Account.Name>"
      }
    }
  ]
}
`
		tpl, err := Parse(strings.NewReader(templateJSON))
		if err != nil {
			t.Error(err)
		}

		got := tpl.notifications(profiles)

		want := []Notification{
			{
				Active: true,
				Callout: Callout{
					Active: true,
					CalloutAuth: CalloutAuth{
						Domain:     "example.com",
						Password:   "verysecret",
						Preemptive: true,
						Username:   "janedoe",
					},
					CalloutBaseURL: "https://example.com/callout",
					CalloutParams: map[string]string{
						"AccountName": "<Account.Name>",
					},
					CalloutRetry:  true,
					Description:   managedNotificationDescription,
					EventTypeName: names[0],
					HTTPMethod:    "POST",
					Name:          names[0],
					RequiredAuth:  true,
				},
				CalloutActive:          true,
				CommunicationProfileID: profiles["Profile A"],
				Description:            managedNotificationDescription,
				EventTypeName:          names[0],
				Name:                   names[0],
			},

			{
				Active: true,
				Callout: Callout{
					Active: true,
					CalloutAuth: CalloutAuth{
						Domain:     "example.com",
						Password:   "verysecret",
						Preemptive: true,
						Username:   "janedoe",
					},
					CalloutBaseURL: "https://example.com/callout",
					CalloutParams: map[string]string{
						"AccountName": "<Account.Name>",
					},
					CalloutRetry:  true,
					Description:   managedNotificationDescription,
					EventTypeName: names[0],
					HTTPMethod:    "POST",
					Name:          names[0],
					RequiredAuth:  true,
				},
				CalloutActive:          true,
				CommunicationProfileID: profiles["Profile B"],
				Description:            managedNotificationDescription,
				EventTypeName:          names[0],
				Name:                   names[0],
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("\ngot:\n%v\nwant:\n%v\ngiven:\n%v", got, want, tpl)
		}
	})

	t.Run("two objects, two triggers, two profiles", func(t *testing.T) {
		templateJSON := `
{
  "callout": {
    "calloutAuth": {
      "domain": "example.com",
      "password": "verysecret",
      "preemptive": true,
      "username": "janedoe"
    },
    "calloutBaseurl": "https://example.com/callout"
  },
  "profiles": ["Profile A", "Profile B"],
  "notifications": [
    {
      "baseObject": "Account",
      "triggers": [
        {
          "name": "insert",
          "condition": "changeType == 'INSERT'"
        },
        {
          "name": "update",
          "condition": "changeType == 'UPDATE'"
        }
      ],
      "calloutParams": {
        "AccountName": "<Account.Name>"
      }
    },
    {
      "baseObject": "Subscription",
      "triggers": [
        {
          "name": "insert",
          "condition": "changeType == 'INSERT'"
        },
        {
          "name": "update",
          "condition": "changeType == 'UPDATE'"
        }
      ],
      "calloutParams": {
	"SubscriptionNumber": "<Subscription.Number>"
      }
    }
  ]
}
`
		tpl, err := Parse(strings.NewReader(templateJSON))
		if err != nil {
			t.Error(err)
		}

		got := tpl.notifications(profiles)

		want := []Notification{
			{
				Active: true,
				Callout: Callout{
					Active: true,
					CalloutAuth: CalloutAuth{
						Domain:     "example.com",
						Password:   "verysecret",
						Preemptive: true,
						Username:   "janedoe",
					},
					CalloutBaseURL: "https://example.com/callout",
					CalloutParams: map[string]string{
						"AccountName": "<Account.Name>",
					},
					CalloutRetry:  true,
					Description:   managedNotificationDescription,
					EventTypeName: names[0],
					HTTPMethod:    "POST",
					Name:          names[0],
					RequiredAuth:  true,
				},
				CalloutActive:          true,
				CommunicationProfileID: profiles["Profile A"],
				Description:            managedNotificationDescription,
				EventTypeName:          names[0],
				Name:                   names[0],
			},

			{
				Active: true,
				Callout: Callout{
					Active: true,
					CalloutAuth: CalloutAuth{
						Domain:     "example.com",
						Password:   "verysecret",
						Preemptive: true,
						Username:   "janedoe",
					},
					CalloutBaseURL: "https://example.com/callout",
					CalloutParams: map[string]string{
						"AccountName": "<Account.Name>",
					},
					CalloutRetry:  true,
					Description:   managedNotificationDescription,
					EventTypeName: names[0],
					HTTPMethod:    "POST",
					Name:          names[0],
					RequiredAuth:  true,
				},
				CalloutActive:          true,
				CommunicationProfileID: profiles["Profile B"],
				Description:            managedNotificationDescription,
				EventTypeName:          names[0],
				Name:                   names[0],
			},
			{
				Active: true,
				Callout: Callout{
					Active: true,
					CalloutAuth: CalloutAuth{
						Domain:     "example.com",
						Password:   "verysecret",
						Preemptive: true,
						Username:   "janedoe",
					},
					CalloutBaseURL: "https://example.com/callout",
					CalloutParams: map[string]string{
						"AccountName": "<Account.Name>",
					},
					CalloutRetry:  true,
					Description:   managedNotificationDescription,
					EventTypeName: names[1],
					HTTPMethod:    "POST",
					Name:          names[1],
					RequiredAuth:  true,
				},
				CalloutActive:          true,
				CommunicationProfileID: profiles["Profile A"],
				Description:            managedNotificationDescription,
				EventTypeName:          names[1],
				Name:                   names[1],
			},

			{
				Active: true,
				Callout: Callout{
					Active: true,
					CalloutAuth: CalloutAuth{
						Domain:     "example.com",
						Password:   "verysecret",
						Preemptive: true,
						Username:   "janedoe",
					},
					CalloutBaseURL: "https://example.com/callout",
					CalloutParams: map[string]string{
						"AccountName": "<Account.Name>",
					},
					CalloutRetry:  true,
					Description:   managedNotificationDescription,
					EventTypeName: names[1],
					HTTPMethod:    "POST",
					Name:          names[1],
					RequiredAuth:  true,
				},
				CalloutActive:          true,
				CommunicationProfileID: profiles["Profile B"],
				Description:            managedNotificationDescription,
				EventTypeName:          names[1],
				Name:                   names[1],
			},

			{
				Active: true,
				Callout: Callout{
					Active: true,
					CalloutAuth: CalloutAuth{
						Domain:     "example.com",
						Password:   "verysecret",
						Preemptive: true,
						Username:   "janedoe",
					},
					CalloutBaseURL: "https://example.com/callout",
					CalloutParams: map[string]string{
						"SubscriptionNumber": "<Subscription.Number>",
					},
					CalloutRetry:  true,
					Description:   managedNotificationDescription,
					EventTypeName: names[2],
					HTTPMethod:    "POST",
					Name:          names[2],
					RequiredAuth:  true,
				},
				CalloutActive:          true,
				CommunicationProfileID: profiles["Profile A"],
				Description:            managedNotificationDescription,
				EventTypeName:          names[2],
				Name:                   names[2],
			},

			{
				Active: true,
				Callout: Callout{
					Active: true,
					CalloutAuth: CalloutAuth{
						Domain:     "example.com",
						Password:   "verysecret",
						Preemptive: true,
						Username:   "janedoe",
					},
					CalloutBaseURL: "https://example.com/callout",
					CalloutParams: map[string]string{
						"SubscriptionNumber": "<Subscription.Number>",
					},
					CalloutRetry:  true,
					Description:   managedNotificationDescription,
					EventTypeName: names[2],
					HTTPMethod:    "POST",
					Name:          names[2],
					RequiredAuth:  true,
				},
				CalloutActive:          true,
				CommunicationProfileID: profiles["Profile B"],
				Description:            managedNotificationDescription,
				EventTypeName:          names[2],
				Name:                   names[2],
			},
			{
				Active: true,
				Callout: Callout{
					Active: true,
					CalloutAuth: CalloutAuth{
						Domain:     "example.com",
						Password:   "verysecret",
						Preemptive: true,
						Username:   "janedoe",
					},
					CalloutBaseURL: "https://example.com/callout",
					CalloutParams: map[string]string{
						"SubscriptionNumber": "<Subscription.Number>",
					},
					CalloutRetry:  true,
					Description:   managedNotificationDescription,
					EventTypeName: names[3],
					HTTPMethod:    "POST",
					Name:          names[3],
					RequiredAuth:  true,
				},
				CalloutActive:          true,
				CommunicationProfileID: profiles["Profile A"],
				Description:            managedNotificationDescription,
				EventTypeName:          names[3],
				Name:                   names[3],
			},

			{
				Active: true,
				Callout: Callout{
					Active: true,
					CalloutAuth: CalloutAuth{
						Domain:     "example.com",
						Password:   "verysecret",
						Preemptive: true,
						Username:   "janedoe",
					},
					CalloutBaseURL: "https://example.com/callout",
					CalloutParams: map[string]string{
						"SubscriptionNumber": "<Subscription.Number>",
					},
					CalloutRetry:  true,
					Description:   managedNotificationDescription,
					EventTypeName: names[3],
					HTTPMethod:    "POST",
					Name:          names[3],
					RequiredAuth:  true,
				},
				CalloutActive:          true,
				CommunicationProfileID: profiles["Profile B"],
				Description:            managedNotificationDescription,
				EventTypeName:          names[3],
				Name:                   names[3],
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("\ngot:\n%v\nwant:\n%v\ngiven:\n%v", got, want, tpl)
		}
	})
}

func TestNotificationDiff(t *testing.T) {
	assertEqual := func(got, want NotificationDiff, t *testing.T) {
		t.Helper()

		if len(got.Add) != len(want.Add) {
			t.Errorf("Add: got %v want %v", got.Add, want.Add)
		}

		if len(got.Remove) != len(want.Remove) {
			t.Errorf("Remove: got %v want %v", got.Remove, want.Remove)
		}

		if len(got.Update) != len(want.Update) {
			t.Errorf("Update: got %v want %v", got.Update, want.Update)
		}
	}

	t.Run("remote is empty", func(t *testing.T) {
		template := []Notification{
			{
				CommunicationProfileID: "profile-id-123",
				EventTypeName:          "znt-Account-onUpdate",
			},
		}

		want := NotificationDiff{
			Add: []Notification{template[0]},
		}

		got := NewNotificationDiff(template, []Notification{})

		assertEqual(got, want, t)
	})

	t.Run("remote is greater than template", func(t *testing.T) {
		template := []Notification{
			{
				CommunicationProfileID: "profile-id-123",
				EventTypeName:          "znt-Account-onUpdate",
			},
		}

		remote := []Notification{
			{
				Active:                 true,
				CommunicationProfileID: "profile-id-123",
				EventTypeName:          "znt-Account-onUpdate",
			},
			{
				CommunicationProfileID: "profile-id-234",
				EventTypeName:          "znt-Account-onUpdate",
			},
		}

		want := NotificationDiff{
			Remove: []Notification{remote[1]},
		}

		got := NewNotificationDiff(template, remote)

		assertEqual(got, want, t)
	})
}
