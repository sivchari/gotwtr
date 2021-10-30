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
	spaceLookUpMaxIDs         = 100
	tweetLookUpMaxIDs         = 100
	tweetSearchMaxQueryLength = 512
)

type Client interface {
	// CountsFullArchiveTweet(ctx context.Context, query string, opt ...*TweetCountsOption) (*TweetCountsResponse, error)
	CountsRecentTweet(ctx context.Context, query string, opt ...*TweetCountsOption) (*TweetCountsResponse, error)
	DeleteRetweet(ctx context.Context, id string, stid string) (*DeleteRetweetResponse, error)
	LookUpSpaces(ctx context.Context, ids []string, opt ...*SpaceLookUpOption) (*SpaceLookUpResponse, error)
	LookUpSpaceByID(ctx context.Context, id string, opt ...*SpaceLookUpOption) (*SpaceLookUpByIDResponse, error)
	LookUpTweets(ctx context.Context, ids []string, opt ...*TweetLookUpOption) (*TweetLookUpResponse, error)
	LookUpTweetByID(ctx context.Context, id string, opt ...*TweetLookUpOption) (*TweetLookUpByIDResponse, error)
	PostRetweet(ctx context.Context, uid string, tid string) (*PostRetweetResponse, error)
	RetweetsLookup(ctx context.Context, id string, opt ...*RetweetsLookupOpts) (*RetweetsLookupResponse, error)
	SampledStream(ctx context.Context, opt ...*SampledStreamOpts) (*SampledStreamResponse, error)
	// SearchFullArchiveTweets(ctx context.Context, query string, opt ...*TweetSearchOption) (*TweetSearchResponse, error)
	SearchRecentTweets(ctx context.Context, query string, opt ...*TweetSearchOption) (*TweetSearchResponse, error)
	SearchSpaces(ctx context.Context, query string, opt ...*SearchSpacesOption) (*SearchSpacesResponse, error)
	UserMentionTimeline(ctx context.Context, id string, opt ...*UserMentionTimelineOpts) (*UserMentionTimelineResponse, error)
	UserTweetTimeline(ctx context.Context, id string, opt ...*UserTweetTimelineOpts) (*UserTweetTimelineResponse, error)
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

func (c *client) DeleteRetweet(ctx context.Context, id string, stid string) (*DeleteRetweetResponse, error) {
	return deleteRetweet(ctx, c, id, stid)
}

func (c *client) LookUpSpaces(ctx context.Context, ids []string, opt ...*SpaceLookUpOption) (*SpaceLookUpResponse, error) {
	return lookUpSpaces(ctx, c, ids, opt...)
}

func (c *client) LookUpSpaceByID(ctx context.Context, id string, opt ...*SpaceLookUpOption) (*SpaceLookUpByIDResponse, error) {
	return lookUpSpaceByID(ctx, c, id, opt...)
}

func (c *client) LookUpTweets(ctx context.Context, ids []string, opt ...*TweetLookUpOption) (*TweetLookUpResponse, error) {
	return lookUpTweets(ctx, c, ids, opt...)
}

func (c *client) LookUpTweetByID(ctx context.Context, id string, opt ...*TweetLookUpOption) (*TweetLookUpByIDResponse, error) {
	return lookUpTweetByID(ctx, c, id, opt...)
}

func (c *client) PostRetweet(ctx context.Context, uid string, tid string) (*PostRetweetResponse, error) {
	return postRetweet(ctx, c, uid, tid)
}

func (c *client) RetweetsLookup(ctx context.Context, id string, opt ...*RetweetsLookupOpts) (*RetweetsLookupResponse, error) {
	return retweetsLookup(ctx, c, id, opt...)
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

func (c *client) SearchSpaces(ctx context.Context, query string, opt ...*SearchSpacesOption) (*SearchSpacesResponse, error) {
	return searchSpaces(ctx, c, query, opt...)
}

func (c *client) UserMentionTimeline(ctx context.Context, id string, opt ...*UserMentionTimelineOpts) (*UserMentionTimelineResponse, error) {
	return userMentionTimeline(ctx, c, id, opt...)
}

func (c *client) UserTweetTimeline(ctx context.Context, id string, opt ...*UserTweetTimelineOpts) (*UserTweetTimelineResponse, error) {
	return userTweetTimeline(ctx, c, id, opt...)
}
