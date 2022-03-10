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

func Test_retweetsLookup(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx     context.Context
		client  *http.Client
		tweetID string
		opt     []*gotwtr.RetweetsLookupOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.RetweetsResponse
		wantErr bool
	}{
		{
			name: "200 success",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"id": "1065249714214457345",
								"created_at": "2018-11-21T14:24:58.000Z",
								"name": "Spaces",
								"pinned_tweet_id": "1389270063807598594",
								"description": "Twitter Spaces is where live audio conversations happen.",
								"username": "TwitterSpaces"
							},
							{
								"id": "783214",
								"created_at": "2007-02-20T14:35:54.000Z",
								"name": "Twitter",
								"description": "What's happening?!",
								"username": "Twitter"
							},
							{
								"id": "1526228120",
								"created_at": "2013-06-17T23:57:45.000Z",
								"name": "Twitter Data",
								"description": "Data-driven insights about notable moments and conversations from Twitter, Inc., plus tips and tricks to help you get the most out of Twitter data.",
								"username": "TwitterData"
							},
							{
								"id": "2244994945",
								"created_at": "2013-12-14T04:35:55.000Z",
								"name": "Twitter Dev",
								"pinned_tweet_id": "1354143047324299264",
								"description": "The voice of the #TwitterDev team and your official source for updates, news, and events, related to the #TwitterAPI.",
								"username": "TwitterDev"
							},
							{
								"id": "6253282",
								"created_at": "2007-05-23T06:01:13.000Z",
								"name": "Twitter API",
								"pinned_tweet_id": "1293595870563381249",
								"description": "Tweets about changes and service issues. Follow @TwitterDev for more.",
								"username": "TwitterAPI"
							}
						],
						"includes": {
							"tweets": [
								{
									"id": "1389270063807598594",
									"text": "now, everyone with 600 or more followers can host a Space.nnbased on what we've learned, these accounts are likely to have a good experience hosting because of their existing audience. before bringing the ability to create a Space to everyone, we're focused on a few things. :thread:"
								},
								{
									"id": "1354143047324299264",
									"text": "Academics are one of the biggest groups using the #TwitterAPI to research what's happening. Their work helps make the world (&amp; Twitter) a better place, and now more than ever, we must enable more of it. nIntroducing :drum_with_drumsticks: the Academic Research product track!nhttps://t.co/nOFiGewAV2"
								},
								{
									"id": "1293595870563381249",
									"text": "Twitter API v2: Early Access releasednnToday we announced Early Access to the first endpoints of the new Twitter API!nn#TwitterAPI #EarlyAccess #VersionBump https://t.co/g7v3aeIbtQ"
								}
							]
						},
						"meta": {
							"result_count": 5
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				tweetID: "1354143047324299264",
			},
			want: &gotwtr.RetweetsResponse{
				Users: []*gotwtr.User{
					{
						ID:            "1065249714214457345",
						CreatedAt:     "2018-11-21T14:24:58.000Z",
						Name:          "Spaces",
						PinnedTweetID: "1389270063807598594",
						Description:   "Twitter Spaces is where live audio conversations happen.",
						UserName:      "TwitterSpaces",
					},
					{
						ID:          "783214",
						CreatedAt:   "2007-02-20T14:35:54.000Z",
						Name:        "Twitter",
						Description: "What's happening?!",
						UserName:    "Twitter",
					},
					{
						ID:          "1526228120",
						CreatedAt:   "2013-06-17T23:57:45.000Z",
						Name:        "Twitter Data",
						Description: "Data-driven insights about notable moments and conversations from Twitter, Inc., plus tips and tricks to help you get the most out of Twitter data.",
						UserName:    "TwitterData",
					},
					{
						ID:            "2244994945",
						CreatedAt:     "2013-12-14T04:35:55.000Z",
						Name:          "Twitter Dev",
						PinnedTweetID: "1354143047324299264",
						Description:   "The voice of the #TwitterDev team and your official source for updates, news, and events, related to the #TwitterAPI.",
						UserName:      "TwitterDev",
					},
					{
						ID:            "6253282",
						CreatedAt:     "2007-05-23T06:01:13.000Z",
						Name:          "Twitter API",
						PinnedTweetID: "1293595870563381249",
						Description:   "Tweets about changes and service issues. Follow @TwitterDev for more.",
						UserName:      "TwitterAPI",
					},
				},
				Includes: &gotwtr.TweetIncludes{
					Tweets: []*gotwtr.Tweet{
						{
							ID:   "1389270063807598594",
							Text: "now, everyone with 600 or more followers can host a Space.nnbased on what we've learned, these accounts are likely to have a good experience hosting because of their existing audience. before bringing the ability to create a Space to everyone, we're focused on a few things. :thread:",
						},
						{
							ID:   "1354143047324299264",
							Text: "Academics are one of the biggest groups using the #TwitterAPI to research what's happening. Their work helps make the world (&amp; Twitter) a better place, and now more than ever, we must enable more of it. nIntroducing :drum_with_drumsticks: the Academic Research product track!nhttps://t.co/nOFiGewAV2",
						},
						{
							ID:   "1293595870563381249",
							Text: "Twitter API v2: Early Access releasednnToday we announced Early Access to the first endpoints of the new Twitter API!nn#TwitterAPI #EarlyAccess #VersionBump https://t.co/g7v3aeIbtQ",
						},
					},
				},
				Meta: &gotwtr.RetweetsLookupMeta{
					ResultCount: 5,
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
								"value":"11111111111111111",
								"detail":"Could not find tweet with id: [11111111111111111].",
								"title":"Not Found Error",
								"resource_type":"tweet",
								"parameter":"id",
								"resource_id":"11111111111111111",
								"type":"https://api.twitter.com/2/problems/resource-not-found"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				tweetID: "11111111111111111",
			},
			want: &gotwtr.RetweetsResponse{
				Users: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Value:        "11111111111111111",
						Detail:       "Could not find tweet with id: [11111111111111111].",
						Title:        "Not Found Error",
						ResourceType: "tweet",
						Parameter:    "id",
						ResourceID:   "11111111111111111",
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
			got, err := c.RetweetsLookup(tt.args.ctx, tt.args.tweetID, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.RetweetsLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.RetweetsLookup() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func Test_postRetweet(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx     context.Context
		client  *http.Client
		userID  string
		tweetID string
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.PostRetweetResponse
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
							"retweeted": true
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID:  "2244994945",
				tweetID: "1228393702244134912",
			},
			want: &gotwtr.PostRetweetResponse{
				Retweeted: &gotwtr.Retweeted{
					Retweeted: true,
				},
			},
			wantErr: false,
		},
		{
			name: "400 request failed",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"errors": [
							{
								"message":"Sorry, that page does not exist, code:34"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusBadRequest,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID:  "2244994945",
				tweetID: "1228393702244134912",
			},
			want: &gotwtr.PostRetweetResponse{
				Retweeted: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Message: "Sorry, that page does not exist, code:34",
					},
				},
			},
			wantErr: true,
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
				userID:  "111111111",
				tweetID: "1228393702244134912",
			},
			want: &gotwtr.PostRetweetResponse{
				Retweeted: nil,
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
			name: "404 not found, invalid tweetID",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"errors":[
							{
								"value":"111111111111111",
								"detail":"Could not find tweet with id: [111111111111111].",
								"title":"Not Found Error",
								"resource_type":"tweet",
								"parameter":"id",
								"resource_id":"111111111111111",
								"type":"https://api.twitter.com/2/problems/resource-not-found"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID:  "2244994945",
				tweetID: "111111111111111",
			},
			want: &gotwtr.PostRetweetResponse{
				Retweeted: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Value:        "111111111111111",
						Detail:       "Could not find tweet with id: [111111111111111].",
						Title:        "Not Found Error",
						ResourceType: "tweet",
						Parameter:    "id",
						ResourceID:   "111111111111111",
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
				userID:  "1111111111",
				tweetID: "1228393702244134912",
			},
			want: &gotwtr.PostRetweetResponse{
				Retweeted: nil,
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
		i := i
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.PostRetweet(tt.args.ctx, tt.args.userID, tt.args.tweetID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostRetweet() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("PostRetweet() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}

func Test_undoRetweet(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx           context.Context
		client        *http.Client
		userID        string
		sourceTweetID string
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.UndoRetweetResponse
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
							"retweeted": false
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID:        "2244994945",
				sourceTweetID: "1228393702244134912",
			},
			want: &gotwtr.UndoRetweetResponse{
				Retweeted: &gotwtr.Retweeted{
					Retweeted: false,
				},
			},
			wantErr: false,
		},
		{
			name: "400 request failed",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"errors": [
							{
								"message":"Sorry, that page does not exist, code:34"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusBadRequest,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID:        "2244994945",
				sourceTweetID: "1228393702244134912",
			},
			want: &gotwtr.UndoRetweetResponse{
				Retweeted: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Message: "Sorry, that page does not exist, code:34",
					},
				},
			},
			wantErr: true,
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
				userID:        "111111111",
				sourceTweetID: "1228393702244134912",
			},
			want: &gotwtr.UndoRetweetResponse{
				Retweeted: nil,
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
			name: "404 not found, invalid tweetID",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"errors":[
							{
								"value":"11111111111111111",
								"detail":"Could not find tweet with id: [11111111111111111].",
								"title":"Not Found Error",
								"resource_type":"tweet",
								"parameter":"id",
								"resource_id":"11111111111111111",
								"type":"https://api.twitter.com/2/problems/resource-not-found"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID:        "2244994945",
				sourceTweetID: "11111111111111111",
			},
			want: &gotwtr.UndoRetweetResponse{
				Retweeted: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Value:        "11111111111111111",
						Detail:       "Could not find tweet with id: [11111111111111111].",
						Title:        "Not Found Error",
						ResourceType: "tweet",
						Parameter:    "id",
						ResourceID:   "11111111111111111",
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
				userID:        "1111111111",
				sourceTweetID: "1228393702244134912",
			},
			want: &gotwtr.UndoRetweetResponse{
				Retweeted: nil,
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
		i := i
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.UndoRetweet(tt.args.ctx, tt.args.userID, tt.args.sourceTweetID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UndoRetweet() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("UndoRetweet() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
