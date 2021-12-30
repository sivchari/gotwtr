package gotwtr

import (
	"context"
	"net/http"
)

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
	CreateNewList(ctx context.Context, listName string, opt ...*CreateNewListOption) (*CreateNewListResponse, error)
	DeleteList(ctx context.Context, listID string) (*DeleteListResponse, error)
	LookUpList(ctx context.Context, listID string, opt ...*LookUpListOption) (*ListResponse, error)
	LookUpAllListsOwned(ctx context.Context, userID string, opt ...*AllListsOwnedOption) (*AllListsOwnedResponse, error)
	LookUpListTweets(ctx context.Context, listID string, opt ...*ListTweetsOption) (*ListTweetsResponse, error)
	ListMembers(ctx context.Context, listID string, opt ...*ListMembersOption) (*ListMembersResponse, error)
	ListsSpecifiedUser(ctx context.Context, userID string, opt ...*ListsSpecifiedUserOption) (*ListsSpecifiedUserResponse, error)
	LookUpListFollowers(ctx context.Context, listID string, opt ...*ListFollowersOption) (*ListFollowersResponse, error)
	LookUpAllListsUserFollows(ctx context.Context, userID string, opt ...*ListFollowsOption) (*AllListsUserFollowsResponse, error)
	UpdateMetaDataForList(ctx context.Context, listID string, opt ...*UpdateMetaDataForListOption) (*UpdateMetaDataForListResponse, error)
}

// Twtr is a main interface for all Twitter API calls.
type Twtr interface {
	Tweets
	Users
	Spaces
	Lists
}

type client struct {
	bearerToken string
	client      *http.Client
}

// Client is an API client for Twitter v2 API.
type Client struct {
	*client
}

var _ Twtr = (*Client)(nil)

type ClientOption func(*client)

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *client) {
		c.client = httpClient
	}
}

