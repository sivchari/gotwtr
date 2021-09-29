package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type client struct {
	bearerToken string
	client      *http.Client
}

const (
	tweetLookUpMaxIDs = 100
)

type Client interface {
	LookUpTweets(ctx context.Context, ids []string, opt ...*TweetOption) (*TweetLookUpResponse, error)
	// User
	// Media
	// Poll
	// Place
}

var _ Client = (*client)(nil)

type ClientOption func(*client)

func New(bearerToken string, opts ...ClientOption) Client {
	c := &client{
		bearerToken: bearerToken,
		client:      http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *client) {
		c.client = httpClient
	}
}

func (c *client) LookUpTweets(ctx context.Context, ids []string, opt ...*TweetOption) (*TweetLookUpResponse, error) {
	// check ids
	switch {
	case len(ids) == 0:
		return nil, errors.New("tweet lookup: ids parameter is required")
	case len(ids) > tweetLookUpMaxIDs:
		return nil, errors.New("tweet lookup: ids parameter must be less than or equal to 100")
	default:
	}

	// join ids to url
	for i, id := range ids {
		if i+1 < len(ids) {
			tweetLookUp += fmt.Sprintf("%s,", id)
		} else {
			tweetLookUp += id
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, tweetLookUp, nil)
	if err != nil {
		return nil, fmt.Errorf("tweet lookup new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt TweetOption
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
		return nil, fmt.Errorf("tweet lookup response: %w", err)
	}
	defer resp.Body.Close()

	var tweet TweetLookUpResponse
	if err := json.NewDecoder(resp.Body).Decode(&tweet); err != nil {
		return nil, fmt.Errorf("tweet lookup decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &tweet, &HTTPError{
			APIName: "tweet lookup",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &tweet, nil
}
