package close

type Role struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Editable    bool     `json:"editable"`
	Permissions []string `json:"permissions"`
	// VisibilityUserLcfIDs []string `json:"visibility_user_lcf_ids"`
	// VisibilityUserLcfBehavior string `json:"visibility_user_lcf_behavior"`
}
