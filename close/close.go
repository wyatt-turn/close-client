package close

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
)

type CloseClient struct {
	BaseURL       string
	encodedAPIKey string
}

func NewCloseClient(apiKey string) *CloseClient {
	return &CloseClient{
		BaseURL:       "https://api.close.com/api/v1",
		encodedAPIKey: base64.StdEncoding.EncodeToString([]byte(apiKey + ":")),
	}
}

func (c *CloseClient) getRequest(model string) (*http.Response, error) {

	req, err := http.NewRequest("GET", c.BaseURL+"/"+model, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+c.encodedAPIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CloseClient) postRequest(model string, record interface{}) (*http.Response, error) {

	payload, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/"+model, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+c.encodedAPIKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CloseClient) putRequest(model string, id string, record interface{}) (*http.Response, error) {

	payload, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", c.BaseURL+"/"+model+"/"+id, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+c.encodedAPIKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CloseClient) deleteRequest(model string, id string) (*http.Response, error) {

	req, err := http.NewRequest("DELETE", c.BaseURL+"/"+model+"/"+id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+c.encodedAPIKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
