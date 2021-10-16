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

func Test_retrieveStreamRules(t *testing.T) {
	type args struct {
		ctx    context.Context
		client *http.Client
		opt    []*gotwtr.RetrieveStreamRulesOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.RetrieveStreamRulesResponse
		wantErr bool
	}{
		{
			name: "Success 200 Retrieve Rules",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					body := `{
					    "data": [
					        {
					            "id": "1273636687768285186",
					            "value": "meme has:images"
					        },
					        {
					            "id": "1273636687768285187",
					            "value": "puppy has:media",
					            "tag": "puppies with media"
					        }
					    ],
					    "meta": {
					        "sent": "2020-06-18T15:21:58.638Z"
					    }
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				opt: []*gotwtr.RetrieveStreamRulesOption{},
			},
			want: &gotwtr.RetrieveStreamRulesResponse{
				Rules: []*gotwtr.FilteredRule{
					{
						ID:    "1273636687768285186",
						Value: "meme has:images",
					},
					{
						ID:    "1273636687768285187",
						Value: "puppy has:media",
						Tag:   "puppies with media",
					},
				},
				Meta: &gotwtr.RetrieveStreamRulesMeta{
					Sent: "2020-06-18T15:21:58.638Z",
				},
			},
			wantErr: false,
		},
		{
			name: "Success 200 Retrieve Rules with Option",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					body := `{
					    "data": [
					        {
					            "id": "1273636687768285186",
					            "value": "meme has:images"
					        }
					    ],
					    "meta": {
					        "sent": "2020-06-18T15:21:58.638Z"
					    }
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				opt: []*gotwtr.RetrieveStreamRulesOption{
					{
						[]string{"1273636687768285186"},
					},
				},
			},
			want: &gotwtr.RetrieveStreamRulesResponse{
				Rules: []*gotwtr.FilteredRule{
					{
						ID:    "1273636687768285186",
						Value: "meme has:images",
					},
				},
				Meta: &gotwtr.RetrieveStreamRulesMeta{
					Sent: "2020-06-18T15:21:58.638Z",
				},
			},
			wantErr: false,
		},
		{
			name: "403 Forbidden Client Not Enrolled",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					body := `{
						"errors": [
							{
							    "client_id": "16340226",
							    "required_enrollment": "Standard Basic",
							    "registration_url": "https://developer.twitter.com/en/account",
							    "title": "Client Forbidden",
							    "detail": "This request must be made using an approved developer account that is enrolled in the requested endpoint. Learn more by visiting our documentation.",
							    "reason": "client-not-enrolled",
							    "type": "https://api.twitter.com/2/problems/client-forbidden"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusForbidden,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				opt: []*gotwtr.RetrieveStreamRulesOption{},
			},
			want: &gotwtr.RetrieveStreamRulesResponse{
				Errors: []*gotwtr.APIResponseError{
					{
						ClientID:           "16340226",
						RequiredEnrollment: "Standard Basic",
						RegistrationURL:    "https://developer.twitter.com/en/account",
						Title:              "Client Forbidden",
						Detail:             "This request must be made using an approved developer account that is enrolled in the requested endpoint. Learn more by visiting our documentation.",
						Reason:             "client-not-enrolled",
						Type:               "https://api.twitter.com/2/problems/client-forbidden",
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
			got, err := c.RetrieveStreamRules(tt.args.ctx, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("searchRecentTweets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("searchRecentTweets() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func Test_AddOrDeleteRules(t *testing.T) {
	type args struct {
		ctx    context.Context
		client *http.Client
		opt    []*gotwtr.AddOrDeleteRulesOption
		body   *gotwtr.AddOrDeleteJSONBody
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.AddOrDeleteRulesResponse
		wantErr bool
	}{
		{
			name: "Success 201 Created List of Rules",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					body := `{
					    "data": [
					        {
					            "value": "meme has:images",
					            "id": "1273636687768285186"
					        },
					        {
					            "value": "puppy has:media",
					            "tag": "puppies with media",
					            "id": "1273636687768285187"
					        }
					    ],
					    "meta": {
					        "sent": "2020-06-18T15:20:24.063Z",
					        "summary": {
					            "created": 2,
					            "not_created": 0,
					            "valid": 2,
					            "invalid": 0
					        }
					    }
					}}`
					return &http.Response{
						StatusCode: http.StatusCreated,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				opt: []*gotwtr.AddOrDeleteRulesOption{},
				body: &gotwtr.AddOrDeleteJSONBody{
					Add: []*gotwtr.Add{
						{
							Value: "puppy has:media",
							Tag:   "puppies with media",
						},
						{
							Value: "meme has:images",
						},
					},
					Delete: &gotwtr.Delete{IDs: []string{}},
				},
			},
			want: &gotwtr.AddOrDeleteRulesResponse{
				Rules: []*gotwtr.FilteredRule{
					{
						ID:    "1273636687768285186",
						Value: "meme has:images",
					},
					{
						ID:    "1273636687768285187",
						Value: "puppy has:media",
						Tag:   "puppies with media",
					},
				},
				Meta: &gotwtr.AddOrDeleteRulesMeta{
					Sent: "2020-06-18T15:20:24.063Z",
					Summary: &gotwtr.AddOrDeleteMetaSummary{
						Created:    2,
						NotCreated: 0,
						Valid:      2,
						Invalid:    0,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "403 Forbidden Client Not Enrolled",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					body := `{
						"errors": [
							{
							    "client_id": "16340226",
							    "required_enrollment": "Standard Basic",
							    "registration_url": "https://developer.twitter.com/en/account",
							    "title": "Client Forbidden",
							    "detail": "This request must be made using an approved developer account that is enrolled in the requested endpoint. Learn more by visiting our documentation.",
							    "reason": "client-not-enrolled",
							    "type": "https://api.twitter.com/2/problems/client-forbidden"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusForbidden,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				opt: []*gotwtr.AddOrDeleteRulesOption{},
				body: &gotwtr.AddOrDeleteJSONBody{
					Add: []*gotwtr.Add{
						{
							Value: "tostones recipe",
						},
					},
					Delete: &gotwtr.Delete{IDs: []string{}},
				},
			},
			want: &gotwtr.AddOrDeleteRulesResponse{
				Errors: []*gotwtr.APIResponseError{
					{
						ClientID:           "16340226",
						RequiredEnrollment: "Standard Basic",
						RegistrationURL:    "https://developer.twitter.com/en/account",
						Title:              "Client Forbidden",
						Detail:             "This request must be made using an approved developer account that is enrolled in the requested endpoint. Learn more by visiting our documentation.",
						Reason:             "client-not-enrolled",
						Type:               "https://api.twitter.com/2/problems/client-forbidden",
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
			got, err := c.AddOrDeleteRules(tt.args.ctx, tt.args.body, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("addOrDelteRules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("addOrDelteRules() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
