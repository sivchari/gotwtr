package gotwtr

import (
	"net/http"
	"strconv"
	"strings"
)

type RetrieveUserOption struct {
	Expansions  []Expansion
	TweetFields []TweetField
	UserFields  []UserField
}

func (r *RetrieveUserOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(r.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(r.Expansions), ","))
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

type FollowOption struct {
	Expansions      []Expansion
	MaxResults      int
	PaginationToken string
	TweetFields     []TweetField
	UserFields      []UserField
}

func (t FollowOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(t.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(t.Expansions), ","))
	}
	if t.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(t.MaxResults))
	}
	if len(t.PaginationToken) > 0 {
		q.Add("pagination_token", t.PaginationToken)
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

func userFieldsToString(ufs []UserField) []string {
	slice := make([]string, len(ufs))
	for i, uf := range ufs {
		slice[i] = string(uf)
	}
	return slice
}