func New(bearerToken string, opts ...ClientOption) *Client {
	c := &client{
		bearerToken: bearerToken,
		client:      http.DefaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	return &Client{
		client: c,
	}
}

// RetrieveMultipleTweets returns a variety of information about the Tweet specified by the requested ID or list of IDs.
func (c *Client) RetrieveMultipleTweets(ctx context.Context, tweetIDs []string, opt ...*RetriveTweetOption) (*TweetsResponse, error) {
	return retrieveMultipleTweets(ctx, c.client, tweetIDs, opt...)
}

// RetrieveSingleTweet returns a variety of information about a single Tweet specified by the requested ID.
func (c *Client) RetrieveSingleTweet(ctx context.Context, tweetID string, opt ...*RetriveTweetOption) (*TweetResponse, error) {
	return retrieveSingleTweet(ctx, c.client, tweetID, opt...)
}

// UserMensionTimeline returns Tweets mentioning a single user specified by the requested userID.
// By default, the most recent ten Tweets are returned per request. Using pagination, up to the most recent 800 Tweets can be retrieved.
func (c *Client) UserMentionTimeline(ctx context.Context, userID string, opt ...*UserMentionTimelineOption) (*UserMentionTimelineResponse, error) {
	return userMentionTimeline(ctx, c.client, userID, opt...)
}

// UserTweetTimeline returns Tweets composed by a single user, specified by the requested userID.
// By default, the most recent ten Tweets are returned per request.
// Using pagination, the most recent 3,200 Tweets can be retrieved.
func (c *Client) UserTweetTimeline(ctx context.Context, userID string, opt ...*UserTweetTimelineOption) (*UserTweetTimelineResponse, error) {
	return userTweetTimeline(ctx, c.client, userID, opt...)
}

// SearchRecentTweets returns Tweets from the last seven days that match a search query.
func (c *Client) SearchRecentTweets(ctx context.Context, tweet string, opt ...*SearchTweetsOption) (*SearchTweetsResponse, error) {
	return searchRecentTweets(ctx, c.client, tweet, opt...)
}

// CountsRecentTweet returns count of Tweets from the last seven days that match a query.
func (c *Client) CountsRecentTweet(ctx context.Context, tweet string, opt ...*TweetCountsOption) (*TweetCountsResponse, error) {
	return countsRecentTweet(ctx, c.client, tweet, opt...)
}

// AddOrDeleteRules To create one or more rules, submit an add JSON body with an array of rules and operators.
// Similarly, to delete one or more rules, submit a delete JSON body with an array of list of existing rule IDs.
func (c *Client) AddOrDeleteRules(ctx context.Context, body *AddOrDeleteJSONBody, opt ...*AddOrDeleteRulesOption) (*AddOrDeleteRulesResponse, error) {
	return addOrDeleteRules(ctx, c.client, body, opt...)
}

// RetriveStreamRules return a list of rules currently active on the streaming endpoint, either as a list or individually.
func (c *Client) RetrieveStreamRules(ctx context.Context, opt ...*RetrieveStreamRulesOption) (*RetrieveStreamRulesResponse, error) {
	return retrieveStreamRules(ctx, c.client, opt...)
}

// ConnectToStream streams Tweets in real-time based on a specific set of filter rules.
func (c *Client) ConnectToStream(ctx context.Context, ch chan<- ConnectToStreamResponse, errCh chan<- error, opt ...*ConnectToStreamOption) *ConnectToStream {
	return connectToStream(ctx, c.client, ch, errCh, opt...)
}

// VolumeStreams streams about 1% of all Tweets in real-time.
func (c *Client) VolumeStreams(ctx context.Context, ch chan<- VolumeStreamsResponse, errCh chan<- error, opt ...*VolumeStreamsOption) *VolumeStreams {
	return volumeStreams(ctx, c.client, ch, errCh, opt...)
}

// RetweetsLookup allows you to get information about who has Retweeted a Tweet.
func (c *Client) RetweetsLookup(ctx context.Context, tweetID string, opt ...*RetweetsLookupOption) (*RetweetsResponse, error) {
	return retweetsLookup(ctx, c.client, tweetID, opt...)
}

// TweetsUserLiked allows you to get information about a Tweet’s liking users.
// You will receive the most recent 100 users who liked the specified Tweet.
func (c *Client) TweetsUserLiked(ctx context.Context, userID string, opt ...*TweetsUserLikedOpts) (*TweetsUserLikedResponse, error) {
	return tweetsUserLiked(ctx, c.client, userID, opt...)
}

// UsersLikingTweet allows you to get information about a user’s liked Tweets.
func (c *Client) UsersLikingTweet(ctx context.Context, tweetID string, opt ...*UsersLikingTweetOption) (*UsersLikingTweetResponse, error) {
	return usersLikingTweet(ctx, c.client, tweetID, opt...)
}

// RetrieveMultipleUsersWithIDs returns a variety of information about one or more users specified by the requested userIDs.
func (c *Client) RetrieveMultipleUsersWithIDs(ctx context.Context, userIDs []string, opt ...*RetrieveUserOption) (*UsersResponse, error) {
	return retrieveMultipleUsersWithIDs(ctx, c.client, userIDs, opt...)
}

// RetrieveSingleWithID returns a variety of information about a single user specified by the requested userID.
func (c *Client) RetrieveSingleUserWithID(ctx context.Context, userID string, opt ...*RetrieveUserOption) (*UserResponse, error) {
	return retrieveSingleUserWithID(ctx, c.client, userID, opt...)
}

// RetrieveMultipleUsersWithUserNames returns a variety of information about one or more users specified by their usernames.
func (c *Client) RetrieveMultipleUsersWithUserNames(ctx context.Context, userNames []string, opt ...*RetrieveUserOption) (*UsersResponse, error) {
	return retrieveMultipleUsersWithUserNames(ctx, c.client, userNames, opt...)
}

// RetrieveSingleUserWithUserName returns a variety of information about one or more users specified by their username.
func (c *Client) RetrieveSingleUserWithUserName(ctx context.Context, userName string, opt ...*RetrieveUserOption) (*UserResponse, error) {
	return retrieveSingleUserWithUserName(ctx, c.client, userName, opt...)
}

// Followers returns a list of users who are followers of the specified userID.
func (c *Client) Followers(ctx context.Context, userID string, opt ...*FollowOption) (*FollowersResponse, error) {
	return followers(ctx, c.client, userID, opt...)
}

// Following returns a list of users the specified userID is following.
func (c *Client) Following(ctx context.Context, userID string, opt ...*FollowOption) (*FollowingResponse, error) {
	return following(ctx, c.client, userID, opt...)
}

// LookUpSpace returns a variety of information about a single Space specified by the requested ID.
func (c *Client) LookUpSpace(ctx context.Context, spaceID string, opt ...*SpaceOption) (*SpaceResponse, error) {
	return lookUpSpace(ctx, c.client, spaceID, opt...)
}

// LookUpSpaces returns details about multiple Spaces. Up to 100 comma-separated SpacesIDs can be looked up using this endpoint.
func (c *Client) LookUpSpaces(ctx context.Context, spaceIDs []string, opt ...*SpaceOption) (*SpacesResponse, error) {
	return lookUpSpaces(ctx, c.client, spaceIDs, opt...)
}

// UsersPurchasedSpaceTicket returns a list of user who purchased a ticket to the requested Space.
// You must authenticate the request using the access token of the creator of the requested Space.
func (c *Client) UsersPurchasedSpaceTicket(ctx context.Context, spaceID string, opt ...*UsersPurchasedSpaceTicketOption) (*UsersPurchasedSpaceTicketResponse, error) {
	return usersPurchasedSpaceTicket(ctx, c.client, spaceID, opt...)
}

// DiscoverSpaces returns live or scheduled Spaces created by the specified userIDs.
// Up to 100 comma-separated IDs can be looked up using this endpoint.
func (c *Client) DiscoverSpaces(ctx context.Context, userIDs []string, opt ...*DiscoverSpacesOption) (*DiscoverSpacesResponse, error) {
	return discoverSpaces(ctx, c.client, userIDs, opt...)
}

// SearchSpaces return live or scheduled Spaces matching your specified search terms.
// This endpoint performs a keyword search, meaning that it will return Spaces that are an exact case-insensitive match of the specified search term.
// The search term will match the original title of the Space.
func (c *Client) SearchSpaces(ctx context.Context, searchTerm string, opt ...*SearchSpacesOption) (*SearchSpacesResponse, error) {
	return searchSpaces(ctx, c.client, searchTerm, opt...)
}

// Enables the authenticated user to create a List.
func (c *Client) CreateNewList(ctx context.Context, listName string, opt ...*CreateNewListOption) (*CreateNewListResponse, error) {
	return createNewList(ctx, c.client, listName, opt...)
}

// DeleteList enables the authenticated user to delete a List that they own.
func (c *Client) DeleteList(ctx context.Context, listID string) (*DeleteListResponse, error) {
	return deleteList(ctx, c.client, listID)
}

// LookUpList returns the details of a specified List.
func (c *Client) LookUpList(ctx context.Context, listID string, opt ...*LookUpListOption) (*ListResponse, error) {
	return lookUpList(ctx, c.client, listID, opt...)
}

// LookUpAllListsOwned returns all Lists owned by the specified user.
func (c *Client) LookUpAllListsOwned(ctx context.Context, userID string, opt ...*AllListsOwnedOption) (*AllListsOwnedResponse, error) {
	return lookUpAllListsOwned(ctx, c.client, userID, opt...)
}

// LookUpListTweets returns a list of Tweets from the specified List.
func (c *Client) LookUpListTweets(ctx context.Context, listID string, opt ...*ListTweetsOption) (*ListTweetsResponse, error) {
	return lookUpListTweets(ctx, c.client, listID, opt...)
}

// ListsSpecifiedUser returns all Lists a specified user is a member of that.
func (c *Client) ListsSpecifiedUser(ctx context.Context, userID string, opt ...*ListsSpecifiedUserOption) (*ListsSpecifiedUserResponse, error) {
	return listsSpecifiedUser(ctx, c.client, userID, opt...)
}

// ListMembers returns a list of users who are members of the specified List.
func (c *Client) ListMembers(ctx context.Context, listid string, opt ...*ListMembersOption) (*ListMembersResponse, error) {
	return listMembers(ctx, c.client, listid, opt...)
}

// LookUpListFollowers returns a list of users who are followers of the specified List.
func (c *Client) LookUpListFollowers(ctx context.Context, listID string, opt ...*ListFollowersOption) (*ListFollowersResponse, error) {
	return lookUpListFollowers(ctx, c.client, listID, opt...)
}

// LookUpAllListsUserFollows returns all Lists a specified user follows.
func (c *Client) LookUpAllListsUserFollows(ctx context.Context, userID string, opt ...*ListFollowsOption) (*AllListsUserFollowsResponse, error) {
	return lookUpAllListsUserFollows(ctx, c.client, userID, opt...)
}

// Enables the authenticated user to update the meta data of a specified List that they own.
func (c *Client) UpdateMetaDataForList(ctx context.Context, listID string, opt ...*UpdateMetaDataForListOption) (*UpdateMetaDataForListResponse, error) {
	return updateMetaDataForList(ctx, c.client, listID, opt...)
}
