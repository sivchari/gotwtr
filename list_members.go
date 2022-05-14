package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func listMembers(ctx context.Context, c *client, listID string, opt ...*ListMembersOption) (*ListMembersResponse, error) {
	if listID == "" {
		return nil, errors.New("look up list members: id parameter is required")
	}
	lm := fmt.Sprintf(listMembersURL, listID)

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

func listsSpecifiedUser(ctx context.Context, c *client, userID string, opt ...*ListsSpecifiedUserOption) (*ListsSpecifiedUserResponse, error) {
	if userID == "" {
		return nil, errors.New("lists specified user: userID parameter is required")
	}

	lm := fmt.Sprintf(listsSpecifiedUserURL, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, lm, nil)
	if err != nil {
		return nil, fmt.Errorf("lists specified user: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var lopt ListsSpecifiedUserOption
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
		return nil, fmt.Errorf("lists specified user: max results must be between %d and %d", minimumMaxResults, maximumMaxResults)
	}
	lopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("lists specified user: %w", err)
	}

	defer resp.Body.Close()

	var lmr ListsSpecifiedUserResponse
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

func postListMembers(ctx context.Context, c *client, listID string, userID string) (*PostListMembersResponse, error) {
	if listID == "" {
		return nil, errors.New("post list members: listID parameter is required")
	}
	ep := fmt.Sprintf(postListMembersURL, listID)

	if userID == "" {
		return nil, errors.New("post list members: userID parameter is required")
	}
	body := &ListMembersBody{
		UserID: userID,
	}
	j, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("post list members: can not marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep, bytes.NewBuffer(j))
	if err != nil {
		return nil, fmt.Errorf("post list members new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("post list members response: %w", err)
	}
	defer resp.Body.Close()

	var postListMembers PostListMembersResponse
	if err := json.NewDecoder(resp.Body).Decode(&postListMembers); err != nil {
		return nil, fmt.Errorf("post list members decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &postListMembers, &HTTPError{
			APIName: "post list members",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &postListMembers, nil
}

func undoListMembers(ctx context.Context, c *client, listID string, userID string) (*UndoListMembersResponse, error) {
	if listID == "" {
		return nil, errors.New("undo list members: listID parameter is required")
	}
	if userID == "" {
		return nil, errors.New("undo list members: userID parameter is required")
	}
	ep := fmt.Sprintf(undoListMembersURL, listID, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("undo list members new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("undo list members response: %w", err)
	}
	defer resp.Body.Close()

	var undoListMembers UndoListMembersResponse
	if err := json.NewDecoder(resp.Body).Decode(&undoListMembers); err != nil {
		return nil, fmt.Errorf("undo list members decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &undoListMembers, &HTTPError{
			APIName: "undo list members",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &undoListMembers, nil
}
