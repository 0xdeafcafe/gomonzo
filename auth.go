package gomonzo

import (
	"fmt"

	"github.com/0xdeafcafe/gomonzo/helpers"
	"github.com/0xdeafcafe/gomonzo/models"
)

const (
	authorizationURL      = "https://auth.getmondo.co.uk/?client_id=%s&redirect_uri=%s&response_type=code&state=%s"
	requestAccessTokenURL = "oauth2/token"
)

// CreateAuthorizationURL creates the URL the user is redirected towards to grant an application
// access to there account.
func (monzo *GoMonzo) CreateAuthorizationURL(state string) string {
	return fmt.Sprintf(authorizationURL, monzo.options.ClientID, monzo.options.RedirectURL, state)
}

// RequestAccessToken requests a new access token and refresh token from an authorization
// code.
func (monzo *GoMonzo) RequestAccessToken(code string) (*models.Token, *models.MonzoError, error) {
	form := map[string]string{
		"grant_type":    "authorization_code",
		"client_id":     monzo.options.ClientID,
		"client_secret": monzo.options.ClientSecret,
		"redirect_uri":  monzo.options.RedirectURL,
		"code":          code,
	}

	resp, monzoError, err := monzo.httpHelper.PostForm(requestAccessTokenURL, make(map[string]string), form)
	if monzoError != nil || err != nil {
		return nil, monzoError, err
	}

	var token models.Token
	err = helpers.UnmarshalJSON(resp.Body, &token)
	if err != nil {
		return nil, nil, err
	}
	return &token, nil, nil
}
