package gotwtr_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sivchari/gotwtr"
)

func Test_lookUp(t *testing.T) {
	type args struct {
		ctx    context.Context
		client *http.Client
		ids    []string
		opt    []*gotwtr.TweetOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.TweetLookUpResponse
		wantErr bool
	}{
		{
			name: "success lookup tweets, no option and single ID",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"id": "123456789",
								"text": "Hello, world!"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				ids: []string{"123456789"},
				opt: []*gotwtr.TweetOption{},
			},
			want: &gotwtr.TweetLookUpResponse{
				Tweets: []*gotwtr.Tweet{
					{
						ID:   "123456789",
						Text: "Hello, world!",
					},
				},
				Includes: nil,
				Errors:   nil,
			},
			wantErr: false,
		},
		{
			name: "success lookup tweets, option and single ID",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"author_id": "11111111",
								"id": "123456789",
								"created_at": "2020-01-01T00:00:00Z",
								"text": "Hello, world!"
							}
						],
						"includes": {
							"users": [
								{
									"id": "11111111",
									"username": "sivchari :D",
									"name": "sivchari",
									"verified": true
								}
							]
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				ids: []string{"123456789"},
				opt: []*gotwtr.TweetOption{
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
			want: &gotwtr.TweetLookUpResponse{
				Tweets: []*gotwtr.Tweet{
					{
						AuthorID:  "11111111",
						ID:        "123456789",
						CreatedAt: "2020-01-01T00:00:00Z",
						Text:      "Hello, world!",
					},
				},
				Includes: &gotwtr.TweetIncludes{
					Users: []*gotwtr.User{
						{
							ID:       "11111111",
							UserName: "sivchari :D",
							Name:     "sivchari",
							Verified: true,
						},
					},
				},
				Errors: nil,
			},
			wantErr: false,
		},
		{
			name: "success lookup tweets, no option and multiple IDs",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"id": "123456789",
								"text": "Hello, world!"
							},
							{
								"id": "987654321",
								"text": "Hello, Go!"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				ids: []string{"123456789", "987654321"},
				opt: []*gotwtr.TweetOption{},
			},
			want: &gotwtr.TweetLookUpResponse{
				Tweets: []*gotwtr.Tweet{
					{
						ID:   "123456789",
						Text: "Hello, world!",
					},
					{
						ID:   "987654321",
						Text: "Hello, Go!",
					},
				},
				Includes: nil,
				Errors:   nil,
			},
			wantErr: false,
		},
		{
			name: "success lookup tweets, option and multiple IDs",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"author_id": "11111111",
								"id": "123456789",
								"created_at": "2020-01-01T00:00:00Z",
								"text": "Hello, world!"
							},
							{
								"author_id": "22222222",
								"id": "987654321",
								"created_at": "2020-01-02T00:00:00Z",
								"text": "Hello, Go!"
							}
						],
						"includes": {
							"users": [
								{
									"id": "11111111",
									"username": "sivchari :D",
									"name": "sivchari",
									"verified": true
								},
								{
									"id": "22222222",
									"username": "twitter :D",
									"name": "twitter",
									"verified": true
								}
							]
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				ids: []string{"123456789", "987654321"},
				opt: []*gotwtr.TweetOption{
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
			want: &gotwtr.TweetLookUpResponse{
				Tweets: []*gotwtr.Tweet{
					{
						AuthorID:  "11111111",
						ID:        "123456789",
						CreatedAt: "2020-01-01T00:00:00Z",
						Text:      "Hello, world!",
					},
					{
						AuthorID:  "22222222",
						ID:        "987654321",
						CreatedAt: "2020-01-02T00:00:00Z",
						Text:      "Hello, Go!",
					},
				},
				Includes: &gotwtr.TweetIncludes{
					Users: []*gotwtr.User{
						{
							ID:       "11111111",
							UserName: "sivchari :D",
							Name:     "sivchari",
							Verified: true,
						},
						{
							ID:       "22222222",
							UserName: "twitter :D",
							Name:     "twitter",
							Verified: true,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "success 1 is valid, 1 is deleted",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"id": "20",
								"text": "just setting up my twttr"
							}
						],
						"errors": [
							{
								"detail": "Could not find tweet with ids: [1276230436478386177].",
								"title": "Not Found Error",
								"resource_type": "tweet",
								"parameter": "ids",
								"value": "1276230436478386177",
								"type": "https://api.twitter.com/2/problems/resource-not-found"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				ids: []string{"20", "1276230436478386177"},
				opt: []*gotwtr.TweetOption{},
			},
			want: &gotwtr.TweetLookUpResponse{
				Tweets: []*gotwtr.Tweet{
					{
						ID:   "20",
						Text: "just setting up my twttr",
					},
				},
				Errors: []*gotwtr.APIResponseError{
					{
						Detail:       "Could not find tweet with ids: [1276230436478386177].",
						Title:        "Not Found Error",
						ResourceType: "tweet",
						Parameter:    "ids",
						Value:        "1276230436478386177",
						Type:         "https://api.twitter.com/2/problems/resource-not-found",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "error not found",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					ids := "`ids`"
					body := fmt.Sprintf(`{
						"errors": [
							{
								"parameters": {
									"ids": [
										"123456789"
									]
								},
								"message": "The %v query parameter value [14421240904714485799] does not match ^[0-9]{1,19}$"
							}
						],
						"title": "Invalid Request",
						"detail": "One or more parameters to your request was invalid.",
						"type": "https://api.twitter.com/2/problems/invalid-request"
					}`, ids)
					return &http.Response{
						StatusCode: http.StatusBadRequest,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				ids: []string{"123456789"},
				opt: []*gotwtr.TweetOption{},
			},
			want: &gotwtr.TweetLookUpResponse{
				Tweets: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Parameters: gotwtr.Parameter{
							IDs: []string{"123456789"},
						},
						Message: "The `ids` query parameter value [14421240904714485799] does not match ^[0-9]{1,19}$",
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
			got, err := c.LookUpTweets(tt.args.ctx, tt.args.ids, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.LookUpTweets() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("client.LookUpTweets() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
