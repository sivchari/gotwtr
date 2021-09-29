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
	UserFieldWithHeld        UserField = "withheld"
)

type User struct {
	ID              string             `json:"id"`
	Name            string             `json:"name"`
	UserName        string             `json:"username"`
	CreatedAt       string             `json:"created_at,omitempty"`
	Description     string             `json:"description,omitempty"`
	Entities        []*UserEntity      `json:"entities,omitempty"`
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
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
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
	Copyright    string   `json:"copyright"`
	CountryCodes []string `json:"country_codes"`
}

func userFieldsToString(ufs []UserField) []string {
	slice := make([]string, len(ufs))
	for i, uf := range ufs {
		slice[i] = string(uf)
	}
	return slice
}
