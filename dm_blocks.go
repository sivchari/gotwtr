package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func postDMBlocking(ctx context.Context, c *client, userID string, targetUserID string) (*PostDMBlockingResponse, error) {
	if userID == "" {
		return nil, errors.New("post dm blocking: userID parameter is required")
	}
	if targetUserID == "" {
		return nil, errors.New("post dm blocking: targetUserID parameter is required")
	}

	ep := fmt.Sprintf(postDMBlockingURL, userID)

	body := &DMBlockingBody{
		TargetUserID: targetUserID,
	}

	j, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("post dm blocking json marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep, bytes.NewReader(j))
	if err != nil {
		return nil, fmt.Errorf("post dm blocking new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("post dm blocking response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	var pdbr PostDMBlockingResponse
	if err := json.NewDecoder(resp.Body).Decode(&pdbr); err != nil {
		return nil, fmt.Errorf("post dm blocking decode: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return &pdbr, &HTTPError{
			APIName: "post dm blocking",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &pdbr, nil
}

func undoDMBlocking(ctx context.Context, c *client, userID string, targetUserID string) (*UndoDMBlockingResponse, error) {
	if userID == "" {
		return nil, errors.New("undo dm blocking: userID parameter is required")
	}
	if targetUserID == "" {
		return nil, errors.New("undo dm blocking: targetUserID parameter is required")
	}

	ep := fmt.Sprintf(undoDMBlockingURL, userID)

	body := &DMBlockingBody{
		TargetUserID: targetUserID,
	}

	j, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("undo dm blocking json marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, ep, bytes.NewReader(j))
	if err != nil {
		return nil, fmt.Errorf("undo dm blocking new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("undo dm blocking response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	var udbr UndoDMBlockingResponse
	if err := json.NewDecoder(resp.Body).Decode(&udbr); err != nil {
		return nil, fmt.Errorf("undo dm blocking decode: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &udbr, &HTTPError{
			APIName: "undo dm blocking",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &udbr, nil
}