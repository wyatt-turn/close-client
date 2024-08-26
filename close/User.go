package close

type User struct {
	ID               string       `json:"id"`
	FirstName        string       `json:"first_name"`
	LastName         string       `json:"last_name"`
	Email            string       `json:"email"`
	Image            string       `json:"image"`
	LastUsedTimezone string       `json:"last_used_timezone"`
	DateCreated      string       `json:"date_created"`
	DateUpdated      string       `json:"date_updated"`
	Memberships      []Membership `json:"memberships"`
	Organizations    []string     `json:"organizations"` //TODO make this the ORganizations struct
	// EmailAccounts    []interface{} `json:"email_accounts"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
}
