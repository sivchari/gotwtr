package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func lookUpListFollowers(ctx context.Context, c *client, listID string, opt ...*ListFollowersOption) (*ListFollowersResponse, error) {
	if listID == "" {
		return nil, errors.New("look up list followers: listID parameter is required")
	}
	ep := fmt.Sprintf(lookUpListFollowersURL, listID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("look up list followers new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var lopt ListFollowersOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		lopt = *opt[0]
	default:
		return nil, errors.New("look up list followers: only one option is allowed")
	}
	lopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("look up list followers: %w", err)
	}
	defer resp.Body.Close()

	var lfr ListFollowersResponse
	if err := json.NewDecoder(resp.Body).Decode(&lfr); err != nil {
		return nil, fmt.Errorf("look up list followers: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &lfr, &HTTPError{
			APIName: "look up list followers",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &lfr, nil
}

func lookUpAllListsUserFollows(ctx context.Context, c *client, userID string, opt ...*ListFollowsOption) (*AllListsUserFollowsResponse, error) {
	if userID == "" {
		return nil, errors.New("look up all lists user follows: userID parameter is required")
	}
	ep := fmt.Sprintf(lookUpAllListsUserFollowsURL, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("look up all lists user follows new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var lopt ListFollowsOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		lopt = *opt[0]
	default:
		return nil, errors.New("look up all lists user follows: only one option is allowed")
	}
	lopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("look up all lists user follows: %w", err)
	}
	defer resp.Body.Close()

	var alufr AllListsUserFollowsResponse
	if err := json.NewDecoder(resp.Body).Decode(&alufr); err != nil {
		return nil, fmt.Errorf("look up all lists user follows: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &alufr, &HTTPError{
			APIName: "look up all lists user follows",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &alufr, nil
}
