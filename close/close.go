package close

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
)

type CloseClient struct {
	BaseURL       string
	encodedAPIKey string
	jsonDecoder   *json.Decoder
	httpClient    *http.Client
}

func NewCloseClient(apiKey string) *CloseClient {
	return &CloseClient{
		BaseURL:       "https://api.close.com/api/v1",
		encodedAPIKey: base64.StdEncoding.EncodeToString([]byte(apiKey + ":")),
		httpClient:    http.DefaultClient,
	}
}

func (c *CloseClient) addDefaultHeaders(req *http.Request) {
	req.Header.Add("Authorization", "Basic "+c.encodedAPIKey)
	req.Header.Add("Content-Type", "application/json")
}

func (c *CloseClient) processRequest(req *http.Request) (*http.Response, error) {

	c.addDefaultHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
