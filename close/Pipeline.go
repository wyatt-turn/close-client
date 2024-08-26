package close

type Pipeline struct {
	ID             string              `json:"id"`
	Name           string              `json:"name"`
	Statuses       []OpportunityStatus `json:"statuses"`
	CreatedBy      string              `json:"created_by"`
	UpdatedBy      string              `json:"updated_by"`
	DateCreated    string              `json:"date_created"`
	DateUpdated    string              `json:"date_updated"`
	OrganizationID string              `json:"organization_id"`
}
