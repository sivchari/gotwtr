package gotwtr

import (
	"context"
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
	LookUpTweetByID(ctx context.Context, id string, opt ...*TweetOption) (*TweetLookUpByIDResponse, error)
	UserTweetTimeline(ctx context.Context, id string, opt ...*TweetOption) (*UserTweetTimelineResponse, error)
	UserMensionTimeline(ctx context.Context, id string, opt ...*TweetOption) (*UserMensionTimelineResponse, error)
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
	return lookUp(ctx, c, ids, opt...)
}

func (c *client) LookUpTweetByID(ctx context.Context, id string, opt ...*TweetOption) (*TweetLookUpByIDResponse, error) {
	return lookUpByID(ctx, c, id, opt...)
}

func (c *client) UserTweetTimeline(ctx context.Context, id string, opt ...*TweetOption) (*UserTweetTimelineResponse, error) {
	return userTweetTimeline(ctx, c, id, opt...)
}

func (c *client) UserMensionTimeline(ctx context.Context, id string, opt ...*TweetOption) (*UserMensionTimelineResponse, error) {
	return userMensionTimeline(ctx, c, id, opt...)
}
