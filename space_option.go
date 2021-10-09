package gotwtr

import (
	"net/http"
	"strconv"
	"strings"
)

type SearchSpacesOption struct {
	Expansions  []Expansion
	MaxResults  int
	SpaceFields []SpaceField
	State       []StateOption
	UserFields  []UserField
}

func (s *SearchSpacesOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(s.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(s.Expansions), ","))
	}
	if s.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(s.MaxResults))
	}
	if len(s.SpaceFields) > 0 {
		q.Add("space.fields", strings.Join(spaceFieldsToString(s.SpaceFields), ","))
	}
	if len(s.State) > 0 {
		q.Add("state", strings.Join(stateOptionToString(s.State), ","))
	}
	if len(s.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(s.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type StateOption string

const (
	SpaceFieldAll       StateOption = "all"
	SpaceFieldLive      StateOption = "live"
	SpaceFieldScheduled StateOption = "scheduled"
)

func stateOptionToString(sopt []StateOption) []string {
	slice := make([]string, len(sopt))
	for i, s := range sopt {
		slice[i] = string(s)
	}
	return slice
}
