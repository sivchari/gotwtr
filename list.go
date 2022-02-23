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

type AllListsOwnedResponse struct {
	Lists    []*List             `json:"data"`
	Includes *ListIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *ListMeta           `json:"meta"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type ListResponse struct {
	List     *List               `json:"data"`
	Includes *ListIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type ListTweetsResponse struct {
	Tweets   []*Tweet            `json:"data"`
	Includes *ListIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *ListMeta           `json:"meta"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type ListFollowersResponse struct {
	Users    []*User             `json:"data"`
	Includes *ListIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *ListMeta           `json:"meta"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type AllListsUserFollowsResponse struct {
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

type ListsSpecifiedUserResponse struct {
	Lists    []*List             `json:"data"`
	Includes *ListIncludes       `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *ListMeta           `json:"meta"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type PostListMembersResponse struct {
	IsMember *IsMember           `json:"data"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
}

type UndoListMembersResponse struct {
	IsMember *IsMember           `json:"data"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
}

type IsMember struct {
	IsMember bool `json:"is_member"`
}

type ListMembersBody struct {
	UserID string `json:"user_id"`
}

type CreateNewListResponse struct {
	Data *CreateNewListData `json:"data"`
}

type CreateNewListData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateNewListBody struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Private     bool   `json:"private,omitempty"`
}

type DeleteListResponse struct {
	Data *DeleteListData `json:"data"`
}

type DeleteListData struct {
	Deleted bool `json:"deleted"`
}

type UpdateMetaDataForListResponse struct {
	Data *UpdateMetaDataForListData `json:"data"`
}

type UpdateMetaDataForListData struct {
	Updated bool `json:"updated"`
}

type UpdateMetaDataForListBody struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Private     bool   `json:"private,omitempty"`
}
