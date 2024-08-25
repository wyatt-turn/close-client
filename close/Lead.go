package close

type Lead struct {
	ID          string    `json:"id"`
	Contacts    []Contact `json:"contacts"`
	DisplayName string    `json:"display_name"`
	Name        string    `json:"name"`
	DateCreated string    `json:"date_created"`
	DateUpdated string    `json:"date_updated"`
}
