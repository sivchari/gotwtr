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
	discoverSpacesMaxIDs        = 100
	filteredStreamRuleMaxLength = 512
	spaceLookUpMaxIDs           = 100
	tweetLookUpMaxIDs           = 100
	searchTweetMaxQueryLength   = 512
	userLookUpMaxIDs            = 100
)

type Tweets interface {
	RetrieveMultipleTweets(ctx context.Context, tweetIDs []string, opt ...*RetriveTweetOption) (*TweetsResponse, error)
	RetrieveSingleTweet(ctx context.Context, tweetID string, opt ...*RetriveTweetOption) (*TweetResponse, error)
	UserMentionTimeline(ctx context.Context, userID string, opt ...*UserMentionTimelineOption) (*UserMentionTimelineResponse, error)
	UserTweetTimeline(ctx context.Context, userID string, opt ...*UserTweetTimelineOption) (*UserTweetTimelineResponse, error)
	SearchRecentTweets(ctx context.Context, tweet string, opt ...*SearchTweetsOption) (*SearchTweetsResponse, error)
	CountsRecentTweet(ctx context.Context, tweet string, opt ...*TweetCountsOption) (*TweetCountsResponse, error)
	AddOrDeleteRules(ctx context.Context, body *AddOrDeleteJSONBody, opt ...*AddOrDeleteRulesOption) (*AddOrDeleteRulesResponse, error)
	RetrieveStreamRules(ctx context.Context, opt ...*RetrieveStreamRulesOption) (*RetrieveStreamRulesResponse, error)
	ConnectToStream(ctx context.Context, ch chan<- ConnectToStreamResponse, errCh chan<- error, opt ...*ConnectToStreamOption) *ConnectToStream
	VolumeStreams(ctx context.Context, ch chan<- VolumeStreamsResponse, errCh chan<- error, opt ...*VolumeStreamsOption) *VolumeStreams
	RetweetsLookup(ctx context.Context, tweetID string, opt ...*RetweetsLookupOption) (*RetweetsResponse, error)
	TweetsUserLiked(ctx context.Context, userID string, opt ...*TweetsUserLikedOpts) (*TweetsUserLikedResponse, error)
	UsersLikingTweet(ctx context.Context, tweetID string, opt ...*UsersLikingTweetOption) (*UsersLikingTweetResponse, error)
}

type Users interface {
	RetrieveMultipleUsersWithIDs(ctx context.Context, userIDs []string, opt ...*RetrieveUserOption) (*UsersResponse, error)
	RetrieveSingleUserWithID(ctx context.Context, userID string, opt ...*RetrieveUserOption) (*UserResponse, error)
	RetrieveMultipleUsersWithUserNames(ctx context.Context, userNames []string, opt ...*RetrieveUserOption) (*UsersResponse, error)
	RetrieveSingleUserWithUserName(ctx context.Context, userName string, opt ...*RetrieveUserOption) (*UserResponse, error)
	Followers(ctx context.Context, userID string, opt ...*FollowOption) (*FollowersResponse, error)
	Following(ctx context.Context, userID string, opt ...*FollowOption) (*FollowingResponse, error)
}

type Spaces interface {
	LookUpSpace(ctx context.Context, spaceID string, opt ...*SpaceOption) (*SpaceResponse, error)
	LookUpSpaces(ctx context.Context, spaceIDs []string, opt ...*SpaceOption) (*SpacesResponse, error)
	UsersPurchasedSpaceTicket(ctx context.Context, spaceID string, opt ...*UsersPurchasedSpaceTicketOption) (*UsersPurchasedSpaceTicketResponse, error)
	DiscoverSpaces(ctx context.Context, userIDs []string, opt ...*DiscoverSpacesOption) (*DiscoverSpacesResponse, error)
	SearchSpaces(ctx context.Context, searchTerm string, opt ...*SearchSpacesOption) (*SearchSpacesResponse, error)
}

type Lists interface {
	LookUpList(ctx context.Context, listID string, opt ...*LookUpListOption) (*ListResponse, error)
	LookUpAllListsOwned(ctx context.Context, userID string, opt ...*AllListsOwnedOption) (*AllListsOwnedResponse, error)
	LookUpListTweets(ctx context.Context, listID string, opt ...*ListTweetsOption) (*ListTweetsResponse, error)
	ListMembers(ctx context.Context, listID string, opt ...*ListMembersOption) (*ListMembersResponse, error)
	LookUpListFollowers(ctx context.Context, listID string, opt ...*ListFollowersOption) (*ListFollowersResponse, error)
	LookUpAllListsUserFollows(ctx context.Context, userID string, opt ...*ListFollowsOption) (*AllListsUserFollowsResponse, error)
}

