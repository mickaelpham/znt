package diff

import (
	"encoding/json"
	"io"
)

// Template represents the intended state
type Template struct {
	Callout Callout

	Profiles []string

	Notifications []struct {
		BaseObject string
		Triggers   []struct {
			Name      string
			Condition string
		}
		CalloutParams map[string]string
	}
}

// Parse the input template file
func Parse(r io.Reader) (*Template, error) {
	dec := json.NewDecoder(r)

	var template Template
	err := dec.Decode(&template)
	if err != nil {
		return nil, err
	}

	return &template, nil
}
