package gotwtr

import (
	"net/http"
	"strconv"
	"strings"
)

// SearchPostsEligibleForNotesOption represents options for searching posts eligible for notes
type SearchPostsEligibleForNotesOption struct {
	TestMode      bool                   `json:"test_mode"`
	MaxResults    int                    `json:"max_results,omitempty"`
	PaginationToken string               `json:"pagination_token,omitempty"`
	Expansions    []Expansion            `json:"expansions,omitempty"`
	NoteFields    []CommunityNoteField   `json:"note.fields,omitempty"`
	PostFields    []TweetField           `json:"post.fields,omitempty"`
	UserFields    []UserField            `json:"user.fields,omitempty"`
}

func (s *SearchPostsEligibleForNotesOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	
	// test_mode is required
	q.Add("test_mode", "true")
	
	if s.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(s.MaxResults))
	}
	
	if s.PaginationToken != "" {
		q.Add("pagination_token", s.PaginationToken)
	}
	
	if len(s.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(s.Expansions), ","))
	}
	
	if len(s.NoteFields) > 0 {
		q.Add("note.fields", strings.Join(communityNoteFieldsToString(s.NoteFields), ","))
	}
	
	if len(s.PostFields) > 0 {
		q.Add("post.fields", strings.Join(tweetFieldsToString(s.PostFields), ","))
	}
	
	if len(s.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(s.UserFields), ","))
	}
	
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

// SearchNotesWrittenOption represents options for searching notes written by user
type SearchNotesWrittenOption struct {
	TestMode        bool                 `json:"test_mode"`
	MaxResults      int                  `json:"max_results,omitempty"`
	PaginationToken string               `json:"pagination_token,omitempty"`
	Expansions      []Expansion          `json:"expansions,omitempty"`
	NoteFields      []CommunityNoteField `json:"note.fields,omitempty"`
	PostFields      []TweetField         `json:"post.fields,omitempty"`
	UserFields      []UserField          `json:"user.fields,omitempty"`
}

func (s *SearchNotesWrittenOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	
	// test_mode is required
	q.Add("test_mode", "true")
	
	if s.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(s.MaxResults))
	}
	
	if s.PaginationToken != "" {
		q.Add("pagination_token", s.PaginationToken)
	}
	
	if len(s.Expansions) > 0 {
		q.Add("expansions", strings.Join(expansionsToString(s.Expansions), ","))
	}
	
	if len(s.NoteFields) > 0 {
		q.Add("note.fields", strings.Join(communityNoteFieldsToString(s.NoteFields), ","))
	}
	
	if len(s.PostFields) > 0 {
		q.Add("post.fields", strings.Join(tweetFieldsToString(s.PostFields), ","))
	}
	
	if len(s.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldsToString(s.UserFields), ","))
	}
	
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}