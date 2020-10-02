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
