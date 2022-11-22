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

//go:embed testdata/me.json
var me []byte

func Test_me(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		opt    []*gotwtr.MeOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.MeResponse
		wantErr bool
	}{
		{
			name: "200 ok default payload",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(string(me))),
					}
				}),
				opt: []*gotwtr.MeOption{},
			},
			want: &gotwtr.MeResponse{
				Me: &gotwtr.Me{
					ID:       "2244994945",
					Name:     "TwitterDev",
					UserName: "Twitter Dev",
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
			got, err := c.Me(tt.args.ctx, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Me() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.Me() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
