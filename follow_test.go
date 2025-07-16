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

func Test_followers(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		userID string
		opt    []*gotwtr.FollowOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.FollowersResponse
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
								"id": "6253282",
								"name": "Twitter API",
								"username": "TwitterAPI"
							},
							{
								"id": "2244994945",
								"name": "Twitter Dev",
								"username": "TwitterDev"
							},
							{
								"id": "783214",
								"name": "Twitter",
								"username": "Twitter"
							},
							{
								"id": "95731075",
								"name": "Twitter Safety",
								"username": "TwitterSafety"
							},
							{
								"id": "3260518932",
								"name": "Twitter Moments",
								"username": "TwitterMoments"
							},
							{
								"id": "373471064",
								"name": "Twitter Music",
								"username": "TwitterMusic"
							},
							{
								"id": "791978718",
								"name": "Twitter Official Partner",
								"username": "OfficialPartner"
							},
							{
								"id": "17874544",
								"name": "Twitter Support",
								"username": "TwitterSupport"
							},
							{
								"id": "234489024",
								"name": "Twitter Comms",
								"username": "TwitterComms"
							},
							{
								"id": "1526228120",
								"name": "Twitter Data",
								"username": "TwitterData"
							}
						],
						"meta": {
							"result_count": 10,
							"next_token": "DFEDBNRFT3MHCZZZ"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID: "2244994945",
				opt:    []*gotwtr.FollowOption{},
			},
			want: &gotwtr.FollowersResponse{
				Users: []*gotwtr.User{
					{
						ID:       "6253282",
						Name:     "Twitter API",
						UserName: "TwitterAPI",
					},
					{
						ID:       "2244994945",
						Name:     "Twitter Dev",
						UserName: "TwitterDev",
					},
					{
						ID:       "783214",
						Name:     "Twitter",
						UserName: "Twitter",
					},
					{
						ID:       "95731075",
						Name:     "Twitter Safety",
						UserName: "TwitterSafety",
					},
					{
						ID:       "3260518932",
						Name:     "Twitter Moments",
						UserName: "TwitterMoments",
					},
					{
						ID:       "373471064",
						Name:     "Twitter Music",
						UserName: "TwitterMusic",
					},
					{
						ID:       "791978718",
						Name:     "Twitter Official Partner",
						UserName: "OfficialPartner",
					},
					{
						ID:       "17874544",
						Name:     "Twitter Support",
						UserName: "TwitterSupport",
					},
					{
						ID:       "234489024",
						Name:     "Twitter Comms",
						UserName: "TwitterComms",
					},
					{
						ID:       "1526228120",
						Name:     "Twitter Data",
						UserName: "TwitterData",
					},
				},
				Meta: &gotwtr.FollowsMeta{
					ResultCount: 10,
					NextToken:   "DFEDBNRFT3MHCZZZ",
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
								"pinned_tweet_id": "1293595870563381249",
								"id": "6253282",
								"username": "TwitterAPI",
								"name": "Twitter API"
							},
							{
								"pinned_tweet_id": "1293593516040269825",
								"id": "2244994945",
								"username": "TwitterDev",
								"name": "Twitter Dev"
							},
							{
								"id": "783214",
								"username": "Twitter",
								"name": "Twitter"
							},
							{
								"pinned_tweet_id": "1271186240323432452",
								"id": "95731075",
								"username": "TwitterSafety",
								"name": "Twitter Safety"
							},
							{
								"id": "3260518932",
								"username": "TwitterMoments",
								"name": "Twitter Moments"
							},
							{
								"pinned_tweet_id": "1293216056274759680",
								"id": "373471064",
								"username": "TwitterMusic",
								"name": "Twitter Music"
							},
							{
								"id": "791978718",
								"username": "OfficialPartner",
								"name": "Twitter Official Partner"
							},
							{
								"pinned_tweet_id": "1289000334497439744",
								"id": "17874544",
								"username": "TwitterSupport",
								"name": "Twitter Support"
							},
							{
								"pinned_tweet_id": "1283543147444711424",
								"id": "234489024",
								"username": "TwitterComms",
								"name": "Twitter Comms"
							},
							{
								"id": "1526228120",
								"username": "TwitterData",
								"name": "Twitter Data"
							}
						],
						"includes": {
							"tweets": [
								{
									"context_annotations": [
										{
											"domain": {
												"id": "46",
												"name": "Brand Category",
												"description": "Categories within Brand Verticals that narrow down the scope of Brands"
											},
											"entity": {
												"id": "781974596752842752",
												"name": "Services"
											}
										},
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
										},
										{
											"domain": {
												"id": "65",
												"name": "Interests and Hobbies Vertical",
												"description": "Top level interests and hobbies groupings, like Food or Travel"
											},
											"entity": {
												"id": "848920371311001600",
												"name": "Technology",
												"description": "Technology and computing"
											}
										},
										{
											"domain": {
												"id": "66",
												"name": "Interests and Hobbies Category",
												"description": "A grouping of interests and hobbies entities, like Novelty Food or Destinations"
											},
											"entity": {
												"id": "848921413196984320",
												"name": "Computer programming",
												"description": "Computer programming"
											}
										},
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
									"id": "1293595870563381249",
									"text": "Twitter API v2: Early Access releasednnToday we announced Early Access to the first endpoints of the new Twitter API!nn#TwitterAPI #EarlyAccess #VersionBump https://t.co/g7v3aeIbtQ"
								},
								{
									"context_annotations": [
										{
											"domain": {
												"id": "46",
												"name": "Brand Category",
												"description": "Categories within Brand Verticals that narrow down the scope of Brands"
											},
											"entity": {
												"id": "781974596752842752",
												"name": "Services"
											}
										},
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
										},
										{
											"domain": {
												"id": "65",
												"name": "Interests and Hobbies Vertical",
												"description": "Top level interests and hobbies groupings, like Food or Travel"
											},
											"entity": {
												"id": "848920371311001600",
												"name": "Technology",
												"description": "Technology and computing"
											}
										},
										{
											"domain": {
												"id": "66",
												"name": "Interests and Hobbies Category",
												"description": "A grouping of interests and hobbies entities, like Novelty Food or Destinations"
											},
											"entity": {
												"id": "848921413196984320",
												"name": "Computer programming",
												"description": "Computer programming"
											}
										}
									],
									"id": "1293593516040269825",
									"text": "It’s finally here! 🥁 Say hello to the new #TwitterAPI.nnWe’re rebuilding the Twitter API v2 from the ground up to better serve our developer community. And today’s launch is only the beginning.nnhttps://t.co/32VrwpGaJw https://t.co/KaFSbjWUA8"
								},
								{
									"id": "1271186240323432452",
									"text": "We’re disclosing new state-linked information operations to our public archive — the only one of its kind in the industry. Originating from the People’s Republic of China (PRC), Russia, and Turkey, all associated accounts and content have been removed. https://t.co/obRqr96iYm"
								},
								{
									"id": "1293216056274759680",
									"text": "say howdy to your new yeehaw king @orvillepeck—our #ArtistToFollow this month 🤠 https://t.co/3pk9fYcPHb"
								},
								{
									"context_annotations": [
										{
											"domain": {
												"id": "46",
												"name": "Brand Category",
												"description": "Categories within Brand Verticals that narrow down the scope of Brands"
											},
											"entity": {
												"id": "781974596752842752",
												"name": "Services"
											}
										},
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
									"id": "1289000334497439744",
									"text": "We’ve significantly limited access to our internal tools and systems. Until we can safely resume normal operations, our response times to some support needs and reports will be slower. Thank you for your patience as we work through this."
								},
								{
									"context_annotations": [
										{
											"domain": {
												"id": "46",
												"name": "Brand Category",
												"description": "Categories within Brand Verticals that narrow down the scope of Brands"
											},
											"entity": {
												"id": "781974596752842752",
												"name": "Services"
											}
										},
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
									"id": "1283543147444711424",
									"text": "Follow @TwitterSupport for the latest on the security incident ⬇️ https://t.co/7FKKksJqxV"
								}
							],
							"meta": {
								"result_count": 10,
								"next_token": "DFEDBNRFT3MHCZZZ"
							}
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID: "2244994945",
				opt: []*gotwtr.FollowOption{
					{
						Expansions: []gotwtr.Expansion{gotwtr.ExpansionPinnedTweetID, gotwtr.ExpansionContextAnnotations},
						TweetFields: []gotwtr.TweetField{
							gotwtr.TweetFieldCreatedAt,
							gotwtr.TweetFieldMaxResults,
						},
					},
				},
			},
			want: &gotwtr.FollowersResponse{
				Users: []*gotwtr.User{
					{
						PinnedTweetID: "1293595870563381249",
						ID:            "6253282",
						UserName:      "TwitterAPI",
						Name:          "Twitter API",
					},
					{
						PinnedTweetID: "1293593516040269825",
						ID:            "2244994945",
						UserName:      "TwitterDev",
						Name:          "Twitter Dev",
					},
					{
						ID:       "783214",
						UserName: "Twitter",
						Name:     "Twitter",
					},
					{
						PinnedTweetID: "1271186240323432452",
						ID:            "95731075",
						UserName:      "TwitterSafety",
						Name:          "Twitter Safety",
					},
					{
						ID:       "3260518932",
						UserName: "TwitterMoments",
						Name:     "Twitter Moments",
					},
					{
						PinnedTweetID: "1293216056274759680",
						ID:            "373471064",
						UserName:      "TwitterMusic",
						Name:          "Twitter Music",
					},
					{
						ID:       "791978718",
						UserName: "OfficialPartner",
						Name:     "Twitter Official Partner",
					},
					{
						PinnedTweetID: "1289000334497439744",
						ID:            "17874544",
						UserName:      "TwitterSupport",
						Name:          "Twitter Support",
					},
					{
						PinnedTweetID: "1283543147444711424",
						ID:            "234489024",
						UserName:      "TwitterComms",
						Name:          "Twitter Comms",
					},
					{
						ID:       "1526228120",
						UserName: "TwitterData",
						Name:     "Twitter Data",
					},
				},
				Includes: &gotwtr.UserIncludes{
					Tweets: []*gotwtr.Tweet{
						{
							ContextAnnotations: []*gotwtr.TweetContextAnnotation{
								{
									Domain: &gotwtr.TweetContextObj{
										ID:          "46",
										Name:        "Brand Category",
										Description: "Categories within Brand Verticals that narrow down the scope of Brands",
									},
									Entity: &gotwtr.TweetContextObj{
										ID:   "781974596752842752",
										Name: "Services",
									},
								},
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
								{
									Domain: &gotwtr.TweetContextObj{
										ID:          "65",
										Name:        "Interests and Hobbies Vertical",
										Description: "Top level interests and hobbies groupings, like Food or Travel",
									},
									Entity: &gotwtr.TweetContextObj{
										ID:          "848920371311001600",
										Name:        "Technology",
										Description: "Technology and computing",
									},
								},
								{
									Domain: &gotwtr.TweetContextObj{
										ID:          "66",
										Name:        "Interests and Hobbies Category",
										Description: "A grouping of interests and hobbies entities, like Novelty Food or Destinations",
									},
									Entity: &gotwtr.TweetContextObj{
										ID:          "848921413196984320",
										Name:        "Computer programming",
										Description: "Computer programming",
									},
								},
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
							ID:   "1293595870563381249",
							Text: "Twitter API v2: Early Access releasednnToday we announced Early Access to the first endpoints of the new Twitter API!nn#TwitterAPI #EarlyAccess #VersionBump https://t.co/g7v3aeIbtQ",
						},
						{
							ContextAnnotations: []*gotwtr.TweetContextAnnotation{
								{
									Domain: &gotwtr.TweetContextObj{
										ID:          "46",
										Name:        "Brand Category",
										Description: "Categories within Brand Verticals that narrow down the scope of Brands",
									},
									Entity: &gotwtr.TweetContextObj{
										ID:   "781974596752842752",
										Name: "Services",
									},
								},
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
								{
									Domain: &gotwtr.TweetContextObj{
										ID:          "65",
										Name:        "Interests and Hobbies Vertical",
										Description: "Top level interests and hobbies groupings, like Food or Travel",
									},
									Entity: &gotwtr.TweetContextObj{
										ID:          "848920371311001600",
										Name:        "Technology",
										Description: "Technology and computing",
									},
								},
								{
									Domain: &gotwtr.TweetContextObj{
										ID:          "66",
										Name:        "Interests and Hobbies Category",
										Description: "A grouping of interests and hobbies entities, like Novelty Food or Destinations",
									},
									Entity: &gotwtr.TweetContextObj{
										ID:          "848921413196984320",
										Name:        "Computer programming",
										Description: "Computer programming",
									},
								},
							},
							ID:   "1293593516040269825",
							Text: "It’s finally here! 🥁 Say hello to the new #TwitterAPI.nnWe’re rebuilding the Twitter API v2 from the ground up to better serve our developer community. And today’s launch is only the beginning.nnhttps://t.co/32VrwpGaJw https://t.co/KaFSbjWUA8",
						},
						{
							ID:   "1271186240323432452",
							Text: "We’re disclosing new state-linked information operations to our public archive — the only one of its kind in the industry. Originating from the People’s Republic of China (PRC), Russia, and Turkey, all associated accounts and content have been removed. https://t.co/obRqr96iYm",
						},
						{
							ID:   "1293216056274759680",
							Text: "say howdy to your new yeehaw king @orvillepeck—our #ArtistToFollow this month 🤠 https://t.co/3pk9fYcPHb",
						},
						{
							ContextAnnotations: []*gotwtr.TweetContextAnnotation{
								{
									Domain: &gotwtr.TweetContextObj{
										ID:          "46",
										Name:        "Brand Category",
										Description: "Categories within Brand Verticals that narrow down the scope of Brands",
									},
									Entity: &gotwtr.TweetContextObj{
										ID:   "781974596752842752",
										Name: "Services",
									},
								},
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
							ID:   "1289000334497439744",
							Text: "We’ve significantly limited access to our internal tools and systems. Until we can safely resume normal operations, our response times to some support needs and reports will be slower. Thank you for your patience as we work through this.",
						},
						{
							ContextAnnotations: []*gotwtr.TweetContextAnnotation{
								{
									Domain: &gotwtr.TweetContextObj{
										ID:          "46",
										Name:        "Brand Category",
										Description: "Categories within Brand Verticals that narrow down the scope of Brands",
									},
									Entity: &gotwtr.TweetContextObj{
										ID:   "781974596752842752",
										Name: "Services",
									},
								},
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
							ID:   "1283543147444711424",
							Text: "Follow @TwitterSupport for the latest on the security incident ⬇️ https://t.co/7FKKksJqxV",
						},
					},
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
				opt:    []*gotwtr.FollowOption{},
			},
			want: &gotwtr.FollowersResponse{
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
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.Followers(tt.args.ctx, tt.args.userID, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Followers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.Followers() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func Test_following(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		userID string
		opt    []*gotwtr.FollowOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.FollowingResponse
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
								"id": "6253282",
								"name": "Twitter API",
								"username": "TwitterAPI"
							},
							{
								"id": "2244994945",
								"name": "Twitter Dev",
								"username": "TwitterDev"
							},
							{
								"id": "783214",
								"name": "Twitter",
								"username": "Twitter"
							},
							{
								"id": "95731075",
								"name": "Twitter Safety",
								"username": "TwitterSafety"
							},
							{
								"id": "3260518932",
								"name": "Twitter Moments",
								"username": "TwitterMoments"
							},
							{
								"id": "373471064",
								"name": "Twitter Music",
								"username": "TwitterMusic"
							},
							{
								"id": "791978718",
								"name": "Twitter Official Partner",
								"username": "OfficialPartner"
							},
							{
								"id": "17874544",
								"name": "Twitter Support",
								"username": "TwitterSupport"
							},
							{
								"id": "234489024",
								"name": "Twitter Comms",
								"username": "TwitterComms"
							},
							{
								"id": "1526228120",
								"name": "Twitter Data",
								"username": "TwitterData"
							}
						],
						"meta": {
							"result_count": 10,
							"next_token": "DFEDBNRFT3MHCZZZ"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID: "2244994945",
				opt:    []*gotwtr.FollowOption{},
			},
			want: &gotwtr.FollowingResponse{
				Users: []*gotwtr.User{
					{
						ID:       "6253282",
						Name:     "Twitter API",
						UserName: "TwitterAPI",
					},
					{
						ID:       "2244994945",
						Name:     "Twitter Dev",
						UserName: "TwitterDev",
					},
					{
						ID:       "783214",
						Name:     "Twitter",
						UserName: "Twitter",
					},
					{
						ID:       "95731075",
						Name:     "Twitter Safety",
						UserName: "TwitterSafety",
					},
					{
						ID:       "3260518932",
						Name:     "Twitter Moments",
						UserName: "TwitterMoments",
					},
					{
						ID:       "373471064",
						Name:     "Twitter Music",
						UserName: "TwitterMusic",
					},
					{
						ID:       "791978718",
						Name:     "Twitter Official Partner",
						UserName: "OfficialPartner",
					},
					{
						ID:       "17874544",
						Name:     "Twitter Support",
						UserName: "TwitterSupport",
					},
					{
						ID:       "234489024",
						Name:     "Twitter Comms",
						UserName: "TwitterComms",
					},
					{
						ID:       "1526228120",
						Name:     "Twitter Data",
						UserName: "TwitterData",
					},
				},
				Meta: &gotwtr.FollowsMeta{
					ResultCount: 10,
					NextToken:   "DFEDBNRFT3MHCZZZ",
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
								"pinned_tweet_id": "1293595870563381249",
								"id": "6253282",
								"username": "TwitterAPI",
								"name": "Twitter API"
							},
							{
								"pinned_tweet_id": "1293593516040269825",
								"id": "2244994945",
								"username": "TwitterDev",
								"name": "Twitter Dev"
							},
							{
								"id": "783214",
								"username": "Twitter",
								"name": "Twitter"
							},
							{
								"pinned_tweet_id": "1271186240323432452",
								"id": "95731075",
								"username": "TwitterSafety",
								"name": "Twitter Safety"
							},
							{
								"id": "3260518932",
								"username": "TwitterMoments",
								"name": "Twitter Moments"
							},
							{
								"pinned_tweet_id": "1293216056274759680",
								"id": "373471064",
								"username": "TwitterMusic",
								"name": "Twitter Music"
							},
							{
								"id": "791978718",
								"username": "OfficialPartner",
								"name": "Twitter Official Partner"
							},
							{
								"pinned_tweet_id": "1289000334497439744",
								"id": "17874544",
								"username": "TwitterSupport",
								"name": "Twitter Support"
							},
							{
								"pinned_tweet_id": "1283543147444711424",
								"id": "234489024",
								"username": "TwitterComms",
								"name": "Twitter Comms"
							},
							{
								"id": "1526228120",
								"username": "TwitterData",
								"name": "Twitter Data"
							}
						],
						"includes": {
							"tweets": [
								{
									"context_annotations": [
										{
											"domain": {
												"id": "46",
												"name": "Brand Category",
												"description": "Categories within Brand Verticals that narrow down the scope of Brands"
											},
											"entity": {
												"id": "781974596752842752",
												"name": "Services"
											}
										},
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
										},
										{
											"domain": {
												"id": "65",
												"name": "Interests and Hobbies Vertical",
												"description": "Top level interests and hobbies groupings, like Food or Travel"
											},
											"entity": {
												"id": "848920371311001600",
												"name": "Technology",
												"description": "Technology and computing"
											}
										},
										{
											"domain": {
												"id": "66",
												"name": "Interests and Hobbies Category",
												"description": "A grouping of interests and hobbies entities, like Novelty Food or Destinations"
											},
											"entity": {
												"id": "848921413196984320",
												"name": "Computer programming",
												"description": "Computer programming"
											}
										},
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
									"id": "1293595870563381249",
									"text": "Twitter API v2: Early Access releasednnToday we announced Early Access to the first endpoints of the new Twitter API!nn#TwitterAPI #EarlyAccess #VersionBump https://t.co/g7v3aeIbtQ"
								},
								{
									"context_annotations": [
										{
											"domain": {
												"id": "46",
												"name": "Brand Category",
												"description": "Categories within Brand Verticals that narrow down the scope of Brands"
											},
											"entity": {
												"id": "781974596752842752",
												"name": "Services"
											}
										},
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
										},
										{
											"domain": {
												"id": "65",
												"name": "Interests and Hobbies Vertical",
												"description": "Top level interests and hobbies groupings, like Food or Travel"
											},
											"entity": {
												"id": "848920371311001600",
												"name": "Technology",
												"description": "Technology and computing"
											}
										},
										{
											"domain": {
												"id": "66",
												"name": "Interests and Hobbies Category",
												"description": "A grouping of interests and hobbies entities, like Novelty Food or Destinations"
											},
											"entity": {
												"id": "848921413196984320",
												"name": "Computer programming",
												"description": "Computer programming"
											}
										}
									],
									"id": "1293593516040269825",
									"text": "It’s finally here! 🥁 Say hello to the new #TwitterAPI.nnWe’re rebuilding the Twitter API v2 from the ground up to better serve our developer community. And today’s launch is only the beginning.nnhttps://t.co/32VrwpGaJw https://t.co/KaFSbjWUA8"
								},
								{
									"id": "1271186240323432452",
									"text": "We’re disclosing new state-linked information operations to our public archive — the only one of its kind in the industry. Originating from the People’s Republic of China (PRC), Russia, and Turkey, all associated accounts and content have been removed. https://t.co/obRqr96iYm"
								},
								{
									"id": "1293216056274759680",
									"text": "say howdy to your new yeehaw king @orvillepeck—our #ArtistToFollow this month 🤠 https://t.co/3pk9fYcPHb"
								},
								{
									"context_annotations": [
										{
											"domain": {
												"id": "46",
												"name": "Brand Category",
												"description": "Categories within Brand Verticals that narrow down the scope of Brands"
											},
											"entity": {
												"id": "781974596752842752",
												"name": "Services"
											}
										},
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
									"id": "1289000334497439744",
									"text": "We’ve significantly limited access to our internal tools and systems. Until we can safely resume normal operations, our response times to some support needs and reports will be slower. Thank you for your patience as we work through this."
								},
								{
									"context_annotations": [
										{
											"domain": {
												"id": "46",
												"name": "Brand Category",
												"description": "Categories within Brand Verticals that narrow down the scope of Brands"
											},
											"entity": {
												"id": "781974596752842752",
												"name": "Services"
											}
										},
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
									"id": "1283543147444711424",
									"text": "Follow @TwitterSupport for the latest on the security incident ⬇️ https://t.co/7FKKksJqxV"
								}
							],
							"meta": {
								"result_count": 10,
								"next_token": "DFEDBNRFT3MHCZZZ"
							}
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID: "2244994945",
				opt: []*gotwtr.FollowOption{
					{
						Expansions: []gotwtr.Expansion{gotwtr.ExpansionPinnedTweetID, gotwtr.ExpansionContextAnnotations},
						TweetFields: []gotwtr.TweetField{
							gotwtr.TweetFieldCreatedAt,
							gotwtr.TweetFieldMaxResults,
						},
					},
				},
			},
			want: &gotwtr.FollowingResponse{
				Users: []*gotwtr.User{
					{
						PinnedTweetID: "1293595870563381249",
						ID:            "6253282",
						UserName:      "TwitterAPI",
						Name:          "Twitter API",
					},
					{
						PinnedTweetID: "1293593516040269825",
						ID:            "2244994945",
						UserName:      "TwitterDev",
						Name:          "Twitter Dev",
					},
					{
						ID:       "783214",
						UserName: "Twitter",
						Name:     "Twitter",
					},
					{
						PinnedTweetID: "1271186240323432452",
						ID:            "95731075",
						UserName:      "TwitterSafety",
						Name:          "Twitter Safety",
					},
					{
						ID:       "3260518932",
						UserName: "TwitterMoments",
						Name:     "Twitter Moments",
					},
					{
						PinnedTweetID: "1293216056274759680",
						ID:            "373471064",
						UserName:      "TwitterMusic",
						Name:          "Twitter Music",
					},
					{
						ID:       "791978718",
						UserName: "OfficialPartner",
						Name:     "Twitter Official Partner",
					},
					{
						PinnedTweetID: "1289000334497439744",
						ID:            "17874544",
						UserName:      "TwitterSupport",
						Name:          "Twitter Support",
					},
					{
						PinnedTweetID: "1283543147444711424",
						ID:            "234489024",
						UserName:      "TwitterComms",
						Name:          "Twitter Comms",
					},
					{
						ID:       "1526228120",
						UserName: "TwitterData",
						Name:     "Twitter Data",
					},
				},
				Includes: &gotwtr.UserIncludes{
					Tweets: []*gotwtr.Tweet{
						{
							ContextAnnotations: []*gotwtr.TweetContextAnnotation{
								{
									Domain: &gotwtr.TweetContextObj{
										ID:          "46",
										Name:        "Brand Category",
										Description: "Categories within Brand Verticals that narrow down the scope of Brands",
									},
									Entity: &gotwtr.TweetContextObj{
										ID:   "781974596752842752",
										Name: "Services",
									},
								},
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
								{
									Domain: &gotwtr.TweetContextObj{
										ID:          "65",
										Name:        "Interests and Hobbies Vertical",
										Description: "Top level interests and hobbies groupings, like Food or Travel",
									},
									Entity: &gotwtr.TweetContextObj{
										ID:          "848920371311001600",
										Name:        "Technology",
										Description: "Technology and computing",
									},
								},
								{
									Domain: &gotwtr.TweetContextObj{
										ID:          "66",
										Name:        "Interests and Hobbies Category",
										Description: "A grouping of interests and hobbies entities, like Novelty Food or Destinations",
									},
									Entity: &gotwtr.TweetContextObj{
										ID:          "848921413196984320",
										Name:        "Computer programming",
										Description: "Computer programming",
									},
								},
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
							ID:   "1293595870563381249",
							Text: "Twitter API v2: Early Access releasednnToday we announced Early Access to the first endpoints of the new Twitter API!nn#TwitterAPI #EarlyAccess #VersionBump https://t.co/g7v3aeIbtQ",
						},
						{
							ContextAnnotations: []*gotwtr.TweetContextAnnotation{
								{
									Domain: &gotwtr.TweetContextObj{
										ID:          "46",
										Name:        "Brand Category",
										Description: "Categories within Brand Verticals that narrow down the scope of Brands",
									},
									Entity: &gotwtr.TweetContextObj{
										ID:   "781974596752842752",
										Name: "Services",
									},
								},
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
								{
									Domain: &gotwtr.TweetContextObj{
										ID:          "65",
										Name:        "Interests and Hobbies Vertical",
										Description: "Top level interests and hobbies groupings, like Food or Travel",
									},
									Entity: &gotwtr.TweetContextObj{
										ID:          "848920371311001600",
										Name:        "Technology",
										Description: "Technology and computing",
									},
								},
								{
									Domain: &gotwtr.TweetContextObj{
										ID:          "66",
										Name:        "Interests and Hobbies Category",
										Description: "A grouping of interests and hobbies entities, like Novelty Food or Destinations",
									},
									Entity: &gotwtr.TweetContextObj{
										ID:          "848921413196984320",
										Name:        "Computer programming",
										Description: "Computer programming",
									},
								},
							},
							ID:   "1293593516040269825",
							Text: "It’s finally here! 🥁 Say hello to the new #TwitterAPI.nnWe’re rebuilding the Twitter API v2 from the ground up to better serve our developer community. And today’s launch is only the beginning.nnhttps://t.co/32VrwpGaJw https://t.co/KaFSbjWUA8",
						},
						{
							ID:   "1271186240323432452",
							Text: "We’re disclosing new state-linked information operations to our public archive — the only one of its kind in the industry. Originating from the People’s Republic of China (PRC), Russia, and Turkey, all associated accounts and content have been removed. https://t.co/obRqr96iYm",
						},
						{
							ID:   "1293216056274759680",
							Text: "say howdy to your new yeehaw king @orvillepeck—our #ArtistToFollow this month 🤠 https://t.co/3pk9fYcPHb",
						},
						{
							ContextAnnotations: []*gotwtr.TweetContextAnnotation{
								{
									Domain: &gotwtr.TweetContextObj{
										ID:          "46",
										Name:        "Brand Category",
										Description: "Categories within Brand Verticals that narrow down the scope of Brands",
									},
									Entity: &gotwtr.TweetContextObj{
										ID:   "781974596752842752",
										Name: "Services",
									},
								},
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
							ID:   "1289000334497439744",
							Text: "We’ve significantly limited access to our internal tools and systems. Until we can safely resume normal operations, our response times to some support needs and reports will be slower. Thank you for your patience as we work through this.",
						},
						{
							ContextAnnotations: []*gotwtr.TweetContextAnnotation{
								{
									Domain: &gotwtr.TweetContextObj{
										ID:          "46",
										Name:        "Brand Category",
										Description: "Categories within Brand Verticals that narrow down the scope of Brands",
									},
									Entity: &gotwtr.TweetContextObj{
										ID:   "781974596752842752",
										Name: "Services",
									},
								},
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
							ID:   "1283543147444711424",
							Text: "Follow @TwitterSupport for the latest on the security incident ⬇️ https://t.co/7FKKksJqxV",
						},
					},
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
				opt:    []*gotwtr.FollowOption{},
			},
			want: &gotwtr.FollowingResponse{
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
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.Following(tt.args.ctx, tt.args.userID, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Following() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.Following() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func Test_postFollowing(t *testing.T) {
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
		want    *gotwtr.PostFollowingResponse
		wantErr bool
	}{
		{
			name: "200 success public user",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					if req.Method != http.MethodPost {
						t.Errorf("the method is not correct got %s want %s", req.Method, http.MethodPost)
					}
					body := `{
						"data": {
							"following": true,
							"pending_follow": false
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
			want: &gotwtr.PostFollowingResponse{
				Following: &gotwtr.Following{
					Following:     true,
					PendingFollow: false,
				},
			},
			wantErr: false,
		},
		{
			name: "200 success protected user",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					if req.Method != http.MethodPost {
						t.Errorf("the method is not correct got %s want %s", req.Method, http.MethodPost)
					}
					body := `{
						"data": {
							"following": false,
							"pending_follow": true
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
			want: &gotwtr.PostFollowingResponse{
				Following: &gotwtr.Following{
					Following:     false,
					PendingFollow: true,
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
								"message":"Sorry, that page does not exist, code:34."
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusBadRequest,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				userID:       "2244994945",
				targetUserID: "1228393702244134912",
			},
			want: &gotwtr.PostFollowingResponse{
				Errors: []*gotwtr.APIResponseError{
					{
						Message: "Sorry, that page does not exist, code:34.",
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
								"detail": "Authenticating with OAuth 2.0 Application-Only is forbidden for this endpoint. Supported authentication types are [OAuth 1.0a User Context, OAuth 2.0 User Context].",
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
			want: &gotwtr.PostFollowingResponse{
				Following: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Title:  "Unsupported Authentication",
						Detail: "Authenticating with OAuth 2.0 Application-Only is forbidden for this endpoint. Supported authentication types are [OAuth 1.0a User Context, OAuth 2.0 User Context].",
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
								"value":"111111111122",
								"detail":"Could not find user with id: [111111111122].",
								"title":"Not Found Error",
								"resource_type":"user",
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
				userID:       "111111111122",
				targetUserID: "1228393702244134912",
			},
			want: &gotwtr.PostFollowingResponse{
				Following: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Value:        "111111111122",
						Detail:       "Could not find user with id: [111111111122].",
						Title:        "Not Found Error",
						ResourceType: "user",
						Parameter:    "id",
						ResourceID:   "111111111122",
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
			got, err := c.PostFollowing(tt.args.ctx, tt.args.userID, tt.args.targetUserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostFollowing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("PostFollowing() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func Test_undoFollowing(t *testing.T) {
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
		want    *gotwtr.UndoFollowingResponse
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
							"following": false
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
			want: &gotwtr.UndoFollowingResponse{
				Following: &gotwtr.Following{
					Following: false,
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
								"message":"Sorry, that page does not exist,   code:34"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusBadRequest,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				sourceUserID: "2244994945",
				targetUserID: "1228393702244134912",
			},
			want: &gotwtr.UndoFollowingResponse{
				Errors: []*gotwtr.APIResponseError{
					{
						Message: "Sorry, that page does not exist,   code:34",
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
								"detail": "Authenticating with OAuth 2.0 Application-Only is forbidden for this endpoint.   Supported authentication types are [OAuth 1.0a User Context, OAuth 2.0 User Context].",
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
			want: &gotwtr.UndoFollowingResponse{
				Following: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Title:  "Unsupported Authentication",
						Detail: "Authenticating with OAuth 2.0 Application-Only is forbidden for this endpoint.   Supported authentication types are [OAuth 1.0a User Context, OAuth 2.0 User Context].",
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
				sourceUserID: "111111111133",
				targetUserID: "1228393702244134912",
			},
			want: &gotwtr.UndoFollowingResponse{
				Following: nil,
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
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.UndoFollowing(tt.args.ctx, tt.args.sourceUserID, tt.args.targetUserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UndoFollowing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("UndoFollowing() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
