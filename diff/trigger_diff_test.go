package diff

import "testing"

func TestTriggerDiff(t *testing.T) {
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

		got := NewTriggerDiff(template, []Trigger{})

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
				Active:     true,
				BaseObject: "Subscription",
				Condition:  "changeType == 'UPDATE'",
			},
		}

		want := TriggerDiff{
			Remove: []Trigger{remote[0]},
		}

		got := NewTriggerDiff(template, remote)

		assertEqual(got, want, t)
	})

	t.Run("remote is different than template", func(t *testing.T) {
		template := []Trigger{
			{
				BaseObject: "Subscription",
				Condition:  "changeType == 'INSERT'",
			},
		}

		remote := []Trigger{
			{
				BaseObject: "Subscription",
				Condition:  "changeType == 'UPDATE'",
			},
		}

		want := TriggerDiff{
			Add:    template,
			Remove: remote,
		}

		got := NewTriggerDiff(template, remote)

		assertEqual(got, want, t)
	})
}
