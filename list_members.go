package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func listMembers(ctx context.Context, c *client, listid string, opt ...*ListMembersOption) (*ListMembersResponse, error) {
	if listid == "" {
		return nil, errors.New("look up list members: id parameter is required")
	}
	lm := fmt.Sprintf(listMembersURL, listid)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, lm, nil)
	if err != nil {
		return nil, fmt.Errorf("look up list members: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var lopt ListMembersOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		lopt = *opt[0]
	default:
		return nil, errors.New("look up list members: only one option is allowed")
	}
	const (
		minimumMaxResults = 1
		maximumMaxResults = 100
		defaultMaxResults = 100
	)
	if lopt.MaxResults == 0 {
		lopt.MaxResults = defaultMaxResults
	}
	if lopt.MaxResults < minimumMaxResults || lopt.MaxResults > maximumMaxResults {
		return nil, fmt.Errorf("look up list members: max results must be between %d and %d", minimumMaxResults, maximumMaxResults)
	}
	lopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("look up list members: %w", err)
	}

	defer resp.Body.Close()

	var lmr ListMembersResponse
	if err := json.NewDecoder(resp.Body).Decode(&lmr); err != nil {
		return nil, fmt.Errorf("look up list members: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &lmr, &HTTPError{
			APIName: "owned lists lookup by id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &lmr, nil
}

func listSpecifiedUser(ctx context.Context, c *client, userID string, opt ...*ListSpecifiedUserOption) (*ListSpecifiedUserResponse, error) {
	if userID == "" {
		return nil, errors.New("lists specified user: userID parameter is required")
	}

	lm := fmt.Sprintf(listSpecifiedUserURL, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, lm, nil)
	if err != nil {
		return nil, fmt.Errorf("lists specified user: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var lopt ListSpecifiedUserOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		lopt = *opt[0]
	default:
		return nil, errors.New("lists specified user: only one option is allowed")
	}
	const (
		minimumMaxResults = 1
		maximumMaxResults = 100
		defaultMaxResults = 100
	)
	if lopt.MaxResults == 0 {
		lopt.MaxResults = defaultMaxResults
	}
	if lopt.MaxResults < minimumMaxResults || lopt.MaxResults > maximumMaxResults {
		return nil, fmt.Errorf("ists specified user: max results must be between %d and %d", minimumMaxResults, maximumMaxResults)
	}
	lopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("lists specified user: %w", err)
	}

	defer resp.Body.Close()

	var lmr ListSpecifiedUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&lmr); err != nil {
		return nil, fmt.Errorf("lists specified user: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &lmr, &HTTPError{
			APIName: "lists specified user",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &lmr, nil
}
