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

func Test_searchRecentTweets(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		query  string
		opt    []*gotwtr.TweetSearchOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.TweetSearchResponse
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
								"id": "1444834989024202755",
								"text": "RT @_ShamGod: This is most \"affordable housing\" deals in luxury buildings. Most of the  ones in NYC housing connect won't let you access th…"
							},
							{
								"id": "1444834986008457217",
								"text": "Harry en NYC - Octubre 3 #LoveOnTourNewYork #LOTNewYork #LOTNewYorkN1 https://t.co/up6J75RSpP"
							},
							{
								"id": "1444834985135927296",
								"text": "RT @ABC7: New statue of #GeorgeFloyd vandalized with paint in NYC; suspect flees on skateboard\nhttps://t.co/zTG1vvkG1b"
							},
							{
								"id": "1444834983231868930",
								"text": "RT @BTB_MikeII: 100 YEARS AGO AT THE POLO GROUNDS 10/3/1921: New York's Managers Plant Their League Championship Flags... https://t.co/SR69…"
							},
							{
								"id": "1444834982023901189",
								"text": "Why am I not surprised ? When you have parents teaching their children to be #Racists … allegedly and systematic Racism is rampant! #NYC #TEACHING #PARENTS #SchoolStrike2021 #NYCDOE #QUEENS #GeorgeFloyd #Homeschool #Teachers #nyc #GENTRIFICATION #Progressives https://t.co/PiPaL70dqG"
							},
							{
								"id": "1444834981579087873",
								"text": "RT @amyjharris: NEW INVESTIGATION: As homelessness has spiked in New York City, the people who run nonprofit shelters have found ways to en…"
							},
							{
								"id": "1444834976848039941",
								"text": "RT @EllenBarkin: Well done NYC!"
							},
							{
								"id": "1444834976143273986",
								"text": "RT @hernameiskayIa: Nyc and la shows always win idc"
							},
							{
								"id": "1444834975883407361",
								"text": "go✍️ to✍️ a✍️ harry✍️ nyc✍️ show✍️"
							},
							{
								"id": "1444834974646116359",
								"text": "RT @elhammohamud: girlies be like nyc is not ready for us like nyc has never seen 10 harry styles fans in head to toe shein walking around…"
							}
						],
						"meta": {
							"newest_id": "1444834989024202755",
							"oldest_id": "1444834974646116359",
							"result_count": 10,
							"next_token": "b26v89c19zqg8o3fpds7phdmycbzydb96f7kaqov4kuwt"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				query: "nyc",
				opt:   []*gotwtr.TweetSearchOption{},
			},
			want: &gotwtr.TweetSearchResponse{
				Tweets: []*gotwtr.Tweet{
					{
						ID:   "1444834989024202755",
						Text: "RT @_ShamGod: This is most \"affordable housing\" deals in luxury buildings. Most of the  ones in NYC housing connect won't let you access th…",
					},
					{
						ID:   "1444834986008457217",
						Text: "Harry en NYC - Octubre 3 #LoveOnTourNewYork #LOTNewYork #LOTNewYorkN1 https://t.co/up6J75RSpP",
					},
					{
						ID:   "1444834985135927296",
						Text: "RT @ABC7: New statue of #GeorgeFloyd vandalized with paint in NYC; suspect flees on skateboard\nhttps://t.co/zTG1vvkG1b",
					},
					{
						ID:   "1444834983231868930",
						Text: "RT @BTB_MikeII: 100 YEARS AGO AT THE POLO GROUNDS 10/3/1921: New York's Managers Plant Their League Championship Flags... https://t.co/SR69…",
					},
					{
						ID:   "1444834982023901189",
						Text: "Why am I not surprised ? When you have parents teaching their children to be #Racists … allegedly and systematic Racism is rampant! #NYC #TEACHING #PARENTS #SchoolStrike2021 #NYCDOE #QUEENS #GeorgeFloyd #Homeschool #Teachers #nyc #GENTRIFICATION #Progressives https://t.co/PiPaL70dqG",
					},
					{
						ID:   "1444834981579087873",
						Text: "RT @amyjharris: NEW INVESTIGATION: As homelessness has spiked in New York City, the people who run nonprofit shelters have found ways to en…",
					},
					{
						ID:   "1444834976848039941",
						Text: "RT @EllenBarkin: Well done NYC!",
					},
					{
						ID:   "1444834976143273986",
						Text: "RT @hernameiskayIa: Nyc and la shows always win idc",
					},
					{
						ID:   "1444834975883407361",
						Text: "go✍️ to✍️ a✍️ harry✍️ nyc✍️ show✍️",
					},
					{
						ID:   "1444834974646116359",
						Text: "RT @elhammohamud: girlies be like nyc is not ready for us like nyc has never seen 10 harry styles fans in head to toe shein walking around…",
					},
				},
				Meta: &gotwtr.TweetSearchMeta{
					NewestID:    "1444834989024202755",
					OldestID:    "1444834974646116359",
					ResultCount: 10,
					NextToken:   "b26v89c19zqg8o3fpds7phdmycbzydb96f7kaqov4kuwt",
				},
			},
			wantErr: false,
		},
		// TODO: Add test cases with request fieds (tweet, user, etc.)
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.SearchRecentTweets(tt.args.ctx, tt.args.query, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.SearchRecentTweets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.SearchRecentTweets() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
