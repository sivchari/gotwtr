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

func Test_searchSpaces(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		query  string
		opt    []*gotwtr.SearchSpacesOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.SearchSpacesResponse
		wantErr bool
	}{
		{
			name: "200 ok the request was successful",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"host_ids": [
									"2244994945"
								],
								"id": "1DXxyRYNejbKM",
								"state": "live",
								"title": "hello world 👋"
							},
							{
								"host_ids": [
									"6253282"
								],
								"id": "1nAJELYEEPvGL",
								"state": "scheduled",
								"title": "Say hello to the Spaces endpoints"
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
				query: "hello",
				opt: []*gotwtr.SearchSpacesOption{
					{
						SpaceFields: []gotwtr.SpaceField{
							gotwtr.SpaceFieldHostIDs,
							gotwtr.SpaceFieldTitle,
						},
					},
				},
			},
			want: &gotwtr.SearchSpacesResponse{
				Spaces: []*gotwtr.Space{
					{
						HostIDs: []string{"2244994945"},
						ID:      "1DXxyRYNejbKM",
						State:   "live",
						Title:   "hello world 👋",
					},
					{
						HostIDs: []string{"6253282"},
						ID:      "1nAJELYEEPvGL",
						State:   "scheduled",
						Title:   "Say hello to the Spaces endpoints",
					},
				},
				Meta: &gotwtr.SearchSpacesMeta{
					ResultCount: 2,
				},
			},
			wantErr: false,
		},
		{
			name: "200 ok the request was successful (result_count is 0)",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					body := `{
						"meta": {
							"result_count": 0
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				query: "hello",
				opt:   []*gotwtr.SearchSpacesOption{},
			},
			want: &gotwtr.SearchSpacesResponse{
				Meta: &gotwtr.SearchSpacesMeta{
					ResultCount: 0,
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
			got, err := c.SearchSpaces(tt.args.ctx, tt.args.query, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.SearchSpaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.SearchSpaces() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
