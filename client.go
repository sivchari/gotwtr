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

type OAuth interface {
	GenerateAppOnlyBearerToken(ctx context.Context) (bool, error)
	// InvalidatingBearerToken(ctx context.Context) (bool, error)
	// RefreshToken() (string, error)
	// RevokeToken() (bool, error)
}

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
	PostRetweet(ctx context.Context, userID string, tweetID string) (*PostRetweetResponse, error)
	UndoRetweet(ctx context.Context, userID string, sourceTweetID string) (*UndoRetweetResponse, error)
	TweetsUserLiked(ctx context.Context, userID string, opt ...*TweetsUserLikedOption) (*TweetsUserLikedResponse, error)
	UsersLikingTweet(ctx context.Context, tweetID string, opt ...*UsersLikingTweetOption) (*UsersLikingTweetResponse, error)
	PostUsersLikingTweet(ctx context.Context, userID string, tweetID string) (*PostUsersLikingTweetResponse, error)
	UndoUsersLikingTweet(ctx context.Context, userID string, tweetID string) (*UndoUsersLikingTweetResponse, error)
	SearchAllTweets(ctx context.Context, tweet string, opt ...*SearchTweetsOption) (*SearchTweetsResponse, error)
	PostTweet(ctx context.Context, body *PostTweetOption) (*PostTweetResponse, error)
	DeleteTweet(ctx context.Context, tweetID string) (*DeleteTweetResponse, error)
}

type Users interface {
	RetrieveMultipleUsersWithIDs(ctx context.Context, userIDs []string, opt ...*RetrieveUserOption) (*UsersResponse, error)
	RetrieveSingleUserWithID(ctx context.Context, userID string, opt ...*RetrieveUserOption) (*UserResponse, error)
	RetrieveMultipleUsersWithUserNames(ctx context.Context, userNames []string, opt ...*RetrieveUserOption) (*UsersResponse, error)
	RetrieveSingleUserWithUserName(ctx context.Context, userName string, opt ...*RetrieveUserOption) (*UserResponse, error)
	Followers(ctx context.Context, userID string, opt ...*FollowOption) (*FollowersResponse, error)
	Following(ctx context.Context, userID string, opt ...*FollowOption) (*FollowingResponse, error)
	PostFollowing(ctx context.Context, userID string, targetUserID string) (*PostFollowingResponse, error)
	UndoFollowing(ctx context.Context, sourceUserID string, targetUserID string) (*UndoFollowingResponse, error)
	Blocking(ctx context.Context, userID string, opt ...*BlockOption) (*BlockingResponse, error)
	PostBlocking(ctx context.Context, userID string, targetUserID string) (*PostBlockingResponse, error)
	UndoBlocking(ctx context.Context, sourceUserID string, targetUserID string) (*UndoBlockingResponse, error)
	Muting(ctx context.Context, userID string, opt ...*MuteOption) (*MutingResponse, error)
	PostMuting(ctx context.Context, userID string, targetUserID string) (*PostMutingResponse, error)
	UndoMuting(ctx context.Context, sourceUserID string, targetUserID string) (*UndoMutingResponse, error)
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
	ListsSpecifiedUser(ctx context.Context, userID string, opt ...*ListsSpecifiedUserOption) (*ListsSpecifiedUserResponse, error)
	PostListMembers(ctx context.Context, listID string, userID string) (*PostListMembersResponse, error)
	UndoListMembers(ctx context.Context, listID string, userID string) (*UndoListMembersResponse, error)
	ListFollowers(ctx context.Context, listID string, opt ...*ListFollowersOption) (*ListFollowersResponse, error)
	AllListsUserFollows(ctx context.Context, userID string, opt ...*ListFollowsOption) (*AllListsUserFollowsResponse, error)
	PostListFollows(ctx context.Context, listID string, userID string) (*PostListFollowsResponse, error)
	UndoListFollows(ctx context.Context, listID string, userID string) (*UndoListFollowsResponse, error)
	PinnedLists(ctx context.Context, userID string, opt ...*PinnedListsOption) (*PinnedListsResponse, error)
	PostPinnedLists(ctx context.Context, listID string, userID string) (*PostPinnedListsResponse, error)
	UndoPinnedLists(ctx context.Context, listID string, userID string) (*UndoPinnedListsResponse, error)
}

type Compliances interface {
	ComplianceJobs(ctx context.Context, opt *ComplianceJobsOption) (*ComplianceJobsResponse, error)
}

// Twtr is a main interface for all Twitter API calls.
type Twtr interface {
	OAuth
	Tweets
	Users
	Spaces
	Lists
	Compliances
}

