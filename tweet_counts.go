package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// This endpoint is only available to those users who have been approved for the Academic Research product track.
func countFullArchiveTweet(ctx context.Context, c *client, query string, opt ...*TweetCountsOption) (*TweetCountsResponse, error) {
	return nil, nil
}

func countsRecentTweet(ctx context.Context, c *client, query string, opt ...*TweetCountsOption) (*TweetCountsResponse, error) {
	switch {
	case len(query) == 0:
		return nil, errors.New("tweet counts: query parameter is required")
	case tweetSearchMaxQueryLength < len(query):
		return nil, errors.New("tweet counts: query parameter length is less than 512")
	default:
	}

	recentTweetCounts := recentTweetCounts + "?query=" + query

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, recentTweetCounts, nil)
	if err != nil {
		return nil, fmt.Errorf("tweet counts new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt TweetCountsOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("tweet counts: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("tweet counts response: %w", err)
	}
	defer resp.Body.Close()

	var tweet TweetCountsResponse
	if err := json.NewDecoder(resp.Body).Decode(&tweet); err != nil {
		return nil, fmt.Errorf("tweet counts decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &tweet, &HTTPError{
			APIName: "tweet counts",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &tweet, nil
}
