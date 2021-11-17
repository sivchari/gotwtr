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
	spaceLookUpMaxIDs           = 100
	tweetLookUpMaxIDs           = 100
	userLookUpMaxIDs            = 100
	tweetSearchMaxQueryLength   = 512
	filteredStreamRuleMaxLength = 512
)

// TODO: Add HideReplies interface
// HideReplises does not handled Twitter v2 API, yet.

type Client interface {
	// CountsFullArchiveTweet(ctx context.Context, query string, opt ...*TweetCountsOption) (*TweetCountsResponse, error)
	AddOrDeleteRules(ctx context.Context, body *AddOrDeleteJSONBody, opt ...*AddOrDeleteRulesOption) (*AddOrDeleteRulesResponse, error)
	CountsRecentTweet(ctx context.Context, query string, opt ...*TweetCountsOption) (*TweetCountsResponse, error)
	ConnectToStream(ctx context.Context, ch chan<- ConnectToStreamResponse, errCh chan<- error, opt ...*ConnectToStreamOption) *ConnectToStream
	Followers(ctx context.Context, id string, opt ...*FollowOption) (*FollowersResponse, error)
	Following(ctx context.Context, id string, opt ...*FollowOption) (*FollowingResponse, error)
	LookUpSpaces(ctx context.Context, ids []string, opt ...*SpaceLookUpOption) (*SpaceLookUpResponse, error)
	LookUpSpaceByID(ctx context.Context, id string, opt ...*SpaceLookUpOption) (*SpaceLookUpByIDResponse, error)
	LookUpTweets(ctx context.Context, ids []string, opt ...*TweetLookUpOption) (*TweetLookUpResponse, error)
	LookUpTweetByID(ctx context.Context, id string, opt ...*TweetLookUpOption) (*TweetLookUpByIDResponse, error)
	RetrieveStreamRules(ctx context.Context, opt ...*RetrieveStreamRulesOption) (*RetrieveStreamRulesResponse, error)
	LookUpUsers(ctx context.Context, ids []string, opt ...*UserLookUpOption) (*UserLookUpResponse, error)
	LookUpUserByID(ctx context.Context, id string, opt ...*UserLookUpOption) (*UserLookUpByIDResponse, error)
	LookUpUserByUserName(ctx context.Context, name string, opt ...*UserLookUpOption) (*UserLookUpByUserNameResponse, error)
	LookUpUsersByUserNames(ctx context.Context, names []string, opt ...*UserLookUpOption) (*UsersLookUpByUserNamesResponse, error)
	// PostFollowing(ctx context.Context, id string, tuid string) (*PostFollowingResponse, error)
	// PostRetweet(ctx context.Context, uid string, tid string) (*PostRetweetResponse, error)
	RetweetsLookup(ctx context.Context, id string, opt ...*RetweetsLookupOpts) (*RetweetsLookupResponse, error)
	SampledStream(ctx context.Context, ch chan<- SampledStreamResponse, errCh chan<- error, opt ...*SampledStreamOpts) *StreamResponse
	// SearchFullArchiveTweets(ctx context.Context, query string, opt ...*TweetSearchOption) (*TweetSearchResponse, error)
	SearchRecentTweets(ctx context.Context, query string, opt ...*TweetSearchOption) (*TweetSearchResponse, error)
	SearchSpaces(ctx context.Context, query string, opt ...*SearchSpacesOption) (*SearchSpacesResponse, error)
	// UndoFollowing(ctx context.Context, suid string, tuid string) (*UndoFollowingResponse, error)
	// UndoRetweet(ctx context.Context, id string, stid string) (*UndoRetweetResponse, error)
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

func (c *client) AddOrDeleteRules(ctx context.Context, body *AddOrDeleteJSONBody, opt ...*AddOrDeleteRulesOption) (*AddOrDeleteRulesResponse, error) {
	return addOrDeleteRules(ctx, c, body, opt...)
}

func (c *client) CountsRecentTweet(ctx context.Context, query string, opt ...*TweetCountsOption) (*TweetCountsResponse, error) {
	return countsRecentTweet(ctx, c, query, opt...)
}

func (c *client) ConnectToStream(ctx context.Context, ch chan<- ConnectToStreamResponse, errCh chan<- error, opt ...*ConnectToStreamOption) *ConnectToStream {
	return connectToStream(ctx, c, ch, errCh, opt...)
}

func (c *client) Followers(ctx context.Context, id string, opt ...*FollowOption) (*FollowersResponse, error) {
	return followers(ctx, c, id, opt...)
}

func (c *client) Following(ctx context.Context, id string, opt ...*FollowOption) (*FollowingResponse, error) {
	return following(ctx, c, id, opt...)
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

func (c *client) LookUpUsers(ctx context.Context, ids []string, opt ...*UserLookUpOption) (*UserLookUpResponse, error) {
	return lookUpUsers(ctx, c, ids, opt...)
}

func (c *client) LookUpUserByID(ctx context.Context, id string, opt ...*UserLookUpOption) (*UserLookUpByIDResponse, error) {
	return lookUpUserByID(ctx, c, id, opt...)
}

func (c *client) LookUpUserByUserName(ctx context.Context, name string, opt ...*UserLookUpOption) (*UserLookUpByUserNameResponse, error) {
	return lookUpUserByUserName(ctx, c, name, opt...)
}

func (c *client) LookUpUsersByUserNames(ctx context.Context, names []string, opt ...*UserLookUpOption) (*UsersLookUpByUserNamesResponse, error) {
	return lookUpUsersByUserNames(ctx, c, names, opt...)
}

func (c *client) PostFollowing(ctx context.Context, id string, tuid string) (*PostFollowingResponse, error) {
	return postFollowing(ctx, c, id, tuid)
}

func (c *client) PostRetweet(ctx context.Context, uid string, tid string) (*PostRetweetResponse, error) {
	return postRetweet(ctx, c, uid, tid)
}

func (c *client) RetweetsLookup(ctx context.Context, id string, opt ...*RetweetsLookupOpts) (*RetweetsLookupResponse, error) {
	return retweetsLookup(ctx, c, id, opt...)
}

func (c *client) SampledStream(ctx context.Context, ch chan<- SampledStreamResponse, errCh chan<- error, opt ...*SampledStreamOpts) *StreamResponse {
	return sampledStream(ctx, c, ch, errCh, opt...)
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

func (c *client) UndoFollowing(ctx context.Context, suid string, tuid string) (*UndoFollowingResponse, error) {
	return undoFollowing(ctx, c, suid, tuid)
}

func (c *client) UndoRetweet(ctx context.Context, id string, stid string) (*UndoRetweetResponse, error) {
	return undoRetweet(ctx, c, id, stid)
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
