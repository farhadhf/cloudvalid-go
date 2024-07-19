package cloudvalid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Base URL for the CloudValid API
const BASE_URL string = "https://api.cloudvalid.com/api/v1"

// Full URL for the `Create DNS Configuration Link` API
const CREATE_DNS_CONFIGURATION_LINK_URL = BASE_URL + "/dns-config/create-dns-configuration-link"

// URL template for the `Get DNS Configuration Link` API
const GET_HOSTED_PAGE_URL = BASE_URL + "/dns-config/get-hosted-page/{id}"

// URL template for the `Get Propagation Status` API
const GET_PROPAGATION_STATUS_URL = BASE_URL + "/dns-config/get-propagation-status/{id}"

func NewCloudValidClient(apiKey string) *CloudValidClient {
	return &CloudValidClient{
		APIKey: apiKey,
		client: &http.Client{},
	}
}

func (cv *CloudValidClient) addAPIKeyParam(link string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	queries := u.Query()
	queries.Add("api_key", cv.APIKey)
	u.RawQuery = queries.Encode()
	return u.String(), nil
}

// CreateDNSConfigurationLink issues a POST request to create a DNS configuration link.
func (cv *CloudValidClient) CreateDNSConfigurationLink(payload CreateDNSConfigurationLinkRequest) (*CreateDNSConfigurationLinkResponse, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	u, err := cv.addAPIKeyParam(CREATE_DNS_CONFIGURATION_LINK_URL)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", u, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := cv.client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("create-dns-configuration-link request failed: received status code %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result CreateDNSConfigurationLinkResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetHostedPage issues a GET request to get the information of a previously created
// DNS Configuration Link
func (cv *CloudValidClient) GetHostedPage(id string) (*HostedPageResponse, error) {
	u, err := cv.addAPIKeyParam(strings.ReplaceAll(GET_HOSTED_PAGE_URL, "{id}", id))
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := cv.client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("get-hosted-page request failed: received status code %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result HostedPageResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Use GetPropagationStatus to get the real-time propagation status of the DNS records.
// `id` is the ID of a DNS Configuration Link previously created.
func (cv *CloudValidClient) GetPropagationStatus(id string) (*PropagationStatusResponse, error) {
	u, err := cv.addAPIKeyParam(strings.ReplaceAll(GET_PROPAGATION_STATUS_URL, "{id}", id))
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := cv.client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("get-propagation-status request failed: received status code %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result PropagationStatusResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelDNSConfigurationLink issues a POST to cancel a DNS Configuration link that has not been propagated yet.
func (cv *CloudValidClient) CancelDNSConfigurationLink(id string) (*CancelDNSConfigurationLinkResponse, error) {
	u, err := cv.addAPIKeyParam(strings.ReplaceAll(GET_PROPAGATION_STATUS_URL, "{id}", id))
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := cv.client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("get-propagation-status request failed: received status code %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result CancelDNSConfigurationLinkResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
