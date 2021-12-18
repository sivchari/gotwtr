package gotwtr_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sivchari/gotwtr"
)

func Test_client_TweetCounts(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		query  string
		opt    []*gotwtr.TweetCountsOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.TweetCountsResponse
		wantErr bool
	}{
		{
			name: "200 ok default",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"end": "2021-09-27T15:00:00.000Z",
								"start": "2021-09-27T14:00:00.000Z",
								"tweet_count": 2
							},
							{
								"end": "2021-09-27T16:00:00.000Z",
								"start": "2021-09-27T15:00:00.000Z",
								"tweet_count": 2
							}
						],
						"meta": {
							"total_tweet_count": 4
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				query: "from:TwitterDev",
				opt:   []*gotwtr.TweetCountsOption{},
			},
			want: &gotwtr.TweetCountsResponse{
				Counts: []*gotwtr.TimeseriesCount{
					{
						Start:      "2021-09-27T14:00:00.000Z",
						End:        "2021-09-27T15:00:00.000Z",
						TweetCount: 2,
					},
					{
						Start:      "2021-09-27T15:00:00.000Z",
						End:        "2021-09-27T16:00:00.000Z",
						TweetCount: 2,
					},
				},
				Meta: &gotwtr.TweetCountMeta{
					TotalTweetCount: 4,
				},
			},
			wantErr: false,
		},
		{
			name: "200 ok specify optional fields",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": [
							{
								"end": "2021-09-28T00:00:00.000Z",
								"start": "2021-09-27T00:00:00.000Z",
								"tweet_count": 4
							},
							{
								"end": "2021-09-29T00:00:00.000Z",
								"start": "2021-09-28T00:00:00.000Z",
								"tweet_count": 3
							}
						],
						"meta": {
							"total_tweet_count": 7
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				query: "from:TwitterDev",
				opt: []*gotwtr.TweetCountsOption{
					{
						StartTime:   time.Now(),
						Granularity: "day",
						SinceID:     "0",
					},
				},
			},
			want: &gotwtr.TweetCountsResponse{
				Counts: []*gotwtr.TimeseriesCount{
					{
						Start:      "2021-09-27T00:00:00.000Z",
						End:        "2021-09-28T00:00:00.000Z",
						TweetCount: 4,
					},
					{
						Start:      "2021-09-28T00:00:00.000Z",
						End:        "2021-09-29T00:00:00.000Z",
						TweetCount: 3,
					},
				},
				Meta: &gotwtr.TweetCountMeta{
					TotalTweetCount: 7,
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
			got, err := c.CountsRecentTweet(tt.args.ctx, tt.args.query, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CountsRecentTweet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.CountsRecentTweet() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
