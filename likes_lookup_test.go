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

func Test_usersLikingTweet(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		opt    []*gotwtr.UsersLikingTweetOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.UsersLikingTweetResponse
		wantErr bool
	}{
		{
			name: "200 success with no optional fileds",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"id": "1065249714214457345",
								"name": "Spaces",
								"username": "TwitterSpaces"
							},
							{
								"id": "783214",
								"name": "Twitter",
								"username": "Twitter"
							},
							{
								"id": "1526228120",
								"name": "Twitter Data",
								"username": "TwitterData"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
			},
			want: &gotwtr.UsersLikingTweetResponse{
				Users: []*gotwtr.LookUpUsersWhoLiked{
					{
						ID:       "1065249714214457345",
						Name:     "Spaces",
						UserName: "TwitterSpaces",
					},
					{
						ID:       "783214",
						Name:     "Twitter",
						UserName: "Twitter",
					},
					{
						ID:       "1526228120",
						Name:     "Twitter Data",
						UserName: "TwitterData",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "200 success with optional fileds",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
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
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				opt: []*gotwtr.UsersLikingTweetOption{
					{
						Expansions:  []gotwtr.Expansion{gotwtr.ExpansionPinnedTweetID},
						UserFields:  []gotwtr.UserField{gotwtr.UserFieldCreatedAt},
						TweetFields: []gotwtr.TweetField{gotwtr.TweetFieldCreatedAt},
					},
				},
			},
			want: &gotwtr.UsersLikingTweetResponse{
				Users: []*gotwtr.LookUpUsersWhoLiked{
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
				Includes: &gotwtr.LookUpUsersWhoLikedIncludes{
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
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New(gotwtr.WithBearerToken("key"), gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.UsersLikingTweet(tt.args.ctx, "tweet_id", tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.UsersLikingTweet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.UsersLikingTweet() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func Test_tweetsUserLiked(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		opt    []*gotwtr.TweetsUserLikedOpts
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.TweetsUserLikedResponse
		wantErr bool
	}{
		{
			name: "200 success with no optional fileds",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
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
						]
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
			},
			want: &gotwtr.TweetsUserLikedResponse{
				Tweets: []*gotwtr.TweetsUserLiked{
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
			},
			wantErr: false,
		},
		{
			name: "200 success with optional fileds",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
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
						]
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				opt: []*gotwtr.TweetsUserLikedOpts{
					{
						TweetFields: []gotwtr.TweetField{gotwtr.TweetFieldCreatedAt, gotwtr.TweetFieldSource},
					},
				},
			},
			want: &gotwtr.TweetsUserLikedResponse{
				Tweets: []*gotwtr.TweetsUserLiked{
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
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New(gotwtr.WithBearerToken("key"), gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.TweetsUserLiked(tt.args.ctx, "user_id", tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("clientTweetsUserLiked(). error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.UTweetsUserLiked() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
