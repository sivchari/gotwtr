package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func retrieveMultipleTweets(ctx context.Context, c *client, tweetIDs []string, opt ...*RetriveTweetOption) (*TweetsResponse, error) {
	switch {
	case len(tweetIDs) == 0:
		return nil, errors.New("retrieve multiple tweets: tweet ids parameter is required")
	case len(tweetIDs) > tweetLookUpMaxIDs:
		return nil, errors.New("retrieve multiple tweets: tweet ids parameter must be less than or equal to 100")
	default:
	}
	ep := retrieveMultipleTweetsURL
	for i, tid := range tweetIDs {
		if i+1 < len(tweetIDs) {
			ep += fmt.Sprintf("%s,", tid)
		} else {
			ep += tid
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve multiple tweets new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var ropt RetriveTweetOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		ropt = *opt[0]
	default:
		return nil, errors.New("retrieve multiple tweets: only one option is allowed")
	}
	ropt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("retrieve multiple tweets response: %w", err)
	}
	defer resp.Body.Close()

	var tweet TweetsResponse
	if err := json.NewDecoder(resp.Body).Decode(&tweet); err != nil {
		return nil, fmt.Errorf("retrieve multiple tweets: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &tweet, &HTTPError{
			APIName: "retrieve multiple tweets",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &tweet, nil
}

func retrieveSingleTweet(ctx context.Context, c *client, tweetID string, opt ...*RetriveTweetOption) (*TweetResponse, error) {
	if tweetID == "" {
		return nil, errors.New("retrieve single tweet: tweet id parameter is required")
	}
	ep := fmt.Sprintf(retrieveSingleTweetURL, tweetID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve single tweet new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var ropt RetriveTweetOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		ropt = *opt[0]
	default:
		return nil, errors.New("retrieve single tweet: only one option is allowed")
	}
	ropt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("retrieve single tweet response: %w", err)
	}

	defer resp.Body.Close()

	var tweet TweetResponse
	if err := json.NewDecoder(resp.Body).Decode(&tweet); err != nil {
		return nil, fmt.Errorf("retrieve single tweet: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &tweet, &HTTPError{
			APIName: "retrieve single tweet",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &tweet, nil
}
