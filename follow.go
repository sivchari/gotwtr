package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func followers(ctx context.Context, c *client, userID string, opt ...*FollowOption) (*FollowersResponse, error) {
	if userID == "" {
		return nil, errors.New("followers: userID parameter is required")
	}
	ep := fmt.Sprintf(followersURL, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("followers new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var fopt FollowOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		fopt = *opt[0]
	default:
		return nil, errors.New("followers: only one option is allowed")
	}
	fopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("followers response: %w", err)
	}

	defer resp.Body.Close()

	var f FollowersResponse
	if err := json.NewDecoder(resp.Body).Decode(&f); err != nil {
		return nil, fmt.Errorf("followers by id decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &f, &HTTPError{
			APIName: "followers",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &f, nil
}

func following(ctx context.Context, c *client, userID string, opt ...*FollowOption) (*FollowingResponse, error) {
	if userID == "" {
		return nil, errors.New("following: userID parameter is required")
	}
	ep := fmt.Sprintf(followingURL, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("following new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var fopt FollowOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		fopt = *opt[0]
	default:
		return nil, errors.New("following: only one option is allowed")
	}
	fopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("following response: %w", err)
	}

	defer resp.Body.Close()

	var f FollowingResponse
	if err := json.NewDecoder(resp.Body).Decode(&f); err != nil {
		return nil, fmt.Errorf("following: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &f, &HTTPError{
			APIName: "following",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &f, nil
}
