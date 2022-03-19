package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func hideReplies(ctx context.Context, c *client, tweetID string, hidden bool) (*HideRepliesResponse, error) {
	if tweetID == "" {
		return nil, errors.New("hide replies: tweetID parameter is required")
	}
	ep := fmt.Sprintf(hideRepliesURL, tweetID)

	body := &hideRepliesBody{
		Hidden: hidden,
	}
	j, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("hide replies: failed to marshal body")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, ep, bytes.NewBuffer(j))
	if err != nil {
		return nil, fmt.Errorf("hide replies: failed to create request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("hide replies: failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var hideReplies HideRepliesResponse
	if err := json.NewDecoder(resp.Body).Decode(&hideReplies); err != nil {
		return nil, fmt.Errorf("hide replies: failed to decode response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &hideReplies, &HTTPError{
			APIName: "hide replies",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}
	return &hideReplies, nil
}
