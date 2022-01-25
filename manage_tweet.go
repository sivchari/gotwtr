package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func postTweet(ctx context.Context, c *client, body *PostTweetOption) (*PostTweetResponse, error) {
	j, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("postTweet: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, postTweetURL, bytes.NewBuffer(j))
	if err != nil {
		return nil, fmt.Errorf("postTweet new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("post tweet response: %w", err)
	}
	defer resp.Body.Close()

	var postTweet PostTweetResponse
	if err := json.NewDecoder(resp.Body).Decode(&postTweet); err != nil {
		return nil, fmt.Errorf("post tweet decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &postTweet, &HTTPError{
			APIName: "post tweet",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}
	return &postTweet, nil
}

func deleteTweet(ctx context.Context, c *client, tweetID string) (*DeleteTweetResponse, error) {
	if tweetID == "" {
		return nil, errors.New("delete tweet: tweetID parameter is required")
	}
	ep := fmt.Sprintf(deleteTweetURL, tweetID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("delete tweet new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("delete tweet response: %w", err)
	}
	defer resp.Body.Close()

	var deleteTweet DeleteTweetResponse
	if err := json.NewDecoder(resp.Body).Decode(&deleteTweet); err != nil {
		return nil, fmt.Errorf("delete tweet decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &deleteTweet, &HTTPError{
			APIName: "delete tweet",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}
	return &deleteTweet, nil
}
