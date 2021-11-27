package gotwtr

import (
	"net/http"
	"strconv"
	"strings"
)

type ListLookUpOption struct {
	Expansions      []Expansion
	ListFields      []ListField
	MaxResults      int
	PaginationToken string
	TweetFields     []TweetField
	UserFields      []UserField
}

func (l *ListLookUpOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(l.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(l.Expansions), ","))
	}
	if len(l.ListFields) > 0 {
		q.Add("list.fields", strings.Join(listFieldsToString(l.ListFields), "."))
	}
	if l.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(l.MaxResults))
	}
	if l.PaginationToken != "" {
		q.Add("pagination_token", l.PaginationToken)
	}
	if len(l.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(l.TweetFields), "."))
	}
	if len(l.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(l.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

func listFieldsToString(lfs []ListField) []string {
	slice := make([]string, len(lfs))
	for i, lf := range lfs {
		slice[i] = string(lf)
	}
	return slice
}

type ListMembersOption struct {
	Expansions      []Expansion
	MaxResults      int
	PaginationToken string
	TweetFields     []TweetField
	UserFields      []UserField
}

func (l *ListMembersOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(l.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(l.Expansions), ","))
	}
	if l.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(l.MaxResults))
	}
	if l.PaginationToken != "" {
		q.Add("pagination_token", l.PaginationToken)
	}
	if len(l.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(l.TweetFields), "."))
	}
	if len(l.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(l.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}
