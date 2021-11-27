package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func searchRecentTweets(ctx context.Context, c *client, tweet string, opt ...*SearchTweetsOption) (*SearchTweetsResponse, error) {
	switch {
	case tweet == "":
		return nil, errors.New("search recent tweets: tweet parameter is required")
	case len(tweet) > searchTweetMaxQueryLength:
		return nil, errors.New("search recent tweets: tweet parameter must be less than or equal to 512 characters")
	}

	ep := fmt.Sprintf(searchRecentTweetsURL, tweet)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("search recent tweets new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var sopt SearchTweetsOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		sopt = *opt[0]
	default:
		return nil, errors.New("search recent tweets: only one option is allowed")
	}
	sopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("search recent tweets: %w", err)
	}
	defer resp.Body.Close()

	var str SearchTweetsResponse
	if err := json.NewDecoder(resp.Body).Decode(&str); err != nil {
		return nil, fmt.Errorf("search recent tweets: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &str, &HTTPError{
			APIName: "search recent tweets",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &str, nil
}
