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

func Test_lookUpListFollowers(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		id     string
		opt    []*gotwtr.ListFollowersOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.ListFollowersResponse
		wantErr bool
	}{
		{
			name: "200 ok no option",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"id": "1324848235714736129",
								"name": "Alan Lee",
								"username": "alanbenlee"
							},
							{
								"id": "1328359963937259520",
								"name": "Wilson Chong",
								"username": "xo_chong"
							},
							{
								"id": "1451609880113070085",
								"name": "Sumira Nazir",
								"username": "SumiraNazir"
							},
							{
								"id": "1420055293082415107",
								"name": "Bořek Šindelka(he/him)",
								"username": "JustBorek"
							},
							{
								"id": "1409136449803325441",
								"name": "金井憲司",
								"username": "w22ZccksRpafZAx"
							}
						],
						"meta": {
							"result_count": 5,
							"next_token": "1714209892546977900"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id:  "84839422",
				opt: []*gotwtr.ListFollowersOption{},
			},
			want: &gotwtr.ListFollowersResponse{
				Users: []*gotwtr.User{
					{
						ID:       "1324848235714736129",
						Name:     "Alan Lee",
						UserName: "alanbenlee",
					},
					{
						ID:       "1328359963937259520",
						Name:     "Wilson Chong",
						UserName: "xo_chong",
					},
					{
						ID:       "1451609880113070085",
						Name:     "Sumira Nazir",
						UserName: "SumiraNazir",
					},
					{
						ID:       "1420055293082415107",
						Name:     "Bořek Šindelka(he/him)",
						UserName: "JustBorek",
					},
					{
						ID:       "1409136449803325441",
						Name:     "金井憲司",
						UserName: "w22ZccksRpafZAx",
					},
				},
				Meta: &gotwtr.ListMeta{
					ResultCount: 5,
					NextToken:   "1714209892546977900",
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
							"username": "alanbenlee",
							"name": "Alan Lee",
							"created_at": "2020-11-06T22:56:55.000Z",
							"id": "1324848235714736129"
						},
						{
							"pinned_tweet_id": "1452599033625657359",
							"username": "xo_chong",
							"name": "Wilson Chong",
							"created_at": "2020-11-16T15:31:22.000Z",
							"id": "1328359963937259520"
						},
						{
							"username": "SumiraNazir",
							"name": "Sumira Nazir",
							"created_at": "2021-10-22T18:02:31.000Z",
							"id": "1451609880113070085"
						},
						{
							"pinned_tweet_id": "1442182396523257861",
							"username": "JustBorek",
							"name": "Bořek Šindelka(he/him)",
							"created_at": "2021-07-27T16:16:23.000Z",
							"id": "1420055293082415107"
						},
						{
							"username": "w22ZccksRpafZAx",
							"name": "金井憲司",
							"created_at": "2021-06-27T13:09:10.000Z",
							"id": "1409136449803325441"
						}
					],
					"includes": {
						"tweets": [
							{
								"created_at": "2021-10-25T11:32:52.000Z",
								"id": "1452599033625657359",
								"text": "https://t.co/aEuBQLXeuL"
							},
							{
								"created_at": "2021-09-26T17:40:52.000Z",
								"id": "1442182396523257861",
								"text": "Yes couple of days back nI want to kill my self I'm still here because of some amazing people please share this is important to talk about #mentalhealth @JustBorek #wheelchair #DisabilityTwitter #MedTwitter @heatherpsyd @Tweetinggoddess @NashaterS @msfatale @castleDD https://t.co/9hkSPV9NB1"
							}
						]
					},
						"meta": {
							"result_count": 5,
							"next_token": "1714209892546977900"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id: "84839422",
				opt: []*gotwtr.ListFollowersOption{
					{
						Expansions: []gotwtr.Expansion{
							gotwtr.ExpansionPinnedTweetID,
						},
						UserFields: []gotwtr.UserField{
							gotwtr.UserFieldCreatedAt,
						},
						TweetFields: []gotwtr.TweetField{
							gotwtr.TweetFieldCreatedAt,
						},
					},
				},
			},
			want: &gotwtr.ListFollowersResponse{
				Users: []*gotwtr.User{
					{
						UserName:  "alanbenlee",
						Name:      "Alan Lee",
						CreatedAt: "2020-11-06T22:56:55.000Z",
						ID:        "1324848235714736129",
					},
					{
						PinnedTweetID: "1452599033625657359",
						UserName:      "xo_chong",
						Name:          "Wilson Chong",
						CreatedAt:     "2020-11-16T15:31:22.000Z",
						ID:            "1328359963937259520",
					},
					{
						UserName:  "SumiraNazir",
						Name:      "Sumira Nazir",
						CreatedAt: "2021-10-22T18:02:31.000Z",
						ID:        "1451609880113070085",
					},
					{
						PinnedTweetID: "1442182396523257861",
						UserName:      "JustBorek",
						Name:          "Bořek Šindelka(he/him)",
						CreatedAt:     "2021-07-27T16:16:23.000Z",
						ID:            "1420055293082415107",
					},
					{
						UserName:  "w22ZccksRpafZAx",
						Name:      "金井憲司",
						CreatedAt: "2021-06-27T13:09:10.000Z",
						ID:        "1409136449803325441",
					},
				},
				Includes: &gotwtr.ListIncludes{
					Tweets: []*gotwtr.Tweet{
						{
							CreatedAt: "2021-10-25T11:32:52.000Z",
							ID:        "1452599033625657359",
							Text:      "https://t.co/aEuBQLXeuL",
						},
						{
							CreatedAt: "2021-09-26T17:40:52.000Z",
							ID:        "1442182396523257861",
							Text:      "Yes couple of days back nI want to kill my self I'm still here because of some amazing people please share this is important to talk about #mentalhealth @JustBorek #wheelchair #DisabilityTwitter #MedTwitter @heatherpsyd @Tweetinggoddess @NashaterS @msfatale @castleDD https://t.co/9hkSPV9NB1",
						},
					},
				},
				Meta: &gotwtr.ListMeta{
					ResultCount: 5,
					NextToken:   "1714209892546977900",
				},
				Errors: nil,
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
									"id":[
										"111111111111111111111111111111111111"
									]
								},
								"message":"The id query parameter value [111111111111111111111111111111111111] is not valid"
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
				id:  "111111111111111111111111111111111111",
				opt: []*gotwtr.ListFollowersOption{},
			},
			want: &gotwtr.ListFollowersResponse{
				Users: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Parameters: gotwtr.Parameter{
							ID: []string{"111111111111111111111111111111111111"},
						},
						Message: "The id query parameter value [111111111111111111111111111111111111] is not valid",
					},
				},
				Meta:   nil,
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
			c := gotwtr.New("key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.LookUpListFollowers(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.LookUpListFollowersByID() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.LookUpListFollowersByID() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}

func Test_lookUpAllListsUserFollows(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		id     string
		opt    []*gotwtr.ListFollowsOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.AllListsUserFollowsResponse
		wantErr bool
	}{
		{
			name: "200 ok no option",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"id": "1630685563471",
								"name": "Test List"
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
				id:  "2244994945",
				opt: []*gotwtr.ListFollowsOption{},
			},
			want: &gotwtr.AllListsUserFollowsResponse{
				Lists: []*gotwtr.List{
					{
						ID:   "1630685563471",
						Name: "Test List",
					},
				},
				Meta: &gotwtr.ListMeta{
					ResultCount: 1,
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
								"follower_count": 123,
								"id": "1630685563471",
								"name": "Test List",
								"owner_id": "1324848235714736129"
							}
						],
						"includes": {
							"users": [
								{
									"username": "alanbenlee",
									"id": "1324848235714736129",
									"created_at": "2009-08-28T18:30:45.000Z",
									"name": "Alan Lee"
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
				id: "2244994945",
				opt: []*gotwtr.ListFollowsOption{
					{
						Expansions: []gotwtr.Expansion{
							gotwtr.ExpansionOwnerID,
						},
						ListFields: []gotwtr.ListField{
							gotwtr.ListFollowerCount,
						},
						UserFields: []gotwtr.UserField{
							gotwtr.UserFieldUserName,
						},
					},
				},
			},
			want: &gotwtr.AllListsUserFollowsResponse{
				Lists: []*gotwtr.List{
					{
						FollowerCount: 123,
						ID:            "1630685563471",
						Name:          "Test List",
						OwnerID:       "1324848235714736129",
					},
				},
				Includes: &gotwtr.ListIncludes{
					Users: []*gotwtr.User{
						{
							UserName:  "alanbenlee",
							ID:        "1324848235714736129",
							CreatedAt: "2009-08-28T18:30:45.000Z",
							Name:      "Alan Lee",
						},
					},
				},
				Meta: &gotwtr.ListMeta{
					ResultCount: 1,
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
									"id":[
										"111111111111111111111111111111111111111111111111111111"
									]
								},
								"message":"The id query parameter value [111111111111111111111111111111111111111111111111111111] is not valid"
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
				id:  "111111111111111111111111111111111111111111111111111111",
				opt: []*gotwtr.ListFollowsOption{},
			},
			want: &gotwtr.AllListsUserFollowsResponse{
				Lists: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Parameters: gotwtr.Parameter{
							ID: []string{"111111111111111111111111111111111111111111111111111111"},
						},
						Message: "The id query parameter value [111111111111111111111111111111111111111111111111111111] is not valid",
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
			c := gotwtr.New("key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.LookUpAllListsUserFollows(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.LookUpAllListsUserFollows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.LookUpAllListsUserFollows() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
