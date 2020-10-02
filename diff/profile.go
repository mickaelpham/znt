package diff

// Profile is a Zuora communication profile, each customer account is associated to one
type Profile struct {
	ID   string
	Name string `json:"ProfileName"`
}
