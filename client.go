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
	RetriveMultipleTweets(ctx context.Context, tweetIDs []string, opt ...*RetriveTweetOption) (*TweetsResponse, error)
	RetriveSingleTweet(ctx context.Context, tweetID string, opt ...*RetriveTweetOption) (*TweetResponse, error)
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
	ListSpecifiedUser(ctx context.Context, userID string, opt ...*ListSpecifiedUserOption) (*ListSpecifiedUserResponse, error)
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

func (c *client) RetriveMultipleTweets(ctx context.Context, tweetIDs []string, opt ...*RetriveTweetOption) (*TweetsResponse, error) {
	return retrieveMultipleTweets(ctx, c, tweetIDs, opt...)
}

func (c *client) RetriveSingleTweet(ctx context.Context, tweetID string, opt ...*RetriveTweetOption) (*TweetResponse, error) {
	return retrieveSingleTweet(ctx, c, tweetID, opt...)
}

func (c *client) UserMentionTimeline(ctx context.Context, userID string, opt ...*UserMentionTimelineOption) (*UserMentionTimelineResponse, error) {
	return userMentionTimeline(ctx, c, userID, opt...)
}

func (c *client) UserTweetTimeline(ctx context.Context, userID string, opt ...*UserTweetTimelineOption) (*UserTweetTimelineResponse, error) {
	return userTweetTimeline(ctx, c, userID, opt...)
}

func (c *client) SearchRecentTweets(ctx context.Context, tweet string, opt ...*SearchTweetsOption) (*SearchTweetsResponse, error) {
	return searchRecentTweets(ctx, c, tweet, opt...)
}

func (c *client) CountsRecentTweet(ctx context.Context, tweet string, opt ...*TweetCountsOption) (*TweetCountsResponse, error) {
	return countsRecentTweet(ctx, c, tweet, opt...)
}

func (c *client) AddOrDeleteRules(ctx context.Context, body *AddOrDeleteJSONBody, opt ...*AddOrDeleteRulesOption) (*AddOrDeleteRulesResponse, error) {
	return addOrDeleteRules(ctx, c, body, opt...)
}

func (c *client) RetrieveStreamRules(ctx context.Context, opt ...*RetrieveStreamRulesOption) (*RetrieveStreamRulesResponse, error) {
	return retrieveStreamRules(ctx, c, opt...)
}

func (c *client) ConnectToStream(ctx context.Context, ch chan<- ConnectToStreamResponse, errCh chan<- error, opt ...*ConnectToStreamOption) *ConnectToStream {
	return connectToStream(ctx, c, ch, errCh, opt...)
}

func (c *client) VolumeStreams(ctx context.Context, ch chan<- VolumeStreamsResponse, errCh chan<- error, opt ...*VolumeStreamsOption) *VolumeStreams {
	return volumeStreams(ctx, c, ch, errCh, opt...)
}

func (c *client) RetweetsLookup(ctx context.Context, tweetID string, opt ...*RetweetsLookupOption) (*RetweetsResponse, error) {
	return retweetsLookup(ctx, c, tweetID, opt...)
}

func (c *client) TweetsUserLiked(ctx context.Context, userID string, opt ...*TweetsUserLikedOpts) (*TweetsUserLikedResponse, error) {
	return tweetsUserLiked(ctx, c, userID, opt...)
}

func (c *client) UsersLikingTweet(ctx context.Context, tweetID string, opt ...*UsersLikingTweetOption) (*UsersLikingTweetResponse, error) {
	return usersLikingTweet(ctx, c, tweetID, opt...)
}

func (c *client) RetrieveMultipleUsersWithIDs(ctx context.Context, userIDs []string, opt ...*RetrieveUserOption) (*UsersResponse, error) {
	return retrieveMultipleUsersWithIDs(ctx, c, userIDs, opt...)
}

func (c *client) RetrieveSingleUserWithID(ctx context.Context, userID string, opt ...*RetrieveUserOption) (*UserResponse, error) {
	return retrieveSingleUserWithID(ctx, c, userID, opt...)
}

func (c *client) RetrieveMultipleUsersWithUserNames(ctx context.Context, userNames []string, opt ...*RetrieveUserOption) (*UsersResponse, error) {
	return retrieveMultipleUsersWithUserNames(ctx, c, userNames, opt...)
}

func (c *client) RetrieveSingleUserWithUserName(ctx context.Context, userName string, opt ...*RetrieveUserOption) (*UserResponse, error) {
	return retrieveSingleUserWithUserName(ctx, c, userName, opt...)
}

func (c *client) Followers(ctx context.Context, userID string, opt ...*FollowOption) (*FollowersResponse, error) {
	return followers(ctx, c, userID, opt...)
}

func (c *client) Following(ctx context.Context, userID string, opt ...*FollowOption) (*FollowingResponse, error) {
	return following(ctx, c, userID, opt...)
}

func (c *client) LookUpSpace(ctx context.Context, spaceID string, opt ...*SpaceOption) (*SpaceResponse, error) {
	return lookUpSpace(ctx, c, spaceID, opt...)
}

func (c *client) LookUpSpaces(ctx context.Context, spaceIDs []string, opt ...*SpaceOption) (*SpacesResponse, error) {
	return lookUpSpaces(ctx, c, spaceIDs, opt...)
}

func (c *client) UsersPurchasedSpaceTicket(ctx context.Context, spaceID string, opt ...*UsersPurchasedSpaceTicketOption) (*UsersPurchasedSpaceTicketResponse, error) {
	return usersPurchasedSpaceTicket(ctx, c, spaceID, opt...)
}
func (c *client) DiscoverSpaces(ctx context.Context, userIDs []string, opt ...*DiscoverSpacesOption) (*DiscoverSpacesResponse, error) {
	return discoverSpaces(ctx, c, userIDs, opt...)
}

func (c *client) SearchSpaces(ctx context.Context, searchTerm string, opt ...*SearchSpacesOption) (*SearchSpacesResponse, error) {
	return searchSpaces(ctx, c, searchTerm, opt...)
}

func (c *client) LookUpList(ctx context.Context, listID string, opt ...*LookUpListOption) (*ListResponse, error) {
	return lookUpList(ctx, c, listID, opt...)
}

func (c *client) LookUpAllListsOwned(ctx context.Context, userID string, opt ...*AllListsOwnedOption) (*AllListsOwnedResponse, error) {
	return lookUpAllListsOwned(ctx, c, userID, opt...)
}

func (c *client) LookUpListTweets(ctx context.Context, listID string, opt ...*ListTweetsOption) (*ListTweetsResponse, error) {
	return lookUpListTweets(ctx, c, listID, opt...)
}

func (c *client) ListSpecifiedUser(ctx context.Context, userid string, opt ...*ListSpecifiedUserOption) (*ListSpecifiedUserResponse, error) {
	return listSpecifiedUser(ctx, c, userid, opt...)
}

func (c *client) ListMembers(ctx context.Context, listid string, opt ...*ListMembersOption) (*ListMembersResponse, error) {
	return listMembers(ctx, c, listid, opt...)
}

func (c *client) LookUpListFollowers(ctx context.Context, listID string, opt ...*ListFollowersOption) (*ListFollowersResponse, error) {
	return lookUpListFollowers(ctx, c, listID, opt...)
}

func (c *client) LookUpAllListsUserFollows(ctx context.Context, userID string, opt ...*ListFollowsOption) (*AllListsUserFollowsResponse, error) {
	return lookUpAllListsUserFollows(ctx, c, userID, opt...)
}
