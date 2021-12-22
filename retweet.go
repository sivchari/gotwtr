package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func retweetsLookup(ctx context.Context, c *client, tweetID string, opt ...*RetweetsLookupOption) (*RetweetsResponse, error) {
	if tweetID == "" {
		return nil, errors.New("retweets lookup by tweetID: tweetID parameter is required")
	}
	ep := fmt.Sprintf(retweetsLookupURL, tweetID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("retweets lookup new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var ropt RetweetsLookupOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		ropt = *opt[0]
	default:
		return nil, errors.New("retweets lookup: only one option is allowed")
	}
	ropt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("retweets lookup response: %w", err)
	}
	defer resp.Body.Close()

	var retweetsLookup RetweetsResponse
	if err := json.NewDecoder(resp.Body).Decode(&retweetsLookup); err != nil {
		return nil, fmt.Errorf("retweets lookup decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &retweetsLookup, &HTTPError{
			APIName: "retweets lookup",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &retweetsLookup, nil
}

func postRetweet(ctx context.Context, c *client, userID string, tweetID string) (*PostRetweetResponse, error) {
	if userID == "" {
		return nil, errors.New("post retweet by userID: userID parameter is required")
	}
	ep := fmt.Sprintf(postRetweetURL, userID)

	if tweetID == "" {
		return nil, errors.New("post retweet by tweetID: tweetID parameter is required")
	}
	body := &TweetBody{
		TweetID: tweetID,
	}
	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("post retweet: can not marshal")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("post retweet new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("post retweet response: %w", err)
	}
	defer resp.Body.Close()

	var postRetweet PostRetweetResponse
	if err := json.NewDecoder(resp.Body).Decode(&postRetweet); err != nil {
		return nil, fmt.Errorf("post retweet decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &postRetweet, &HTTPError{
			APIName: "post retweet",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &postRetweet, nil
}

func undoRetweet(ctx context.Context, c *client, userID string, sourceTweetID string) (*UndoRetweetResponse, error) {
	if userID == "" {
		return nil, errors.New("undo retweet by userID: userID parameter is required")
	}
	if sourceTweetID == "" {
		return nil, errors.New("undo retweet by sourceTweetID: sourceTweetID parameter is required")
	}
	ep := fmt.Sprintf(undoRetweetURL, userID, sourceTweetID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("undo retweet new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("undo retweet response: %w", err)
	}
	defer resp.Body.Close()

	var undoRetweet UndoRetweetResponse
	if err := json.NewDecoder(resp.Body).Decode(&undoRetweet); err != nil {
		return nil, fmt.Errorf("undo retweet decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &undoRetweet, &HTTPError{
			APIName: "undo retweet",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &undoRetweet, nil
}
