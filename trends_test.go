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

func Test_trendsByWOEID(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		woeid  string
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.TrendsByWOEIDResponse
		wantErr bool
	}{
		{
			name: "200 ok the request was successful",
			args: args{
				ctx:   context.Background(),
				woeid: "1",
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"data": [
							{
								"trend_name": "Europe",
								"tweet_count": 232408
							},
							{
								"trend_name": "Isak",
								"tweet_count": 2956
							},
							{
								"trend_name": "#MondayMotivation"
							}
						]
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
			},
			want: &gotwtr.TrendsByWOEIDResponse{
				Trends: []*gotwtr.Trend{
					{
						TrendName:  "Europe",
						TweetCount: 232408,
					},
					{
						TrendName:  "Isak",
						TweetCount: 2956,
					},
					{
						TrendName: "#MondayMotivation",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "404 not found - invalid WOEID",
			args: args{
				ctx:   context.Background(),
				woeid: "999999999",
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"title": "Not Found Error",
						"detail": "Could not find trends for WOEID: 999999999",
						"type": "https://api.x.com/2/problems/resource-not-found"
					}`
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
			},
			want: &gotwtr.TrendsByWOEIDResponse{
				Title:  "Not Found Error",
				Detail: "Could not find trends for WOEID: 999999999",
				Type:   "https://api.x.com/2/problems/resource-not-found",
			},
			wantErr: true,
		},
		{
			name: "empty woeid",
			args: args{
				ctx:    context.Background(),
				woeid:  "",
				client: http.DefaultClient,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "rate limit exceeded",
			args: args{
				ctx:   context.Background(),
				woeid: "1",
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"title": "Too Many Requests",
						"detail": "Too Many Requests",
						"type": "https://api.x.com/2/problems/usage-capped"
					}`
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
			},
			want: &gotwtr.TrendsByWOEIDResponse{
				Title:  "Too Many Requests",
				Detail: "Too Many Requests",
				Type:   "https://api.x.com/2/problems/usage-capped",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.TrendsByWOEID(tt.args.ctx, tt.args.woeid)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.TrendsByWOEID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.TrendsByWOEID() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}