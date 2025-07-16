package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func pinnedLists(ctx context.Context, c *client, userID string, opt ...*PinnedListsOption) (*PinnedListsResponse, error) {
	if userID == "" {
		return nil, errors.New("pinned lists: userID parameter is required")
	}
	ep := fmt.Sprintf(pinnedListsURL, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("pinned lists new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var lopt PinnedListsOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		lopt = *opt[0]
	default:
		return nil, errors.New("pinned lists: only one option is allowed")
	}
	lopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("pinned lists response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var plr PinnedListsResponse
	if err := json.NewDecoder(resp.Body).Decode(&plr); err != nil {
		return nil, fmt.Errorf("pinned lists decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &plr, &HTTPError{
			APIName: "pinned lists",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &plr, nil
}

func postPinnedLists(ctx context.Context, c *client, listID string, userID string) (*PostPinnedListsResponse, error) {
	if userID == "" {
		return nil, errors.New("post pinned lists: userID parameter is required")
	}
	ep := fmt.Sprintf(postPinnedListsURL, userID)

	if listID == "" {
		return nil, errors.New("post pinned lists: listID parameter is required")
	}
	body := &PinnedListsBody{
		ListID: listID,
	}
	j, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("post pinned lists: can not marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep, bytes.NewBuffer(j))
	if err != nil {
		return nil, fmt.Errorf("post pinned lists new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("post pinned lists response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var ppl PostPinnedListsResponse
	if err := json.NewDecoder(resp.Body).Decode(&ppl); err != nil {
		return nil, fmt.Errorf("post pinned lists decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &ppl, &HTTPError{
			APIName: "post pinned lists",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}
	return &ppl, nil
}

func undoPinnedLists(ctx context.Context, c *client, listID string, userID string) (*UndoPinnedListsResponse, error) {
	if listID == "" {
		return nil, errors.New("undo pinned lists: listID parameter is required")
	}
	if userID == "" {
		return nil, errors.New("undo pinned lists: userID parameter is required")
	}
	ep := fmt.Sprintf(undoPinnedListsURL, userID, listID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("undo pinned lists new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("undo pinned lists response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var upl UndoPinnedListsResponse
	if err := json.NewDecoder(resp.Body).Decode(&upl); err != nil {
		return nil, fmt.Errorf("undo pinned lists decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &upl, &HTTPError{
			APIName: "undo pinned lists",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}
	return &upl, nil
}
