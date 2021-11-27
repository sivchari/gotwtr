package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func retweetsLookup(ctx context.Context, c *client, tweetID string, opt ...*RetweetsLookupOption) (*RetweetsLookupResponse, error) {
	if tweetID == "" {
		return nil, errors.New("retweets lookup: tweet id parameter is required")
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
