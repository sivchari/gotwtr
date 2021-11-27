package gotwtr

type ListField string

const (
	ListFieldCreatedAt   ListField = "created_at"
	ListFollowerCount    ListField = "follower_count"
	ListMemberCount      ListField = "member_count"
	ListFieldPrivate     ListField = "private"
	ListFieldDescription ListField = "description"
	ListOwnerID          ListField = "owner_id"
)

type List struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	CreatedAt     string `json:"created_at,omitempty"`
	Private       bool   `json:"private,omitempty"`
	FollowerCount int    `json:"follower_count,omitempty"`
	MemberCount   int    `json:"member_count,omitempty"`
	OwnerID       string `json:"owner_id,omitempty"`
	Description   string `json:"description,omitempty"`
}

type ListIncludes struct {
	Tweets []*Tweet
	Users  []*User
}

type ListMeta struct {
	ResultCount   int    `json:"result_count"`
	PreviousToken string `json:"previous_token,omitempty"`
	NextToken     string `json:"next_token,omitempty"`
}

type OwnedListsLookUpByIDResponse struct {
	Lists    []*List             `json:"data"`
	Includes *ListIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *ListMeta           `json:"meta"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type ListLookUpByIDResponse struct {
	List     *List               `json:"data"`
	Includes *ListIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type ListsTweetsLookUpByIDResponse struct {
	Tweets   []*Tweet            `json:"data"`
	Includes *ListIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *ListMeta           `json:"meta"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type ListFollowersLookUpByIDResponse struct {
	Users    []*User             `json:"data"`
	Includes *ListIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *ListMeta           `json:"meta"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type ListsUserFollowingLookUpByIDResponse struct {
	Lists    []*List             `json:"data"`
	Includes *ListIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *ListMeta           `json:"meta"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type ListMembersResponse struct {
	Users    []*User             `json:"data"`
	Includes *ListIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *ListMeta           `json:"meta"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}
