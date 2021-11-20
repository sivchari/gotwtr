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

func Test_lookUpSpaces(t *testing.T) {
	type args struct {
		ctx    context.Context
		client *http.Client
		ids    []string
		opt    []*gotwtr.SpaceLookUpOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.SpaceLookUpResponse
		wantErr bool
	}{
		{
			name: "200 ok the request was successful",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"data": [
							{
								"id": "12345",
								"state": "ended"
							},
							{
								"id": "67890",
								"state": "ended"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				ids: []string{
					"12345",
					"67890",
				},
				opt: []*gotwtr.SpaceLookUpOption{},
			},
			want: &gotwtr.SpaceLookUpResponse{
				Spaces: []*gotwtr.Space{
					{
						ID:    "12345",
						State: "ended",
					},
					{
						ID:    "67890",
						State: "ended",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "500 internal server error the request has failed",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					id := "`id`"
					body := fmt.Sprintf(`{
						"errors": [
							{
								"parameters": {
									"id": [
										"1234567890"
									]
								},
								"message": "The %v query parameter value [1234567890] does not match ^[a-zA-Z0-9]{1,13}$"
							},
							{
								"parameters": {
									"id": [
										"0987654321"
									]
								},
								"message": "The %v query parameter value [0987654321] does not match ^[a-zA-Z0-9]{1,13}$"
							}
						],
						"title": "Invalid Request",
						"detail": "One or more parameters to your request was invalid.",
						"type": "https://api.twitter.com/2/problems/invalid-request"
					}`, id, id)
					return &http.Response{
						StatusCode: http.StatusInternalServerError,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				ids: []string{
					"1234567890",
					"0987654321",
				},
				opt: []*gotwtr.SpaceLookUpOption{},
			},
			want: &gotwtr.SpaceLookUpResponse{
				Errors: []*gotwtr.APIResponseError{
					{
						Parameters: gotwtr.Parameter{
							ID: []string{"1234567890"},
						},
						Message: "The `id` query parameter value [1234567890] does not match ^[a-zA-Z0-9]{1,13}$",
					},
					{
						Parameters: gotwtr.Parameter{
							ID: []string{"0987654321"},
						},
						Message: "The `id` query parameter value [0987654321] does not match ^[a-zA-Z0-9]{1,13}$",
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
			got, err := c.LookUpSpaces(tt.args.ctx, tt.args.ids, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("LookUpSpaces() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("LookUpSpaces() index = %v diff = %v", i, diff)
				return
			}
		})
	}
}

func Test_lookUpSpaceByID(t *testing.T) {
	type args struct {
		ctx    context.Context
		client *http.Client
		id     string
		opt    []*gotwtr.SpaceLookUpOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.SpaceLookUpByIDResponse
		wantErr bool
	}{
		{
			name: "200 ok the request was successful",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"data": {
							"id": "12345",
							"state": "ended"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				id:  "12345",
				opt: []*gotwtr.SpaceLookUpOption{},
			},
			want: &gotwtr.SpaceLookUpByIDResponse{
				Space: &gotwtr.Space{
					ID:    "12345",
					State: "ended",
				},
			},
			wantErr: false,
		},
		{
			name: "500 internal server error the request has failed",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					id := "`id`"
					body := fmt.Sprintf(`{
						"errors": [
							{
								"parameters": {
									"id": [
										"111111111111111"
									]
								},
								"message": "The %v query parameter value [111111111111111] does not match ^[a-zA-Z0-9]{1,13}$"
							}
						],
						"title": "Invalid Request",
						"detail": "One or more parameters to your request was invalid.",
						"type": "https://api.twitter.com/2/problems/invalid-request"
					}`, id)
					return &http.Response{
						StatusCode: http.StatusInternalServerError,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				id:  "111111111111111",
				opt: []*gotwtr.SpaceLookUpOption{},
			},
			want: &gotwtr.SpaceLookUpByIDResponse{
				Errors: []*gotwtr.APIResponseError{
					{
						Parameters: gotwtr.Parameter{
							ID: []string{"111111111111111"},
						},
						Message: "The `id` query parameter value [111111111111111] does not match ^[a-zA-Z0-9]{1,13}$",
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
			got, err := c.LookUpSpaceByID(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("LookUpSpaceByID() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("LookUpSpaceByID() index = %v diff = %v", i, diff)
				return
			}
		})
	}
}

func Test_lookUpUsersWhoPurchasedSpaceTicket(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		id     string
		opt    []*gotwtr.LookUpUsersWhoPurchasedSpaceTicketOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.LookUpUsersWhoPurchasedSpaceTicketResponse
		wantErr bool
	}{
		{
			name: "200 ok users who bought a ticket to a Space",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"data": [
							{
								"id": "2244994945",
								"username": "TwitterDev",
								"name": "Twitter Dev"
							},
							{
								"id": "783214",
								"username": "Twitter",
								"name": "Twitter"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				id:  "1DXxyRYNejbKM",
				opt: []*gotwtr.LookUpUsersWhoPurchasedSpaceTicketOption{},
			},
			want: &gotwtr.LookUpUsersWhoPurchasedSpaceTicketResponse{
				Users: []*gotwtr.User{
					{
						ID:       "2244994945",
						UserName: "TwitterDev",
						Name:     "Twitter Dev",
					},
					{
						ID:       "783214",
						UserName: "Twitter",
						Name:     "Twitter",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "200 ok users who bought a ticket to a Space with optional fields",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"data": [
							{
								"created_at": "2013-12-14T04:35:55.000Z",
								"username": "TwitterDev",
								"pinned_tweet_id": "1255542774432063488",
								"id": "2244994945",
								"name": "Twitter Dev"
							},
							{
								"created_at": "2007-02-20T14:35:54.000Z",
								"username": "Twitter",
								"pinned_tweet_id": "1274087687469715457",
								"id": "783214",
								"name": "Twitter"
							}
						],
						"includes": {
							"tweets": [
								{
									"created_at": "2020-04-29T17:01:38.000Z",
									"text": "During these unprecedented times, what’s happening on Twitter can help the world better understand &amp; respond to the pandemic. nnWe're launching a free COVID-19 stream endpoint so qualified devs &amp; researchers can study the public conversation in real-time. https://t.co/BPqMcQzhId",
									"id": "1255542774432063488"
								},
								{
									"created_at": "2020-06-19T21:12:30.000Z",
									"text": "📍 Minneapolisn🗣️ @FredTJoseph https://t.co/lNTOkyguG1",
									"id": "1274087687469715457"
								}
							]
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				id: "1DXxyRYNejbKM",
				opt: []*gotwtr.LookUpUsersWhoPurchasedSpaceTicketOption{
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
			want: &gotwtr.LookUpUsersWhoPurchasedSpaceTicketResponse{
				Users: []*gotwtr.User{
					{
						ID:            "2244994945",
						UserName:      "TwitterDev",
						Name:          "Twitter Dev",
						CreatedAt:     "2013-12-14T04:35:55.000Z",
						PinnedTweetID: "1255542774432063488",
					},
					{
						ID:            "783214",
						UserName:      "Twitter",
						Name:          "Twitter",
						CreatedAt:     "2007-02-20T14:35:54.000Z",
						PinnedTweetID: "1274087687469715457",
					},
				},
				Includes: &gotwtr.LookUpUsersWhoPurchasedSpaceTicketIncludes{
					Tweets: []*gotwtr.Tweet{
						{
							ID:        "1255542774432063488",
							CreatedAt: "2020-04-29T17:01:38.000Z",
							Text:      "During these unprecedented times, what’s happening on Twitter can help the world better understand &amp; respond to the pandemic. nnWe're launching a free COVID-19 stream endpoint so qualified devs &amp; researchers can study the public conversation in real-time. https://t.co/BPqMcQzhId",
						},
						{
							ID:        "1274087687469715457",
							CreatedAt: "2020-06-19T21:12:30.000Z",
							Text:      "📍 Minneapolisn🗣️ @FredTJoseph https://t.co/lNTOkyguG1",
						},
					},
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
			got, err := c.LookUpUsersWhoPurchasedSpaceTicket(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("LookUpSpaceByID() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("discoverSpacesByUserIDs() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
