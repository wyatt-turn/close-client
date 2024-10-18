package close

import (
	"encoding/json"
	"net/http"
)

type Lead struct {
	ID          string    `json:"id"`
	Contacts    []Contact `json:"contacts"`
	DisplayName string    `json:"display_name"`
	Name        string    `json:"name"`
	DateCreated string    `json:"date_created"`
	DateUpdated string    `json:"date_updated"`
}

type NewLead struct {
	DisplayName string `json:"display_name"`
	Name        string `json:"name"`
}

type UpdateLead struct {
	DisplayName *string `json:"display_name"`
	Name        *string `json:"name"`
}

type leadsResponse struct {
	Data []Lead `json:"data"`
}

type LeadSearchParams struct {
	DisplayName string `json:"display_name"`
}

func (c *CloseClient) GetLeads() ([]Lead, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/lead", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.processRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var leadsResponse leadsResponse
	err = json.NewDecoder(resp.Body).Decode(&leadsResponse)
	if err != nil {
		return nil, err
	}
	return leadsResponse.Data, nil
}

func (c *CloseClient) GetLead(id string) (*Lead, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/lead/"+id, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.processRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var lead Lead
	err = json.NewDecoder(resp.Body).Decode(&lead)
	if err != nil {
		return nil, err
	}
	return &lead, nil
}

func (c *CloseClient) CreateLead(lead NewLead) (*Lead, error) {
	req, err := http.NewRequest("POST", c.BaseURL+"/lead/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.processRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var createdLead Lead
	err = json.NewDecoder(resp.Body).Decode(&createdLead)
	if err != nil {
		return nil, err
	}
	return &createdLead, nil
}

func (c *CloseClient) UpdateLead(id string, lead UpdateLead) (*Lead, error) {
	req, err := http.NewRequest("PUT", c.BaseURL+"/lead/"+id, nil)
	//TODO add payload
	resp, err := c.processRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var updatedLead Lead
	err = json.NewDecoder(resp.Body).Decode(&updatedLead)
	if err != nil {
		return nil, err
	}
	return &updatedLead, nil
}

func (c *CloseClient) DeleteLead(id string) error {
	req, err := http.NewRequest("DELETE", c.BaseURL+"/lead/"+id, nil)
	if err != nil {
		return err
	}

	resp, err := c.processRequest(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}
