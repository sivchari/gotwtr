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
	filteredStreamRuleMaxLength = 512
	tweetLookUpMaxIDs           = 100
	tweetSearchMaxQueryLength   = 512
)

type Client interface {
	// CountsFullArchiveTweet(ctx context.Context, query string, opt ...*TweetCountsOption) (*TweetCountsResponse, error)
	AddOrDeleteRules(ctx context.Context, body *AddOrDeleteJSONBody, opt ...*AddOrDeleteRulesOption) (*AddOrDeleteRulesResponse, error)
	ConnectToStream(ctx context.Context, ch chan<- ConnectToStreamResponse, errCh chan<- error, opt ...*ConnectToStreamOption)
	CountsRecentTweet(ctx context.Context, query string, opt ...*TweetCountsOption) (*TweetCountsResponse, error)
	LookUpSpaceByID(ctx context.Context, id string, opt ...*SpaceLookUpOption) (*SpaceLookUpByIDResponse, error)
	LookUpTweets(ctx context.Context, ids []string, opt ...*TweetLookUpOption) (*TweetLookUpResponse, error)
	LookUpTweetByID(ctx context.Context, id string, opt ...*TweetLookUpOption) (*TweetLookUpByIDResponse, error)
	RetrieveStreamRules(ctx context.Context, opt ...*RetrieveStreamRulesOption) (*RetrieveStreamRulesResponse, error)
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

func (c *client) ConnectToStream(ctx context.Context, ch chan<- ConnectToStreamResponse, errCh chan<- error, opt ...*ConnectToStreamOption) {
	connectToStream(ctx, c, ch, errCh, opt...)
}

func (c *client) LookUpSpaceByID(ctx context.Context, id string, opt ...*SpaceLookUpOption) (*SpaceLookUpByIDResponse, error) {
	return lookUpSpaceByID(ctx, c, id, opt...)
}

func (c *client) LookUpTweets(ctx context.Context, ids []string, opt ...*TweetLookUpOption) (*TweetLookUpResponse, error) {
	return lookUp(ctx, c, ids, opt...)
}

func (c *client) LookUpTweetByID(ctx context.Context, id string, opt ...*TweetLookUpOption) (*TweetLookUpByIDResponse, error) {
	return lookUpByID(ctx, c, id, opt...)
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

func (c *client) RetrieveStreamRules(ctx context.Context, opt ...*RetrieveStreamRulesOption) (*RetrieveStreamRulesResponse, error) {
	return retrieveStreamRules(ctx, c, opt...)
}

func (c *client) AddOrDeleteRules(ctx context.Context, body *AddOrDeleteJSONBody, opt ...*AddOrDeleteRulesOption) (*AddOrDeleteRulesResponse, error) {
	return addOrDeleteRules(ctx, c, body, opt...)
}
