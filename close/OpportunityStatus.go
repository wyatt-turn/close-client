package close

type OpportunityStatus struct {
	ID             string `json:"id"`
	Label          string `json:"label"`
	Type           string `json:"type"`
	PipelineID     string `json:"pipeline_id"`
	OrganizationID string `json:"organization_id"`
}
