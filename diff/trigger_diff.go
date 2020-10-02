package diff

import "strings"

// TriggerDiff contains the differences between the template and the remote environment
type TriggerDiff struct {
	Add    []Trigger
	Remove []Trigger
	Update []Trigger
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

// Apply the trigger diff to the targeted Zuora environment
func (d TriggerDiff) Apply() {
	for _, t := range d.Add {
		t.Insert()
	}

	for _, t := range d.Remove {
		t.Destroy()
	}
}
