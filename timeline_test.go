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

func Test_userTweetTimeline(t *testing.T) {
	type args struct {
		ctx    context.Context
		client *http.Client
		id     string
		opt    []*gotwtr.TweetOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.UserTweetTimelineResponse
		wantErr bool
	}{
		{
			name: "200 success single ID default",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
						  {
							"id": "1338971066773905408",
							"text": "💡 Using Twitter data for academic research? Join our next livestream this Friday @ 9am PT on https://t.co/GrtBOXh5Y1!\n \n@SuhemParack will show how to get started with recent search &amp; filtered stream endpoints on the #TwitterAPI v2, the new Tweet payload, annotations, &amp; more. https://t.co/IraD2Z7wEg"
						  },
						  {
							"id": "1338923691497959425",
							"text": "📈 Live now with @jessicagarson and @i_am_daniele! https://t.co/Y1AFzsTTxb"
						  }
						],
						"meta": {
						  "oldest_id": "1334564488884862976",
						  "newest_id": "1338971066773905408",
						  "result_count": 2,
						  "next_token": "7140dibdnow9c7btw3w29grvxfcgvpb9n9coehpk7xz5i"
						}
					  }`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id: "1338971066773905408",
			},
			want: &gotwtr.UserTweetTimelineResponse{
				Tweets: []*gotwtr.Tweet{
					{
						ID:   "1338971066773905408",
						Text: "💡 Using Twitter data for academic research? Join our next livestream this Friday @ 9am PT on https://t.co/GrtBOXh5Y1!\n \n@SuhemParack will show how to get started with recent search &amp; filtered stream endpoints on the #TwitterAPI v2, the new Tweet payload, annotations, &amp; more. https://t.co/IraD2Z7wEg",
					},
					{
						ID:   "1338923691497959425",
						Text: "📈 Live now with @jessicagarson and @i_am_daniele! https://t.co/Y1AFzsTTxb",
					},
				},
				Meta: &gotwtr.UserTimelineMeta{
					ResultCount: 2,
					OldestID:    "1334564488884862976",
					NewestID:    "1338971066773905408",
					NextToken:   "7140dibdnow9c7btw3w29grvxfcgvpb9n9coehpk7xz5i",
				},
			},
			wantErr: false,
		},
		{
			name: "200 success Optional Fields",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
						  {
							"author_id": "2244994945",
							"conversation_id": "1338971066773905408",
							"text": "💡 Using Twitter data for academic research? Join our next livestream this Friday @ 9am PT on https://t.co/GrtBOXh5Y1!\n \n@SuhemParack will show how to get started with recent search &amp; filtered stream endpoints on the #TwitterAPI v2, the new Tweet payload, annotations, &amp; more. https://t.co/IraD2Z7wEg",
							"context_annotations": [
							  {
								"domain": {
								  "id": "47",
								  "name": "Brand",
								  "description": "Brands and Companies"
								},
								"entity": {
								  "id": "10045225402",
								  "name": "Twitter"
								}
							  }
							],
							"public_metrics": {
							  "retweet_count": 10,
							  "reply_count": 1,
							  "like_count": 41,
							  "quote_count": 4
							},
							"id": "1338971066773905408",
							"created_at": "2020-12-15T22:15:53.000Z"
						  },
						  {
							"author_id": "2244994945",
							"conversation_id": "1338923691497959425",
							"text": "📈 Live now with @jessicagarson and @i_am_daniele! https://t.co/Y1AFzsTTxb",
							"context_annotations": [
							  {
								"domain": {
								  "id": "47",
								  "name": "Brand",
								  "description": "Brands and Companies"
								},
								"entity": {
								  "id": "10026378521",
								  "name": "Google "
								}
							  }
							],
							"public_metrics": {
							  "retweet_count": 3,
							  "reply_count": 0,
							  "like_count": 12,
							  "quote_count": 1
							},
							"id": "1338923691497959425",
							"created_at": "2020-12-15T19:07:38.000Z"
						  }
						],
						"includes": {
						  "users": [
							{
							  "id": "2244994945",
							  "name": "Twitter Dev",
							  "username": "TwitterDev"
							}
						  ]
						},
						"meta": {
						  "oldest_id": "1337122535188652033",
						  "newest_id": "1338971066773905408",
						  "result_count": 2,
						  "next_token": "7140dibdnow9c7btw3w29n4v1mtag9kegr0gr7y26pnw3"
						}
					  }`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id: "2244994945",
				opt: []*gotwtr.TweetOption{
					{
						TweetFields: []gotwtr.TweetField{
							gotwtr.TweetFieldAttachments,
							gotwtr.TweetFieldAuthorID,
							gotwtr.TweetFieldCreatedAt,
							gotwtr.TweetFieldEntities,
							gotwtr.TweetFieldGeo,
							gotwtr.TweetFieldID,
							gotwtr.TweetFieldInReplyToUserID,
							gotwtr.TweetFieldLanguage,
							gotwtr.TweetFieldPossiblySensitve,
							gotwtr.TweetFieldReferencedTweets,
							gotwtr.TweetFieldSource,
							gotwtr.TweetFieldText,
							gotwtr.TweetFieldWithHeld,
						},
						UserFields: []gotwtr.UserField{
							gotwtr.UserFieldUserName,
						},
						Expansions: []gotwtr.Expansion{
							gotwtr.ExpansionAuthorID,
						},
					},
				},
			},
			want: &gotwtr.UserTweetTimelineResponse{
				Tweets: []*gotwtr.Tweet{
					{
						ID:             "1338971066773905408",
						Text:           "💡 Using Twitter data for academic research? Join our next livestream this Friday @ 9am PT on https://t.co/GrtBOXh5Y1!\n \n@SuhemParack will show how to get started with recent search &amp; filtered stream endpoints on the #TwitterAPI v2, the new Tweet payload, annotations, &amp; more. https://t.co/IraD2Z7wEg",
						AuthorID:       "2244994945",
						ConversationID: "1338971066773905408",
						CreatedAt:      "2020-12-15T22:15:53.000Z",
						ContextAnnotations: []*gotwtr.TweetContextAnnotation{
							{
								Domain: &gotwtr.TweetContextObj{
									ID:          "47",
									Name:        "Brand",
									Description: "Brands and Companies",
								},
								Entity: &gotwtr.TweetContextObj{
									ID:   "10045225402",
									Name: "Twitter",
								},
							},
						},
						PublicMetrics: &gotwtr.TweetMetrics{
							RetweetCount: 10,
							ReplyCount:   1,
							LikeCount:    41,
							QuoteCount:   4,
						},
					},
					{
						ID:             "1338923691497959425",
						Text:           "📈 Live now with @jessicagarson and @i_am_daniele! https://t.co/Y1AFzsTTxb",
						AuthorID:       "2244994945",
						ConversationID: "1338923691497959425",
						CreatedAt:      "2020-12-15T19:07:38.000Z",
						ContextAnnotations: []*gotwtr.TweetContextAnnotation{
							{
								Domain: &gotwtr.TweetContextObj{
									ID:          "47",
									Name:        "Brand",
									Description: "Brands and Companies",
								},
								Entity: &gotwtr.TweetContextObj{
									ID:   "10026378521",
									Name: "Google ",
								},
							},
						},
						PublicMetrics: &gotwtr.TweetMetrics{
							RetweetCount: 3,
							ReplyCount:   0,
							LikeCount:    12,
							QuoteCount:   1,
						},
					},
				},
				Includes: &gotwtr.TweetIncludes{
					Users: []*gotwtr.User{
						{
							ID:       "2244994945",
							Name:     "Twitter Dev",
							UserName: "TwitterDev",
						},
					},
				},
				Meta: &gotwtr.UserTimelineMeta{
					ResultCount: 2,
					OldestID:    "1337122535188652033",
					NewestID:    "1338971066773905408",
					NextToken:   "7140dibdnow9c7btw3w29n4v1mtag9kegr0gr7y26pnw3",
				},
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.UserTweetTimeline(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.UserTweetTimeline() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("client.LookUpTweetByID() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}

func Test_userMentionTimeline(t *testing.T) {
	type args struct {
		ctx    context.Context
		client *http.Client
		id     string
		opt    []*gotwtr.TweetOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.UserMentionTimelineResponse
		wantErr bool
	}{
		{
			name: "200 success single ID default",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
						  {
								"id": "1375152598945312768",
								"text": "@LeBraat @TwitterDev @LeGuud There's enough @twitterdev love to go around, @LeBraat"
							},
							{
								"id": "1375152449594523649",
								"text": "Apparently I'm not the only one (of my test accounts) that loves @TwitterDev nn@LeStaache @LeGuud"
							},
							{
								"id": "1375152043455873027",
								"text": "Can I join this @twitterdev love party?!"
							},
							{
								"id": "1375151947360174082",
								"text": "I love me some @twitterdev too!"
							},
							{
								"id": "1375151827189137412",
								"text": "This is a test, but also a good excuse to express my love for @TwitterDev 😍"
							}
						],
						"meta": {
						  "oldest_id": "1375151827189137412",
							"newest_id": "1375152598945312768",
							"result_count": 5,
							"next_token": "7140dibdnow9c7btw3w3y5b6jqjnk3k4g5zkmfkvqwa42"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id: "1338971066773905408",
			},
			want: &gotwtr.UserMentionTimelineResponse{
				Tweets: []*gotwtr.Tweet{
					{
						ID: "1375152598945312768",
						Text: "@LeBraat @TwitterDev @LeGuud There's enough @twitterdev love to go around, @LeBraat",
					},
					{
						ID: "1375152449594523649",
						Text: "Apparently I'm not the only one (of my test accounts) that loves @TwitterDev nn@LeStaache @LeGuud",
					},
					{
						ID: "1375152043455873027",
						Text: "Can I join this @twitterdev love party?!",
					},
					{
						ID: "1375151947360174082",
						Text: "I love me some @twitterdev too!",
					},
					{
						ID: "1375151827189137412",
						Text: "This is a test, but also a good excuse to express my love for @TwitterDev 😍",
					},
				},
				Meta: &gotwtr.UserTimelineMeta{
					ResultCount: 5,
					OldestID:    "1375151827189137412",
					NewestID:    "1375152598945312768",
					NextToken:   "7140dibdnow9c7btw3w3y5b6jqjnk3k4g5zkmfkvqwa42",
				},
			},
			wantErr: false,
		},
		{
			name: "200 success Optional Fields",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"author_id": "1034147061711679488",
								"text": "@LeBraat @TwitterDev @LeGuud There's enough @twitterdev love to go around, @LeBraat",
								"lang": "en",
								"conversation_id": "1375152449594523649",
								"id": "1375152598945312768"
							},
							{
								"author_id": "199566737",
								"text": "Apparently I'm not the only one (of my test accounts) that loves @TwitterDev nn@LeStaache @LeGuud",
								"lang": "en",
								"conversation_id": "1375152449594523649",
								"id": "1375152449594523649"
							},
							{
								"author_id": "930524282358325248",
								"text": "Can I join this @twitterdev love party?!",
								"lang": "en",
								"conversation_id": "1375152043455873027",
								"id": "1375152043455873027"
							},
							{
								"author_id": "1034147061711679488",
								"text": "I love me some @twitterdev too!",
								"lang": "en",
								"conversation_id": "1375151947360174082",
								"id": "1375151947360174082"
							},
							{
								"author_id": "199566737",
								"text": "This is a test, but also a good excuse to express my love for @TwitterDev 😍",
								"lang": "en",
								"conversation_id": "1375151827189137412",
								"id": "1375151827189137412"
							}
						],
						"includes": {
							"users": [
								{
									"name": "LeStache",
									"id": "1034147061711679488",
									"entities": [
										{
											"url": {
												"urls": [
													{
														"start": 0,
														"end": 23,
														"url": "https://t.co/7IDoW8iFLm",
														"expanded_url": "https://twitter.com",
														"display_url": "twitter.com"
													}
												]
											},
											"description": {
												"urls": [
													{
														"start": 21,
														"end": 44,
														"url": "https://t.co/v6nxjDjZR3",
														"expanded_url": "https://google.com",
														"display_url": "google.com"
													}
												],
												"hashtags": [
													{
														"start": 15,
														"end": 20,
														"tag": "test"
													}
												],
												"user_mentions": [
													{
														"start": 0,
														"end": 8,
														"username": "lebraat"
													}
												],
												"cashtags": [
													{
														"start": 9,
														"end": 14,
														"tag": "twtr"
													}
												]
											}
										}
									],
									"username": "LeStaache",
									"created_at": "2018-08-27T18:34:07.000Z"
								},
								{
									"name": "dan dale",
									"id": "199566737",
									"entities": [
										{
											"description": {
												"hashtags": [
													{
														"start": 30,
														"end": 37,
														"tag": "devrel"
													}
												],
												"user_mentions": [
													{
														"start": 18,
														"end": 29,
														"username": "twitterdev"
													}
												]
											}
										}
									],
									"username": "LeBraat",
									"created_at": "2010-10-07T05:36:28.000Z"
								},
								{
									"name": "LeGuud",
									"id": "930524282358325248",
									"entities": [
										{
											"url": {
												"urls": [
													{
														"start": 0,
														"end": 23,
														"url": "https://t.co/8IkCzClPCz",
														"expanded_url": "https://developer.twitter.com",
														"display_url": "developer.twitter.com"
													}
												]
											}
										}
									],
									"username": "LeGuud",
									"created_at": "2017-11-14T19:54:12.000Z"
								}
							]
						},
						"meta": {
							"oldest_id": "1375151827189137412",
							"newest_id": "1375152598945312768",
							"result_count": 5,
							"next_token": "7140dibdnow9c7btw3w3y5b6jqjnk3k4g5zkmfkvqwa42"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id: "2244994945",
				opt: []*gotwtr.TweetOption{
					{
						TweetFields: []gotwtr.TweetField{
							gotwtr.TweetFieldAttachments,
							gotwtr.TweetFieldAuthorID,
							gotwtr.TweetFieldCreatedAt,
							gotwtr.TweetFieldEntities,
							gotwtr.TweetFieldGeo,
							gotwtr.TweetFieldID,
							gotwtr.TweetFieldInReplyToUserID,
							gotwtr.TweetFieldLanguage,
							gotwtr.TweetFieldPossiblySensitve,
							gotwtr.TweetFieldReferencedTweets,
							gotwtr.TweetFieldSource,
							gotwtr.TweetFieldText,
							gotwtr.TweetFieldWithHeld,
						},
						UserFields: []gotwtr.UserField{
							gotwtr.UserFieldUserName,
						},
						Expansions: []gotwtr.Expansion{
							gotwtr.ExpansionAuthorID,
						},
					},
				},
			},
			want: &gotwtr.UserMentionTimelineResponse{
				Tweets: []*gotwtr.Tweet{
					{
						AuthorID: "1034147061711679488",
						Text: "@LeBraat @TwitterDev @LeGuud There's enough @twitterdev love to go around, @LeBraat",
						Lang: "en",
						ConversationID: "1375152449594523649",
						ID: "1375152598945312768",
					},
					{
						AuthorID: "199566737",
						Text: "Apparently I'm not the only one (of my test accounts) that loves @TwitterDev nn@LeStaache @LeGuud",
						Lang: "en",
						ConversationID: "1375152449594523649",
						ID: "1375152449594523649",
					},
					{
						AuthorID: "930524282358325248",
						Text: "Can I join this @twitterdev love party?!",
						Lang: "en",
						ConversationID: "1375152043455873027",
						ID: "1375152043455873027",
					},
					{
						AuthorID: "1034147061711679488",
						Text: "I love me some @twitterdev too!",
						Lang: "en",
						ConversationID: "1375151947360174082",
						ID: "1375151947360174082",
					},
					{
						AuthorID: "199566737",
						Text: "This is a test, but also a good excuse to express my love for @TwitterDev 😍",
						Lang: "en",
						ConversationID: "1375151827189137412",
						ID: "1375151827189137412",
					},
				},
				Includes: &gotwtr.TweetIncludes{
					Users: []*gotwtr.User{
						{
							Name: "LeStache",
							ID: "1034147061711679488",
							Entities: []*gotwtr.UserEntity {
								{
									URL: &gotwtr.UserURL {
										URLs: []*gotwtr.UserURLs {
											{
												Start: 0,
												End: 23,
												URL: "https://t.co/7IDoW8iFLm",
												ExpandedURL: "https://twitter.com",
												DisplayURL: "twitter.com",
											},
										},
									},
									Description: &gotwtr.UserDescription {
										URLs: []*gotwtr.UserURLs {
											{
												Start: 21,
												End: 44,
												URL: "https://t.co/v6nxjDjZR3",
												ExpandedURL: "https://google.com",
												DisplayURL: "google.com",
											},
										},
										Hashtags: []*gotwtr.UserHashtag {
											{
												Start: 15,
												End: 20,
												Tag: "test",
											},
										},
										Mentions: []*gotwtr.UserMention {
											{
												Start: 0,
												End: 8,
												UserName: "lebraat",
											},
										},
										Cashtags: []*gotwtr.UserCashtag {
											{
												Start: 9,
												End: 14,
												Tag: "twtr",
											},
										},
									},
								},
							},
							UserName: "LeStaache",
							CreatedAt: "2018-08-27T18:34:07.000Z",
						},
						{
							Name: "dan dale",
							ID: "199566737",
							Entities: []*gotwtr.UserEntity {
								{
									Description: &gotwtr.UserDescription{
										Hashtags: []*gotwtr.UserHashtag {
											{
												Start: 30,
												End: 37,
												Tag: "devrel",
											},
										},
										Mentions: []*gotwtr.UserMention {
											{
												Start: 18,
												End: 29,
												UserName: "twitterdev",
											},
										},
									},
								},
							},
							UserName: "LeBraat",
							CreatedAt: "2010-10-07T05:36:28.000Z",
						},
						{
							Name: "LeGuud",
							ID: "930524282358325248",
							Entities: []*gotwtr.UserEntity {
								{
									URL: &gotwtr.UserURL{
										URLs: []*gotwtr.UserURLs {
											{
												Start: 0,
												End: 23,
												URL: "https://t.co/8IkCzClPCz",
												ExpandedURL: "https://developer.twitter.com",
												DisplayURL: "developer.twitter.com",
											},
										},
									},
								},
							},
							UserName: "LeGuud",
							CreatedAt: "2017-11-14T19:54:12.000Z",
						},
					},
				},
				Meta: &gotwtr.UserTimelineMeta{
					OldestID: "1375151827189137412",
					NewestID: "1375152598945312768",
					ResultCount: 5,
					NextToken: "7140dibdnow9c7btw3w3y5b6jqjnk3k4g5zkmfkvqwa42",
				},
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.UserMentionTimeline(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.UserMentionTimeline() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("client.UserMentionTimeline() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
