package gotwtr

import (
	"net/http"
	"strconv"
	"strings"
)

type LookUpListOption struct {
	Expansions []Expansion
	ListFields []ListField
	UserFields []UserField
}

func (l LookUpListOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(l.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(l.Expansions), ","))
	}
	if len(l.ListFields) > 0 {
		q.Add("list.fields", strings.Join(listFieldsToString(l.ListFields), "."))
	}
	if len(l.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(l.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type AllListsOwnedOption struct {
	Expansions      []Expansion
	ListFields      []ListField
	MaxResults      int
	PaginationToken string
	UserFields      []UserField
}

func (a AllListsOwnedOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(a.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(a.Expansions), ","))
	}
	if len(a.ListFields) > 0 {
		q.Add("list.fields", strings.Join(listFieldsToString(a.ListFields), "."))
	}
	if a.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(a.MaxResults))
	}
	if a.PaginationToken != "" {
		q.Add("pagination_token", a.PaginationToken)
	}
	if len(a.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(a.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type ListTweetsOption struct {
	Expansions      []Expansion
	MaxResults      int
	PaginationToken string
	TweetFields     []TweetField
	UserFields      []UserField
}

func (l ListTweetsOption) addQuery(req *http.Request) {
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

type ListMembersOption struct {
	Expansions      []Expansion
	MaxResults      int
	PaginationToken string
	TweetFields     []TweetField
	UserFields      []UserField
}

func (l ListMembersOption) addQuery(req *http.Request) {
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

type ListFollowsOption struct {
	Expansions      []Expansion
	MaxResults      int
	PaginationToken string
	TweetFields     []TweetField
	UserFields      []UserField
}

func (l ListFollowsOption) addQuery(req *http.Request) {
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

type ListFollowersOption struct {
	Expansions      []Expansion
	MaxResults      int
	PaginationToken string
	TweetFields     []TweetField
	UserFields      []UserField
}

func (l ListFollowersOption) addQuery(req *http.Request) {
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

func listFieldsToString(lfs []ListField) []string {
	slice := make([]string, len(lfs))
	for i, lf := range lfs {
		slice[i] = string(lf)
	}
	return slice
}
