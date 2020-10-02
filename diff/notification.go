package diff

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

func (t *Template) notifications() []Notification {
	result := make([]Notification, 0)

	return result
}

func (t *Template) callout() Callout {
	return Callout{}
}
