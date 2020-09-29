package diff

import (
	"encoding/json"
	"io"
)

type Template struct {
	Callout struct {
		CalloutAuth struct {
			Domain     string
			Password   string
			Preemptive bool
			Username   string
		}
		CalloutBaseurl string
	}

	Profiles []string

	Notifications []struct {
		BaseObject    string
		Triggers      map[string]string
		CalloutParams map[string]string
	}
}

func Parse(r io.Reader) (*Template, error) {
	dec := json.NewDecoder(r)

	var template Template
	err := dec.Decode(&template)
	if err != nil {
		return nil, err
	}

	return &template, nil
}
