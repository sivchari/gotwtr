package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func retweetsLookup(ctx context.Context, c *client, id string, opt ...*RetweetsLookupOpts) (*RetweetsLookupResponse, error) {
	if id == "" {
		return nil, errors.New("retweets lookup by id: id parameter is required")
	}
	retweetsLookupPath := baseTweetPath + "/" + id + "/retweeted_by"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, retweetsLookupPath, nil)
	if err != nil {
		return nil, fmt.Errorf("retweets lookup new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var ropt RetweetsLookupOpts
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

	var retweetsLookup RetweetsLookupResponse
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

// uid = "user_id" tid = "tweet_id"
func postRetweet(ctx context.Context, c *client, uid string, tid string) (*PostRetweetResponse, error) {
	return nil, nil
}

// stid = "source_tweet_id"
func undoRetweet(ctx context.Context, c *client, id string, stid string) (*UndoRetweetResponse, error) {
	return nil, nil
}
