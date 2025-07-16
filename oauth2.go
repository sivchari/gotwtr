package gotwtr

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type oauth struct {
	tokenType   string
	accessToken string
	errors      []*APIResponseError
}

func (o *oauth) unmarshalJSON(b io.ReadCloser) error {
	var d struct {
		TokenType   string              `json:"token_type,omitempty"`
		AccessToken string              `json:"access_token,omitempty"`
		Errors      []*APIResponseError `json:"errors,omitempty"`
	}
	if err := json.NewDecoder(b).Decode(&d); err != nil {
		return err
	}
	o.tokenType = d.TokenType
	o.accessToken = d.AccessToken
	o.errors = d.Errors
	return nil
}

func generateAppOnlyBearerToken(ctx context.Context, c *client) (bool, error) {
	ck := c.consumerKey
	cs := c.consumerSecret
	if ck == "" {
		return false, errors.New("consumer key is empty")
	}
	if cs == "" {
		return false, errors.New("consumer secret is empty")
	}
	credentials := ck + ":" + cs
	b64credentials := base64.StdEncoding.EncodeToString([]byte(credentials))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, generateAppOnlyBearerTokenURL, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", "Basic "+b64credentials)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	resp, err := c.client.Do(req)
	if err != nil {
		return false, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return false, &HTTPError{
			APIName: "generate app only bearer token",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	var o oauth
	err = o.unmarshalJSON(resp.Body)
	if err != nil {
		return false, err
	}

	c.bearerToken = o.accessToken

	return true, nil
}
