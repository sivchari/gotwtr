package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	tweetCountsRecent = "https://api.twitter.com/2/tweets/counts/recent?query="
)

type TweetCountsOption struct {
	StartTime   string
	EndTime     string
	SinceId     string
	UntilId     string
	Granularity string
}

const MaxQueryLength int = 512

type TweetCount struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	TweetCount int    `json:"tweet_count"`
}

type MetaField struct {
	TotalTweetCount int `json:"total_tweet_count"`
}

type TweetCountsResponse struct {
	Tweets []*TweetCount `json:"data"`
	Meta   *MetaField    `json:"meta"`
}

func (t *TweetCountsOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(t.StartTime) > 0 {
		q.Add("expansions", t.StartTime)
	}
	if len(t.EndTime) > 0 {
		q.Add("media.fields", t.EndTime)
	}
	if len(t.SinceId) > 0 {
		q.Add("place.fields", t.SinceId)
	}
	if len(t.UntilId) > 0 {
		q.Add("poll.fields", t.UntilId)
	}
	if len(t.Granularity) > 0 {
		q.Add("tweet.fields", t.Granularity)
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

// query specified search parameter, for example, if you want to understand tweet that contains "golang" volume,
// query is "golang" .
func tweetCounts(ctx context.Context, c *client, query string, opt ...*TweetCountsOption) (*TweetCountsResponse, error) {
	switch {
	case len(query) == 0:
		return nil, errors.New("tweet counts: query parameter is required")
	case MaxQueryLength < len(query):
		return nil, errors.New("tweet counts: query parameter length is less than 512")
	default:
	}

	tweetCountsRecent += fmt.Sprintf("%s ", query)

	/*
		// join ids to url
		for i, q := range query {
			if i+1 < len(query) {
				tweetCountsRecent += fmt.Sprintf("%s ", q)
			} else {
				tweetCountsRecent += q
			}
		}
	*/

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, tweetCountsRecent, nil)
	if err != nil {
		return nil, fmt.Errorf("tweet counts new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt TweetCountsOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("tweet counts: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("tweet counts response: %w", err)
	}
	defer resp.Body.Close()

	var tweet TweetCountsResponse
	if err := json.NewDecoder(resp.Body).Decode(&tweet); err != nil {
		return nil, fmt.Errorf("tweet counts decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &tweet, &HTTPError{
			APIName: "tweet counts",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &tweet, nil
}
