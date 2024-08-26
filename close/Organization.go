package close

type Organization struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	LeadStatuses []LeadStatus `json:"lead_statuses"`
	// Pipelines []Pipeline `json:"pipelines"`
	CreatedBy           string       `json:"created_by"`
	Memberships         []Membership `json:"memberships"`
	InactiveMemberships []Membership `json:"inactive_memberships"`
	DateUpdated         string       `json:"date_updated"`
	DateCreated         string       `json:"date_created"`
	PlanType            string       `json:"plan_type"`
	UpdatedBy           string       `json:"updated_by"`
}
