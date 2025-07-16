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

func Test_postDMBlocking(t *testing.T) {
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
		want    *gotwtr.PostDMBlockingResponse
		wantErr bool
	}{
		{
			name: "200 ok dm blocking successful",
			args: args{
				ctx:          context.Background(),
				userID:       "123456789",
				targetUserID: "987654321",
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"data": {
							"dm_blocking": true
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
			},
			want: &gotwtr.PostDMBlockingResponse{
				DMBlocking: &gotwtr.DMBlocking{
					DMBlocking: true,
				},
			},
			wantErr: false,
		},
		{
			name: "400 bad request - already blocked",
			args: args{
				ctx:          context.Background(),
				userID:       "123456789",
				targetUserID: "987654321",
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"title": "Bad Request",
						"detail": "User is already blocked for DMs",
						"type": "https://api.twitter.com/2/problems/invalid-request"
					}`
					return &http.Response{
						StatusCode: http.StatusBadRequest,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
			},
			want: &gotwtr.PostDMBlockingResponse{
				Title:  "Bad Request",
				Detail: "User is already blocked for DMs",
				Type:   "https://api.twitter.com/2/problems/invalid-request",
			},
			wantErr: true,
		},
		{
			name: "empty userID",
			args: args{
				ctx:          context.Background(),
				userID:       "",
				targetUserID: "987654321",
				client:       http.DefaultClient,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty targetUserID",
			args: args{
				ctx:          context.Background(),
				userID:       "123456789",
				targetUserID: "",
				client:       http.DefaultClient,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.PostDMBlocking(tt.args.ctx, tt.args.userID, tt.args.targetUserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.PostDMBlocking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.PostDMBlocking() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func Test_undoDMBlocking(t *testing.T) {
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
		want    *gotwtr.UndoDMBlockingResponse
		wantErr bool
	}{
		{
			name: "200 ok dm unblocking successful",
			args: args{
				ctx:          context.Background(),
				userID:       "123456789",
				targetUserID: "987654321",
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"data": {
							"dm_blocking": false
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
			},
			want: &gotwtr.UndoDMBlockingResponse{
				DMBlocking: &gotwtr.DMBlocking{
					DMBlocking: false,
				},
			},
			wantErr: false,
		},
		{
			name: "400 bad request - not blocked",
			args: args{
				ctx:          context.Background(),
				userID:       "123456789",
				targetUserID: "987654321",
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"title": "Bad Request",
						"detail": "User is not blocked for DMs",
						"type": "https://api.twitter.com/2/problems/invalid-request"
					}`
					return &http.Response{
						StatusCode: http.StatusBadRequest,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
			},
			want: &gotwtr.UndoDMBlockingResponse{
				Title:  "Bad Request",
				Detail: "User is not blocked for DMs",
				Type:   "https://api.twitter.com/2/problems/invalid-request",
			},
			wantErr: true,
		},
		{
			name: "empty userID",
			args: args{
				ctx:          context.Background(),
				userID:       "",
				targetUserID: "987654321",
				client:       http.DefaultClient,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty targetUserID",
			args: args{
				ctx:          context.Background(),
				userID:       "123456789",
				targetUserID: "",
				client:       http.DefaultClient,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "404 not found - invalid user",
			args: args{
				ctx:          context.Background(),
				userID:       "123456789",
				targetUserID: "invalid_user",
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"title": "Not Found Error",
						"detail": "Could not find user with id: [invalid_user].",
						"type": "https://api.twitter.com/2/problems/resource-not-found"
					}`
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
			},
			want: &gotwtr.UndoDMBlockingResponse{
				Title:  "Not Found Error",
				Detail: "Could not find user with id: [invalid_user].",
				Type:   "https://api.twitter.com/2/problems/resource-not-found",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.UndoDMBlocking(tt.args.ctx, tt.args.userID, tt.args.targetUserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.UndoDMBlocking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.UndoDMBlocking() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}