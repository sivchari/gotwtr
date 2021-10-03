package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func searchRecentTweets(ctx context.Context, c *client, query string, opt ...*TweetSearchOption) (*TweetSearchResponse, error) {
	switch {
	case len(query) == 0:
		return nil, errors.New("tweets search recent: query is required")
	case len(query) > tweetSearchMaxQueryLength:
		return nil, errors.New("tweets search recent: query is too long")
	default:
	}

	tweetRecent := tweetRecentSearch + "?query=" + query

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, tweetRecent, nil)
	if err != nil {
		return nil, fmt.Errorf("tweets search recent: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt TweetSearchOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("tweet lookup: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("tweets search recent: %w", err)
	}
	defer resp.Body.Close()

	var tsr TweetSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&tsr); err != nil {
		return nil, fmt.Errorf("tweets search recent: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			APIName: "tweets search recent",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &tsr, nil
}
