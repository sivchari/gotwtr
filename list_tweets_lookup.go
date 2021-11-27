package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func lookUpListTweets(ctx context.Context, c *client, listID string, opt ...*ListTweetsOption) (*ListTweetsResponse, error) {
	if listID == "" {
		return nil, errors.New("look up list tweets: listID parameter is required")
	}
	ep := fmt.Sprintf(lookUpListTweetsURL, listID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("look up list tweets new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var lopt ListTweetsOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		lopt = *opt[0]
	default:
		return nil, errors.New("look up list tweets: only one option is allowed")
	}
	lopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("look up list tweets: %w", err)
	}
	defer resp.Body.Close()

	var ltr ListTweetsResponse
	if err := json.NewDecoder(resp.Body).Decode(&ltr); err != nil {
		return nil, fmt.Errorf("look up list tweets: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &ltr, &HTTPError{
			APIName: "look up list tweets",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &ltr, nil
}
