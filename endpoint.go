package gotwtr

// EndpointURL is the base URL for the Twitter V2 API.
const (
	generateAppOnlyBearerTokenURL = "https://api.twitter.com/oauth2/token?grant_type=client_credentials"
	// TODO: oauth2/invalidate_token
)

const (
	// Bookmarks
	removeBookmarkOfTweetURL = "https://api.twitter.com/2/users/%v/bookmarks/%v"
	lookupUserBookmarksURL   = "https://api.twitter.com/2/users/%v/bookmarks"
	bookmarkTweetURL         = "https://api.twitter.com/2/users/%v/bookmarks"
	// Filtered stream
	connectToStreamURL     = "https://api.twitter.com/2/tweets/search/stream"
	retrieveStreamRulesURL = "https://api.twitter.com/2/tweets/search/stream/rules"
	addOrDeleteRulesURL    = "https://api.twitter.com/2/tweets/search/stream/rules"
	// Hide replies
	hideRepliesURL = "https://api.twitter.com/2/tweets/%v/hidden"
	// Likes
	undoUsersLikingTweetURL = "https://api.twitter.com/2/users/%v/likes/%v"
	usersLikingTweetURL     = "https://api.twitter.com/2/tweets/%v/liking_users"
	tweetsUserLikedURL      = "https://api.twitter.com/2/users/%v/liked_tweets"
	postUsersLikingTweetURL = "https://api.twitter.com/2/users/%v/likes"
	// Manage Tweets
	deleteTweetURL = "https://api.twitter.com/2/tweets/%v"
	postTweetURL   = "https://api.twitter.com/2/tweets"
	// TODO: /2/tweets/:id/quote_tweets
	// Retweets
	undoRetweetURL    = "https://api.twitter.com/2/users/%v/retweets/%v"
	retweetsLookupURL = "https://api.twitter.com/2/tweets/%v/retweeted_by"
	postRetweetURL    = "https://api.twitter.com/2/users/%v/retweets"
	// Search Tweets
	searchAllTweetsURL    = "https://api.twitter.com/2/tweets/search/all"
	searchRecentTweetsURL = "https://api.twitter.com/2/tweets/search/recent"
	// Timelines
	userMentionTimelineURL = "https://api.twitter.com/2/users/%v/mentions"
	// TODO: /2/users/:id/timelines/reverse_chronological
	userTweetTimelineURL = "https://api.twitter.com/2/users/%v/tweets"
	// Tweet counts
	countsAllTweetsURL    = "https://api.twitter.com/2/tweets/counts/all"
	countsRecentTweetsURL = "https://api.twitter.com/2/tweets/counts/recent"
	// Tweets lookup
	retrieveMultipleTweetsURL = "https://api.twitter.com/2/tweets?ids="
	retrieveSingleTweetURL    = "https://api.twitter.com/2/tweets/%v"
	// Volume streams
	volumeStreamsURL = "https://api.twitter.com/2/tweets/sample/stream"
	// TODO: /2/tweets/sample10/stream
)

const (
	// Blocks
	undoBlockingURL = "https://api.twitter.com/2/users/%v/blocking/%v"
	blockingURL     = "https://api.twitter.com/2/users/%v/blocking"
	postBlockingURL = "https://api.twitter.com/2/users/%v/blocking"
	// Follows
	undoFollowingURL = "https://api.twitter.com/2/users/%v/following/%v"
	followersURL     = "https://api.twitter.com/2/users/%v/followers"
	followingURL     = "https://api.twitter.com/2/users/%v/following"
	postFollowingURL = "https://api.twitter.com/2/users/%v/following"
	// Mutes
	undoMutingURL = "https://api.twitter.com/2/users/%v/muting/%v"
	mutingURL     = "https://api.twitter.com/2/users/%v/muting"
	postMutingURL = "https://api.twitter.com/2/users/%v/muting"
	// Users lookup
	retrieveMultipleUsersWithIDsURL       = "https://api.twitter.com/2/users?ids="
	retrieveSingleUserWithIDURL           = "https://api.twitter.com/2/users/%v"
	retrieveMultipleUsersWithUserNamesURL = "https://api.twitter.com/2/users/by?usernames="
	retrieveSingleUserWithUserNameURL     = "https://api.twitter.com/2/users/by/username/%v"
	meURL                                 = "https://api.twitter.com/2/users/me"
)

const (
	// Search Spaces
	searchSpacesURL = "https://api.twitter.com/2/spaces/search"
	// Spaces lookup
	spacesURL                    = "https://api.twitter.com/2/spaces?ids="
	spaceURL                     = "https://api.twitter.com/2/spaces/%v"
	usersPurchasedSpaceTicketURL = "https://api.twitter.com/2/spaces/%v/buyers"
	// TODO: /2/spaces/:id/tweets
	discoverSpacesURL = "https://api.twitter.com/2/spaces/by/creator_ids?user_ids="
)

const (
	// List Tweets lookup
	lookUpListTweetsURL = "https://api.twitter.com/2/lists/%v/tweets"
	// List follows
	undoListFollowsURL     = "https://api.twitter.com/2/users/%v/followed_lists/%v"
	listFollowersURL       = "https://api.twitter.com/2/lists/%v/followers"
	allListsUserFollowsURL = "https://api.twitter.com/2/users/%v/followed_lists"
	postListFollowsURL     = "https://api.twitter.com/2/users/%v/followed_lists"
	// List lookup
	lookUpListURL          = "https://api.twitter.com/2/lists/%v"
	lookUpAllListsOwnedURL = "https://api.twitter.com/2/users/%v/owned_lists"
	// List members
	undoListMembersURL    = "https://api.twitter.com/2/lists/%v/members/%v"
	listMembersURL        = "https://api.twitter.com/2/lists/%v/members"
	listsSpecifiedUserURL = "https://api.twitter.com/2/users/%v/list_memberships"
	postListMembersURL    = "https://api.twitter.com/2/lists/%v/members"
	// Manage Lists
	deleteListURL            = "https://api.twitter.com/2/lists/%v"
	updateMetaDataForListURL = "https://api.twitter.com/2/lists/%v"
	createNewListURL         = "https://api.twitter.com/2/lists"
	// Pinned Lists
	undoPinnedListsURL = "https://api.twitter.com/2/users/%v/pinned_lists/%v"
	pinnedListsURL     = "https://api.twitter.com/2/users/%v/pinned_lists"
	postPinnedListsURL = "https://api.twitter.com/2/users/%v/pinned_lists"
)

const (
	// Batch compliance
	complianceJobsURL      = "https://api.twitter.com/2/compliance/jobs"
	complianceJobURL       = "https://api.twitter.com/2/compliance/jobs/%v"
	createComplianceJobURL = "https://api.twitter.com/2/compliance/jobs"
)
