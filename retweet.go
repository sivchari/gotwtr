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

// uid = "user_id" tid = "tweet_id"
func postRetweet(ctx context.Context, c *client, uid string, tid string) (*PostRetweetResponse, error) {
	if uid == "" {
		return nil, errors.New("post retweet by uid: uid parameter is required")
	}
	ep := fmt.Sprintf(postRetweetURL, uid)

	if tid == "" {
		return nil, errors.New("post retweet by tid: tid parameter is required")
	}
	body := &TweetBody{
		TweetID: tid,
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

// stid = "source_tweet_id"
func undoRetweet(ctx context.Context, c *client, id string, stid string) (*UndoRetweetResponse, error) {
	if id == "" {
		return nil, errors.New("undo retweet by id: id parameter is required")
	}
	if stid == "" {
		return nil, errors.New("undo retweet by source_tweet_id: source_tweet_id parameter is required")
	}
	ep := fmt.Sprintf(undoRetweetURL, id, stid)

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
