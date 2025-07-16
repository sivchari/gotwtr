package gotwtr

// Community Notes field enum for request parameters
type CommunityNoteField string

const (
	CommunityNoteFieldID                  CommunityNoteField = "id"
	CommunityNoteFieldText                CommunityNoteField = "text"
	CommunityNoteFieldCreatedAt           CommunityNoteField = "created_at"
	CommunityNoteFieldAuthorID            CommunityNoteField = "author_id"
	CommunityNoteFieldPostID              CommunityNoteField = "post_id"
	CommunityNoteFieldClassification      CommunityNoteField = "classification"
	CommunityNoteFieldBelievable          CommunityNoteField = "believable"
	CommunityNoteFieldHarmful             CommunityNoteField = "harmful"
	CommunityNoteFieldValidationDifficult CommunityNoteField = "validation_difficult"
	CommunityNoteFieldMisleadingOther     CommunityNoteField = "misleading_other"
	CommunityNoteFieldNotMisleading       CommunityNoteField = "not_misleading"
)

// CommunityNote represents a Community Note structure
type CommunityNote struct {
	ID                     string   `json:"id"`
	Text                   string   `json:"text"`
	CreatedAt              string   `json:"created_at,omitempty"`
	AuthorID               string   `json:"author_id,omitempty"`
	PostID                 string   `json:"post_id,omitempty"`
	Classification         string   `json:"classification,omitempty"`
	Believable             *bool    `json:"believable,omitempty"`
	Harmful                *bool    `json:"harmful,omitempty"`
	ValidationDifficult    *bool    `json:"validation_difficult,omitempty"`
	MisleadingOther        *bool    `json:"misleading_other,omitempty"`
	NotMisleading          *bool    `json:"not_misleading,omitempty"`
	TrustworthySources     []string `json:"trustworthy_sources,omitempty"`
}

// Post represents a post/tweet eligible for Community Notes
type Post struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	AuthorID  string `json:"author_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

// CommunityNotesIncludes represents included data in Community Notes responses
type CommunityNotesIncludes struct {
	Users []*User  `json:"users,omitempty"`
	Posts []*Post  `json:"posts,omitempty"`
	Notes []*CommunityNote `json:"notes,omitempty"`
}

// SearchPostsEligibleForNotesResponse represents the response for searching posts eligible for notes
type SearchPostsEligibleForNotesResponse struct {
	Data     []*Post                 `json:"data,omitempty"`
	Includes *CommunityNotesIncludes `json:"includes,omitempty"`
	Meta     *SearchNotesMeta        `json:"meta,omitempty"`
	Errors   []*APIResponseError     `json:"errors,omitempty"`
	Title    string                  `json:"title,omitempty"`
	Detail   string                  `json:"detail,omitempty"`
	Type     string                  `json:"type,omitempty"`
}

// SearchNotesWrittenResponse represents the response for searching notes written by user
type SearchNotesWrittenResponse struct {
	Data     []*CommunityNote        `json:"data,omitempty"`
	Includes *CommunityNotesIncludes `json:"includes,omitempty"`
	Meta     *SearchNotesMeta        `json:"meta,omitempty"`
	Errors   []*APIResponseError     `json:"errors,omitempty"`
	Title    string                  `json:"title,omitempty"`
	Detail   string                  `json:"detail,omitempty"`
	Type     string                  `json:"type,omitempty"`
}

// CreateCommunityNoteResponse represents the response for creating a community note
type CreateCommunityNoteResponse struct {
	Data   *CommunityNote          `json:"data,omitempty"`
	Errors []*APIResponseError     `json:"errors,omitempty"`
	Title  string                  `json:"title,omitempty"`
	Detail string                  `json:"detail,omitempty"`
	Type   string                  `json:"type,omitempty"`
}

// SearchNotesMeta represents metadata for search responses
type SearchNotesMeta struct {
	ResultCount   int    `json:"result_count,omitempty"`
	NextToken     string `json:"next_token,omitempty"`
	PreviousToken string `json:"previous_token,omitempty"`
}

// CreateCommunityNoteBody represents the request body for creating a community note
type CreateCommunityNoteBody struct {
	TestMode bool                       `json:"test_mode"`
	PostID   string                     `json:"post_id"`
	Info     *CommunityNoteInfoRequest  `json:"info"`
}

// CommunityNoteInfoRequest represents the info section of a create note request
type CommunityNoteInfoRequest struct {
	Text               string   `json:"text"`
	Classification     string   `json:"classification"`
	MisleadingTags     []string `json:"misleading_tags,omitempty"`
	TrustworthySources []string `json:"trustworthy_sources,omitempty"`
}

func communityNoteFieldsToString(fields []CommunityNoteField) []string {
	slice := make([]string, len(fields))
	for i, field := range fields {
		slice[i] = string(field)
	}
	return slice
}