package close

type EventLog struct {
	ID             string  `json:"id"`
	DateCreated    string  `json:"date_created"`
	DateUpdated    string  `json:"date_updated"`
	OrganizationID string  `json:"organization_id"`
	UserID         *string `json:"user_id"`
	RequestID      string  `json:"request_id"`
	ApiKeyID       *string `json:"api_key_id"`
	ObjectType     string  `json:"object_type"`
	ObjectID       string  `json:"object_id"`
	LeadID         *string `json:"lead_id"`
	Action         string  `json:"action"`
	// ChangedFields  []CustomFields  `json:"changed_fields"`
	// Data interface{} `json:"data"`
	// PreviousData interface{} `json:"previous_data"`
	// Meta interface{} `json:"meta"`
}
