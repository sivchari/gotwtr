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

func Test_muting(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		userID string
		opt    []*gotwtr.MuteOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.MutingResponse
		wantErr bool
	}{
		{
			name: "200 ok default",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"id": "2244994945",
								"name": "Twitter Dev",
								"username": "TwitterDev"
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
				userID: "2244994945",
				opt:    []*gotwtr.MuteOption{},
			},
			want: &gotwtr.MutingResponse{
				Users: []*gotwtr.User{
					{
						ID:       "2244994945",
						Name:     "Twitter Dev",
						UserName: "TwitterDev",
					},
				},
				Meta: &gotwtr.MutesMeta{
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
								"username": "TwitterDev",
								"created_at": "2013-12-14T04:35:55.000Z",
								"id": "2244994945",
								"name": "Twitter Dev",
								"pinned_tweet_id": "1430984356139470849"
							}
						],
						"includes": {
							"tweets": [
								{
									"created_at": "2021-08-26T20:03:51.000Z",
									"id": "1430984356139470849",
									"text": "Help us build a better Twitter Developer Platform!n nTake the annual developer survey &gt;&gt;&gt; https://t.co/9yTbEKlJHH https://t.co/fYIwKPzqua"
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
				userID: "2244994945",
				opt: []*gotwtr.MuteOption{
					{
						TweetFields: []gotwtr.TweetField{
							gotwtr.TweetFieldCreatedAt,
							gotwtr.TweetFieldID,
							gotwtr.TweetFieldText,
						},
					},
				},
			},
			want: &gotwtr.MutingResponse{
				Users: []*gotwtr.User{
					{
						UserName:      "TwitterDev",
						CreatedAt:     "2013-12-14T04:35:55.000Z",
						ID:            "2244994945",
						Name:          "Twitter Dev",
						PinnedTweetID: "1430984356139470849",
					},
				},
				Includes: &gotwtr.UserIncludes{
					Tweets: []*gotwtr.Tweet{
						{
							CreatedAt: "2021-08-26T20:03:51.000Z",
							ID:        "1430984356139470849",
							Text:      "Help us build a better Twitter Developer Platform!n nTake the annual developer survey &gt;&gt;&gt; https://t.co/9yTbEKlJHH https://t.co/fYIwKPzqua",
						},
					},
				},
				Meta: &gotwtr.MutesMeta{
					ResultCount: 1,
				},
			},
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
										"1111111111111111111111"
									]
								},
								"message":"The id query parameter value [1111111111111111111111] is not valid"
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
				userID: "1111111111111111111111",
				opt:    []*gotwtr.MuteOption{},
			},
			want: &gotwtr.MutingResponse{
				Users: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Parameters: gotwtr.Parameter{
							ID: []string{"1111111111111111111111"},
						},
						Message: "The id query parameter value [1111111111111111111111] is not valid",
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
			got, err := c.Muting(tt.args.ctx, tt.args.userID, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Muting() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.Muting() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}

func Test_postMuting(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx          context.Context
		client       *http.Client
		userID       string
		targetUserID string
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.PostMutingResponse
		wantErr bool
	}{
		{
			name: "200 success",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					if req.Method != http.MethodPost {
						t.Fatalf("the method is not correct got %s want %s", req.Method, http.MethodPost)
					}
					body := `{
						"data": {
							"muting": true
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID:       "6253282",
				targetUserID: "2244994945",
			},
			want: &gotwtr.PostMutingResponse{
				Muting: &gotwtr.Muting{
					Muting: true,
				},
			},
			wantErr: false,
		},
		{
			name: "403 authentication error",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"errors": [
							{
								"title": "Unsupported Authentication",
								"detail": "Authenticating with OAuth 2.0 Application-Only is forbidden for this endpoint.  Supported authentication types are [OAuth 1.0a User Context, OAuth 2.0 User Context].",
								"type": "https://api.twitter.com/2/problems/unsupported-authentication",
								"status": 403
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusForbidden,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID:       "111111111",
				targetUserID: "1228393702244134912",
			},
			want: &gotwtr.PostMutingResponse{
				Muting: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Title:  "Unsupported Authentication",
						Detail: "Authenticating with OAuth 2.0 Application-Only is forbidden for this endpoint.  Supported authentication types are [OAuth 1.0a User Context, OAuth 2.0 User Context].",
						Type:   "https://api.twitter.com/2/problems/unsupported-authentication",
						Status: 403,
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
								"value":"1111111111",
								"detail":"Could not find user with id: [1111111111].",
								"title":"Not Found Error",
								"resource_type":"user",
								"parameter":"id",
								"resource_id":"1111111111",
								"type":"https://api.twitter.com/2/problems/resource-not-found"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID:       "1111111111",
				targetUserID: "1228393702244134912",
			},
			want: &gotwtr.PostMutingResponse{
				Muting: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Value:        "1111111111",
						Detail:       "Could not find user with id: [1111111111].",
						Title:        "Not Found Error",
						ResourceType: "user",
						Parameter:    "id",
						ResourceID:   "1111111111",
						Type:         "https://api.twitter.com/2/problems/resource-not-found",
					},
				},
			},
			wantErr: true,
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.PostMuting(tt.args.ctx, tt.args.userID, tt.args.targetUserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostMuting() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("PostMuting() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}

func Test_undoMuting(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx          context.Context
		client       *http.Client
		sourceUserID string
		targetUserID string
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.UndoMutingResponse
		wantErr bool
	}{
		{
			name: "200 success",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					if req.Method != http.MethodDelete {
						t.Fatalf("the method is not correct got %s want %s", req.Method, http.MethodDelete)
					}
					body := `{
						"data": {
							"muting": false
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				sourceUserID: "2244994945",
				targetUserID: "6253282",
			},
			want: &gotwtr.UndoMutingResponse{
				Muting: &gotwtr.Muting{
					Muting: false,
				},
			},
			wantErr: false,
		},
		{
			name: "403 authentication error",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"errors": [
							{
								"title": "Unsupported Authentication",
								"detail": "Authenticating with OAuth 2.0 Application-Only is forbidden for this endpoint.  Supported authentication types are [OAuth 1.0a User Context, OAuth 2.0 User Context].",
								"type": "https://api.twitter.com/2/problems/unsupported-authentication",
								"status": 403
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusForbidden,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				sourceUserID: "111111111",
				targetUserID: "1228393702244134912",
			},
			want: &gotwtr.UndoMutingResponse{
				Muting: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Title:  "Unsupported Authentication",
						Detail: "Authenticating with OAuth 2.0 Application-Only is forbidden for this endpoint.  Supported authentication types are [OAuth 1.0a User Context, OAuth 2.0 User Context].",
						Type:   "https://api.twitter.com/2/problems/unsupported-authentication",
						Status: 403,
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
								"value":"1111111111",
								"detail":"Could not find user with id: [1111111111].",
								"title":"Not Found Error",
								"resource_type":"user",
								"parameter":"id",
								"resource_id":"1111111111",
								"type":"https://api.twitter.com/2/problems/resource-not-found"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				sourceUserID: "1111111111",
				targetUserID: "1228393702244134912",
			},
			want: &gotwtr.UndoMutingResponse{
				Muting: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Value:        "1111111111",
						Detail:       "Could not find user with id: [1111111111].",
						Title:        "Not Found Error",
						ResourceType: "user",
						Parameter:    "id",
						ResourceID:   "1111111111",
						Type:         "https://api.twitter.com/2/problems/resource-not-found",
					},
				},
			},
			wantErr: true,
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.UndoMuting(tt.args.ctx, tt.args.sourceUserID, tt.args.targetUserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UndoMuting() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("UndoMuting() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
