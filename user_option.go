package gotwtr

import (
	"net/http"
	"strings"
)

type UserLookUpOption struct {
	Expansions  []Expansion
	TweetFields []TweetField
	UserFields  []UserField
}

func (t *UserLookUpOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(t.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(t.Expansions), ","))
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
