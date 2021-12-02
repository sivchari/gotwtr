package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func lookUpList(ctx context.Context, c *client, listID string, opt ...*LookUpListOption) (*ListResponse, error) {
	if listID == "" {
		return nil, errors.New("look up list: listID parameter is required")
	}
	ep := fmt.Sprintf(lookUpListURL, listID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("look up list new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var lopt LookUpListOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		lopt = *opt[0]
	default:
		return nil, errors.New("look up list: only one option is allowed")
	}
	lopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("look up list response: %w", err)
	}
	defer resp.Body.Close()

	var lr ListResponse
	if err := json.NewDecoder(resp.Body).Decode(&lr); err != nil {
		return nil, fmt.Errorf("look up list decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &lr, &HTTPError{
			APIName: "look up list",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &lr, nil
}

func lookUpAllListsOwned(ctx context.Context, c *client, userID string, opt ...*AllListsOwnedOption) (*AllListsOwnedResponse, error) {
	if userID == "" {
		return nil, errors.New("look up all lists owned: userID parameter is required")
	}
	ep := fmt.Sprintf(lookUpAllListsOwnedURL, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("look up all lists owned new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var aopt AllListsOwnedOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		aopt = *opt[0]
	default:
		return nil, errors.New("look up all lists owned: only one option is allowed")
	}
	const (
		minimumMaxResults = 1
		maximumMaxResults = 100
		defaultMaxResults = 100
	)
	if aopt.MaxResults == 0 {
		aopt.MaxResults = defaultMaxResults
	}
	if aopt.MaxResults < minimumMaxResults || aopt.MaxResults > maximumMaxResults {
		return nil, fmt.Errorf("look up all lists owned: max results must be between %d and %d", minimumMaxResults, maximumMaxResults)
	}
	aopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("look up all lists owned response: %w", err)
	}

	defer resp.Body.Close()

	var alor AllListsOwnedResponse
	if err := json.NewDecoder(resp.Body).Decode(&alor); err != nil {
		return nil, fmt.Errorf("look up all lists owned decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &alor, &HTTPError{
			APIName: "look up all lists owned",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &alor, nil
}
