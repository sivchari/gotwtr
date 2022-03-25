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

var (
	//go:embed embed/hide_replies_hidden.json
	hideRepliesHidden []byte
	//go:embed embed/hide_replies_unhidden.json
	hideRepliesUnhidden []byte
)

func Test_hideReplies(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx     context.Context
		client  *http.Client
		tweetID string
		hidden  bool
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.HideRepliesResponse
		wantErr bool
	}{
		{
			name: "200 ok default hidden",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(string(hideRepliesHidden))),
					}
				}),
				tweetID: "12345",
				hidden:  true,
			},
			want: &gotwtr.HideRepliesResponse{
				HideRepliesResponseData: &gotwtr.HideRepliesResponseData{
					Hidden: true,
				},
			},
			wantErr: false,
		},
		{
			name: "200 ok default unhidden",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(string(hideRepliesUnhidden))),
					}
				}),
				tweetID: "12345",
				hidden:  false,
			},
			want: &gotwtr.HideRepliesResponse{
				HideRepliesResponseData: &gotwtr.HideRepliesResponseData{
					Hidden: false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.HideReplies(tt.args.ctx, tt.args.tweetID, tt.args.hidden)
			if (err != nil) != tt.wantErr {
				t.Errorf("hideReplies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("hideReplies() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
