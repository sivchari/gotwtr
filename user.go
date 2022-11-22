package gotwtr

type UserField string

/*
	You must also pass one of the user expansions to return the desired user fields.

	- expansions=author_id
	- expansions=entities.mentions.username
	- expansions=in_reply_to_user_id
	- expansions=referenced_tweets.id.author_id
*/
const (
	UserFieldCreatedAt       UserField = "created_at"
	UserFieldDescription     UserField = "description"
	UserFieldEntities        UserField = "entities"
	UserFieldID              UserField = "id"
	UserFieldLocation        UserField = "location"
	UserFieldName            UserField = "name"
	UserFieldPinnedTweetID   UserField = "pinned_tweet_id"
	UserFieldProfileImageURL UserField = "profile_image_url"
	UserFieldProtected       UserField = "protected"
	UserFieldPublicMetrics   UserField = "public_metrics"
	UserFieldURL             UserField = "url"
	UserFieldUserName        UserField = "username"
	UserFieldVerified        UserField = "verified"
	UserFieldWithHeld        UserField = "withhel"
)

type User struct {
	ID              string             `json:"id"`
	Name            string             `json:"name"`
	UserName        string             `json:"username"`
	CreatedAt       string             `json:"created_at,omitempty"`
	Description     string             `json:"description,omitempty"`
	Entities        *UserEntity        `json:"entities,omitempty"`
	Location        string             `json:"location,omitempty"`
	PinnedTweetID   string             `json:"pinned_tweet_id,omitempty"`
	ProfileImageURL string             `json:"profile_image_url,omitempty"`
	Protected       bool               `json:"protected,omitempty"`
	PublicMetrics   *UserPublicMetrics `json:"public_metrics,omitempty"`
	URL             string             `json:"url,omitempty"`
	Verified        bool               `json:"verified,omitempty"`
	Withheld        *UserWithheld      `json:"withheld,omitempty"`
}

type UserEntity struct {
	URL         *UserURL         `json:"url"`
	Description *UserDescription `json:"description"`
}

type UserURL struct {
	URLs []*UserURLs `json:"urls"`
}

type UserURLs struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	URL         string `json:"url"`
	ExpandedURL string `json:"expanded_url"`
	DisplayURL  string `json:"display_url"`
}

type UserDescription struct {
	URLs     []*UserURLs    `json:"urls"`
	Hashtags []*UserHashtag `json:"hashtags"`
	Mentions []*UserMention `json:"user_mentions"`
	Cashtags []*UserCashtag `json:"cashtags"`
}

type UserHashtag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type UserMention struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	UserName string `json:"username"`
}

type UserCashtag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type UserPublicMetrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
}

type UserWithheld struct {
	Copyright    bool     `json:"copyright"`
	CountryCodes []string `json:"country_codes"`
}

type UserIncludes struct {
	Users  []*User
	Tweets []*Tweet
}

type UsersResponse struct {
	Users    []*User             `json:"data"`
	Includes *UserIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type UserResponse struct {
	User     *User               `json:"data"`
	Includes *UserIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type FollowingResponse struct {
	Users    []*User             `json:"data"`
	Includes *UserIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *FollowsMeta        `json:"meta"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type FollowersResponse struct {
	Users    []*User             `json:"data"`
	Includes *UserIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *FollowsMeta        `json:"meta"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type PostFollowingResponse struct {
	Following *Following          `json:"data"`
	Errors    []*APIResponseError `json:"errors,omitempty"`
}

type UndoFollowingResponse struct {
	Following *Following          `json:"data"`
	Errors    []*APIResponseError `json:"errors,omitempty"`
}

type Following struct {
	Following     bool `json:"following"`
	PendingFollow bool `json:"pending_follow,omitempty"`
}

type FollowsMeta struct {
	ResultCount   int    `json:"result_count"`
	PreviousToken string `json:"previous_token,omitempty"`
	NextToken     string `json:"next_token,omitempty"`
}

type FollowingBody struct {
	TargetUserID string `json:"target_user_id"`
}

type BlockingResponse struct {
	Users    []*User             `json:"data"`
	Includes *UserIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *BlocksMeta         `json:"meta"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type PostBlockingResponse struct {
	Blocking *Blocking           `json:"data"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
}

type UndoBlockingResponse struct {
	Blocking *Blocking           `json:"data"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
}

type Blocking struct {
	Blocking bool `json:"blocking"`
}

type BlocksMeta struct {
	ResultCount   int    `json:"result_count"`
	PreviousToken string `json:"previous_token,omitempty"`
	NextToken     string `json:"next_token,omitempty"`
}

type BlockingBody struct {
	TargetUserID string `json:"target_user_id"`
}

type MutingResponse struct {
	Users    []*User             `json:"data"`
	Includes *UserIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *MutesMeta          `json:"meta"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type PostMutingResponse struct {
	Muting *Muting             `json:"data"`
	Errors []*APIResponseError `json:"errors,omitempty"`
}

type UndoMutingResponse struct {
	Muting *Muting             `json:"data"`
	Errors []*APIResponseError `json:"errors,omitempty"`
}

type Muting struct {
	Muting bool `json:"muting"`
}

type MutesMeta struct {
	ResultCount   int    `json:"result_count"`
	PreviousToken string `json:"previous_token,omitempty"`
	NextToken     string `json:"next_token,omitempty"`
}

type MutingBody struct {
	TargetUserID string `json:"target_user_id"`
}

type MeResponse struct {
	Me     *Me                 `json:"data"`
	Errors []*APIResponseError `json:"errors,omitempty"`
}

type Me struct {
	ID              string             `json:"id"`
	Name            string             `json:"name"`
	UserName        string             `json:"username"`
	CreatedAt       string             `json:"created_at,omitempty"`
	Protected       bool               `json:"protected,omitempty"`
	Withheld        *MeWithheld        `json:"withheld,omitempty"`
	Location        string             `json:"location,omitempty"`
	URL             string             `json:"url,omitempty"`
	Description     string             `json:"description,omitempty"`
	Verified        bool               `json:"verified,omitempty"`
	Entities        *UserEntity        `json:"entities,omitempty"`
	ProfileImageURL string             `json:"profile_image_url,omitempty"`
	PublicMetrics   *UserPublicMetrics `json:"public_metrics,omitempty"`
	PinnedTweetID   string             `json:"pinned_tweet_id,omitempty"`
	Includes        *MeIncludes        `json:"includes,omitempty"`
}

type MeWithheld struct {
	CountryCodes []string `json:"country_codes"`
	Scope        string   `json:"scope"`
}

type MeIncludes struct {
	Tweets []*Tweet
}
