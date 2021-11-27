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

func Test_lookUpListTweets(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		id     string
		opt    []*gotwtr.ListTweetsOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.ListTweetsResponse
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
								"id": "1067094924124872705",
								"text": "Just getting started with Twitter APIs? Find out what you need in order to build an app. Watch this video! https://t.co/Hg8nkfoizN"
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
				opt: []*gotwtr.ListTweetsOption{},
			},
			want: &gotwtr.ListTweetsResponse{
				Tweets: []*gotwtr.Tweet{
					{
						ID:   "1067094924124872705",
						Text: "Just getting started with Twitter APIs? Find out what you need in order to build an app. Watch this video! https://t.co/Hg8nkfoizN",
					},
				},
				Meta: &gotwtr.ListMeta{
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
								"author_id": "2244994945",
								"created_at": "2018-11-26T16:37:10.000Z",
								"text": "Just getting started with Twitter APIs? Find out what you need in order to build an app. Watch this video! https://t.co/Hg8nkfoizN",
								"id": "1067094924124872705"
							}
						],
						"includes": {
							"users": [
								{
									"verified": true,
									"username": "TwitterDev",
									"id": "2244994945",
									"name": "Twitter Dev"
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
				id: "84839422",
				opt: []*gotwtr.ListTweetsOption{
					{
						UserFields: []gotwtr.UserField{
							gotwtr.UserFieldVerified,
							gotwtr.UserFieldUserName,
							gotwtr.UserFieldID,
							gotwtr.UserFieldName,
						},
					},
				},
			},
			want: &gotwtr.ListTweetsResponse{
				Tweets: []*gotwtr.Tweet{
					{
						AuthorID:  "2244994945",
						CreatedAt: "2018-11-26T16:37:10.000Z",
						Text:      "Just getting started with Twitter APIs? Find out what you need in order to build an app. Watch this video! https://t.co/Hg8nkfoizN",
						ID:        "1067094924124872705",
					},
				},
				Includes: &gotwtr.ListIncludes{
					Users: []*gotwtr.User{
						{
							Verified: true,
							UserName: "TwitterDev",
							ID:       "2244994945",
							Name:     "Twitter Dev",
						},
					},
				},
				Meta: &gotwtr.ListMeta{
					ResultCount: 1,
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
										"111111111111111111111111"
									]
								},
								"message":"The id query parameter value [111111111111111111111111] is not valid"
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
				id:  "111111111111111111111111",
				opt: []*gotwtr.ListTweetsOption{},
			},
			want: &gotwtr.ListTweetsResponse{
				Tweets: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Parameters: gotwtr.Parameter{
							ID: []string{"111111111111111111111111"},
						},
						Message: "The id query parameter value [111111111111111111111111] is not valid",
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
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.LookUpListTweets(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.LookUpListsTweetsByID() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.LookUpListsTweetsByID() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
