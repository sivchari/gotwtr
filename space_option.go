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
	TopicFields []TopicField
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
	if len(s.TopicFields) > 0 {
		q.Add("topic.fields", strings.Join(topicFieldsToString(s.TopicFields), ","))
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

type SpaceLookUpOption struct {
	Expansions  []Expansion
	SpaceFields []SpaceField
	TopicFields []TopicField
	UserFields  []UserField
}

func (s *SpaceLookUpOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(s.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(s.Expansions), ","))
	}
	if len(s.SpaceFields) > 0 {
		q.Add("space.fields", strings.Join(spaceFieldsToString(s.SpaceFields), ","))
	}
	if len(s.TopicFields) > 0 {
		q.Add("topic.fields", strings.Join(topicFieldsToString(s.TopicFields), ","))
	}
	if len(s.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(s.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type DiscoverSpacesOption struct {
	Expansions  []Expansion
	SpaceFields []SpaceField
	TopicFields []TopicField
	UserFields  []UserField
}

func (s *DiscoverSpacesOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(s.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(s.Expansions), ","))
	}
	if len(s.SpaceFields) > 0 {
		q.Add("space.fields", strings.Join(spaceFieldsToString(s.SpaceFields), ","))
	}
	if len(s.TopicFields) > 0 {
		q.Add("topic.fields", strings.Join(topicFieldsToString(s.TopicFields), ","))
	}
	if len(s.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(s.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type LookUpUsersWhoPurchasedSpaceTicketOption struct {
	Expansions  []Expansion
	MediaFields []MediaField
	PlaceFields []PlaceField
	PollFields  []PollField
	TweetFields []TweetField
	UserFields  []UserField
}

func (o *LookUpUsersWhoPurchasedSpaceTicketOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(o.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(o.Expansions), ","))
	}
	if len(o.MediaFields) > 0 {
		q.Add("media.fields", strings.Join(mediaFieldsToString(o.MediaFields), ","))
	}
	if len(o.PlaceFields) > 0 {
		q.Add("place.fields", strings.Join(placeFieldsToString(o.PlaceFields), ","))
	}
	if len(o.PollFields) > 0 {
		q.Add("poll.fields", strings.Join(pollFieldsToString(o.PollFields), ","))
	}
	if len(o.TweetFields) > 0 {
		q.Add("tweet.fields", strings.Join(tweetFieldsToString(o.TweetFields), ","))
	}
	if len(o.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(o.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}
