package close

type Task struct {
	Type           string `json:"_type"`
	View           string `json:"view"`
	AssignedTo     string `json:"assigned_to"`
	AssignedToName string `json:"assigned_to_name"`
	ContactID      string `json:"contact_id"`
	ContactName    string `json:"contact_name"`
	CreatedBy      string `json:"created_by"`
	CreatedByName  string `json:"created_by_name"`
	Date           string `json:"date"`
	DateCreated    string `json:"date_created"`
	DateUpdated    string `json:"date_updated"`
	ID             string `json:"id"`
	IsComplete     bool   `json:"is_complete"`
	IsDateless     bool   `json:"is_dateless"`
	LeadID         string `json:"lead_id"`
	LeadName       string `json:"lead_name"`
	ObjectID       string `json:"object_id"`
	ObjectType     string `json:"object_type"`
	OrganizationID string `json:"organization_id"`
	Text           string `json:"text"`
	UpdatedBy      string `json:"updated_by"`
	UpdatedByName  string `json:"updated_by_name"`
}
