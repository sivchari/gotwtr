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

func Test_LookupUserBookmarks(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		userID string
		opt    *gotwtr.LookupUserBookmarksOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.LookupUserBookmarksResponse
		wantErr bool
	}{
		{
			name: "200 Lookup user bookmarks",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"id": "1362449997430542337",
								"text": "Honored to be the first developer to be featured in @TwitterDev's love fest 🥰♥️😍 https://t.co/g8TsPoZsij"
							},
							{
								"id": "1365416026435854338",
								"text": "We're so happy for our Official Partner @Brandwatch and their big news. https://t.co/3DwWBNSq0o https://t.co/bDUGbgPkKO"
							},
							{
								"id": "1296487407475462144",
								"text": "Check out this feature on @TwitterDev to learn more about how we're mining social media data to make sense of this evolving #publichealth crisis https://t.co/sIFLXRSvEX."
							},
							{
								"id": "1294346980072624128",
								"text": "I awake from five years of slumber https://t.co/OEPVyAFcfB"
							},
							{
								"id": "1283153843367206912",
								"text": "@wongmjane Wish we could tell you more, but I’m only a teapot 👀"
							}
						],
						"meta": {
							"result_count": 5,
							"next_token": "zldjwdz3w6sba13nbs0mbravfipbtqvbiqplg9h0p4k"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID: "2244994945",
				opt:    &gotwtr.LookupUserBookmarksOption{},
			},
			want: &gotwtr.LookupUserBookmarksResponse{
				Tweets: []*gotwtr.Tweet{
					{
						ID:   "1362449997430542337",
						Text: "Honored to be the first developer to be featured in @TwitterDev's love fest 🥰♥️😍 https://t.co/g8TsPoZsij",
					},
					{
						ID:   "1365416026435854338",
						Text: "We're so happy for our Official Partner @Brandwatch and their big news. https://t.co/3DwWBNSq0o https://t.co/bDUGbgPkKO",
					},
					{
						ID:   "1296487407475462144",
						Text: "Check out this feature on @TwitterDev to learn more about how we're mining social media data to make sense of this evolving #publichealth crisis https://t.co/sIFLXRSvEX.",
					},
					{
						ID:   "1294346980072624128",
						Text: "I awake from five years of slumber https://t.co/OEPVyAFcfB",
					},
					{
						ID:   "1283153843367206912",
						Text: "@wongmjane Wish we could tell you more, but I’m only a teapot 👀",
					},
				},
				Meta: &gotwtr.LookupUserBookmarksMeta{
					ResultCount: 5,
					NextToken:   "zldjwdz3w6sba13nbs0mbravfipbtqvbiqplg9h0p4k",
				},
			},
			wantErr: false,
		},
		{
			name: "200 Lookup user bookmarks with Option",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
					  "data": [
					    {
					      "created_at": "2021-02-18T17:12:47.000Z",
					      "source": "Twitter Web App",
					      "id": "1362449997430542337",
					      "text": "Honored to be the first developer to be featured in @TwitterDev's love fest 🥰♥️😍 https://t.co/g8TsPoZsij"
					    },
					    {
					      "created_at": "2021-02-26T21:38:43.000Z",
					      "source": "Twitter Web App",
					      "id": "1365416026435854338",
					      "text": "We're so happy for our Official Partner @Brandwatch and their big news. https://t.co/3DwWBNSq0o https://t.co/bDUGbgPkKO"
					    },
					    {
					      "created_at": "2020-08-20T16:41:00.000Z",
					      "source": "Twitter Web App",
					      "id": "1296487407475462144",
					      "text": "Check out this feature on @TwitterDev to learn more about how we're mining social media data to make sense of this evolving #publichealth crisis https://t.co/sIFLXRSvEX."
					    },
					    {
					      "created_at": "2020-08-14T18:55:42.000Z",
					      "source": "Twitter for Android",
					      "id": "1294346980072624128",
					      "text": "I awake from five years of slumber https://t.co/OEPVyAFcfB"
					    },
					    {
					      "created_at": "2020-07-14T21:38:10.000Z",
					      "source": "Twitter for  iPhone",
					      "id": "1283153843367206912",
					      "text": "@wongmjane Wish we could tell you more, but I’m only a teapot 👀"
					    }
					  ],
					  "meta": {
					    "result_count": 5,
					    "next_token": "zldjwdz3w6sba13nbs0mbravfipbtqvbiqplg9h0p4k"
					  }
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID: "2244994945",
				opt: &gotwtr.LookupUserBookmarksOption{
					MaxResults: 5,
					Expansions: []gotwtr.Expansion{
						gotwtr.ExpansionAttachmentsMediaKeys,
					},
					MediaFields: []gotwtr.MediaField{
						gotwtr.MediaFieldType,
						gotwtr.MediaFieldDurationMS,
					},
				},
			},
			want: &gotwtr.LookupUserBookmarksResponse{
				Tweets: []*gotwtr.Tweet{
					{
						CreatedAt: "2021-02-18T17:12:47.000Z",
						Source:    "Twitter Web App",
						ID:        "1362449997430542337",
						Text:      "Honored to be the first developer to be featured in @TwitterDev's love fest 🥰♥️😍 https://t.co/g8TsPoZsij",
					},
					{
						CreatedAt: "2021-02-26T21:38:43.000Z",
						Source:    "Twitter Web App",
						ID:        "1365416026435854338",
						Text:      "We're so happy for our Official Partner @Brandwatch and their big news. https://t.co/3DwWBNSq0o https://t.co/bDUGbgPkKO",
					},
					{
						CreatedAt: "2020-08-20T16:41:00.000Z",
						Source:    "Twitter Web App",
						ID:        "1296487407475462144",
						Text:      "Check out this feature on @TwitterDev to learn more about how we're mining social media data to make sense of this evolving #publichealth crisis https://t.co/sIFLXRSvEX.",
					},
					{
						CreatedAt: "2020-08-14T18:55:42.000Z",
						Source:    "Twitter for Android",
						ID:        "1294346980072624128",
						Text:      "I awake from five years of slumber https://t.co/OEPVyAFcfB",
					},
					{
						CreatedAt: "2020-07-14T21:38:10.000Z",
						Source:    "Twitter for  iPhone",
						ID:        "1283153843367206912",
						Text:      "@wongmjane Wish we could tell you more, but I’m only a teapot 👀",
					},
				},
				Meta: &gotwtr.LookupUserBookmarksMeta{
					ResultCount: 5,
					NextToken:   "zldjwdz3w6sba13nbs0mbravfipbtqvbiqplg9h0p4k",
				},
			},
			wantErr: false,
		},
		{
			name: "404 not found, invalid user id",
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
				userID: "1228393702244134912",
				opt:    &gotwtr.LookupUserBookmarksOption{},
			},
			want: &gotwtr.LookupUserBookmarksResponse{
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
			got, err := c.LookupUserBookmarks(tt.args.ctx, tt.args.userID, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.LookupUserBookmarks() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.LookupUserBookmarks() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}

func Test_BookmarkTweet(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		userID string
		body   *gotwtr.BookmarkTweetBody
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.BookmarkTweetResponse
		wantErr bool
	}{
		{
			name: "200 Bookmark tweet",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
                      "data": {
                        "bookmarked": true
                       }
                    }`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID: "1441162269824405510",
				body: &gotwtr.BookmarkTweetBody{
					TweetID: "1441162269824405528",
				},
			},
			want: &gotwtr.BookmarkTweetResponse{
				BookmarkTweetData: &gotwtr.BookmarkTweetData{
					Bookmarked: true,
				},
			},
			wantErr: false,
		},
		{
			name: "404 not found, invalid user id",
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
				userID: "1228393702244134912",
				body: &gotwtr.BookmarkTweetBody{
					TweetID: "1228393702244134912",
				},
			},
			want: &gotwtr.BookmarkTweetResponse{
				BookmarkTweetData: nil,
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
			got, err := c.BookmarkTweet(tt.args.ctx, tt.args.userID, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.BookmarkTweet() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.BookmarkTweet() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}

func Test_RemoveBookmarkOfTweet(t *testing.T) {
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
		want    *gotwtr.RemoveBookmarkOfTweetResponse
		wantErr bool
	}{
		{
			name: "200 remove bookmark",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
                        "data": {
                            "bookmarks": false
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
			want: &gotwtr.RemoveBookmarkOfTweetResponse{
				RemoveBookmarkOfTweetData: &gotwtr.RemoveBookmarkOfTweetData{
					Bookmarks: false,
				},
			},
			wantErr: false,
		},
		{
			name: "404 not found, invalid tweet id",
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
				userID:  "2244994945",
				tweetID: "11111111111111111",
			},
			want: &gotwtr.RemoveBookmarkOfTweetResponse{
				RemoveBookmarkOfTweetData: nil,
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
			name: "404 not found, invalid user id",
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
				userID:  "1228393702244134912",
				tweetID: "11111111111111111",
			},
			want: &gotwtr.RemoveBookmarkOfTweetResponse{
				RemoveBookmarkOfTweetData: nil,
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
			got, err := c.RemoveBookmarkOfTweet(tt.args.ctx, tt.args.userID, tt.args.tweetID)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.RemoveBookmarkOfTweet() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.RemoveBookmarkOfTweet() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
