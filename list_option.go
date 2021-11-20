package gotwtr

import (
	"net/http"
	"strconv"
	"strings"
)

type ListLookUpOption struct {
	Expansions []Expansion
	ListFields []ListField
	MaxResults int
	PaginationToken string
	UserFields []UserField
}

func (t *ListLookUpOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(t.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(t.Expansions), ","))
	}
	if len(t.ListFields) > 0 {
		q.Add("list.fields", strings.Join(listFieldsToString(t.ListFields), "."))
	}
	if t.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(t.MaxResults))
	}
	if len(t.PaginationToken) > 0 {
		q.Add("pagination_token", t.PaginationToken)
	}
	if len(t.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(t.UserFields), ","))
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
