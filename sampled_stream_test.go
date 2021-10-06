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

func Test_sampledStream(t *testing.T) {
	type args struct {
		ctx    context.Context
		client *http.Client
		opt    []*gotwtr.SampledStreamOpts
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.SampledStreamResponse
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
								"author_id": "2244994945",
								"created_at": "2020-02-14T19:00:55.000Z",
								"id": "1228393702244134912",
								"text": "What did the developer write in their Valentine’s card?\n  \nwhile(true) {\n    I = Love(You);  \n}"
							},
							{
								"author_id": "2244994945",
								"created_at": "2020-02-12T17:09:56.000Z",
								"id": "1227640996038684673",
								"text": "Doctors: Googling stuff online does not make you a doctor\n\nDevelopers: https://t.co/mrju5ypPkb"
							},
							{
								"author_id": "2244994945",
								"created_at": "2019-11-27T20:26:41.000Z",
								"id": "1199786642791452673",
								"text": "C#"
							}
						],
						"includes": {
							"users": [
								{
									"created_at": "2013-12-14T04:35:55.000Z",
									"id": "2244994945",
									"name": "Twitter Dev",
									"username": "TwitterDev"
								}
							]
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				opt:   []*gotwtr.SampledStreamOpts{},
			},
			want: &gotwtr.SampledStreamResponse{
				Tweets: []*gotwtr.Tweet{
					{
						AuthorID: "2244994945",
						CreatedAt: "2020-02-14T19:00:55.000Z",
						ID: "1228393702244134912",
						Text: "What did the developer write in their Valentine’s card?\n  \nwhile(true) {\n    I = Love(You);  \n}",
					},
					{
						AuthorID: "2244994945",
						CreatedAt: "2020-02-12T17:09:56.000Z",
						ID: "1227640996038684673",
						Text: "Doctors: Googling stuff online does not make you a doctor\n\nDevelopers: https://t.co/mrju5ypPkb",
					},
					{
						AuthorID: "2244994945",
						CreatedAt: "2019-11-27T20:26:41.000Z",
						ID: "1199786642791452673",
						Text: "C#",
					},
				},
				Includes: &gotwtr.TweetIncludes{
					Users: []*gotwtr.User{
						{
							CreatedAt: "2013-12-14T04:35:55.000Z",
							ID: "2244994945",
							Name: "Twitter Dev",
							UserName: "TwitterDev",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.SampledStream(tt.args.ctx, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SampledStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("SampledStream() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}