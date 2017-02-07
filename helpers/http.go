package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"strings"

	"encoding/json"

	"github.com/0xdeafcafe/gomonzo/models"
)

// HTTPHelper ..
type HTTPHelper struct {
	BaseURL string
	Client  *http.Client
}

const (
	authorizationHeaderContent = "Bearer %s"
)

// Post makes a POST request
func (helper HTTPHelper) Post(endpoint string, params map[string]string, form map[string]string, token *models.Token) (*http.Response, *models.MonzoError, error) {
	return helper.makeRequest("POST", endpoint, params, form, token)
}

// Get ..
func (helper HTTPHelper) Get(endpoint string, params map[string]string, token *models.Token) (*http.Response, *models.MonzoError, error) {
	return helper.makeRequest("GET", endpoint, params, nil, token)
}

// makeRequest ..
func (helper HTTPHelper) makeRequest(method, endpoint string, params map[string]string, form map[string]string, token *models.Token) (*http.Response, *models.MonzoError, error) {
	reqURL, err := formatURL(helper.BaseURL, endpoint, params)
	if err != nil {
		return nil, nil, err
	}

	req, err := createRequest(method, reqURL, form)
	if err != nil {
		return nil, nil, err
	}

	if token != nil {
		req.Header.Add("Authorization", fmt.Sprintf(authorizationHeaderContent, token.AccessToken))
	}

	resp, err := helper.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode == http.StatusOK {
		return resp, nil, nil
	}

	if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		var monzoError models.MonzoError
		UnmarshalJSON(resp.Body, &monzoError)
		return nil, &monzoError, errors.New(monzoError.Code)
	}

	return resp, nil, errors.New(resp.Status)
}

// NewHTTPHelper ..
func NewHTTPHelper() *HTTPHelper {
	return &HTTPHelper{
		BaseURL: "https://api.monzo.com/%s",
		Client:  http.DefaultClient,
	}
}

// UnmarshalJSON takes in a HTTP response body and decodes it into a JSON interface
func UnmarshalJSON(body io.ReadCloser, v interface{}) error {
	defer body.Close()
	err := json.NewDecoder(body).Decode(&v)
	if err != nil {
		return err
	}

	return nil
}

// formatURL creates a URL with base, endpoint and optional params
func formatURL(baseURL, endpoint string, params map[string]string) (*url.URL, error) {
	reqURL, err := url.Parse(fmt.Sprintf(baseURL, endpoint))
	if err != nil {
		return nil, err
	}

	// Add Query Params if they exist
	if len(params) > 0 {
		q := reqURL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		reqURL.RawQuery = q.Encode()
	}

	return reqURL, nil
}

// createRequest creates a request with method, url, and optional form data
func createRequest(method string, reqURL *url.URL, form map[string]string) (*http.Request, error) {
	if len(form) <= 0 {
		return http.NewRequest(method, reqURL.String(), nil)
	}

	// Create form content
	formContent := url.Values{}
	for k, v := range form {
		formContent.Add(k, v)
	}

	// Create Request
	req, err := http.NewRequest(method, reqURL.String(), bytes.NewBufferString(formContent.Encode()))
	if err != nil {
		return nil, err
	}

	// Add form content type
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}
