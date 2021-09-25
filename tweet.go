package twitter

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
	Attachments        []*TweetAttachment        `json:"attachments,omitempty"`
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
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type TweetURL struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	URL         string `json:"url"`
	ExpandedURL string `json:"expanded_url"`
	DisplayURL  string `json:"display_url"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UnwoundURL  string `json:"unwound_url"`
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
	Copyright    string   `json:"copyright"`
	CountryCodes []string `json:"country_codes"`
}
