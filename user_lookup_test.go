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

func Test_retrieveMultipleUsersWithIDs(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		ids    []string
		opt    []*gotwtr.RetrieveUserOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.UsersResponse
		wantErr bool
	}{
		{
			name: "200 ok no option and single ID",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"id": "2244994945",
								"username": "TwitterDev",
								"name": "Twitter Dev"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				ids: []string{"2244994945"},
				opt: []*gotwtr.RetrieveUserOption{},
			},
			want: &gotwtr.UsersResponse{
				Users: []*gotwtr.User{
					{
						ID:       "2244994945",
						UserName: "TwitterDev",
						Name:     "Twitter Dev",
					},
				},
				Includes: nil,
				Errors:   nil,
			},
			wantErr: false,
		},
		{
			name: "200 ok option and single ID",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"created_at": "2013-12-14T04:35:55.000Z",
								"username": "TwitterDev",
								"pinned_tweet_id": "1255542774432063488",
								"id": "2244994945",
								"name": "Twitter Dev"
							}
						],
						"includes": {
							"tweets": [
								{
									"created_at": "2020-04-29T17:01:38.000Z",
									"text": "During these unprecedented times, what’s happening on Twitter can help the world better understand &amp; respond to the pandemic. nnWe're launching a free COVID-19 stream endpoint so qualified devs &amp; researchers can study the public conversation in real-time. https://t.co/BPqMcQzhId",
									"id": "1255542774432063488"
								},
								{
									"created_at": "2020-06-19T21:12:30.000Z",
									"text": "📍 Minneapolisn🗣️ @FredTJoseph https://t.co/lNTOkyguG1",
									"id": "1274087687469715457"
								}
							]
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				ids: []string{"2244994945"},
				opt: []*gotwtr.RetrieveUserOption{
					{
						Expansions: []gotwtr.Expansion{gotwtr.ExpansionAuthorID},
						TweetFields: []gotwtr.TweetField{
							gotwtr.TweetFieldCreatedAt,
						},
						UserFields: []gotwtr.UserField{
							gotwtr.UserFieldUserName,
							gotwtr.UserFieldVerified,
						},
					},
				},
			},
			want: &gotwtr.UsersResponse{
				Users: []*gotwtr.User{
					{
						CreatedAt:     "2013-12-14T04:35:55.000Z",
						UserName:      "TwitterDev",
						PinnedTweetID: "1255542774432063488",
						ID:            "2244994945",
						Name:          "Twitter Dev",
					},
				},
				Includes: &gotwtr.UserIncludes{
					Tweets: []*gotwtr.Tweet{
						{
							CreatedAt: "2020-04-29T17:01:38.000Z",
							Text:      "During these unprecedented times, what’s happening on Twitter can help the world better understand &amp; respond to the pandemic. nnWe're launching a free COVID-19 stream endpoint so qualified devs &amp; researchers can study the public conversation in real-time. https://t.co/BPqMcQzhId",
							ID:        "1255542774432063488",
						},
						{
							CreatedAt: "2020-06-19T21:12:30.000Z",
							Text:      "📍 Minneapolisn🗣️ @FredTJoseph https://t.co/lNTOkyguG1",
							ID:        "1274087687469715457",
						},
					},
				},
				Errors: nil,
			},
			wantErr: false,
		},
		{
			name: "200 ok no option and multiple IDs",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"id": "2244994945",
								"username": "TwitterDev",
								"name": "Twitter Dev"
							},
							{
								"id": "783214",
								"username": "Twitter",
								"name": "Twitter"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				ids: []string{"2244994945", "783214"},
				opt: []*gotwtr.RetrieveUserOption{},
			},
			want: &gotwtr.UsersResponse{
				Users: []*gotwtr.User{
					{
						ID:       "2244994945",
						UserName: "TwitterDev",
						Name:     "Twitter Dev",
					},
					{
						ID:       "783214",
						UserName: "Twitter",
						Name:     "Twitter",
					},
				},
				Includes: nil,
				Errors:   nil,
			},
			wantErr: false,
		},
		{
			name: "200 ok option and multiple IDs",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"created_at": "2013-12-14T04:35:55.000Z",
								"username": "TwitterDev",
								"pinned_tweet_id": "1255542774432063488",
								"id": "2244994945",
								"name": "Twitter Dev"
							},
							{
								"created_at": "2007-02-20T14:35:54.000Z",
								"username": "Twitter",
								"pinned_tweet_id": "1274087687469715457",
								"id": "783214",
								"name": "Twitter"
							}
						],
						"includes": {
							"tweets": [
								{
									"created_at": "2020-04-29T17:01:38.000Z",
									"text": "During these unprecedented times, what’s happening on Twitter can help the world better understand &amp; respond to the pandemic. nnWe're launching a free COVID-19 stream endpoint so qualified devs &amp; researchers can study the public conversation in real-time. https://t.co/BPqMcQzhId",
									"id": "1255542774432063488"
								},
								{
									"created_at": "2020-06-19T21:12:30.000Z",
									"text": "📍 Minneapolisn🗣️ @FredTJoseph https://t.co/lNTOkyguG1",
									"id": "1274087687469715457"
								}
							]
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				ids: []string{"2244994945", "783214"},
				opt: []*gotwtr.RetrieveUserOption{
					{
						Expansions: []gotwtr.Expansion{gotwtr.ExpansionAuthorID},
						TweetFields: []gotwtr.TweetField{
							gotwtr.TweetFieldCreatedAt,
						},
						UserFields: []gotwtr.UserField{
							gotwtr.UserFieldUserName,
							gotwtr.UserFieldVerified,
						},
					},
				},
			},
			want: &gotwtr.UsersResponse{
				Users: []*gotwtr.User{
					{
						CreatedAt:     "2013-12-14T04:35:55.000Z",
						UserName:      "TwitterDev",
						PinnedTweetID: "1255542774432063488",
						ID:            "2244994945",
						Name:          "Twitter Dev",
					},
					{
						CreatedAt:     "2007-02-20T14:35:54.000Z",
						UserName:      "Twitter",
						PinnedTweetID: "1274087687469715457",
						ID:            "783214",
						Name:          "Twitter",
					},
				},
				Includes: &gotwtr.UserIncludes{
					Tweets: []*gotwtr.Tweet{
						{
							CreatedAt: "2020-04-29T17:01:38.000Z",
							Text:      "During these unprecedented times, what’s happening on Twitter can help the world better understand &amp; respond to the pandemic. nnWe're launching a free COVID-19 stream endpoint so qualified devs &amp; researchers can study the public conversation in real-time. https://t.co/BPqMcQzhId",
							ID:        "1255542774432063488",
						},
						{
							CreatedAt: "2020-06-19T21:12:30.000Z",
							Text:      "📍 Minneapolisn🗣️ @FredTJoseph https://t.co/lNTOkyguG1",
							ID:        "1274087687469715457",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "200 ok 1 is valid 1 is deleted",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"id":"6253282",
								"name":"Twitter API",
								"username":"TwitterAPI"
							}
						],
						"errors": [
							{
								"value":"11111111111",
								"detail":"Could not find user with ids: [11111111111].",
								"title":"Not Found Error",
								"resource_type":"user",
								"parameter":"ids",
								"resource_id":"11111111111",
								"type":"https://api.twitter.com/2/problems/resource-not-found"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				ids: []string{"6253282", "11111111111"},
				opt: []*gotwtr.RetrieveUserOption{},
			},
			want: &gotwtr.UsersResponse{
				Users: []*gotwtr.User{
					{
						ID:       "6253282",
						Name:     "Twitter API",
						UserName: "TwitterAPI",
					},
				},
				Errors: []*gotwtr.APIResponseError{
					{
						Value:        "11111111111",
						Detail:       "Could not find user with ids: [11111111111].",
						Title:        "Not Found Error",
						ResourceType: "user",
						Parameter:    "ids",
						ResourceID:   "11111111111",
						Type:         "https://api.twitter.com/2/problems/resource-not-found",
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
						"errors": [
							{
								"value":"11111111111",
								"detail":"Could not find user with ids: [11111111111].",
								"title":"Not Found Error",
								"resource_type":"user",
								"parameter":"ids",
								"resource_id":"11111111111",
								"type":"https://api.twitter.com/2/problems/resource-not-found"
							},
							{
								"value":"22222222222",
								"detail":"Could not find user with ids: [22222222222].",
								"title":"Not Found Error",
								"resource_type":"user",
								"parameter":"ids",
								"resource_id":"22222222222",
								"type":"https://api.twitter.com/2/problems/resource-not-found"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				ids: []string{"11111111111", "22222222222"},
				opt: []*gotwtr.RetrieveUserOption{},
			},
			want: &gotwtr.UsersResponse{
				Users: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Value:        "11111111111",
						Detail:       "Could not find user with ids: [11111111111].",
						Title:        "Not Found Error",
						ResourceType: "user",
						Parameter:    "ids",
						ResourceID:   "11111111111",
						Type:         "https://api.twitter.com/2/problems/resource-not-found",
					},
					{
						Value:        "22222222222",
						Detail:       "Could not find user with ids: [22222222222].",
						Title:        "Not Found Error",
						ResourceType: "user",
						Parameter:    "ids",
						ResourceID:   "22222222222",
						Type:         "https://api.twitter.com/2/problems/resource-not-found",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.RetrieveMultipleUsersWithIDs(tt.args.ctx, tt.args.ids, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.RetriveMultipleUsersWithIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.RetriveMultipleUsersWithIDs() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func Test_retrieveSingleUserWithID(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		id     string
		opt    []*gotwtr.RetrieveUserOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.UserResponse
		wantErr bool
	}{
		{
			name: "200 ok default payload",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": {
							"id": "2244994945",
							"name": "Twitter Dev",
							"username": "TwitterDev"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id:  "2244994945",
				opt: []*gotwtr.RetrieveUserOption{},
			},
			want: &gotwtr.UserResponse{
				User: &gotwtr.User{
					ID:       "2244994945",
					Name:     "Twitter Dev",
					UserName: "TwitterDev",
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
							"username": "TwitterDev",
							"created_at": "2013-12-14T04:35:55.000Z",
							"pinned_tweet_id": "1255542774432063488",
							"id": "2244994945",
							"name": "Twitter Dev"
						},
						"includes": {
							"tweets": [
								{
									"text": "During these unprecedented times, what’s happening on Twitter can help the world better understand &amp; respond to the pandemic. nnWe're launching a free COVID-19 stream endpoint so qualified devs &amp; researchers can study the public conversation in real-time. https://t.co/BPqMcQzhId",
									"created_at": "2020-04-29T17:01:38.000Z",
									"id": "1255542774432063488"
								}
							]
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id: "2244994945",
				opt: []*gotwtr.RetrieveUserOption{
					{
						Expansions: []gotwtr.Expansion{gotwtr.ExpansionAuthorID},
						TweetFields: []gotwtr.TweetField{
							gotwtr.TweetFieldCreatedAt,
						},
						UserFields: []gotwtr.UserField{
							gotwtr.UserFieldUserName,
							gotwtr.UserFieldVerified,
						},
					},
				},
			},
			want: &gotwtr.UserResponse{
				User: &gotwtr.User{
					UserName:      "TwitterDev",
					CreatedAt:     "2013-12-14T04:35:55.000Z",
					PinnedTweetID: "1255542774432063488",
					ID:            "2244994945",
					Name:          "Twitter Dev",
				},
				Includes: &gotwtr.UserIncludes{
					Tweets: []*gotwtr.Tweet{
						{
							Text:      "During these unprecedented times, what’s happening on Twitter can help the world better understand &amp; respond to the pandemic. nnWe're launching a free COVID-19 stream endpoint so qualified devs &amp; researchers can study the public conversation in real-time. https://t.co/BPqMcQzhId",
							CreatedAt: "2020-04-29T17:01:38.000Z",
							ID:        "1255542774432063488",
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
								"value":"11111111111",
								"detail":"Could not find user with id: [11111111111].",
								"title":"Not Found Error",
								"resource_type":"user",
								"parameter":"id",
								"resource_id":"11111111111",
								"type":"https://api.twitter.com/2/problems/resource-not-found"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id:  "11111111111",
				opt: []*gotwtr.RetrieveUserOption{},
			},
			want: &gotwtr.UserResponse{
				User: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Value:        "11111111111",
						Detail:       "Could not find user with id: [11111111111].",
						Title:        "Not Found Error",
						ResourceType: "user",
						Parameter:    "id",
						ResourceID:   "11111111111",
						Type:         "https://api.twitter.com/2/problems/resource-not-found",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.RetrieveSingleUserWithID(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.RetriveSingleUserWithID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.RetriveSingleUserWithID() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func Test_retrieveSingleUserWithUserName(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		name   string
		opt    []*gotwtr.RetrieveUserOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.UserResponse
		wantErr bool
	}{
		{
			name: "200 ok default payload",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": {
							"id": "2244994945",
							"name": "Twitter Dev",
							"username": "TwitterDev"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				name: "TwitterDev",
				opt:  []*gotwtr.RetrieveUserOption{},
			},
			want: &gotwtr.UserResponse{
				User: &gotwtr.User{
					ID:       "2244994945",
					Name:     "Twitter Dev",
					UserName: "TwitterDev",
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
							"username": "TwitterDev",
							"created_at": "2013-12-14T04:35:55.000Z",
							"pinned_tweet_id": "1255542774432063488",
							"id": "2244994945",
							"name": "Twitter Dev"
						},
						"includes": {
							"tweets": [
								{
									"text": "During these unprecedented times, what’s happening on Twitter can help the world better understand &amp; respond to the pandemic. nnWe're launching a free COVID-19 stream endpoint so qualified devs &amp; researchers can study the public conversation in real-time. https://t.co/BPqMcQzhId",
									"created_at": "2020-04-29T17:01:38.000Z",
									"id": "1255542774432063488"
								}
							]
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				name: "TwitterDev",
				opt: []*gotwtr.RetrieveUserOption{
					{
						Expansions: []gotwtr.Expansion{gotwtr.ExpansionAuthorID},
						TweetFields: []gotwtr.TweetField{
							gotwtr.TweetFieldCreatedAt,
						},
						UserFields: []gotwtr.UserField{
							gotwtr.UserFieldUserName,
							gotwtr.UserFieldVerified,
						},
					},
				},
			},
			want: &gotwtr.UserResponse{
				User: &gotwtr.User{
					UserName:      "TwitterDev",
					CreatedAt:     "2013-12-14T04:35:55.000Z",
					PinnedTweetID: "1255542774432063488",
					ID:            "2244994945",
					Name:          "Twitter Dev",
				},
				Includes: &gotwtr.UserIncludes{
					Tweets: []*gotwtr.Tweet{
						{
							Text:      "During these unprecedented times, what’s happening on Twitter can help the world better understand &amp; respond to the pandemic. nnWe're launching a free COVID-19 stream endpoint so qualified devs &amp; researchers can study the public conversation in real-time. https://t.co/BPqMcQzhId",
							CreatedAt: "2020-04-29T17:01:38.000Z",
							ID:        "1255542774432063488",
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
									"username":[
										"aaaaaaaaaaaaaaaaaa"
									]
								},
								"message":"The username query parameter value [aaaaaaaaaaaaaaaaaa] does not match ^[A-Za-z0-9_]{1,15}$"
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
				name: "aaaaaaaaaaaaaaaaaa",
				opt:  []*gotwtr.RetrieveUserOption{},
			},
			want: &gotwtr.UserResponse{
				User: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Parameters: gotwtr.Parameter{
							UserName: []string{"aaaaaaaaaaaaaaaaaa"},
						},
						Message: "The username query parameter value [aaaaaaaaaaaaaaaaaa] does not match ^[A-Za-z0-9_]{1,15}$",
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
			got, err := c.RetrieveSingleUserWithUserName(tt.args.ctx, tt.args.name, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.RetrieveSingleUserWithUserName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("client.RetrieveSingleUserWithUserName() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func Test_retrieveMultipleUsersWithUserNames(t *testing.T) {
	type args struct {
		ctx    context.Context
		client *http.Client
		names  []string
		opt    []*gotwtr.RetrieveUserOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.UsersResponse
		wantErr bool
	}{
		{
			name: "200 ok default payload",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"id": "2244994945",
								"username": "TwitterDev",
								"name": "Twitter Dev"
							},
							{
								"id": "783214",
								"username": "Twitter",
								"name": "Twitter"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				names: []string{"TwitterDev", "Twitter"},
				opt:   []*gotwtr.RetrieveUserOption{},
			},
			want: &gotwtr.UsersResponse{
				Users: []*gotwtr.User{
					{
						ID:       "2244994945",
						Name:     "Twitter Dev",
						UserName: "TwitterDev",
					},
					{
						ID:       "783214",
						UserName: "Twitter",
						Name:     "Twitter",
					},
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
								"created_at": "2013-12-14T04:35:55.000Z",
								"username": "TwitterDev",
								"pinned_tweet_id": "1255542774432063488",
								"id": "2244994945",
								"name": "Twitter Dev"
							},
							{
								"created_at": "2007-02-20T14:35:54.000Z",
								"username": "Twitter",
								"pinned_tweet_id": "1274087687469715457",
								"id": "783214",
								"name": "Twitter"
							}
						],
						"includes": {
							"tweets": [
								{
									"created_at": "2020-04-29T17:01:38.000Z",
									"text": "During these unprecedented times, what’s happening on Twitter can help the world better understand &amp; respond to the pandemic. nnWe're launching a free COVID-19 stream endpoint so qualified devs &amp; researchers can study the public conversation in real-time. https://t.co/BPqMcQzhId",
									"id": "1255542774432063488"
								},
								{
									"created_at": "2020-06-19T21:12:30.000Z",
									"text": "📍 Minneapolisn🗣️ @FredTJoseph https://t.co/lNTOkyguG1",
									"id": "1274087687469715457"
								}
							]
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				names: []string{"TwitterDev", "Twitter"},
				opt: []*gotwtr.RetrieveUserOption{
					{
						Expansions: []gotwtr.Expansion{gotwtr.ExpansionAuthorID},
						TweetFields: []gotwtr.TweetField{
							gotwtr.TweetFieldCreatedAt,
						},
						UserFields: []gotwtr.UserField{
							gotwtr.UserFieldUserName,
							gotwtr.UserFieldVerified,
						},
					},
				},
			},
			want: &gotwtr.UsersResponse{
				Users: []*gotwtr.User{
					{
						CreatedAt:     "2013-12-14T04:35:55.000Z",
						UserName:      "TwitterDev",
						PinnedTweetID: "1255542774432063488",
						ID:            "2244994945",
						Name:          "Twitter Dev",
					},
					{
						CreatedAt:     "2007-02-20T14:35:54.000Z",
						UserName:      "Twitter",
						PinnedTweetID: "1274087687469715457",
						ID:            "783214",
						Name:          "Twitter",
					},
				},
				Includes: &gotwtr.UserIncludes{
					Tweets: []*gotwtr.Tweet{
						{
							CreatedAt: "2020-04-29T17:01:38.000Z",
							Text:      "During these unprecedented times, what’s happening on Twitter can help the world better understand &amp; respond to the pandemic. nnWe're launching a free COVID-19 stream endpoint so qualified devs &amp; researchers can study the public conversation in real-time. https://t.co/BPqMcQzhId",
							ID:        "1255542774432063488",
						},
						{
							CreatedAt: "2020-06-19T21:12:30.000Z",
							Text:      "📍 Minneapolisn🗣️ @FredTJoseph https://t.co/lNTOkyguG1",
							ID:        "1274087687469715457",
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
									"usernames":[
										"aaaaaaaaaaaaaaaaaaaaaa","bbbbbbbbbbbbbbbbbbbbbbbb"
									]
								},
								"message":"The usernames query parameter value [aaaaaaaaaaaaaaaaaaaaaa] does not match ^[A-Za-z0-9_]{1,15}$"
							},
							{
								"parameters":{
									"usernames":[
										"aaaaaaaaaaaaaaaaaaaaaa","bbbbbbbbbbbbbbbbbbbbbbbb"
									]
								},
								"message":"The usernames query parameter value [bbbbbbbbbbbbbbbbbbbbbbbb] does not match ^[A-Za-z0-9_]{1,15}$"
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
				names: []string{"aaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbbbbb"},
				opt:   []*gotwtr.RetrieveUserOption{},
			},
			want: &gotwtr.UsersResponse{
				Users: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Parameters: gotwtr.Parameter{
							UserNames: []string{"aaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbbbbb"},
						},
						Message: "The usernames query parameter value [aaaaaaaaaaaaaaaaaaaaaa] does not match ^[A-Za-z0-9_]{1,15}$",
					},
					{
						Parameters: gotwtr.Parameter{
							UserNames: []string{"aaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbbbbb"},
						},
						Message: "The usernames query parameter value [bbbbbbbbbbbbbbbbbbbbbbbb] does not match ^[A-Za-z0-9_]{1,15}$",
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
			got, err := c.RetrieveMultipleUsersWithUserNames(tt.args.ctx, tt.args.names, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.RetrieveMultipleUsersWithUserNames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.RetrieveMultipleUsersWithUserNames() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
