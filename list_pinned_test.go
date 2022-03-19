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

func Test_pinnedLists(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		id     string
		opt    []*gotwtr.PinnedListsOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.PinnedListsResponse
		wantErr bool
	}{
		{
			name: "200 ok no option",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"id": "1451305624956858369",
								"name": "Test List"
							}
						],
						"meta": {
							"result_count": 1
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id:  "2244994945",
				opt: []*gotwtr.PinnedListsOption{},
			},
			want: &gotwtr.PinnedListsResponse{
				Lists: []*gotwtr.List{
					{
						ID:   "1451305624956858369",
						Name: "Test List",
					},
				},
				Meta: &gotwtr.ListMeta{
					ResultCount: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "200 ok and option",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"follower_count": 0,
								"id": "1451305624956858369",
								"name": "Test List",
								"owner_id": "2244994945"
							}
						],
						"includes": {
							"users": [
								{
									"username": "TwitterDev",
									"id": "2244994945",
									"created_at": "2013-12-14T04:35:55.000Z",
									"name": "Twitter Dev"
								}
							]
						},
						"meta": {
							"result_count": 1
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id: "84839422",
				opt: []*gotwtr.PinnedListsOption{
					{
						UserFields: []gotwtr.UserField{
							gotwtr.UserFieldUserName,
							gotwtr.UserFieldID,
							gotwtr.UserFieldCreatedAt,
							gotwtr.UserFieldName,
						},
					},
				},
			},
			want: &gotwtr.PinnedListsResponse{
				Lists: []*gotwtr.List{
					{
						FollowerCount: 0,
						ID:            "1451305624956858369",
						Name:          "Test List",
						OwnerID:       "2244994945",
					},
				},
				Includes: &gotwtr.ListIncludes{
					Users: []*gotwtr.User{
						{
							UserName:  "TwitterDev",
							ID:        "2244994945",
							CreatedAt: "2013-12-14T04:35:55.000Z",
							Name:      "Twitter Dev",
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
			name: "404 not found",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"errors":[
							{
								"parameters":{
									"id":[
										"111111111111111111111111"
									]
								},
								"message":"The id query parameter value [111111111111111111111111] is not valid"
							}
						],
						"title":"Invalid Request",
						"detail":"One or more parameters to your request was invalid.",
						"type":"https://api.twitter.com/2/problems/invalid-request"
					}`
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id:  "111111111111111111111111",
				opt: []*gotwtr.PinnedListsOption{},
			},
			want: &gotwtr.PinnedListsResponse{
				Lists: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Parameters: gotwtr.Parameter{
							ID: []string{"111111111111111111111111"},
						},
						Message: "The id query parameter value [111111111111111111111111] is not valid",
					},
				},
				Title:  "Invalid Request",
				Detail: "One or more parameters to your request was invalid.",
				Type:   "https://api.twitter.com/2/problems/invalid-request",
			},
			wantErr: true,
		},
	}
	for i, tt := range tests {
		i := i
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.PinnedLists(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.PinendLists() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.PinnedLists() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}

func Test_postPinnedLists(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		listID string
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.PostPinnedListsResponse
		wantErr bool
	}{
		{
			name: "200 success",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					if req.Method != http.MethodPost {
						t.Errorf("the method is not correct got %s want %s", req.Method, http.MethodPost)
					}
					body := `{
						"data": {
							"pinned": true
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				listID: "6253282",
				userID: "2244994945",
			},
			want: &gotwtr.PostPinnedListsResponse{
				Pinned: &gotwtr.Pinned{
					Pinned: true,
				},
			},
			wantErr: false,
		},
		{
			name: "404 not found, invalid listID",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"errors":[
							{
								"value":"111111111122",
								"detail":"Could not find list with id: [111111111122].",
								"title":"Not Found Error",
								"resource_type":"list",
								"parameter":"id",
								"resource_id":"111111111122",
								"type":"https://api.twitter.com/2/problems/resource-not-found"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				listID: "111111111122",
				userID: "1228393702244134912",
			},
			want: &gotwtr.PostPinnedListsResponse{
				Pinned: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Value:        "111111111122",
						Detail:       "Could not find list with id: [111111111122].",
						Title:        "Not Found Error",
						ResourceType: "list",
						Parameter:    "id",
						ResourceID:   "111111111122",
						Type:         "https://api.twitter.com/2/problems/resource-not-found",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "404 not found, invalid userID",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"errors":[
							{
								"value":"1228393702244134912",
								"detail":"Could not find user with id: [1228393702244134912].",
								"title":"Not Found Error",
								"resource_type":"user",
								"parameter":"id",
								"resource_id":"1228393702244134912",
								"type":"https://api.twitter.com/2/problems/resource-not-found"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				listID: "111111111122",
				userID: "1228393702244134912",
			},
			want: &gotwtr.PostPinnedListsResponse{
				Pinned: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Value:        "1228393702244134912",
						Detail:       "Could not find user with id: [1228393702244134912].",
						Title:        "Not Found Error",
						ResourceType: "user",
						Parameter:    "id",
						ResourceID:   "1228393702244134912",
						Type:         "https://api.twitter.com/2/problems/resource-not-found",
					},
				},
			},
			wantErr: true,
		},
	}
	for i, tt := range tests {
		i := i
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.PostPinnedLists(tt.args.ctx, tt.args.listID, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostPinnedLists() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("PostPinnedLists() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}

func Test_undoPinnedLists(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		listID string
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.UndoPinnedListsResponse
		wantErr bool
	}{
		{
			name: "200 success",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					if req.Method != http.MethodDelete {
						t.Errorf("the method is not correct got %s want %s", req.Method, http.MethodDelete)
					}
					body := `{
						"data": {
							"Pinned": false
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				listID: "2244994945",
				userID: "6253282",
			},
			want: &gotwtr.UndoPinnedListsResponse{
				Pinned: &gotwtr.Pinned{
					Pinned: false,
				},
			},
			wantErr: false,
		},
		{
			name: "404 not found, invalid listID",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"errors":[
							{
								"value":"111111111122",
								"detail":"Could not find list with id: [111111111122].",
								"title":"Not Found Error",
								"resource_type":"list",
								"parameter":"id",
								"resource_id":"111111111122",
								"type":"https://api.twitter.com/2/problems/resource-not-found"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				listID: "111111111122",
				userID: "1228393702244134912",
			},
			want: &gotwtr.UndoPinnedListsResponse{
				Pinned: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Value:        "111111111122",
						Detail:       "Could not find list with id: [111111111122].",
						Title:        "Not Found Error",
						ResourceType: "list",
						Parameter:    "id",
						ResourceID:   "111111111122",
						Type:         "https://api.twitter.com/2/problems/resource-not-found",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "404 not found, invalid userID",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"errors":[
							{
								"value":"111111111133",
								"detail":"Could not find user with id: [111111111133].",
								"title":"Not Found Error",
								"resource_type":"user",
								"parameter":"id",
								"resource_id":"111111111133",
								"type":"https://api.twitter.com/2/problems/resource-not-found"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				listID: "111111111133",
				userID: "1228393702244134912",
			},
			want: &gotwtr.UndoPinnedListsResponse{
				Pinned: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Value:        "111111111133",
						Detail:       "Could not find user with id: [111111111133].",
						Title:        "Not Found Error",
						ResourceType: "user",
						Parameter:    "id",
						ResourceID:   "111111111133",
						Type:         "https://api.twitter.com/2/problems/resource-not-found",
					},
				},
			},
			wantErr: true,
		},
	}
	for i, tt := range tests {
		i := i
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.UndoPinnedLists(tt.args.ctx, tt.args.listID, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UndoPinnedLists() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("UndoPinnedLists() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
