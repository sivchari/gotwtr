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

func (f *FollowOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(f.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(f.Expansions), ","))
	}
	if f.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(f.MaxResults))
	}
	if len(f.PaginationToken) > 0 {
		q.Add("pagination_token", f.PaginationToken)
	}
	if len(f.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(f.TweetFields), ","))
	}
	if len(f.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(f.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type BlockOption struct {
	Expansions      []Expansion
	MaxResults      int
	PaginationToken string
	TweetFields     []TweetField
	UserFields      []UserField
}

func (b *BlockOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(b.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(b.Expansions), ","))
	}
	if b.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(b.MaxResults))
	}
	if len(b.PaginationToken) > 0 {
		q.Add("pagination_token", b.PaginationToken)
	}
	if len(b.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(b.TweetFields), ","))
	}
	if len(b.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(b.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type MuteOption struct {
	Expansions      []Expansion
	MaxResults      int
	PaginationToken string
	TweetFields     []TweetField
	UserFields      []UserField
}

func (m *MuteOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(m.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(m.Expansions), ","))
	}
	if m.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(m.MaxResults))
	}
	if len(m.PaginationToken) > 0 {
		q.Add("pagination_token", m.PaginationToken)
	}
	if len(m.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(m.TweetFields), ","))
	}
	if len(m.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(m.UserFields), ","))
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

type MeOption struct {
	Expansions  []Expansion
	TweetFields []TweetField
	UserFields  []UserField
}

func (m *MeOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(m.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(m.Expansions), ","))
	}
	if len(m.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(m.TweetFields), ","))
	}
	if len(m.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(m.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type SearchUsersOption struct {
	Expansions      []Expansion
	MaxResults      int
	PaginationToken string
	TweetFields     []TweetField
	UserFields      []UserField
}

func (s *SearchUsersOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(s.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(s.Expansions), ","))
	}
	if s.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(s.MaxResults))
	}
	if len(s.PaginationToken) > 0 {
		q.Add("pagination_token", s.PaginationToken)
	}
	if len(s.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(s.TweetFields), ","))
	}
	if len(s.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(s.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}
