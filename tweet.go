package gotwtr

import (
	"net/http"
	"sync"
)

type TweetField string

/*
	Tweet field will only return
	if you've also included the expansions=referenced_tweets.id query parameter in your request.
*/
const (
	TweetFieldAttachments        TweetField = "attachments"
	TweetFieldAuthorID           TweetField = "author_id"
	TweetFieldCreatedAt          TweetField = "created_at"
	TweetFieldConversationID     TweetField = "conversation_id"
	TweetFieldContextAnnotations TweetField = "context_annotations"
	TweetFieldEntities           TweetField = "entities"
	TweetFieldGeo                TweetField = "geo"
	TweetFieldID                 TweetField = "id"
	TweetFieldInReplyToUserID    TweetField = "in_reply_to_user_id"
	TweetFieldLanguage           TweetField = "lang"
	TweetFieldMaxResults         TweetField = "max_results"
	TweetFieldNonPublicMetrics   TweetField = "non_public_metrics"
	TweetFieldPublicMetrics      TweetField = "public_metrics"
	TweetFieldOrganicMetrics     TweetField = "organic_metrics"
	TweetFieldPromotedMetrics    TweetField = "promoted_metrics"
	TweetFieldPossiblySensitve   TweetField = "possibly_sensitive"
	TweetFieldReferencedTweets   TweetField = "referenced_tweets"
	TweetFieldSource             TweetField = "source"
	TweetFieldText               TweetField = "text"
	TweetFieldWithHeld           TweetField = "withheld"
)

type Tweet struct {
	ID                 string                    `json:"id"`
	Text               string                    `json:"text"`
	Attachments        *TweetAttachment          `json:"attachments,omitempty"`
	AuthorID           string                    `json:"author_id,omitempty"`
	ContextAnnotations []*TweetContextAnnotation `json:"context_annotations,omitempty"`
	ConversationID     string                    `json:"conversation_id,omitempty"`
	CreatedAt          string                    `json:"created_at"`
	Entities           *TweetEntity              `json:"entities,omitempty"`
	Geo                *TweetGeo                 `json:"geo,omitempty"`
	InReplyToUserID    string                    `json:"in_reply_to_user_id,omitempty"`
	Lang               string                    `json:"lang,omitempty"`
	NonPublicMetrics   *TweetMetrics             `json:"non_public_metrics,omitempty"`
	OrganicMetrics     *TweetMetrics             `json:"organic_metrics,omitempty"`
	PossiblySensitive  bool                      `json:"possibly_sensitive,omitempty"`
	PromotedMetrics    *TweetMetrics             `json:"promoted_metrics,omitempty"`
	PublicMetrics      *TweetMetrics             `json:"public_metrics,omitempty"`
	ReferencedTweets   []*TweetReferencedTweet   `json:"referenced_tweets,omitempty"`
	ReplySettings      string                    `json:"reply_settings,omitempty"`
	Source             string                    `json:"source,omitempty"`
	Withheld           *TweetWithheld            `json:"withheld,omitempty"`
}

type TweetAttachment struct {
	PollIDs   []string `json:"poll_ids"`
	MediaKeys []string `json:"media_keys"`
}

type TweetContextAnnotation struct {
	Domain *TweetContextObj `json:"domain"`
	Entity *TweetContextObj `json:"entity"`
}

type TweetContextObj struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TweetEntity struct {
	Annotations []*TweetAnnotation `json:"annotations"`
	Cashtags    []*TweetCashtag    `json:"cashtags"`
	Hashtags    []*TweetHashtag    `json:"hashtags"`
	Mentions    []*TweetMention    `json:"mentions"`
	URLs        []*TweetURL        `json:"urls"`
}

type TweetAnnotation struct {
	Start          int     `json:"start"`
	End            int     `json:"end"`
	Probability    float64 `json:"probability"`
	Type           string  `json:"type"`
	NormalizedText string  `json:"normalized_text"`
}

type TweetCashtag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type TweetHashtag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type TweetMention struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	UserName string `json:"user_name"`
}

type TweetURL struct {
	Start       int           `json:"start"`
	End         int           `json:"end"`
	URL         string        `json:"url"`
	ExpandedURL string        `json:"expanded_url"`
	DisplayURL  string        `json:"display_url"`
	Images      []*TweetImage `json:"images"`
	Status      int           `json:"status"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	UnwoundURL  string        `json:"unwound_url"`
}

type TweetImage struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type TweetGeo struct {
	Coordinates *TweetCoordinates `json:"coordinates"`
	PlaceID     string            `json:"place_id"`
}

type TweetCoordinates struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type TweetMetrics struct {
	ImpressionCount   int `json:"impression_count"`
	LikeCount         int `json:"like_count"`
	ReplyCount        int `json:"reply_count"`
	RetweetCount      int `json:"retweet_count"`
	URLLinkClicks     int `json:"url_link_clicks"`
	UserProfileClicks int `json:"user_profile_clicks"`
	QuoteCount        int `json:"quote_count"`
}

type TweetReferencedTweet struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type TweetWithheld struct {
	Copyright    bool     `json:"copyright"`
	CountryCodes []string `json:"country_codes"`
}

type TweetLookUpResponse struct {
	Tweets   []*Tweet            `json:"data"`
	Includes *TweetIncludes      `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type TweetLookUpByIDResponse struct {
	Tweet    *Tweet              `json:"data"`
	Includes *TweetIncludes      `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
}

type TweetIncludes struct {
	Media  []*Media
	Places []*Place
	Polls  []*Poll
	Tweets []*Tweet
	Users  []*User
}

type UserTweetTimelineResponse struct {
	Tweets   []*Tweet            `json:"data"`
	Includes *TweetIncludes      `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *UserTimelineMeta   `json:"meta"`
}

type UserMentionTimelineResponse struct {
	Tweets   []*Tweet            `json:"data"`
	Includes *TweetIncludes      `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *UserTimelineMeta   `json:"meta"`
}

type UserTimelineMeta struct {
	ResultCount int    `json:"result_count"`
	NewestID    string `json:"newest_id"`
	OldestID    string `json:"oldest_id"`
	NextToken   string `json:"next_token"`
}

type TweetSearchResponse struct {
	Tweets   []*Tweet            `json:"data"`
	Includes *TweetIncludes      `json:"includes,omitempty"`
	Meta     *TweetSearchMeta    `json:"meta"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
}

type TweetSearchMeta struct {
	ResultCount int    `json:"result_count"`
	NewestID    string `json:"newest_id"`
	OldestID    string `json:"oldest_id"`
	NextToken   string `json:"next_token,omitempty"`
}

type TimeseriesCount struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	TweetCount int    `json:"tweet_count"`
}

