package gotwtr_test

import (
	"context"
	_ "embed"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sivchari/gotwtr"
)

//go:embed testdata/lookup_all_one_to_one_dm.json
var lookUpAllOneToOneDM []byte

//go:embed testdata/lookup_all_one_to_one_dm_option.json
var lookUpAllOneToOneDMOption []byte

//go:embed testdata/lookup_all_dm.json
var lookUpAllDM []byte

//go:embed testdata/lookup_all_dm_option.json
var lookUpAllDMOption []byte

//go:embed testdata/lookup_dm.json
var lookUpDM []byte

//go:embed testdata/lookup_dm_option.json
var lookUpDMOption []byte

//go:embed testdata/403.json
var forbidden []byte

func Test_LookUpAllOneToOneDM(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx           context.Context
		client        *http.Client
		participantID string
		opt           []*gotwtr.DirectMessageOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.LookUpAllOneToOneDMResponse
		wantErr bool
	}{
		{
			name: "200 ok",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(string(lookUpAllOneToOneDM))),
					}
				}),
				participantID: "1346889436626259968",
			},
			want: &gotwtr.LookUpAllOneToOneDMResponse{
				Message: []*gotwtr.DirectMessage{
					{
						EventType: "MessageCreate",
						ID:        "1346889436626259968",
						Text:      "Hello just you...",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "200 ok and option",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(string(lookUpAllOneToOneDMOption))),
					}
				}),
				participantID: "1585321444547837956",
				opt: []*gotwtr.DirectMessageOption{
					{
						DMEventFields: []gotwtr.DMEventField{
							gotwtr.DirectMessageFieldDMConversationID,
							gotwtr.DirectMessageFieldCreatedAt,
							gotwtr.DirectMessageFieldSenderID,
						},
					},
				},
			},
			want: &gotwtr.LookUpAllOneToOneDMResponse{
				Message: []*gotwtr.DirectMessage{
					{
						ID:               "1585321444547837956",
						Text:             "Another photo https://t.co/J5KotyeIyd",
						EventType:        "MessageCreate",
						DMConversationID: "1585094756761149440",
						CreatedAt:        "2022-10-26T17:24:21.000Z",
						SenderID:         "906948460078698496",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "403 authentication error",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusForbidden,
						Body:       io.NopCloser(strings.NewReader(string(forbidden))),
					}
				}),
				participantID: "2244994945",
			},
			want: &gotwtr.LookUpAllOneToOneDMResponse{
				Errors: []*gotwtr.APIResponseError{
					{
						Title:  "Unsupported Authentication",
						Detail: "Authenticating with OAuth 2.0 Application-Only is forbidden for this endpoint.  Supported authentication types are [OAuth 1.0a User Context, OAuth 2.0 User Context].",
						Type:   "https://api.twitter.com/2/problems/unsupported-authentication",
						Status: 403,
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
			got, err := c.LookUpAllOneToOneDM(tt.args.ctx, tt.args.participantID, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.LookUpAllOneToOneDM() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.LookUpAllOneToOneDM() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
func Test_LookUpDM(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx            context.Context
		client         *http.Client
		conversationID string
		opt            []*gotwtr.DirectMessageOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.LookUpDMResponse
		wantErr bool
	}{
		{
			name: "200 ok",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(string(lookUpDM))),
					}
				}),
				conversationID: "1346889436626259968",
				opt:            nil,
			},
			want: &gotwtr.LookUpDMResponse{
				Message: []*gotwtr.DirectMessage{
					{
						EventType: "MessageCreate",
						ID:        "1346889436626259968",
						Text:      "Hello just you...",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "200 ok and option",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(string(lookUpDMOption))),
					}
				}),
				conversationID: "1346889436626259968",
				opt: []*gotwtr.DirectMessageOption{
					{
						DMEventFields: []gotwtr.DMEventField{
							gotwtr.DirectMessageFieldDMConversationID,
							gotwtr.DirectMessageFieldCreatedAt,
							gotwtr.DirectMessageFieldSenderID,
						},
					},
				},
			},
			want: &gotwtr.LookUpDMResponse{
				Message: []*gotwtr.DirectMessage{
					{
						ID:               "1585321444547837956",
						Text:             "Another photo https://t.co/J5KotyeIyd",
						EventType:        "MessageCreate",
						DMConversationID: "1585094756761149440",
						CreatedAt:        "2022-10-26T17:24:21.000Z",
						SenderID:         "906948460078698496",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "403 authentication error",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusForbidden,
						Body:       io.NopCloser(strings.NewReader(string(forbidden))),
					}
				}),
				conversationID: "1346889436626259968",
				opt:            nil,
			},
			want: &gotwtr.LookUpDMResponse{
				Message: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Title:  "Unsupported Authentication",
						Detail: "Authenticating with OAuth 2.0 Application-Only is forbidden for this endpoint.  Supported authentication types are [OAuth 1.0a User Context, OAuth 2.0 User Context].",
						Type:   "https://api.twitter.com/2/problems/unsupported-authentication",
						Status: 403,
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
			got, err := c.LookUpDM(tt.args.ctx, tt.args.conversationID, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.LookUpDM() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.LookUpDM() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
func Test_LookUpAllDM(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		opt    []*gotwtr.DirectMessageOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.LookUpAllDMResponse
		wantErr bool
	}{
		{
			name: "200 ok",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(string(lookUpAllDM))),
					}
				}),
				opt: nil,
			},
			want: &gotwtr.LookUpAllDMResponse{
				Message: []*gotwtr.DirectMessage{
					{
						EventType: "MessageCreate",
						ID:        "1346889436626259968",
						Text:      "Hello just you...",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "200 ok and option",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(string(lookUpAllDMOption))),
					}
				}),
				opt: []*gotwtr.DirectMessageOption{
					{
						DMEventFields: []gotwtr.DMEventField{
							gotwtr.DirectMessageFieldDMConversationID,
							gotwtr.DirectMessageFieldCreatedAt,
							gotwtr.DirectMessageFieldSenderID,
						},
					},
				},
			},
			want: &gotwtr.LookUpAllDMResponse{
				Message: []*gotwtr.DirectMessage{
					{
						ID:               "1585321444547837956",
						Text:             "Another photo https://t.co/J5KotyeIyd",
						EventType:        "MessageCreate",
						DMConversationID: "1585094756761149440",
						CreatedAt:        "2022-10-26T17:24:21.000Z",
						SenderID:         "906948460078698496",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "403 authentication error",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusForbidden,
						Body:       io.NopCloser(strings.NewReader(string(forbidden))),
					}
				}),
				opt: nil,
			},
			want: &gotwtr.LookUpAllDMResponse{
				Message: nil,
				Errors: []*gotwtr.APIResponseError{
					{
						Title:  "Unsupported Authentication",
						Detail: "Authenticating with OAuth 2.0 Application-Only is forbidden for this endpoint.  Supported authentication types are [OAuth 1.0a User Context, OAuth 2.0 User Context].",
						Type:   "https://api.twitter.com/2/problems/unsupported-authentication",
						Status: 403,
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
			got, err := c.LookUpAllDM(tt.args.ctx, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.LookUpAllDM() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.LookUpAllDM() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
