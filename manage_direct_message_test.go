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

//go:embed testdata/create_one_to_one_dm.json
var createOneToOneDM []byte

//go:embed testdata/create_new_group_dm.json
var createNewGroupDM []byte

//go:embed testdata/post_dm.json
var postDM []byte

//go:embed testdata/403.json
var forbidden []byte

func Test_CreateOneToOneDM(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx           context.Context
		client        *http.Client
		participantID string
		body          *gotwtr.CreateOneToOneDMBody
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.CreateOneToOneDMResponse
		wantErr bool
	}{
		{
			name: "201 Create a one to one dm",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusCreated,
						Body:       io.NopCloser(strings.NewReader(string(createOneToOneDM))),
					}
				}),
				participantID: "2244994945",
				body: &gotwtr.CreateOneToOneDMBody{
					Text: "This is a one-to-one Direct Message with an attachment",
					Attachments: []gotwtr.MessageAttachment{
						{
							MediaID: "1455952740635586573",
						},
					},
				},
			},
			want: &gotwtr.CreateOneToOneDMResponse{
				DMConversationID: "1346889436626259968",
				DMEventID:        "128341038123",
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
				body: &gotwtr.CreateOneToOneDMBody{
					Text: "This is a one-to-one Direct Message with an attachment",
					Attachments: []gotwtr.MessageAttachment{
						{
							MediaID: "1455952740635586573",
						},
					},
				},
			},
			want: &gotwtr.CreateOneToOneDMResponse{
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
			got, err := c.CreateOneToOneDM(tt.args.ctx, tt.args.participantID, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CreateOneToOneDM() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.CreateNewOneToOneDM() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
func Test_CreateNewGroupDM(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx            context.Context
		client         *http.Client
		conversationID string
		body           *gotwtr.CreateNewGroupDMBody
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.CreateNewGroupDMResponse
		wantErr bool
	}{
		{
			name: "201 Create New Group",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusCreated,
						Body:       io.NopCloser(strings.NewReader(string(createNewGroupDM))),
					}
				}),
				conversationID: "1346889436626259968",
				body: &gotwtr.CreateNewGroupDMBody{
					Text: "Adding a Direct Message to a conversation by referencing the conversation ID. This method supports both one-to-one and group conversations.",
				},
			},
			want: &gotwtr.CreateNewGroupDMResponse{
				DMConversationID: "1346889436626259968",
				DMEventID:        "128341038123",
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
				body: &gotwtr.CreateNewGroupDMBody{
					Text: "Adding a Direct Message to a conversation by referencing the conversation ID. This method supports both one-to-one and group conversations.",
				},
			},
			want: &gotwtr.CreateNewGroupDMResponse{
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
			got, err := c.CreateNewGroupDM(tt.args.ctx, tt.args.conversationID, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CreateNewGroupDM() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.CreateNewGroupDM() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
func Test_PostDM(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		body   *gotwtr.PostDMBody
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.PostDMResponse
		wantErr bool
	}{
		{
			name: "201 Post DM",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusCreated,
						Body:       io.NopCloser(strings.NewReader(string(postDM))),
					}
				}),
				body: &gotwtr.PostDMBody{
					ConversationType: "Group",
					ParticipantIDs: []string{
						"944480690",
						"906948460078698496",
					},
					Message: &gotwtr.Message{
						Text: "Hello to you two, this is a new group conversation",
					},
				},
			},
			want: &gotwtr.PostDMResponse{
				DMConversationID: "1346889436626259968",
				DMEventID:        "128341038123",
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
				body: &gotwtr.PostDMBody{
					ConversationType: "Group",
					ParticipantIDs: []string{
						"944480690",
						"906948460078698496",
					},
					Message: &gotwtr.Message{
						Text: "Hello to you two, this is a new group conversation",
					},
				},
			},
			want: &gotwtr.PostDMResponse{
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
			got, err := c.PostDM(tt.args.ctx, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.PostDM() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.PostDM() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
