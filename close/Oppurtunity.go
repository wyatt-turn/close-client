package close

type Oppurtunity struct {
	ID                      string `json:"id"`
	OrganizationID          string `json:"organization_id"`
	LeadID                  string `json:"lead_id"`
	LeadName                string `json:"lead_name"`
	DateWon                 string `json:"date_won"`
	Confidence              string `json:"confidence"`
	Value                   string `json:"value"`
	ValuePeriod             string `json:"value_period"`
	ValueFormatted          string `json:"value_formatted"`
	ValueCurrency           string `json:"value_currency"`
	ExpectedValue           string `json:"expected_value"`
	AnnualizedValue         string `json:"annualized_value"`
	AnnualizedExpectedValue string `json:"annualized_expected_value"`
	Note                    string `json:"note"`
	StatusID                string `json:"status_id"`
	StatusLabel             string `json:"status_label"`
	StatusType              string `json:"status_type"`
	PipelineID              string `json:"pipeline_id"`
	PipelineName            string `json:"pipeline_name"`
	ContactID               string `json:"contact_id"`
	UserID                  string `json:"user_id"`
	UserName                string `json:"user_name"`
	CreatedBy               string `json:"created_by"`
	UpdatedBy               string `json:"updated_by"`
	DateCreated             string `json:"date_created"`
	DateUpdated             string `json:"date_updated"`
}