type TweetCountMeta struct {
	TotalTweetCount int `json:"total_tweet_count"`
}

type TweetCountsResponse struct {
	Counts []*TimeseriesCount `json:"data"`
	Meta   *TweetCountMeta    `json:"meta"`
}

type RetweetsLookupResponse struct {
	Users    []*User             `json:"data"`
	Includes *TweetIncludes      `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *RetweetsLookupMeta `json:"meta"`
}

type RetweetsLookupMeta struct {
	ResultCount int `json:"result_count"`
}

type RetrieveStreamRulesResponse struct {
	Rules  []*FilteredRule          `json:"data"`
	Meta   *RetrieveStreamRulesMeta `json:"meta"`
	Errors []*APIResponseError      `json:"errors,omitempty"`
}

type FilteredRule struct {
	ID    string `json:"id"`
	Value string `json:"value"`
	Tag   string `json:"tag,omitempty"`
}

type RetrieveStreamRulesMeta struct {
	Sent string // TODO: Is it number ?
}

type AddOrDeleteJSONBody struct {
	Add    []*AddRule  `json:"add,omitempty"`
	Delete *DeleteRule `json:"delete,omitempty"`
}

type AddRule struct {
	Value string `json:"value"`
	Tag   string `json:"tag,omitempty"`
}

type DeleteRule struct {
	IDs []string `json:"ids"`
}

type AddOrDeleteRulesResponse struct {
	Rules  []*FilteredRule       `json:"data"`
	Meta   *AddOrDeleteRulesMeta `json:"meta"`
	Errors []*APIResponseError   `json:"errors,omitempty"`
}

type AddOrDeleteRulesMeta struct {
	Sent    string                  `json:"sent"` // TODO: Is it number ?
	Summary *AddOrDeleteMetaSummary `json:"summary"`
}

type AddOrDeleteMetaSummary struct {
	Created    int `json:"created"`
	NotCreated int `json:"not_created"`
	Deleted    int `json:"deleted"`
	NotDeleted int `json:"not_deleted"`
	Valid      int `json:"valid"`
	Invalid    int `json:"invalid"`
}

type ConnectToStreamResponse struct {
	Tweet         *Tweet          `json:"data"`
	Includes      *TweetIncludes  `json:"includes,omitempty"`
	MatchingRules []*MatchingRule `json:"matching_rules"`
}

type MatchingRule struct {
	ID  string `json:"id"`
	Tag string `json:"tag"`
}

type ConnectToStream struct {
	client *http.Client
	errCh  chan<- error
	ch     chan<- ConnectToStreamResponse
	done   chan struct{}
	wg     *sync.WaitGroup
}

type PostRetweetResponse struct {
	Retweeted *Retweeted          `json:"data"`
	Errors    []*APIResponseError `json:"errors,omitempty"`
}

type UndoRetweetResponse struct {
	Retweeted *Retweeted          `json:"data"`
	Errors    []*APIResponseError `json:"errors,omitempty"`
}

type Retweeted struct {
	Retweeted bool `json:"retweeted"`
}

type TweetBody struct {
	TweetID string `json:"tweet_id"`
}

type SampledStreamResponse struct {
	Tweet    *Tweet              `json:"data"`
	Includes *TweetIncludes      `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
}

type StreamResponse struct {
	client *http.Client
	errCh  chan<- error
	ch     chan<- SampledStreamResponse
	done   chan struct{}
	wg     *sync.WaitGroup
}
type LikesLookUpPublicMetrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
}

type LikeLookUpUser struct {
	ID              string                    `json:"id"`
	Name            string                    `json:"name"`
	UserName        string                    `json:"username"`
	CreatedAt       string                    `json:"created_at,omitempty"`
	Protected       bool                      `json:"protected,omitempty"`
	Withheld        *TweetWithheld            `json:"withheld,omitempty"`
	Location        string                    `json:"location,omitempty"`
	URL             string                    `json:"url,omitempty"`
	Description     string                    `json:"description,omitempty"`
	Verified        bool                      `json:"verified,omitempty"`
	Entities        *TweetEntity              `json:"entities,omitempty"`
	ProfileImageUrl string                    `json:"profile_image_url,omitempty"`
	PublicMetrics   *LikesLookUpPublicMetrics `json:"public_metrics,omitempty"`
	PinnedTweetID   string                    `json:"pinned_tweet_id,omitempty"`
}

type LikesLookUpByTweetResponse struct {
	Users    []*LikeLookUpUser   `json:"data"`
	Includes *TweetIncludes      `json:"includes,omitempty"`
	Meta     *TweetSearchMeta    `json:"meta"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
}
