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

// TODO
// Twitter look up API's response may be changed. Need to check the response.
// The sample response now definitely exists.
// A few might have been added.
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
			name: "200 success lookup tweets, no option and single ID",
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
			name: "200 success lookup tweets, option and single ID",
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
			name: "200 success lookup tweets, no option and multiple IDs",
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
			name: "200 success lookup tweets, option and multiple IDs",
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
			name: "200 success 1 is valid, 1 is deleted",
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
			name: "404 error not found",
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
						StatusCode: http.StatusNotFound,
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

func Test_lookUpByID(t *testing.T) {
	type args struct {
		ctx    context.Context
		client *http.Client
		id     string
		opt    []*gotwtr.TweetOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.TweetLookUpByIDResponse
		wantErr bool
	}{
		{
			name: "200 success default payload",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": {
							"id": "20",
							"text": "just setting up my twttr"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id:  "20",
				opt: []*gotwtr.TweetOption{},
			},
			want: &gotwtr.TweetLookUpByIDResponse{
				Tweet: &gotwtr.Tweet{
					ID:   "20",
					Text: "just setting up my twttr",
				},
			},
			wantErr: false,
		},
		{
			name: "200 success request tweet fields",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": {
							"author_id": "2244994945",
							"created_at": "2020-06-24T16:28:14.000Z",
							"entities": {
								"urls": [
									{
										"start": 140,
										"end": 163,
										"url": "https://t.co/IKM3zo6ngu",
										"expanded_url": "https://blog.twitter.com/developer/en_us/topics/tips/2020/how-to-analyze-the-sentiment-of-your-own-tweets.html",
										"display_url": "blog.twitter.com/developer/en_u…",
										"images": [
											{
												"url": "https://pbs.twimg.com/news_img/1275828115110060033/WIbBrSld?format=jpg&name=orig",
												"width": 1600,
												"height": 600
											},
											{
												"url": "https://pbs.twimg.com/news_img/1275828115110060033/WIbBrSld?format=jpg&name=150x150",
												"width": 150,
												"height": 150
											}
										],
										"status": 200,
										"title": "How to analyze the sentiment of your own Tweets",
										"description": "This post helps developers try out sentiment analysis by analyzing their own past Tweets.",
										"unwound_url": "https://blog.twitter.com/developer/en_us/topics/tips/2020/how-to-analyze-the-sentiment-of-your-own-tweets.html"
									}
								],
								"annotations": [
									{
										"start": 59,
										"end": 73,
										"probability": 0.9028,
										"type": "Product",
										"normalized_text": "Microsoft Azure"
									},
									{
										"start": 76,
										"end": 81,
										"probability": 0.382,
										"type": "Product",
										"normalized_text": "Python"
									},
									{
										"start": 88,
										"end": 109,
										"probability": 0.3541,
										"type": "Product",
										"normalized_text": "Twitter Developer Labs"
									}
								]
							},
							"id": "1275828087666679809",
							"lang": "en",
							"possibly_sensitive": false,
							"source": "Twitter Web App",
							"text": "Learn how to create a sentiment score for your Tweets with Microsoft Azure, Python, and Twitter Developer Labs recent search functionality.\nhttps://t.co/IKM3zo6ngu"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id: "1275828087666679809",
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
					},
				},
			},
			want: &gotwtr.TweetLookUpByIDResponse{
				Tweet: &gotwtr.Tweet{
					AuthorID:  "2244994945",
					CreatedAt: "2020-06-24T16:28:14.000Z",
					Entities: &gotwtr.TweetEntity{
						URLs: []*gotwtr.TweetURL{
							{
								Start:       140,
								End:         163,
								URL:         "https://t.co/IKM3zo6ngu",
								ExpandedURL: "https://blog.twitter.com/developer/en_us/topics/tips/2020/how-to-analyze-the-sentiment-of-your-own-tweets.html",
								DisplayURL:  "blog.twitter.com/developer/en_u…",
								Images: []*gotwtr.TweetImage{
									{
										URL:    "https://pbs.twimg.com/news_img/1275828115110060033/WIbBrSld?format=jpg&name=orig",
										Width:  1600,
										Height: 600,
									},
									{
										URL:    "https://pbs.twimg.com/news_img/1275828115110060033/WIbBrSld?format=jpg&name=150x150",
										Width:  150,
										Height: 150,
									},
								},
								Status:      200,
								Title:       "How to analyze the sentiment of your own Tweets",
								Description: "This post helps developers try out sentiment analysis by analyzing their own past Tweets.",
								UnwoundURL:  "https://blog.twitter.com/developer/en_us/topics/tips/2020/how-to-analyze-the-sentiment-of-your-own-tweets.html",
							},
						},
						Annotations: []*gotwtr.TweetAnnotation{
							{
								Start:          59,
								End:            73,
								Probability:    0.9028,
								Type:           "Product",
								NormalizedText: "Microsoft Azure",
							},
							{
								Start:          76,
								End:            81,
								Probability:    0.382,
								Type:           "Product",
								NormalizedText: "Python",
							},
							{
								Start:          88,
								End:            109,
								Probability:    0.3541,
								Type:           "Product",
								NormalizedText: "Twitter Developer Labs",
							},
						},
					},
					ID:                "1275828087666679809",
					Lang:              "en",
					PossiblySensitive: false,
					Source:            "Twitter Web App",
					Text:              "Learn how to create a sentiment score for your Tweets with Microsoft Azure, Python, and Twitter Developer Labs recent search functionality.\nhttps://t.co/IKM3zo6ngu",
				},
			},
			wantErr: false,
		},
		{
			name: "200 success deleted tweet",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"errors": [
							{
								"value": "1276230436478386177",
								"detail": "Could not find tweet with id: [1276230436478386177].",
								"title": "Not Found Error",
								"resource_type": "tweet",
								"parameter": "id",
								"resource_id": "1276230436478386177",
								"type": "https://api.twitter.com/2/problems/resource-not-found"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id: "1276230436478386177",
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
					},
				},
			},
			want: &gotwtr.TweetLookUpByIDResponse{
				Errors: []*gotwtr.APIResponseError{
					{
						Value:        "1276230436478386177",
						Detail:       "Could not find tweet with id: [1276230436478386177].",
						Title:        "Not Found Error",
						ResourceType: "tweet",
						Parameter:    "id",
						ResourceID:   "1276230436478386177",
						Type:         "https://api.twitter.com/2/problems/resource-not-found",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "200 success request place fields",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": {
							"geo": {
								"place_id": "01a9a39529b27f36"
							},
							"id": "1136017751028449283",
							"text": "We’re getting ready to #TapIntoTwitter with our NYC developer community! See you soon @TwitterNYC https://t.co/5rEn5dhsAq"
						},
						"includes": {
							"places": [
								{
									"geo": {
										"type": "Feature",
										"bbox": [
											-74.026675,
											40.683935,
											-73.910408,
											40.877483
										],
										"properties": {}
									},
									"id": "01a9a39529b27f36",
									"country_code": "US",
									"full_name": "Manhattan, NY",
									"name": "Manhattan",
									"place_type": "city",
									"country": "United States"
								}
							]
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id: "1136017751028449283",
				opt: []*gotwtr.TweetOption{
					{
						Expansions: []gotwtr.Expansion{
							gotwtr.ExpansionGeoPlaceID,
						},
						PlaceFields: []gotwtr.PlaceField{
							gotwtr.PlaceFieldContainedWithin,
							gotwtr.PlaceFieldCountry,
							gotwtr.PlaceFieldCountryCode,
							gotwtr.PlaceFieldFullName,
							gotwtr.PlaceFieldGeo,
							gotwtr.PlaceFieldID,
							gotwtr.PlaceFieldName,
							gotwtr.PlaceFieldPlaceType,
						},
					},
				},
			},
			want: &gotwtr.TweetLookUpByIDResponse{
				Tweet: &gotwtr.Tweet{
					Geo: &gotwtr.TweetGeo{
						PlaceID: "01a9a39529b27f36",
					},
					ID:   "1136017751028449283",
					Text: "We’re getting ready to #TapIntoTwitter with our NYC developer community! See you soon @TwitterNYC https://t.co/5rEn5dhsAq",
				},
				Includes: &gotwtr.TweetIncludes{
					Places: []*gotwtr.Place{
						{
							Geo: &gotwtr.PlaceGeo{
								Type: "Feature",
								BBox: []float64{
									-74.026675,
									40.683935,
									-73.910408,
									40.877483,
								},
								Properties: map[string]interface{}{},
							},
							ID:          "01a9a39529b27f36",
							CountryCode: "US",
							FullName:    "Manhattan, NY",
							Name:        "Manhattan",
							PlaceType:   "city",
							Country:     "United States",
						},
					},
				},
			},
		},
		{
			name: "200 success request poll fields",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": {
							"attachments": {
								"poll_ids": [
									"1199786642468413448"
								]
							},
							"id": "1199786642791452673",
							"text": "C#"
						},
						"includes": {
							"polls": [
								{
									"end_datetime": "2019-11-28T20:26:41.000Z",
									"options": [
										{
											"position": 1,
											"label": "“C Sharp”",
											"votes": 795
										},
										{
											"position": 2,
											"label": "“C Hashtag”",
											"votes": 156
										}
									],
									"duration_minutes": 1440,
									"id": "1199786642468413448",
									"voting_status": "closed"
								}
							]
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id: "1199786642791452673",
				opt: []*gotwtr.TweetOption{
					{
						Expansions: []gotwtr.Expansion{
							gotwtr.ExpansionAttachmentsPollIDs,
						},
						PollFields: []gotwtr.PollField{
							gotwtr.PollFieldDurationMinutes,
							gotwtr.PollFieldEndDateTime,
							gotwtr.PollFieldID,
							gotwtr.PollFieldOptions,
							gotwtr.PollFieldVotingStatus,
						},
					},
				},
			},
			want: &gotwtr.TweetLookUpByIDResponse{
				Tweet: &gotwtr.Tweet{
					Attachments: &gotwtr.TweetAttachment{
						PollIDs: []string{
							"1199786642468413448",
						},
					},
					ID:   "1199786642791452673",
					Text: "C#",
				},
				Includes: &gotwtr.TweetIncludes{
					Polls: []*gotwtr.Poll{
						{
							EndDatetime: "2019-11-28T20:26:41.000Z",
							Options: []*gotwtr.PollOption{
								{
									Position: 1,
									Label:    "“C Sharp”",
									Votes:    795,
								},
								{
									Position: 2,
									Label:    "“C Hashtag”",
									Votes:    156,
								},
							},
							DurationMinutes: 1440,
							ID:              "1199786642468413448",
							VotingStatus:    "closed",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "200 success request media fields",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": {
							"attachments": {
								"media_keys": [
									"13_1263145212760805376"
								]
							},
							"id": "1263145271946551300",
							"created_at": "2020-05-20T16:31:15.000Z",
							"lang": "en",
							"entities": {
								"urls": [
									{
										"start": 154,
										"end": 177,
										"url": "https://t.co/pV53mvjAVT",
										"expanded_url": "https://twitter.com/Twitter/status/1263145271946551300/video/1",
										"display_url": "pic.twitter.com/pV53mvjAVT"
									}
								]
							},
							"source": "Sprinklr",
							"possibly_sensitive": false,
							"author_id": "783214",
							"text": "Testing, testing...\n\nA new way to have a convo with exactly who you want. We’re starting with a small % globally, so keep your 👀 out to see it in action. https://t.co/pV53mvjAVT"
						},
						"includes": {
							"media": [
								{
									"media_key": "13_1263145212760805376",
									"width": 1920,
									"preview_image_url": "https://pbs.twimg.com/media/EYeX7akWsAIP1_1.jpg",
									"public_metrics": {
										"view_count": 7578411
									},
									"duration_ms": 46947,
									"type": "video",
									"height": 1080
								}
							]
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id: "1263145271946551300",
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
						Expansions: []gotwtr.Expansion{
							gotwtr.ExpansionAttachmentsMediaKeys,
						},
						MediaFields: []gotwtr.MediaField{
							gotwtr.MediaFieldDurationMS,
							gotwtr.MediaFieldHeight,
							gotwtr.MediaFieldMediaKey,
							gotwtr.MediaFieldPreviewImageURL,
							gotwtr.MediaFieldPublicMetrics,
							gotwtr.MediaFieldType,
							gotwtr.MediaFieldURL,
							gotwtr.MediaFieldWidth,
						},
					},
				},
			},
			want: &gotwtr.TweetLookUpByIDResponse{
				Tweet: &gotwtr.Tweet{
					Attachments: &gotwtr.TweetAttachment{
						MediaKeys: []string{
							"13_1263145212760805376",
						},
					},
					ID:        "1263145271946551300",
					CreatedAt: "2020-05-20T16:31:15.000Z",
					Lang:      "en",
					Entities: &gotwtr.TweetEntity{
						URLs: []*gotwtr.TweetURL{
							{
								Start:       154,
								End:         177,
								URL:         "https://t.co/pV53mvjAVT",
								ExpandedURL: "https://twitter.com/Twitter/status/1263145271946551300/video/1",
								DisplayURL:  "pic.twitter.com/pV53mvjAVT",
							},
						},
					},
					Source:            "Sprinklr",
					PossiblySensitive: false,
					AuthorID:          "783214",
					Text:              "Testing, testing...\n\nA new way to have a convo with exactly who you want. We’re starting with a small % globally, so keep your 👀 out to see it in action. https://t.co/pV53mvjAVT",
				},
				Includes: &gotwtr.TweetIncludes{
					Media: []*gotwtr.Media{
						{
							MediaKey:        "13_1263145212760805376",
							Width:           1920,
							PreviewImageURL: "https://pbs.twimg.com/media/EYeX7akWsAIP1_1.jpg",
							PublicMetrics: &gotwtr.MediaMetrics{
								ViewCount: 7578411,
							},
							DurationMs: 46947,
							Type:       "video",
							Height:     1080,
						},
					},
				},
			},
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.LookUpTweetByID(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.LookUpTweetByID() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("client.LookUpTweetByID() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
