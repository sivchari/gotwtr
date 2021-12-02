package gotwtr

import (
	"net/http"
	"strings"
)

type SearchSpacesOption struct {
	Expansions  []Expansion
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

type SpaceOption struct {
	Expansions  []Expansion
	SpaceFields []SpaceField
	TopicFields []TopicField
	UserFields  []UserField
}

func (s *SpaceOption) addQuery(req *http.Request) {
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

func (d *DiscoverSpacesOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(d.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(d.Expansions), ","))
	}
	if len(d.SpaceFields) > 0 {
		q.Add("space.fields", strings.Join(spaceFieldsToString(d.SpaceFields), ","))
	}
	if len(d.TopicFields) > 0 {
		q.Add("topic.fields", strings.Join(topicFieldsToString(d.TopicFields), ","))
	}
	if len(d.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(d.UserFields), ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type UsersPurchasedSpaceTicketOption struct {
	Expansions  []Expansion
	MediaFields []MediaField
	PlaceFields []PlaceField
	PollFields  []PollField
	TweetFields []TweetField
	UserFields  []UserField
}

func (u *UsersPurchasedSpaceTicketOption) addQuery(req *http.Request) {
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
