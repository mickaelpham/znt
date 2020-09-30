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

		want := []Trigger{
			{
				Active:      true,
				BaseObject:  "Account",
				Condition:   "changeType == 'INSERT'",
				Description: "trigger managed by znt",
				EventType: EventType{
					Description: "event managed by znt",
					DisplayName: names[0],
					Name:        names[0],
				},
			},
			{
				Active:      true,
				BaseObject:  "Account",
				Condition:   "changeType == 'UPDATE'",
				Description: "trigger managed by znt",
				EventType: EventType{
					Description: "event managed by znt",
					DisplayName: names[1],
					Name:        names[1],
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

		want := []Trigger{
			{
				Active:      true,
				BaseObject:  "Account",
				Condition:   "changeType == 'INSERT'",
				Description: "trigger managed by znt",
				EventType: EventType{
					Description: "event managed by znt",
					DisplayName: names[0],
					Name:        names[0],
				},
			},
			{
				Active:      true,
				BaseObject:  "Account",
				Condition:   "changeType == 'UPDATE'",
				Description: "trigger managed by znt",
				EventType: EventType{
					Description: "event managed by znt",
					DisplayName: names[1],
					Name:        names[1],
				},
			},
			{
				Active:      true,
				BaseObject:  "Subscription",
				Condition:   "changeType == 'INSERT'",
				Description: "trigger managed by znt",
				EventType: EventType{
					Description: "event managed by znt",
					DisplayName: names[2],
					Name:        names[2],
				},
			},
			{
				Active:      true,
				BaseObject:  "Subscription",
				Condition:   "changeType == 'UPDATE'",
				Description: "trigger managed by znt",
				EventType: EventType{
					Description: "event managed by znt",
					DisplayName: names[3],
					Name:        names[3],
				},
			},
		}

		got := template.triggers()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v want %v given %v", got, want, template)
		}
	})
}

func TestDiff(t *testing.T) {
	assertEqual := func(got, want TriggerDiff, t *testing.T) {
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
		template := []Trigger{
			{
				BaseObject: "Subscription",
				Condition:  "changeType == 'INSERT'",
			},
		}

		want := TriggerDiff{
			Add: []Trigger{template[0]},
		}

		got := NewDiff(template, []Trigger{})

		assertEqual(got, want, t)
	})

	t.Run("remote is greater than template", func(t *testing.T) {
		template := []Trigger{
			{
				BaseObject: "Subscription",
				Condition:  "changeType == 'UPDATE'",
			},
		}

		remote := []Trigger{
			{
				BaseObject: "Subscription",
				Condition:  "changeType == 'INSERT'",
			},
			{
				BaseObject: "Subscription",
				Condition:  "changeType == 'UPDATE'",
			},
		}

		want := TriggerDiff{
			Remove: []Trigger{remote[0]},
		}

		got := NewDiff(template, remote)

		assertEqual(got, want, t)
	})
}
