package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func countRecentTweets(ctx context.Context, c *client, tweet string, opt ...*TweetCountsOption) (*TweetCountsResponse, error) {
	if tweet == "" {
		return nil, errors.New("count of recent tweets: tweet parameter is required")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, countsRecentTweetsURL, nil)
	if err != nil {
		return nil, fmt.Errorf("count of recent tweets new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt TweetCountsOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("count of recent tweets: only one option is allowed")
	}
	topt.addQuery(req, tweet)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("count of recent tweets: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var tcr TweetCountsResponse
	if err := json.NewDecoder(resp.Body).Decode(&tcr); err != nil {
		return nil, fmt.Errorf("count of recent tweets: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &tcr, &HTTPError{
			APIName: "count of recent tweets",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &tcr, nil
}

func countAllTweets(ctx context.Context, c *client, tweet string, opt ...*TweetCountsAllOption) (*TweetCountsResponse, error) {
	if tweet == "" {
		return nil, errors.New("count of all tweets: tweet parameter is required")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, countsAllTweetsURL, nil)
	if err != nil {
		return nil, fmt.Errorf("count of all tweets new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt TweetCountsAllOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("count of all tweets: only one option is allowed")
	}
	topt.addQuery(req, tweet)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("count of all tweets: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var tcr TweetCountsResponse
	if err := json.NewDecoder(resp.Body).Decode(&tcr); err != nil {
		return nil, fmt.Errorf("count of all tweets: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &tcr, &HTTPError{
			APIName: "count of all tweets",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}
	return &tcr, nil
}
