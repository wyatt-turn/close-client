package close

type Membership struct {
	ID                    string   `json:"id"`
	PrimaryPersonalNumber string   `json:"primary_personal_number"`
	UserID                string   `json:"user_id"`
	PlanType              string   `json:"plan_type"`
	RoleID                string   `json:"role_id"`
	OrganizationID        string   `json:"organization_id"`
	PermissionsGranted    []string `json:"permissions_granted"`
	RecordCalls           bool     `json:"record_calls"`
}
