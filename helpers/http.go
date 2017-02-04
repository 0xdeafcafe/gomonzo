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
	"io/ioutil"

	"github.com/0xdeafcafe/gomonzo/models"
)

// HTTPHelper ..
type HTTPHelper struct {
	BaseURL string
	Client  *http.Client
}

// PostForm ..
func (helper HTTPHelper) PostForm(endpoint string, params map[string]string, form map[string]string) (*http.Response, *models.MonzoError, error) {
	requestURL, err := url.Parse(fmt.Sprintf(helper.BaseURL, endpoint))
	if err != nil {
		return nil, nil, err
	}

	for k, v := range params {
		requestURL.Query().Add(k, v)
	}

	formContent := url.Values{}
	for k, v := range form {
		formContent.Add(k, v)
	}

	req, err := http.NewRequest("POST", requestURL.String(), bytes.NewBufferString(formContent.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, nil, err
	}

	resp, err := helper.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode == http.StatusOK {
		return resp, nil, nil
	}

	if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, nil, err
		}

		var monzoError models.MonzoError
		json.Unmarshal(body, &monzoError)
		return nil, &monzoError, errors.New(monzoError.Error)
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

// UnmarshalJSON takes in a Http Response Body and Unmarshals it into it's JSON interface
func UnmarshalJSON(body io.ReadCloser, v interface{}) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	return nil
}
