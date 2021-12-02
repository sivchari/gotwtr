package gotwtr_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/sivchari/gotwtr"
)

func Test_listMembers(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		id     string
		opt    []*gotwtr.ListMembersOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.ListMembersResponse
		wantErr bool
	}{
		{
			name: "200 ok default payload",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"data": [
							{
								"id": "1319036828964454402",
								"name": "Birdwatch",
								"username": "birdwatch"
							},
							{
								"id": "1244731491088809984",
								"name": "Twitter Thailand",
								"username": "TwitterThailand"
							},
							{
								"id": "1194267639100723200",
								"name": "Twitter Retweets",
								"username": "TwitterRetweets"
							},
							{
								"id": "1168976680867762177",
								"name": "Twitter Able",
								"username": "TwitterAble"
							},
							{
								"id": "1065249714214457345",
								"name": "Spaces",
								"username": "TwitterSpaces"
							}
						],
						"meta": {
							"result_count": 5,
							"next_token": "5676935732641845249"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				id:  "84839422",
				opt: []*gotwtr.ListMembersOption{},
			},
			want: &gotwtr.ListMembersResponse{
				Users: []*gotwtr.User{
					{
						ID:       "1319036828964454402",
						Name:     "Birdwatch",
						UserName: "birdwatch",
					},
					{
						ID:       "1244731491088809984",
						Name:     "Twitter Thailand",
						UserName: "TwitterThailand",
					},
					{
						ID:       "1194267639100723200",
						Name:     "Twitter Retweets",
						UserName: "TwitterRetweets",
					},
					{
						ID:       "1168976680867762177",
						Name:     "Twitter Able",
						UserName: "TwitterAble",
					},
					{
						ID:       "1065249714214457345",
						Name:     "Spaces",
						UserName: "TwitterSpaces",
					},
				},
				Meta: &gotwtr.ListMeta{
					ResultCount: 5,
					NextToken:   "5676935732641845249",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.ListMembers(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.ListMembers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.ListMembers() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func Test_ListSpecifiedUserMemberOf(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		id     string
		opt    []*gotwtr.ListSpecifiedUserMemberOfOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.ListSpecifiedUserMemberOfResponse
		wantErr bool
	}{
		{
			name: "200 ok default payload",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"data": [
							   {
							     "id": "1451951974291689472",
							     "name": "Twitter"
							   },
							   {
							     "id": "1451812298184540161",
							     "name": "Updates"
							   },
							   {
							     "id": "1450519480132509697",
							     "name": "Twitter"
							   }
                        ],
					    "meta": {
							"result_count": 3
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				id:  "84839422",
				opt: []*gotwtr.ListSpecifiedUserMemberOfOption{},
			},
			want: &gotwtr.ListSpecifiedUserMemberOfResponse{
				Lists: []*gotwtr.List{
					{
						ID:   "1451951974291689472",
						Name: "Twitter",
					},
					{
						ID:   "1451812298184540161",
						Name: "Updates",
					},
					{
						ID:   "1450519480132509697",
						Name: "Twitter",
					},
				},
				Meta: &gotwtr.ListMeta{
					ResultCount: 3,
				},
			},
			wantErr: false,
		},
		{
			name: "200 ok option",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"data": [
							{
								"follower_count": 5,
								"id": "1451951974291689472",
								"name": "Twitter",
								"owner_id": "1227213680120479745"
							}
						],
						"includes": {
							"users": [
								{
									"name": "구돆",
									"created_at": "2020-02-11T12:52:11.000Z",
									"id": "1227213680120479745",
									"username": "Follow__Y0U"
								}
							]
						},
						"meta": {
							"result_count": 1
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				id:  "84839422",
				opt: []*gotwtr.ListSpecifiedUserMemberOfOption{},
			},
			want: &gotwtr.ListSpecifiedUserMemberOfResponse{
				Lists: []*gotwtr.List{
					{
						ID:            "1451951974291689472",
						Name:          "Twitter",
						FollowerCount: 5,
						OwnerID:       "1227213680120479745",
					},
				},
				Includes: &gotwtr.ListIncludes{
					Users: []*gotwtr.User{
						{
							Name:      "구돆",
							CreatedAt: "2020-02-11T12:52:11.000Z",
							ID:        "1227213680120479745",
							UserName:  "Follow__Y0U",
						},
					},
				},
				Meta: &gotwtr.ListMeta{
					ResultCount: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "User Not Found",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"errors":[
							{
								"value":"849422",
								"detail":"Could not find user with id: [849422].",
								"title":"Not Found Error",
								"resource_type":"user",
								"parameter":"id",
								"resource_id":"849422",
								"type":"https://api.twitter.com/2/problems/resource-not-found"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				id:  "849422",
				opt: []*gotwtr.ListSpecifiedUserMemberOfOption{},
			},
			want: &gotwtr.ListSpecifiedUserMemberOfResponse{
				Errors: []*gotwtr.APIResponseError{
					{
						Value:        "849422",
						Detail:       "Could not find user with id: [849422].",
						Title:        "Not Found Error",
						ResourceType: "user",
						Parameter:    "id",
						ResourceID:   "849422",
						Type:         "https://api.twitter.com/2/problems/resource-not-found",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid Request",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"errors":[
							{
								"parameters":{
									"id": ["8488877666666666666666666666666622839422"]
								},
								"message":"The 'id' query parameter value [8488877666666666666666666666666622839422] is not valid"
							}
						],
						"title":"Invalid Request",
						"detail":"One or more parameters to your request was invalid.",
						"type":"https://api.twitter.com/2/problems/invalid-request"
					}`
					return &http.Response{
						StatusCode: http.StatusBadRequest,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				id:  "8488877666666666666666666666666622839422",
				opt: []*gotwtr.ListSpecifiedUserMemberOfOption{},
			},
			want: &gotwtr.ListSpecifiedUserMemberOfResponse{
				Errors: []*gotwtr.APIResponseError{
					{
						Parameters: gotwtr.Parameter{
							ID: []string{"8488877666666666666666666666666622839422"},
						},
						Message: "The 'id' query parameter value [8488877666666666666666666666666622839422] is not valid",
					},
				},
				Title:  "Invalid Request",
				Detail: "One or more parameters to your request was invalid.",
				Type:   "https://api.twitter.com/2/problems/invalid-request",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.ListSpecifiedUserMemberOf(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.ListSpecifiedUserMemberOf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.ListSpecifiedUserMemberOf() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
