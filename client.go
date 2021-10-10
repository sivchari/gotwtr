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
	tweetLookUpMaxIDs         = 100
	tweetSearchMaxQueryLength = 512
)

type Client interface {
	// CountsFullArchiveTweet(ctx context.Context, query string, opt ...*TweetCountsOption) (*TweetCountsResponse, error)
	CountsRecentTweet(ctx context.Context, query string, opt ...*TweetCountsOption) (*TweetCountsResponse, error)
	LookUpTweets(ctx context.Context, ids []string, opt ...*TweetOption) (*TweetLookUpResponse, error)
	LookUpTweetByID(ctx context.Context, id string, opt ...*TweetOption) (*TweetLookUpByIDResponse, error)
	SampledStream(ctx context.Context, opt ...*SampledStreamOpts) (*SampledStreamResponse, error)
	// SearchFullArchiveTweets(ctx context.Context, query string, opt ...*TweetSearchOption) (*TweetSearchResponse, error)
	SearchRecentTweets(ctx context.Context, query string, opt ...*TweetSearchOption) (*TweetSearchResponse, error)
	UserMentionTimeline(ctx context.Context, id string, opt ...*UserMentionTimelineOpts) (*UserMentionTimelineResponse, error)
	UserTweetTimeline(ctx context.Context, id string, opt ...*UserTweetTimelineOpts) (*UserTweetTimelineResponse, error)
	RetweetsLookup(ctx context.Context, id string, opt ...*RetweetsLookupOpts) (*RetweetsLookupResponse, error)
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

// func (c *client) CountsFullArchiveTweet(ctx context.Context, query string, opt ...*TweetCountsOption) (*TweetCountsResponse, error) {
// 	return countFullArchiveTweet(ctx, c, query, opt...)
// }

func (c *client) CountsRecentTweet(ctx context.Context, query string, opt ...*TweetCountsOption) (*TweetCountsResponse, error) {
	return countsRecentTweet(ctx, c, query, opt...)
}

func (c *client) LookUpTweets(ctx context.Context, ids []string, opt ...*TweetOption) (*TweetLookUpResponse, error) {
	return lookUp(ctx, c, ids, opt...)
}

func (c *client) LookUpTweetByID(ctx context.Context, id string, opt ...*TweetOption) (*TweetLookUpByIDResponse, error) {
	return lookUpByID(ctx, c, id, opt...)
}

func (c *client) SampledStream(ctx context.Context, opt ...*SampledStreamOpts) (*SampledStreamResponse, error) {
	return sampledStream(ctx, c, opt...)
}

// func (c *client) SearchFullArchiveTweets(ctx context.Context, query string, opt ...*TweetSearchOption) (*TweetSearchResponse, error) {
// 	return searchFullArchiveTweets(ctx, c, query, opt...)
// }

func (c *client) SearchRecentTweets(ctx context.Context, query string, opt ...*TweetSearchOption) (*TweetSearchResponse, error) {
	return searchRecentTweets(ctx, c, query, opt...)
}

func (c *client) UserMentionTimeline(ctx context.Context, id string, opt ...*UserMentionTimelineOpts) (*UserMentionTimelineResponse, error) {
	return userMentionTimeline(ctx, c, id, opt...)
}

func (c *client) UserTweetTimeline(ctx context.Context, id string, opt ...*UserTweetTimelineOpts) (*UserTweetTimelineResponse, error) {
	return userTweetTimeline(ctx, c, id, opt...)
}

func (c *client) RetweetsLookup(ctx context.Context, id string, opt ...*RetweetsLookupOpts) (*RetweetsLookupResponse, error) {
	return retweetsLookup(ctx, c, id, opt...)
}
