package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func listFollowers(ctx context.Context, c *client, listID string, opt ...*ListFollowersOption) (*ListFollowersResponse, error) {
	if listID == "" {
		return nil, errors.New("list followers: listID parameter is required")
	}
	ep := fmt.Sprintf(listFollowersURL, listID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("list followers new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var lopt ListFollowersOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		lopt = *opt[0]
	default:
		return nil, errors.New("list followers: only one option is allowed")
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
		return nil, fmt.Errorf("list followers: max results must be between %d and %d", minimumMaxResults, maximumMaxResults)
	}
	lopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("list followers: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var lfr ListFollowersResponse
	if err := json.NewDecoder(resp.Body).Decode(&lfr); err != nil {
		return nil, fmt.Errorf("list followers: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &lfr, &HTTPError{
			APIName: "list followers",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &lfr, nil
}

func allListsUserFollows(ctx context.Context, c *client, userID string, opt ...*ListFollowsOption) (*AllListsUserFollowsResponse, error) {
	if userID == "" {
		return nil, errors.New("all lists user follows: userID parameter is required")
	}
	ep := fmt.Sprintf(allListsUserFollowsURL, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("all lists user follows new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var lopt ListFollowsOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		lopt = *opt[0]
	default:
		return nil, errors.New("all lists user follows: only one option is allowed")
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
		return nil, fmt.Errorf("all lists user follows: max results must be between %d and %d", minimumMaxResults, maximumMaxResults)
	}
	lopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("all lists user follows: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var alufr AllListsUserFollowsResponse
	if err := json.NewDecoder(resp.Body).Decode(&alufr); err != nil {
		return nil, fmt.Errorf("all lists user follows: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &alufr, &HTTPError{
			APIName: "all lists user follows",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &alufr, nil
}

func postListFollows(ctx context.Context, c *client, listID string, userID string) (*PostListFollowsResponse, error) {
	if userID == "" {
		return nil, errors.New("post list follows: userID parameter is required")
	}
	ep := fmt.Sprintf(postListFollowsURL, listID)

	if listID == "" {
		return nil, errors.New("post list follows: listID parameter is required")
	}
	body := &ListFollowsBody{
		ListID: listID,
	}
	j, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("post list follows: can not marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep, bytes.NewBuffer(j))
	if err != nil {
		return nil, fmt.Errorf("post list follows new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("post list follows response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var postListFollows PostListFollowsResponse
	if err := json.NewDecoder(resp.Body).Decode(&postListFollows); err != nil {
		return nil, fmt.Errorf("post list follows decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &postListFollows, &HTTPError{
			APIName: "post list follows",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &postListFollows, nil
}

func undoListFollows(ctx context.Context, c *client, listID string, userID string) (*UndoListFollowsResponse, error) {
	if listID == "" {
		return nil, errors.New("undo list follows: listID parameter is required")
	}
	if userID == "" {
		return nil, errors.New("undo list follows: userID parameter is required")
	}
	ep := fmt.Sprintf(undoListFollowsURL, userID, listID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("undo list follows new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("undo list follows response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var undoListFollows UndoListFollowsResponse
	if err := json.NewDecoder(resp.Body).Decode(&undoListFollows); err != nil {
		return nil, fmt.Errorf("undo list follows decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &undoListFollows, &HTTPError{
			APIName: "undo list follows",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &undoListFollows, nil
}
