package diff

// Callout sent by Zuora
type Callout struct {
	Active         bool
	CalloutAuth    map[string]string
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
