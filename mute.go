package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func muting(ctx context.Context, c *client, userID string, opt ...*MuteOption) (*MutingResponse, error) {
	if userID == "" {
		return nil, errors.New("muting: userID parameter is required")
	}
	ep := fmt.Sprintf(mutingURL, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("muting new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var fopt MuteOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		fopt = *opt[0]
	default:
		return nil, errors.New("muting: only one option is allowed")
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
		return nil, fmt.Errorf("muting: maxResults must be between %d and %d", minimumMaxResults, maximumMaxResults)
	}
	fopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("muting response: %w", err)
	}

	defer resp.Body.Close()

	var m MutingResponse
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return nil, fmt.Errorf("muting: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &m, &HTTPError{
			APIName: "muting",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &m, nil
}

func postMuting(ctx context.Context, c *client, userID string, targetUserID string) (*PostMutingResponse, error) {
	if userID == "" {
		return nil, errors.New("post muting: userID parameter is required")
	}
	ep := fmt.Sprintf(postMutingURL, userID)

	if targetUserID == "" {
		return nil, errors.New("post muting: targetUserID parameter is required")
	}
	body := &MutingBody{
		TargetUserID: targetUserID,
	}
	j, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("post muting json marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep, bytes.NewBuffer(j))
	if err != nil {
		return nil, fmt.Errorf("post muting new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("post muting response: %w", err)
	}
	defer resp.Body.Close()

	var postMuting PostMutingResponse
	if err := json.NewDecoder(resp.Body).Decode(&postMuting); err != nil {
		return nil, fmt.Errorf("post muting decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &postMuting, &HTTPError{
			APIName: "post muting",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &postMuting, nil
}

func undoMuting(ctx context.Context, c *client, sourceUserID string, targetUserID string) (*UndoMutingResponse, error) {
	if sourceUserID == "" {
		return nil, errors.New("undo muting: sourceUserID parameter is required")
	}
	if targetUserID == "" {
		return nil, errors.New("undo muting: targetUserID parameter is required")
	}
	ep := fmt.Sprintf(undoMutingURL, sourceUserID, targetUserID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("undo muting new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("undo muting response: %w", err)
	}
	defer resp.Body.Close()

	var undoMuting UndoMutingResponse
	if err := json.NewDecoder(resp.Body).Decode(&undoMuting); err != nil {
		return nil, fmt.Errorf("undo muting decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &undoMuting, &HTTPError{
			APIName: "undo muting",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &undoMuting, nil
}
