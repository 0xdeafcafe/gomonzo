package gomonzo

import (
	"fmt"

	"time"

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

	resp, monzoError, err := monzo.httpHelper.Post(requestAccessTokenURL, make(map[string]string), form, nil)
	if monzoError != nil || err != nil {
		return nil, monzoError, err
	}

	var token models.Token
	err = helpers.UnmarshalJSON(resp.Body, &token)
	if err != nil {
		return nil, nil, err
	}

	// Minus 10 seconds for possible issues when checking if token as expired
	token.ExpiresAt = time.Now().UTC().Add(time.Duration(token.ExpiresIn-10) * time.Second)
	return &token, nil, nil
}

// RefreshAuthentication refreshes authentication by requesting a new access and refresh
// token.
func (monzo *GoMonzo) RefreshAuthentication(token *models.Token) (*models.Token, *models.MonzoError, error) {
	form := map[string]string{
		"grant_type":    "refresh_token",
		"client_id":     monzo.options.ClientID,
		"client_secret": monzo.options.ClientSecret,
		"refresh_token": token.RefreshToken,
	}

	resp, monzoError, err := monzo.httpHelper.Post(requestAccessTokenURL, make(map[string]string), form, nil)
	if monzoError != nil || err != nil {
		return nil, monzoError, err
	}

	var newToken models.Token
	err = helpers.UnmarshalJSON(resp.Body, &newToken)
	if err != nil {
		return nil, nil, err
	}

	// Minus 10 seconds for possible issues when checking if token as expired
	newToken.ExpiresAt = time.Now().UTC().Add(time.Duration(newToken.ExpiresIn-10) * time.Second)
	return &newToken, nil, nil
}

// RefreshAuthenticationIfNeeded checks of the token has expired, and if it has
func (monzo *GoMonzo) RefreshAuthenticationIfNeeded(token *models.Token) (*models.Token, bool, *models.MonzoError, error) {
	if time.Now().UTC().After(token.ExpiresAt) {
		token, monzoErr, err := monzo.RefreshAuthentication(token)
		return token, true, monzoErr, err
	}
	return token, false, nil, nil
}