type client struct {
	consumerKey    string
	consumerSecret string
	bearerToken    string
	client         *http.Client
}

// Client is an API client for Twitter v2 API.
type Client struct {
	*client
}

var _ Twtr = (*Client)(nil)

type ClientOption func(*client)

func WithConsumerKey(consumerKey string) ClientOption {
	return func(c *client) {
		c.consumerKey = consumerKey
	}
}

func WithConsumerSecret(consumerSecret string) ClientOption {
	return func(c *client) {
		c.consumerSecret = consumerSecret
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *client) {
		c.client = httpClient
	}
}

func New(bearerToken string, opts ...ClientOption) *Client {
	c := &client{
		consumerKey:    "",
		consumerSecret: "",
		bearerToken:    bearerToken,
		client:         http.DefaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	return &Client{
		client: c,
	}
}

// GenerateAppOnlyBearerToken generates a bearer token for app-only auth.
func (c *client) GenerateAppOnlyBearerToken(ctx context.Context) (bool, error) {
	return generateAppOnlyBearerToken(ctx, c)
}

// func (c *client) InvalidatingBearerToken(ctx context.Context) (bool, error) {
// 	return invalidatingBearerToken(ctx, c)
// }

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

// SearchAllTweets returns Tweets since the first Tweet was created on March 26, 2006.
// This endpoint is only available to those users who have been approved for Academic Research access.
func (c *client) SearchAllTweets(ctx context.Context, tweet string, opt ...*SearchTweetsOption) (*SearchTweetsResponse, error) {
	return searchAllTweets(ctx, c, tweet, opt...)
}

// PostTweet creates a Tweet on behalf of an authenticated user.
func (c *client) PostTweet(ctx context.Context, body *PostTweetOption) (*PostTweetResponse, error) {
	return postTweet(ctx, c, body)
}

// DeleteTweet allows a user or authenticated user ID to delete a Tweet.
func (c *client) DeleteTweet(ctx context.Context, tweetID string) (*DeleteTweetResponse, error) {
	return deleteTweet(ctx, c, tweetID)
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

// PostRetweet causes the user ID identified in the path parameter to Retweet the target Tweet.
func (c *client) PostRetweet(ctx context.Context, userID string, tweetID string) (*PostRetweetResponse, error) {
	return postRetweet(ctx, c, userID, tweetID)
}

// UndoRetweet allows a user or authenticated user ID to remove the Retweet of a Tweet.
func (c *client) UndoRetweet(ctx context.Context, userID string, sourceTweetID string) (*UndoRetweetResponse, error) {
	return undoRetweet(ctx, c, userID, sourceTweetID)
}

// TweetsUserLiked allows you to get information about a Tweet’s liking users.
// You will receive the most recent 100 users who liked the specified Tweet.
func (c *Client) TweetsUserLiked(ctx context.Context, userID string, opt ...*TweetsUserLikedOption) (*TweetsUserLikedResponse, error) {
	return tweetsUserLiked(ctx, c.client, userID, opt...)
}

// UsersLikingTweet allows you to get information about a user’s liked Tweets.
func (c *Client) UsersLikingTweet(ctx context.Context, tweetID string, opt ...*UsersLikingTweetOption) (*UsersLikingTweetResponse, error) {
	return usersLikingTweet(ctx, c.client, tweetID, opt...)
}

// PostUsersLikingTweet causes the user ID identified in the path parameter to Like the target Tweet.
func (c *client) PostUsersLikingTweet(ctx context.Context, userID string, tweetID string) (*PostUsersLikingTweetResponse, error) {
	return postUsersLikingTweet(ctx, c, userID, tweetID)
}

// UndoUsersLikingTweet allows a user or authenticated user ID to unlike a Tweet.
// The request succeeds with no action when the user sends a request to a user they're not liking the Tweet or have already unliked the Tweet.
func (c *client) UndoUsersLikingTweet(ctx context.Context, userID string, tweetID string) (*UndoUsersLikingTweetResponse, error) {
	return undoUsersLikingTweet(ctx, c, userID, tweetID)
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

// PostFollowing allows a user ID to follow another user.
// If the target user does not have public Tweets, this method will send a follow request.
func (c *Client) PostFollowing(ctx context.Context, userID string, targetUserID string) (*PostFollowingResponse, error) {
	return postFollowing(ctx, c.client, userID, targetUserID)
}

// UndoFollowing allows a user ID to unfollow another user.
func (c *Client) UndoFollowing(ctx context.Context, sourceUserID string, targetUserID string) (*UndoFollowingResponse, error) {
	return undoFollowing(ctx, c.client, sourceUserID, targetUserID)
}

// Blocking returns a list of users who are blocked by the specified user ID.
func (c *Client) Blocking(ctx context.Context, userID string, opt ...*BlockOption) (*BlockingResponse, error) {
	return blocking(ctx, c.client, userID, opt...)
}

// PostBlocking causes the user (in the path) to block the target user.
// The user (in the path) must match the user Access Tokens being used to authorize the request.
func (c *Client) PostBlocking(ctx context.Context, userID string, targetUserID string) (*PostBlockingResponse, error) {
	return postBlocking(ctx, c.client, userID, targetUserID)
}

// UndoBlocking allows a user or authenticated user ID to unblock another user.
func (c *Client) UndoBlocking(ctx context.Context, sourceUserID string, targetUserID string) (*UndoBlockingResponse, error) {
	return undoBlocking(ctx, c.client, sourceUserID, targetUserID)
}

// Muting returns a list of users who are muted by the specified user ID.
func (c *Client) Muting(ctx context.Context, userID string, opt ...*MuteOption) (*MutingResponse, error) {
	return muting(ctx, c.client, userID, opt...)
}

// PostMuting allows an authenticated user ID to mute the target user.
func (c *Client) PostMuting(ctx context.Context, userID string, targetUserID string) (*PostMutingResponse, error) {
	return postMuting(ctx, c.client, userID, targetUserID)
}

// UndoMuting allows an authenticated user ID to unmute the target user.
func (c *Client) UndoMuting(ctx context.Context, sourceUserID string, targetUserID string) (*UndoMutingResponse, error) {
	return undoMuting(ctx, c.client, sourceUserID, targetUserID)
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

// ListMembers returns a list of users who are members of the specified List.
func (c *Client) ListMembers(ctx context.Context, listID string, opt ...*ListMembersOption) (*ListMembersResponse, error) {
	return listMembers(ctx, c.client, listID, opt...)
}

// ListsSpecifiedUser returns all Lists a specified user is a member of that.
func (c *Client) ListsSpecifiedUser(ctx context.Context, userID string, opt ...*ListsSpecifiedUserOption) (*ListsSpecifiedUserResponse, error) {
	return listsSpecifiedUser(ctx, c.client, userID, opt...)
}

// PostListMembers enables the authenticated user to add a member to a List they own.
func (c *Client) PostListMembers(ctx context.Context, listID string, userID string) (*PostListMembersResponse, error) {
	return postListMembers(ctx, c.client, listID, userID)
}

// UndoListMembers enables the authenticated user to remove a member from a List they own.
func (c *Client) UndoListMembers(ctx context.Context, listID string, userID string) (*UndoListMembersResponse, error) {
	return undoListMembers(ctx, c.client, listID, userID)
}

// ListFollowers returns a list of users who are followers of the specified List.
func (c *Client) ListFollowers(ctx context.Context, listID string, opt ...*ListFollowersOption) (*ListFollowersResponse, error) {
	return listFollowers(ctx, c.client, listID, opt...)
}

// AllListsUserFollows returns all Lists a specified user follows.
func (c *Client) AllListsUserFollows(ctx context.Context, userID string, opt ...*ListFollowsOption) (*AllListsUserFollowsResponse, error) {
	return allListsUserFollows(ctx, c.client, userID, opt...)
}

// PostListFollows enables the authenticated user to follow a List.
func (c *Client) PostListFollows(ctx context.Context, listID string, userID string) (*PostListFollowsResponse, error) {
	return postListFollows(ctx, c.client, listID, userID)
}

// UndoListFollows enables the authenticated user to unfollow a List.
func (c *Client) UndoListFollows(ctx context.Context, listID string, userID string) (*UndoListFollowsResponse, error) {
	return undoListFollows(ctx, c.client, listID, userID)
}

// PinnedLists returns the Lists pinned by a specified user.
func (c *Client) PinnedLists(ctx context.Context, userID string, opt ...*PinnedListsOption) (*PinnedListsResponse, error) {
	return pinnedLists(ctx, c.client, userID, opt...)
}

// PostPinnedLists enables the authenticated user to pin a List.
func (c *Client) PostPinnedLists(ctx context.Context, listID string, userID string) (*PostPinnedListsResponse, error) {
	return postPinnedLists(ctx, c.client, listID, userID)
}

// UndoPinnedLists enables the authenticated user to unpin a List.
func (c *Client) UndoPinnedLists(ctx context.Context, listID string, userID string) (*UndoPinnedListsResponse, error) {
	return undoPinnedLists(ctx, c.client, listID, userID)
}

func (c *Client) ComplianceJobs(ctx context.Context, opt *ComplianceJobsOption) (*ComplianceJobsResponse, error) {
	return complianceJobs(ctx, c.client, opt)
}
