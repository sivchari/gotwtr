package gotwtr

import (
	"bytes"
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
	const (
		minimumMaxResults = 1
		maximumMaxResults = 1000
		defaultMaxResults = 100
	)
	if fopt.MaxResults == 0 {
		fopt.MaxResults = defaultMaxResults
	}
	if fopt.MaxResults < minimumMaxResults || fopt.MaxResults > maximumMaxResults {
		return nil, fmt.Errorf("followers: maxResults must be between %d and %d", minimumMaxResults, maximumMaxResults)
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
	const (
		minimumMaxResults = 1
		maximumMaxResults = 1000
		defaultMaxResults = 100
	)
	if fopt.MaxResults == 0 {
		fopt.MaxResults = defaultMaxResults
	}
	if fopt.MaxResults < minimumMaxResults || fopt.MaxResults > maximumMaxResults {
		return nil, fmt.Errorf("following: maxResults must be between %d and %d", minimumMaxResults, maximumMaxResults)
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

// tuid = target_user_id
func postFollowing(ctx context.Context, c *client, id string, tuid string) (*PostFollowingResponse, error) {
	if id == "" {
		return nil, errors.New("post following by id: id parameter is required")
	}
	ep := fmt.Sprintf(postFollowingURL, id)

	if tuid == "" {
		return nil, errors.New("post following by tuid: tuid parameter is required")
	}
	body := &FollowingBody{
		TargetUserID: tuid,
	}
	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("post following: can not marshal")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("post following new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("post following response: %w", err)
	}
	defer resp.Body.Close()

	var postFollowing PostFollowingResponse
	if err := json.NewDecoder(resp.Body).Decode(&postFollowing); err != nil {
		return nil, fmt.Errorf("post following decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &postFollowing, &HTTPError{
			APIName: "post following",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &postFollowing, nil
}

// suid = source_user_id tuid = target_user_id
func undoFollowing(ctx context.Context, c *client, suid string, tuid string) (*UndoFollowingResponse, error) {
	if suid == "" {
		return nil, errors.New("undo following by suid: suid parameter is required")
	}
	if tuid == "" {
		return nil, errors.New("undo following by tuid: tuid parameter is required")
	}
	ep := fmt.Sprintf(undoFollowingURL, suid, tuid)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("undo following new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("undo following response: %w", err)
	}
	defer resp.Body.Close()

	var undoFollowing UndoFollowingResponse
	if err := json.NewDecoder(resp.Body).Decode(&undoFollowing); err != nil {
		return nil, fmt.Errorf("undo following decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &undoFollowing, &HTTPError{
			APIName: "undo following",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &undoFollowing, nil
}
