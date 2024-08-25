package close

type Email struct {
	Email        string `json:"email"`
	IsSubscribed bool   `json:"is_subscribed"`
	Type         string `json:"type"`
}
