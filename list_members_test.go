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

func Test_listMembers(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		id     string
		opt    []*gotwtr.ListMembersOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.ListMembersResponse
		wantErr bool
	}{
		{
			name: "200 ok default payload",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"data": [
							{
								"id": "1319036828964454402",
								"name": "Birdwatch",
								"username": "birdwatch"
							},
							{
								"id": "1244731491088809984",
								"name": "Twitter Thailand",
								"username": "TwitterThailand"
							},
							{
								"id": "1194267639100723200",
								"name": "Twitter Retweets",
								"username": "TwitterRetweets"
							},
							{
								"id": "1168976680867762177",
								"name": "Twitter Able",
								"username": "TwitterAble"
							},
							{
								"id": "1065249714214457345",
								"name": "Spaces",
								"username": "TwitterSpaces"
							}
						],
						"meta": {
							"result_count": 5,
							"next_token": "5676935732641845249"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				id:  "84839422",
				opt: []*gotwtr.ListMembersOption{},
			},
			want: &gotwtr.ListMembersResponse{
				Users: []*gotwtr.User{
					{
						ID:       "1319036828964454402",
						Name:     "Birdwatch",
						UserName: "birdwatch",
					},
					{
						ID:       "1244731491088809984",
						Name:     "Twitter Thailand",
						UserName: "TwitterThailand",
					},
					{
						ID:       "1194267639100723200",
						Name:     "Twitter Retweets",
						UserName: "TwitterRetweets",
					},
					{
						ID:       "1168976680867762177",
						Name:     "Twitter Able",
						UserName: "TwitterAble",
					},
					{
						ID:       "1065249714214457345",
						Name:     "Spaces",
						UserName: "TwitterSpaces",
					},
				},
				Meta: &gotwtr.ListMeta{
					ResultCount: 5,
					NextToken:   "5676935732641845249",
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
			got, err := c.ListMembers(tt.args.ctx, tt.args.id, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.ListMembers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.ListMembers() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
