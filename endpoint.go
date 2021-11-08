package gotwtr

// EndpointURL is the base URL for the Twitter V2 API.
const (
	generateAppOnlyBearerTokenURL = "https://api.twitter.com/oauth2/token?grant_type=client_credentials"
	// invalidatingBearerTokenURL    = "https://api.twitter.com/oauth2/invalidate_token?access_token=%v"
)

const (
	retrieveMultipleTweetsURL = "https://api.twitter.com/2/tweets?ids="
	retrieveSingleTweetURL    = "https://api.twitter.com/2/tweets/%v"
	userTweetTimelineURL      = "https://api.twitter.com/2/users/%v/tweets"
	userMentionTimelineURL    = "https://api.twitter.com/2/users/%v/mentions"
	searchRecentTweetsURL     = "https://api.twitter.com/2/tweets/search/recent?query=%v"
	countsRecentTweetsURL     = "https://api.twitter.com/2/tweets/counts/recent?query=%v"
	addOrDeleteRulesURL       = "https://api.twitter.com/2/tweets/search/stream/rules"
	retrieveStreamRulesURL    = "https://api.twitter.com/2/tweets/search/stream/rules"
	connectToStreamURL        = "https://api.twitter.com/2/tweets/search/stream"
	volumeStreamsURL          = "https://api.twitter.com/2/tweets/sample/stream"
	retweetsLookupURL         = "https://api.twitter.com/2/tweets/%v/retweeted_by"
	usersLikingTweetURL       = "https://api.twitter.com/2/tweets/%v/liking_users"
	tweetsUserLikedURL        = "https://api.twitter.com/2/users/%v/liked_tweets"
	searchAllTweetsURL        = "https://api.twitter.com/2/tweets/search/all?query=%v"
)

const (
	retrieveMultipleUsersWithIDsURL       = "https://api.twitter.com/2/users?ids="
	retrieveSingleUserWithIDURL           = "https://api.twitter.com/2/users/%v"
	retrieveMultipleUsersWithUserNamesURL = "https://api.twitter.com/2/users/by?usernames="
	retrieveSingleUserWithUserNameURL     = "https://api.twitter.com/2/users/by/username/%v"
	followingURL                          = "https://api.twitter.com/2/users/%v/following"
	followersURL                          = "https://api.twitter.com/2/users/%v/followers"
)

const (
	spaceURL                     = "https://api.twitter.com/2/spaces/%v"
	spacesURL                    = "https://api.twitter.com/2/spaces?ids="
	usersPurchasedSpaceTicketURL = "https://api.twitter.com/2/spaces/%v/buyers"
	discoverSpacesURL            = "https://api.twitter.com/2/spaces/by/creator_ids?user_ids="
	searchSpacesURL              = "https://api.twitter.com/2/spaces/search?query=%v"
)

const (
	lookUpListURL                = "https://api.twitter.com/2/lists/%v"
	lookUpAllListsOwnedURL       = "https://api.twitter.com/2/users/%v/owned_lists"
	lookUpListTweetsURL          = "https://api.twitter.com/2/lists/%v/tweets"
	listsSpecifiedUserURL        = "https://api.twitter.com/2/users/%v/list_memberships"
	listMembersURL               = "https://api.twitter.com/2/lists/%v/members"
	lookUpListFollowersURL       = "https://api.twitter.com/2/lists/%v/followers"
	lookUpAllListsUserFollowsURL = "https://api.twitter.com/2/users/%v/followed_lists"
)
