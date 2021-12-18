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

func Test_discoverSpaces(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		ids    []string
		opt    []*gotwtr.DiscoverSpacesOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.DiscoverSpacesResponse
		wantErr bool
	}{
		{
			name: "200 ok default payload",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"id": "1DXxyRYNejbKM",
								"state": "live"
							}
						],
						"meta": {
							"result_count": 2
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				ids: []string{"2244994945,6253282"},
				opt: []*gotwtr.DiscoverSpacesOption{},
			},
			want: &gotwtr.DiscoverSpacesResponse{
				Spaces: []*gotwtr.Space{
					{
						ID:    "1DXxyRYNejbKM",
						State: "live",
					},
				},
				Meta: &gotwtr.DiscoverSpacesMeta{
					ResultCount: 2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New(gotwtr.WithBearerToken("key"), gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.DiscoverSpaces(tt.args.ctx, tt.args.ids, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.DiscoverSpaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.DiscoverSpaces() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
