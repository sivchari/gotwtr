package gotwtr

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RetriveTweetOption struct {
	Expansions  []Expansion
	MediaFields []MediaField
	PlaceFields []PlaceField
	PollFields  []PollField
	TweetFields []TweetField
	UserFields  []UserField
}

func (r RetriveTweetOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(r.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(r.Expansions), ","))
	}
	if len(r.MediaFields) > 0 {
		q.Add("media.fields", strings.Join(mediaFieldsToString(r.MediaFields), ","))
	}
	if len(r.PlaceFields) > 0 {
		q.Add("place.fields", strings.Join(placeFieldsToString(r.PlaceFields), ","))
	}
	if len(r.PollFields) > 0 {
		q.Add("poll.fields", strings.Join(pollFieldsToString(r.PollFields), ","))
	}
	if len(r.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(r.TweetFields), ","))
	}
	if len(r.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(r.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type UserTweetTimelineOption struct {
	EndTime         time.Time
	Exclude         []Exclude
	Expansions      []Expansion
	MaxResults      int
	MediaFields     []MediaField
	PaginationToken string
	PlaceFields     []PlaceField
	PollFields      []PollField
	SinceID         string
	StartTime       time.Time
	TweetFields     []TweetField
	UntilID         string
	UserFields      []UserField
}

func (u UserTweetTimelineOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if !u.EndTime.IsZero() {
		q.Add("end_time", u.EndTime.Format(time.RFC3339))
	}
	if len(u.Exclude) > 0 {
		q.Add("exclude", strings.Join(excludeToString(u.Exclude), ","))
	}
	if len(u.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(u.Expansions), ","))
	}
	if u.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(u.MaxResults))
	}
	if len(u.MediaFields) > 0 {
		q.Add("media.fields", strings.Join(mediaFieldsToString(u.MediaFields), ","))
	}
	if u.PaginationToken != "" {
		q.Add("pagination_token", u.PaginationToken)
	}
	if len(u.PlaceFields) > 0 {
		q.Add("place.fields", strings.Join(placeFieldsToString(u.PlaceFields), ","))
	}
	if len(u.PollFields) > 0 {
		q.Add("poll.fields", strings.Join(pollFieldsToString(u.PollFields), ","))
	}
	if u.SinceID != "" {
		q.Add("since_id", u.SinceID)
	}
	if !u.StartTime.IsZero() {
		q.Add("start_time", u.StartTime.Format(time.RFC3339))
	}
	if len(u.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(u.TweetFields), ","))
	}
	if u.UntilID != "" {
		q.Add("until_id", u.UntilID)
	}
	if len(u.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(u.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type UserMentionTimelineOption struct {
	EndTime         time.Time
	Expansions      []Expansion
	MaxResults      int
	MediaFields     []MediaField
	PaginationToken string
	PlaceFields     []PlaceField
	PollFields      []PollField
	SinceID         string
	StartTime       time.Time
	TweetFields     []TweetField
	UntilID         string
	UserFields      []UserField
}

func (u UserMentionTimelineOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if !u.EndTime.IsZero() {
		q.Add("end_time", u.EndTime.Format(time.RFC3339))
	}
	if len(u.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(u.Expansions), ","))
	}
	if u.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(u.MaxResults))
	}
	if len(u.MediaFields) > 0 {
		q.Add("media.fields", strings.Join(mediaFieldsToString(u.MediaFields), ","))
	}
	if u.PaginationToken != "" {
		q.Add("pagination_token", u.PaginationToken)
	}
	if len(u.PlaceFields) > 0 {
		q.Add("place.fields", strings.Join(placeFieldsToString(u.PlaceFields), ","))
	}
	if len(u.PollFields) > 0 {
		q.Add("poll.fields", strings.Join(pollFieldsToString(u.PollFields), ","))
	}
	if u.SinceID != "" {
		q.Add("since_id", u.SinceID)
	}
	if !u.StartTime.IsZero() {
		q.Add("start_time", u.StartTime.Format(time.RFC3339))
	}
	if len(u.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(u.TweetFields), ","))
	}
	if u.UntilID != "" {
		q.Add("until_id", u.UntilID)
	}
	if len(u.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(u.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type SearchTweetsOption struct {
	EndTime     time.Time
	Expansions  []Expansion
	MaxResults  int
	MediaFields []MediaField
	NextToken   string
	PlaceFields []PlaceField
	PollFields  []PollField
	SinceID     string
	StartTime   time.Time
	TweetFields []TweetField
	UntilID     string
	UserFields  []UserField
}

func (t SearchTweetsOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if !t.EndTime.IsZero() {
		q.Add("end_time", t.EndTime.Format(time.RFC3339))
	}
	if len(t.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(t.Expansions), ","))
	}
	if t.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(t.MaxResults))
	}
	if len(t.MediaFields) > 0 {
		q.Add("media.fields", strings.Join(mediaFieldsToString(t.MediaFields), ","))
	}
	if t.NextToken != "" {
		q.Add("next_token", t.NextToken)
	}
	if len(t.PlaceFields) > 0 {
		q.Add("place.fields", strings.Join(placeFieldsToString(t.PlaceFields), ","))
	}
	if len(t.PollFields) > 0 {
		q.Add("poll.fields", strings.Join(pollFieldsToString(t.PollFields), ","))
	}
	if t.SinceID != "" {
		q.Add("since_id", t.SinceID)
	}
	if !t.StartTime.IsZero() {
		q.Add("start_time", t.StartTime.Format(time.RFC3339))
	}
	if len(t.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(t.TweetFields), ","))
	}
	if t.UntilID != "" {
		q.Add("until_id", t.UntilID)
	}
	if len(t.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(t.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type TweetCountsOption struct {
	StartTime   time.Time
	EndTime     time.Time
	SinceID     string
	UntilID     string
	Granularity string
}

func (t *TweetCountsOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if !t.StartTime.IsZero() {
		// YYYY-MM-DDTHH:mm:ssZ (ISO 8601/RFC 3339).
		q.Add("start_time", t.StartTime.Format(time.RFC3339))
	}
	if !t.EndTime.IsZero() {
		// YYYY-MM-DDTHH:mm:ssZ (ISO 8601/RFC 3339).
		q.Add("end_time", t.EndTime.Format(time.RFC3339))
	}
	if len(t.SinceID) > 0 {
		q.Add("since_id", t.SinceID)
	}
	if len(t.UntilID) > 0 {
		q.Add("until_id", t.UntilID)
	}
	if len(t.Granularity) > 0 {
		q.Add("granularity", t.Granularity)
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type AddOrDeleteRulesOption struct {
	DryRun bool // If it is true, test a the syntax of your rule without submitting it
}

func (t *AddOrDeleteRulesOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if t.DryRun {
		q.Add("dry_run", strconv.FormatBool(t.DryRun))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type RetrieveStreamRulesOption struct {
	IDs []string
}

func (t *RetrieveStreamRulesOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(t.IDs) > 0 {
		q.Add("ids", strings.Join(t.IDs, ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type ConnectToStreamOption struct {
	// BackfillMinutes int // This feature is currently only available to the Academic Research product track.
	Expansions  []Expansion
	MediaFields []MediaField
	PlaceFields []PlaceField
	PollFields  []PollField
	TweetFields []TweetField
	UserFields  []UserField
}

func (t *ConnectToStreamOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(t.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(t.Expansions), ","))
	}
	if len(t.MediaFields) > 0 {
		q.Add("media.fields", strings.Join(mediaFieldsToString(t.MediaFields), ","))
	}
	if len(t.PlaceFields) > 0 {
		q.Add("place.fields", strings.Join(placeFieldsToString(t.PlaceFields), ","))
	}
	if len(t.PollFields) > 0 {
		q.Add("poll.fields", strings.Join(pollFieldsToString(t.PollFields), ","))
	}
	if len(t.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(t.TweetFields), ","))
	}
	if len(t.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(t.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type VolumeStreamsOption struct {
	// BackfillMinutes int // This feature is currently only available to the Academic Research product track.
	Expansions  []Expansion
	MediaFields []MediaField
	PlaceFields []PlaceField
	PollFields  []PollField
	TweetFields []TweetField
	UserFields  []UserField
}

func (v VolumeStreamsOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(v.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(v.Expansions), ","))
	}
	if len(v.MediaFields) > 0 {
		q.Add("media.fields", strings.Join(mediaFieldsToString(v.MediaFields), ","))
	}
	if len(v.PlaceFields) > 0 {
		q.Add("place.fields", strings.Join(placeFieldsToString(v.PlaceFields), ","))
	}
	if len(v.PollFields) > 0 {
		q.Add("poll.fields", strings.Join(pollFieldsToString(v.PollFields), ","))
	}
	if len(v.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(v.TweetFields), ","))
	}
	if len(v.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(v.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type UsersLikingTweetOption struct {
	Expansions  []Expansion
	MediaFields []MediaField
	PlaceFields []PlaceField
	PollFields  []PollField
	TweetFields []TweetField
	UserFields  []UserField
}

func (u *UsersLikingTweetOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(u.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(u.Expansions), ","))
	}
	if len(u.MediaFields) > 0 {
		q.Add("media.fields", strings.Join(mediaFieldsToString(u.MediaFields), ","))
	}
	if len(u.PlaceFields) > 0 {
		q.Add("place.fields", strings.Join(placeFieldsToString(u.PlaceFields), ","))
	}
	if len(u.PollFields) > 0 {
		q.Add("poll.fields", strings.Join(pollFieldsToString(u.PollFields), ","))
	}
	if len(u.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(u.TweetFields), ","))
	}
	if len(u.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(u.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type RetweetsLookupOption struct {
	Expansions  []Expansion
	MediaFields []MediaField
	PlaceFields []PlaceField
	PollFields  []PollField
	TweetFields []TweetField
	UserFields  []UserField
}

func (r RetweetsLookupOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(r.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(r.Expansions), ","))
	}
	if len(r.MediaFields) > 0 {
		q.Add("media.fields", strings.Join(mediaFieldsToString(r.MediaFields), ","))
	}
	if len(r.PlaceFields) > 0 {
		q.Add("place.fields", strings.Join(placeFieldsToString(r.PlaceFields), ","))
	}
	if len(r.PollFields) > 0 {
		q.Add("poll.fields", strings.Join(pollFieldsToString(r.PollFields), ","))
	}
	if len(r.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(r.TweetFields), ","))
	}
	if len(r.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(r.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type TweetsUserLikedOpts struct {
	Expansions      []Expansion
	MediaFields     []MediaField
	PlaceFields     []PlaceField
	PollFields      []PollField
	TweetFields     []TweetField
	UserFields      []UserField
	MaxResults      int
	PaginationToken string
}

func (t *TweetsUserLikedOpts) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(t.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(t.Expansions), ","))
	}
	if len(t.MediaFields) > 0 {
		q.Add("media.fields", strings.Join(mediaFieldsToString(t.MediaFields), ","))
	}
	if len(t.PlaceFields) > 0 {
		q.Add("place.fields", strings.Join(placeFieldsToString(t.PlaceFields), ","))
	}
	if len(t.PollFields) > 0 {
		q.Add("poll.fields", strings.Join(pollFieldsToString(t.PollFields), ","))
	}
	if len(t.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(t.TweetFields), ","))
	}
	if len(t.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(t.UserFields), ","))
	}
	if 10 <= t.MaxResults && t.MaxResults <= 100 {
		q.Add("max_results", strconv.Itoa(t.MaxResults))
	}
	if len(t.PaginationToken) > 0 {
		q.Add("pagination_token", t.PaginationToken)
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

func tweetFieldsToString(tfs []TweetField) []string {
	slice := make([]string, len(tfs))
	for i, tf := range tfs {
		slice[i] = string(tf)
	}
	return slice
}
