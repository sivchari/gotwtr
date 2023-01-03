package gotwtr

import (
	"net/http"
	"strconv"
	"strings"
)

type DirectMessageOption struct {
	DMEventFields   []DMEventField
	EventTypes      EventTypes
	Expansions      []Expansion
	MaxResults      int
	MediaFields     []MediaField
	PaginationToken string
	TweetFields     []TweetField
	UserFields      []UserField
}

func (d *DirectMessageOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(d.DMEventFields) > 0 {
		q.Add("dm_event.fields", strings.Join(dmEventFieldsToString(d.DMEventFields), ","))
	}
	if len(d.EventTypes) > 0 {
		q.Add("event_types", string(d.EventTypes))
	}
	if len(d.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(d.Expansions), ","))
	}
	if d.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(d.MaxResults))
	}
	if len(d.MediaFields) > 0 {
		q.Add("media.fields", strings.Join(mediaFieldsToString(d.MediaFields), ","))
	}
	if d.PaginationToken != "" {
		q.Add("pagination_token", d.PaginationToken)
	}
	if len(d.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(d.TweetFields), ","))
	}
	if len(d.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(d.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

func dmEventFieldsToString(defs []DMEventField) []string {
	slice := make([]string, len(defs))
	for i, def := range defs {
		slice[i] = string(def)
	}
	return slice
}
