package gotwtr

// EndpointURL is the base URL for the Twitter V2 API.
const (
	generateAppOnlyBearerTokenURL = "https://api.twitter.com/oauth2/token?grant_type=client_credentials"
)

const (
	retrieveMultipleTweetsURL = "https://api.twitter.com/2/tweets?ids="
	retrieveSingleTweetURL    = "https://api.twitter.com/2/tweets/%v"
	userTweetTimelineURL      = "https://api.twitter.com/2/users/%v/tweets"
	userMentionTimelineURL    = "https://api.twitter.com/2/users/%v/mentions"
	searchRecentTweetsURL     = "https://api.twitter.com/2/tweets/search/recent"
	countsRecentTweetsURL     = "https://api.twitter.com/2/tweets/counts/recent"
	countsAllTweetsURL        = "https://api.twitter.com/2/tweets/counts/all"
	addOrDeleteRulesURL       = "https://api.twitter.com/2/tweets/search/stream/rules"
	retrieveStreamRulesURL    = "https://api.twitter.com/2/tweets/search/stream/rules"
	connectToStreamURL        = "https://api.twitter.com/2/tweets/search/stream"
	volumeStreamsURL          = "https://api.twitter.com/2/tweets/sample/stream"
	retweetsLookupURL         = "https://api.twitter.com/2/tweets/%v/retweeted_by"
	postRetweetURL            = "https://api.twitter.com/2/users/%v/retweets"
	undoRetweetURL            = "https://api.twitter.com/2/users/%v/retweets/%v"
	usersLikingTweetURL       = "https://api.twitter.com/2/tweets/%v/liking_users"
	tweetsUserLikedURL        = "https://api.twitter.com/2/users/%v/liked_tweets"
	postUsersLikingTweetURL   = "https://api.twitter.com/2/users/%v/likes"
	undoUsersLikingTweetURL   = "https://api.twitter.com/2/users/%v/likes/%v"
	searchAllTweetsURL        = "https://api.twitter.com/2/tweets/search/all"
	postTweetURL              = "https://api.twitter.com/2/tweets"
	deleteTweetURL            = "https://api.twitter.com/2/tweets/%v"
	hideRepliesURL            = "https://api.twitter.com/2/tweets/%v/hidden"
	lookupUserBookmarksURL    = "https://api.twitter.com/2/users/%v/bookmarks"
	bookmarkTweetURL          = "https://api.twitter.com/2/users/%v/bookmarks"
	removeBookmarkOfTweetURL  = "https://api.twitter.com/2/users/%v/bookmarks/%v"
)

const (
	retrieveMultipleUsersWithIDsURL       = "https://api.twitter.com/2/users?ids="
	retrieveSingleUserWithIDURL           = "https://api.twitter.com/2/users/%v"
	retrieveMultipleUsersWithUserNamesURL = "https://api.twitter.com/2/users/by?usernames="
	retrieveSingleUserWithUserNameURL     = "https://api.twitter.com/2/users/by/username/%v"
	followingURL                          = "https://api.twitter.com/2/users/%v/following"
	followersURL                          = "https://api.twitter.com/2/users/%v/followers"
	postFollowingURL                      = "https://api.twitter.com/2/users/%v/following"
	undoFollowingURL                      = "https://api.twitter.com/2/users/%v/following/%v"
	blockingURL                           = "https://api.twitter.com/2/users/%v/blocking"
	postBlockingURL                       = "https://api.twitter.com/2/users/%v/blocking"
	undoBlockingURL                       = "https://api.twitter.com/2/users/%v/blocking/%v"
	mutingURL                             = "https://api.twitter.com/2/users/%v/muting"
	postMutingURL                         = "https://api.twitter.com/2/users/%v/muting"
	undoMutingURL                         = "https://api.twitter.com/2/users/%v/muting/%v"
)

const (
	spaceURL                     = "https://api.twitter.com/2/spaces/%v"
	spacesURL                    = "https://api.twitter.com/2/spaces?ids="
	usersPurchasedSpaceTicketURL = "https://api.twitter.com/2/spaces/%v/buyers"
	discoverSpacesURL            = "https://api.twitter.com/2/spaces/by/creator_ids?user_ids="
	searchSpacesURL              = "https://api.twitter.com/2/spaces/search"
)

const (
	lookUpListURL            = "https://api.twitter.com/2/lists/%v"
	lookUpAllListsOwnedURL   = "https://api.twitter.com/2/users/%v/owned_lists"
	lookUpListTweetsURL      = "https://api.twitter.com/2/lists/%v/tweets"
	listMembersURL           = "https://api.twitter.com/2/lists/%v/members"
	listsSpecifiedUserURL    = "https://api.twitter.com/2/users/%v/list_memberships"
	postListMembersURL       = "https://api.twitter.com/2/lists/%v/members"
	undoListMembersURL       = "https://api.twitter.com/2/lists/%v/members/%v"
	listFollowersURL         = "https://api.twitter.com/2/lists/%v/followers"
	allListsUserFollowsURL   = "https://api.twitter.com/2/users/%v/followed_lists"
	postListFollowsURL       = "https://api.twitter.com/2/users/%v/followed_lists"
	undoListFollowsURL       = "https://api.twitter.com/2/users/%v/followed_lists/%v"
	pinnedListsURL           = "https://api.twitter.com/2/users/%v/pinned_lists"
	postPinnedListsURL       = "https://api.twitter.com/2/users/%v/pinned_lists"
	undoPinnedListsURL       = "https://api.twitter.com/2/users/%v/pinned_lists/%v"
	createNewListURL         = "https://api.twitter.com/2/lists"
	updateMetaDataForListURL = "https://api.twitter.com/2/lists/%v"
	deleteListURL            = "https://api.twitter.com/2/lists/%v"
)

const (
	complianceJobsURL      = "https://api.twitter.com/2/compliance/jobs"
	complianceJobURL       = "https://api.twitter.com/2/compliance/jobs/%v"
	createComplianceJobURL = "https://api.twitter.com/2/compliance/jobs"
)
