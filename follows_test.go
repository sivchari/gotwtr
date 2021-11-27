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
		id     string
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
				id:  "2244994945",
				opt: []*gotwtr.FollowOption{},
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
				Meta: gotwtr.FollowsMeta{
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
				id: "2244994945",
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
				id:  "1111111111111111111111",
				opt: []*gotwtr.FollowOption{},
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
	for i, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.LookUpFollowers(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Followers() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.Followers() index = %v mismatch (-want +got):\n%s", i, diff)
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
		id     string
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
				id:  "2244994945",
				opt: []*gotwtr.FollowOption{},
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
				Meta: gotwtr.FollowsMeta{
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
				id: "2244994945",
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
				id:  "1111111111111111111111",
				opt: []*gotwtr.FollowOption{},
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
	for i, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.LookUpFollowing(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Following() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.Following() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