type Client interface {
	Tweets
	Users
	Spaces
	Lists
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

// RetrieveMultipleTweets returns a variety of information about the Tweet specified by the requested ID or list of IDs.
func (c *client) RetrieveMultipleTweets(ctx context.Context, tweetIDs []string, opt ...*RetriveTweetOption) (*TweetsResponse, error) {
	return retrieveMultipleTweets(ctx, c, tweetIDs, opt...)
}

// RetrieveSingleTweet returns a variety of information about a single Tweet specified by the requested ID.
func (c *client) RetrieveSingleTweet(ctx context.Context, tweetID string, opt ...*RetriveTweetOption) (*TweetResponse, error) {
	return retrieveSingleTweet(ctx, c, tweetID, opt...)
}

// UserMensionTimeline returns Tweets mentioning a single user specified by the requested userID.
// By default, the most recent ten Tweets are returned per request. Using pagination, up to the most recent 800 Tweets can be retrieved.
func (c *client) UserMentionTimeline(ctx context.Context, userID string, opt ...*UserMentionTimelineOption) (*UserMentionTimelineResponse, error) {
	return userMentionTimeline(ctx, c, userID, opt...)
}

// UserTweetTimeline returns Tweets composed by a single user, specified by the requested userID.
// By default, the most recent ten Tweets are returned per request.
// Using pagination, the most recent 3,200 Tweets can be retrieved.
func (c *client) UserTweetTimeline(ctx context.Context, userID string, opt ...*UserTweetTimelineOption) (*UserTweetTimelineResponse, error) {
	return userTweetTimeline(ctx, c, userID, opt...)
}

// SearchRecentTweets returns Tweets from the last seven days that match a search query.
func (c *client) SearchRecentTweets(ctx context.Context, tweet string, opt ...*SearchTweetsOption) (*SearchTweetsResponse, error) {
	return searchRecentTweets(ctx, c, tweet, opt...)
}

// CountsRecentTweet returns count of Tweets from the last seven days that match a query.
func (c *client) CountsRecentTweet(ctx context.Context, tweet string, opt ...*TweetCountsOption) (*TweetCountsResponse, error) {
	return countsRecentTweet(ctx, c, tweet, opt...)
}

// AddOrDeleteRules To create one or more rules, submit an add JSON body with an array of rules and operators.
// Similarly, to delete one or more rules, submit a delete JSON body with an array of list of existing rule IDs.
func (c *client) AddOrDeleteRules(ctx context.Context, body *AddOrDeleteJSONBody, opt ...*AddOrDeleteRulesOption) (*AddOrDeleteRulesResponse, error) {
	return addOrDeleteRules(ctx, c, body, opt...)
}

// RetriveStreamRules return a list of rules currently active on the streaming endpoint, either as a list or individually.
func (c *client) RetrieveStreamRules(ctx context.Context, opt ...*RetrieveStreamRulesOption) (*RetrieveStreamRulesResponse, error) {
	return retrieveStreamRules(ctx, c, opt...)
}

// ConnectToStream streams Tweets in real-time based on a specific set of filter rules.
func (c *client) ConnectToStream(ctx context.Context, ch chan<- ConnectToStreamResponse, errCh chan<- error, opt ...*ConnectToStreamOption) *ConnectToStream {
	return connectToStream(ctx, c, ch, errCh, opt...)
}

// VolumeStreams streams about 1% of all Tweets in real-time.
func (c *client) VolumeStreams(ctx context.Context, ch chan<- VolumeStreamsResponse, errCh chan<- error, opt ...*VolumeStreamsOption) *VolumeStreams {
	return volumeStreams(ctx, c, ch, errCh, opt...)
}

// RetweetsLookup allows you to get information about who has Retweeted a Tweet.
func (c *client) RetweetsLookup(ctx context.Context, tweetID string, opt ...*RetweetsLookupOption) (*RetweetsResponse, error) {
	return retweetsLookup(ctx, c, tweetID, opt...)
}

// TweetsUserLiked allows you to get information about a Tweet’s liking users.
// You will receive the most recent 100 users who liked the specified Tweet.
func (c *client) TweetsUserLiked(ctx context.Context, userID string, opt ...*TweetsUserLikedOpts) (*TweetsUserLikedResponse, error) {
	return tweetsUserLiked(ctx, c, userID, opt...)
}

// UsersLikingTweet allows you to get information about a user’s liked Tweets.
func (c *client) UsersLikingTweet(ctx context.Context, tweetID string, opt ...*UsersLikingTweetOption) (*UsersLikingTweetResponse, error) {
	return usersLikingTweet(ctx, c, tweetID, opt...)
}

// RetrieveMultipleUsersWithIDs returns a variety of information about one or more users specified by the requested userIDs.
func (c *client) RetrieveMultipleUsersWithIDs(ctx context.Context, userIDs []string, opt ...*RetrieveUserOption) (*UsersResponse, error) {
	return retrieveMultipleUsersWithIDs(ctx, c, userIDs, opt...)
}

// RetrieveSingleWithID returns a variety of information about a single user specified by the requested userID.
func (c *client) RetrieveSingleUserWithID(ctx context.Context, userID string, opt ...*RetrieveUserOption) (*UserResponse, error) {
	return retrieveSingleUserWithID(ctx, c, userID, opt...)
}

// RetrieveMultipleUsersWithUserNames returns a variety of information about one or more users specified by their usernames.
func (c *client) RetrieveMultipleUsersWithUserNames(ctx context.Context, userNames []string, opt ...*RetrieveUserOption) (*UsersResponse, error) {
	return retrieveMultipleUsersWithUserNames(ctx, c, userNames, opt...)
}

// RetrieveSingleUserWithUserName returns a variety of information about one or more users specified by their username.
func (c *client) RetrieveSingleUserWithUserName(ctx context.Context, userName string, opt ...*RetrieveUserOption) (*UserResponse, error) {
	return retrieveSingleUserWithUserName(ctx, c, userName, opt...)
}

// Followers returns a list of users who are followers of the specified userID.
func (c *client) Followers(ctx context.Context, userID string, opt ...*FollowOption) (*FollowersResponse, error) {
	return followers(ctx, c, userID, opt...)
}

// Following returns a list of users the specified userID is following.
func (c *client) Following(ctx context.Context, userID string, opt ...*FollowOption) (*FollowingResponse, error) {
	return following(ctx, c, userID, opt...)
}

// LookUpSpace returns a variety of information about a single Space specified by the requested ID.
func (c *client) LookUpSpace(ctx context.Context, spaceID string, opt ...*SpaceOption) (*SpaceResponse, error) {
	return lookUpSpace(ctx, c, spaceID, opt...)
}

// LookUpSpaces returns details about multiple Spaces. Up to 100 comma-separated SpacesIDs can be looked up using this endpoint.
func (c *client) LookUpSpaces(ctx context.Context, spaceIDs []string, opt ...*SpaceOption) (*SpacesResponse, error) {
	return lookUpSpaces(ctx, c, spaceIDs, opt...)
}

// UsersPurchasedSpaceTicket returns a list of user who purchased a ticket to the requested Space.
// You must authenticate the request using the access token of the creator of the requested Space.
func (c *client) UsersPurchasedSpaceTicket(ctx context.Context, spaceID string, opt ...*UsersPurchasedSpaceTicketOption) (*UsersPurchasedSpaceTicketResponse, error) {
	return usersPurchasedSpaceTicket(ctx, c, spaceID, opt...)
}

// DiscoverSpaces returns live or scheduled Spaces created by the specified userIDs.
// Up to 100 comma-separated IDs can be looked up using this endpoint.
func (c *client) DiscoverSpaces(ctx context.Context, userIDs []string, opt ...*DiscoverSpacesOption) (*DiscoverSpacesResponse, error) {
	return discoverSpaces(ctx, c, userIDs, opt...)
}

// SearchSpaces return live or scheduled Spaces matching your specified search terms.
// This endpoint performs a keyword search, meaning that it will return Spaces that are an exact case-insensitive match of the specified search term.
// The search term will match the original title of the Space.
func (c *client) SearchSpaces(ctx context.Context, searchTerm string, opt ...*SearchSpacesOption) (*SearchSpacesResponse, error) {
	return searchSpaces(ctx, c, searchTerm, opt...)
}

// LookUpList returns the details of a specified List.
func (c *client) LookUpList(ctx context.Context, listID string, opt ...*LookUpListOption) (*ListResponse, error) {
	return lookUpList(ctx, c, listID, opt...)
}

// LookUpAllListsOwned returns all Lists owned by the specified user.
func (c *client) LookUpAllListsOwned(ctx context.Context, userID string, opt ...*AllListsOwnedOption) (*AllListsOwnedResponse, error) {
	return lookUpAllListsOwned(ctx, c, userID, opt...)
}

// LookUpListTweets returns a list of Tweets from the specified List.
func (c *client) LookUpListTweets(ctx context.Context, listID string, opt ...*ListTweetsOption) (*ListTweetsResponse, error) {
	return lookUpListTweets(ctx, c, listID, opt...)
}

// ListMembers returns a list of users who are members of the specified List.
func (c *client) ListMembers(ctx context.Context, listid string, opt ...*ListMembersOption) (*ListMembersResponse, error) {
	return listMembers(ctx, c, listid, opt...)
}

// LookUpListFollowers returns a list of users who are followers of the specified List.
func (c *client) LookUpListFollowers(ctx context.Context, listID string, opt ...*ListFollowersOption) (*ListFollowersResponse, error) {
	return lookUpListFollowers(ctx, c, listID, opt...)
}

// LookUpAllListsUserFollows returns all Lists a specified user follows.
func (c *client) LookUpAllListsUserFollows(ctx context.Context, userID string, opt ...*ListFollowsOption) (*AllListsUserFollowsResponse, error) {
	return lookUpAllListsUserFollows(ctx, c, userID, opt...)
}
