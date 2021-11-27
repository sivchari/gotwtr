package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func countsRecentTweet(ctx context.Context, c *client, tweet string, opt ...*TweetCountsOption) (*TweetCountsResponse, error) {
	if tweet == "" {
		return nil, errors.New("counts recent tweets: tweet parameter is required")
	}
	ep := fmt.Sprintf(countsRecentTweetsURL, tweet)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("counts recent tweets new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt TweetCountsOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("counts recent tweets: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("counts recent tweets: %w", err)
	}
	defer resp.Body.Close()

	var tcr TweetCountsResponse
	if err := json.NewDecoder(resp.Body).Decode(&tcr); err != nil {
		return nil, fmt.Errorf("tweet counts: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &tcr, &HTTPError{
			APIName: "counts recent tweets",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &tcr, nil
}
