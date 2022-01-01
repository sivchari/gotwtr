package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func blocking(ctx context.Context, c *client, userID string, opt ...*BlockOption) (*BlockingResponse, error) {
	if userID == "" {
		return nil, errors.New("blocking: userID parameter is required")
	}
	ep := fmt.Sprintf(blockingURL, userID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("blocking new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	var ropt BlockOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		ropt = *opt[0]
	default:
		return nil, errors.New("blocking: only one option is allowed")
	}
	ropt.addQuery(req)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("blocking response: %w", err)
	}
	defer resp.Body.Close()
	var blocking BlockingResponse
	if err := json.NewDecoder(resp.Body).Decode(&blocking); err != nil {
		return nil, fmt.Errorf("blocking decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &blocking, &HTTPError{
			APIName: "blocking",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &blocking, nil
}

func postBlocking(ctx context.Context, c *client, userID string, targetUserID string) (*PostBlockingResponse, error) {
	if userID == "" {
		return nil, errors.New("post blocking: userID parameter is required")
	}
	ep := fmt.Sprintf(postBlockingURL, userID)

	if targetUserID == "" {
		return nil, errors.New("post blocking: targetUserID parameter is required")
	}
	body := &BlockingBody{
		TargetUserID: targetUserID,
	}
	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("post blocking: can not marshal")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("post blocking new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("post blocking response: %w", err)
	}
	defer resp.Body.Close()

	var postBlocking PostBlockingResponse
	if err := json.NewDecoder(resp.Body).Decode(&postBlocking); err != nil {
		return nil, fmt.Errorf("post blocking decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &postBlocking, &HTTPError{
			APIName: "post blocking",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &postBlocking, nil
}

func undoBlocking(ctx context.Context, c *client, sourceUserID string, targetUserID string) (*UndoBlockingResponse, error) {
	if sourceUserID == "" {
		return nil, errors.New("undo blocking: sourceUserID parameter is required")
	}
	if targetUserID == "" {
		return nil, errors.New("undo blocking: targetUserID parameter is required")
	}
	ep := fmt.Sprintf(undoBlockingURL, sourceUserID, targetUserID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("undo blocking new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("undo blocking response: %w", err)
	}
	defer resp.Body.Close()

	var undoBlocking UndoBlockingResponse
	if err := json.NewDecoder(resp.Body).Decode(&undoBlocking); err != nil {
		return nil, fmt.Errorf("undo blocking decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &undoBlocking, &HTTPError{
			APIName: "undo blocking",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &undoBlocking, nil
}
