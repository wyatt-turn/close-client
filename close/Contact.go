package close

type Contact struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	Title          string        `json:"title"`
	DateUpdated    string        `json:"date_updated"`
	Phones         []PhoneNumber `json:"phones"`
	CreatedBy      string        `json:"created_by"` //TODO this may be a user struct
	OrganizationID int           `json:"organization_id"`
	DateCreated    string        `json:"date_created"`
	Emails         []Email       `json:"emails"`
}
