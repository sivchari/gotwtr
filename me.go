package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func me(ctx context.Context, c *client, opt ...*MeOption) (*MeResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, meURL, nil)
	if err != nil {
		return nil, fmt.Errorf("me new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var mopt MeOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		mopt = *opt[0]
	default:
		return nil, errors.New("me: only one option is allowed")
	}
	mopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("me response: %w", err)
	}
	defer resp.Body.Close()

	var me MeResponse
	if err := json.NewDecoder(resp.Body).Decode(&me); err != nil {
		return nil, fmt.Errorf("me decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &me, &HTTPError{
			APIName: "me",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}
	return &me, nil
}
