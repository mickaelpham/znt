package diff

// Callout sent by Zuora
type Callout struct {
	Active         bool
	CalloutAuth    CalloutAuth
	CalloutBaseURL string
	CalloutParams  map[string]string
	CalloutRetry   bool
	Description    string
	EventTypeName  string
	HTTPMethod     string
	ID             string
	Name           string
	RequiredAuth   bool
}

// CalloutAuth sent by Zuora
type CalloutAuth struct {
	Domain     string
	Password   string
	Preemptive bool
	Username   string
}
