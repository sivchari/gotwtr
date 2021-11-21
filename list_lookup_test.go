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

func Test_lookUpOwnedListByID(t *testing.T) {
	type args struct {
		ctx    context.Context
		client *http.Client
		id     string
		opt    []*gotwtr.ListLookUpOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.OwnedListsLookUpByIDResponse
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
				opt: []*gotwtr.ListLookUpOption{},
			},
			want: &gotwtr.OwnedListsLookUpByIDResponse{
				Lists: []*gotwtr.List{
					{
						ID:   "1451305624956858369",
						Name: "Test List",
					},
				},
				Meta: &gotwtr.ListMeta{
					ResultCount: 1,
				},
				Includes: nil,
				Errors:   nil,
			},
			wantErr: false,
		},
		{
			name: "200 ok option",
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
				id: "2244994945",
				opt: []*gotwtr.ListLookUpOption{
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
			want: &gotwtr.OwnedListsLookUpByIDResponse{
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
				Errors: nil,
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
				opt: []*gotwtr.ListLookUpOption{},
			},
			want: &gotwtr.OwnedListsLookUpByIDResponse{
				Lists: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Parameters: gotwtr.Parameter{
							ID: []string{"111111111111111111111111"},
						},
						Message: "The id query parameter value [111111111111111111111111] is not valid",
					},
				},
				Meta: nil,
				Title:  "Invalid Request",
				Detail: "One or more parameters to your request was invalid.",
				Type:   "https://api.twitter.com/2/problems/invalid-request",
			},
			wantErr: true,
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.LookUpOwnedListsByID(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.LookUpOwnedListsByID() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("client.LookUpOwnedListsByID() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}

func Test_lookUpListByID(t *testing.T) {
	type args struct {
		ctx    context.Context
		client *http.Client
		id     string
		opt    []*gotwtr.ListLookUpOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.ListLookUpByIDResponse
		wantErr bool
	}{
		{
			name: "200 ok no option",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
					"data": {
						"id": "84839422",
						"name": "Official Twitter Accounts"
					}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id:  "84839422",
				opt: []*gotwtr.ListLookUpOption{},
			},
			want: &gotwtr.ListLookUpByIDResponse{
				List: &gotwtr.List{
					ID:   "84839422",
					Name: "Official Twitter Accounts",
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
						"data": {
							"follower_count": 906,
							"id": "84839422",
							"name": "Official Twitter Accounts",
							"owner_id": "783214"
						},
						"includes": {
							"users": [
								{
									"id": "783214",
									"name": "Twitter",
									"username": "Twitter"
								}
							]
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id: "84839422",
				opt: []*gotwtr.ListLookUpOption{
					{
						UserFields: []gotwtr.UserField{
							gotwtr.UserFieldID,
							gotwtr.UserFieldName,
							gotwtr.UserFieldUserName,
						},
					},
				},
			},
			want: &gotwtr.ListLookUpByIDResponse{
				List: &gotwtr.List{
					FollowerCount: 906,
					ID:            "84839422",
					Name:          "Official Twitter Accounts",
					OwnerID:       "783214",
				},
				Includes: &gotwtr.ListIncludes{
					Users: []*gotwtr.User{
						{
							ID:       "783214",
							Name:     "Twitter",
							UserName: "Twitter",
						},
					},
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
				opt: []*gotwtr.ListLookUpOption{},
			},
			want: &gotwtr.ListLookUpByIDResponse{
				List: nil,
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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.LookUpListByID(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.LookUpListByID() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("client.LookUpListByID() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
