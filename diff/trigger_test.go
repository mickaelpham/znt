package diff

import (
	"reflect"
	"strings"
	"testing"
)

func TestTriggers(t *testing.T) {
	t.Run("given a single object and two triggers", func(t *testing.T) {

		template, err := Parse(strings.NewReader(`
{
  "notifications": [
    {
      "baseObject": "Account",
      "triggers": {
        "insert": "changeType == 'INSERT'",
        "update": "changeType == 'UPDATE'"
      }
    }
  ]
}
`))

		if err != nil {
			t.Error(err)
		}

		names := []string{
			"znt-Account-onInsert",
			"znt-Account-onUpdate",
		}

		want := []trigger{
			{
				active:      true,
				baseObject:  "Account",
				condition:   "changeType == 'INSERT'",
				description: "trigger managed by znt",
				eventType: eventType{
					description: "event managed by znt",
					displayName: names[0],
					name:        names[0],
				},
			},
			{
				active:      true,
				baseObject:  "Account",
				condition:   "changeType == 'UPDATE'",
				description: "trigger managed by znt",
				eventType: eventType{
					description: "event managed by znt",
					displayName: names[1],
					name:        names[1],
				},
			},
		}

		got := template.triggers()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v want %v given %v", got, want, template)
		}
	})

	t.Run("given two objects and two triggers on each", func(t *testing.T) {

		template, err := Parse(strings.NewReader(`
{
  "notifications": [
    {
      "baseObject": "Account",
      "triggers": {
        "insert": "changeType == 'INSERT'",
        "update": "changeType == 'UPDATE'"
      }
    },
    {
      "baseObject": "Subscription",
      "triggers": {
        "insert": "changeType == 'INSERT'",
        "update": "changeType == 'UPDATE'"
      }
    }
  ]
}
`))
		if err != nil {
			t.Error(err)
		}

		names := []string{
			"znt-Account-onInsert",
			"znt-Account-onUpdate",
			"znt-Subscription-onInsert",
			"znt-Subscription-onUpdate",
		}

		want := []trigger{
			{
				active:      true,
				baseObject:  "Account",
				condition:   "changeType == 'INSERT'",
				description: "trigger managed by znt",
				eventType: eventType{
					description: "event managed by znt",
					displayName: names[0],
					name:        names[0],
				},
			},
			{
				active:      true,
				baseObject:  "Account",
				condition:   "changeType == 'UPDATE'",
				description: "trigger managed by znt",
				eventType: eventType{
					description: "event managed by znt",
					displayName: names[1],
					name:        names[1],
				},
			},
			{
				active:      true,
				baseObject:  "Subscription",
				condition:   "changeType == 'INSERT'",
				description: "trigger managed by znt",
				eventType: eventType{
					description: "event managed by znt",
					displayName: names[2],
					name:        names[2],
				},
			},
			{
				active:      true,
				baseObject:  "Subscription",
				condition:   "changeType == 'UPDATE'",
				description: "trigger managed by znt",
				eventType: eventType{
					description: "event managed by znt",
					displayName: names[3],
					name:        names[3],
				},
			},
		}

		got := template.triggers()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v want %v given %v", got, want, template)
		}
	})
}
