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
