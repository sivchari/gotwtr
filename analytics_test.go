package gotwtr_test

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/sivchari/gotwtr"
)

func TestCalculateEngagementRate(t *testing.T) {
	t.Parallel()
	type args struct {
		metrics *gotwtr.TweetMetrics
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "normal engagement calculation",
			args: args{
				metrics: &gotwtr.TweetMetrics{
					ImpressionCount: 1000,
					LikeCount:       50,
					RetweetCount:    20,
					ReplyCount:      10,
					QuoteCount:      5,
					BookmarkCount:   3,
				},
			},
			want: 8.8, // (50+20+10+5+3)/1000 * 100 = 8.8%
		},
		{
			name: "zero impressions",
			args: args{
				metrics: &gotwtr.TweetMetrics{
					ImpressionCount: 0,
					LikeCount:       10,
					RetweetCount:    5,
				},
			},
			want: 0.0,
		},
		{
			name: "nil metrics",
			args: args{
				metrics: nil,
			},
			want: 0.0,
		},
		{
			name: "no engagements",
			args: args{
				metrics: &gotwtr.TweetMetrics{
					ImpressionCount: 1000,
					LikeCount:       0,
					RetweetCount:    0,
					ReplyCount:      0,
					QuoteCount:      0,
					BookmarkCount:   0,
				},
			},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("key")
			got := c.CalculateEngagementRate(tt.args.metrics)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("CalculateEngagementRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateEngagementMetrics(t *testing.T) {
	t.Parallel()
	type args struct {
		metrics *gotwtr.TweetMetrics
	}
	tests := []struct {
		name string
		args args
		want *gotwtr.EngagementMetrics
	}{
		{
			name: "detailed engagement metrics",
			args: args{
				metrics: &gotwtr.TweetMetrics{
					ImpressionCount:    1000,
					LikeCount:          50,
					RetweetCount:       20,
					ReplyCount:         10,
					QuoteCount:         5,
					BookmarkCount:      3,
					URLLinkClicks:      15,
					UserProfileClicks:  8,
				},
			},
			want: &gotwtr.EngagementMetrics{
				EngagementRate:   8.8,  // (50+20+10+5+3)/1000 * 100
				LikeRate:         5.0,  // 50/1000 * 100
				RetweetRate:      2.0,  // 20/1000 * 100
				ReplyRate:        1.0,  // 10/1000 * 100
				QuoteRate:        0.5,  // 5/1000 * 100
				BookmarkRate:     0.3,  // 3/1000 * 100
				ClickThroughRate: 2.3,  // (15+8)/1000 * 100
			},
		},
		{
			name: "nil metrics",
			args: args{
				metrics: nil,
			},
			want: &gotwtr.EngagementMetrics{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("key")
			got := c.CalculateEngagementMetrics(tt.args.metrics)
			opt := cmp.Comparer(func(a, b float64) bool {
				return math.Abs(a-b) < 1e-9
			})
			if diff := cmp.Diff(tt.want, got, opt); diff != "" {
				t.Errorf("CalculateEngagementMetrics() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetTopPerformingTweets(t *testing.T) {
	t.Parallel()
	type args struct {
		tweets []*gotwtr.Tweet
		limit  int
	}
	tests := []struct {
		name string
		args args
		want []*gotwtr.Tweet
	}{
		{
			name: "sort tweets by engagement rate",
			args: args{
				tweets: []*gotwtr.Tweet{
					{
						ID:   "1",
						Text: "Low engagement tweet",
						PublicMetrics: &gotwtr.TweetMetrics{
							ImpressionCount: 1000,
							LikeCount:       10,
							RetweetCount:    5,
						},
					},
					{
						ID:   "2",
						Text: "High engagement tweet",
						PublicMetrics: &gotwtr.TweetMetrics{
							ImpressionCount: 1000,
							LikeCount:       100,
							RetweetCount:    50,
						},
					},
					{
						ID:   "3",
						Text: "Medium engagement tweet",
						PublicMetrics: &gotwtr.TweetMetrics{
							ImpressionCount: 1000,
							LikeCount:       50,
							RetweetCount:    25,
						},
					},
				},
				limit: 2,
			},
			want: []*gotwtr.Tweet{
				{
					ID:   "2",
					Text: "High engagement tweet",
					PublicMetrics: &gotwtr.TweetMetrics{
						ImpressionCount: 1000,
						LikeCount:       100,
						RetweetCount:    50,
					},
				},
				{
					ID:   "3",
					Text: "Medium engagement tweet",
					PublicMetrics: &gotwtr.TweetMetrics{
						ImpressionCount: 1000,
						LikeCount:       50,
						RetweetCount:    25,
					},
				},
			},
		},
		{
			name: "empty tweets slice",
			args: args{
				tweets: []*gotwtr.Tweet{},
				limit:  5,
			},
			want: []*gotwtr.Tweet{},
		},
		{
			name: "limit larger than tweets count",
			args: args{
				tweets: []*gotwtr.Tweet{
					{
						ID:   "1",
						Text: "Only tweet",
						PublicMetrics: &gotwtr.TweetMetrics{
							ImpressionCount: 1000,
							LikeCount:       10,
						},
					},
				},
				limit: 5,
			},
			want: []*gotwtr.Tweet{
				{
					ID:   "1",
					Text: "Only tweet",
					PublicMetrics: &gotwtr.TweetMetrics{
						ImpressionCount: 1000,
						LikeCount:       10,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("key")
			got := c.GetTopPerformingTweets(tt.args.tweets, tt.args.limit)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("GetTopPerformingTweets() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCompareMetrics(t *testing.T) {
	t.Parallel()
	type args struct {
		current  *gotwtr.TweetMetrics
		previous *gotwtr.TweetMetrics
	}
	tests := []struct {
		name string
		args args
		want *gotwtr.AnalyticsComparison
	}{
		{
			name: "positive growth",
			args: args{
				current: &gotwtr.TweetMetrics{
					ImpressionCount: 1000,
					LikeCount:       60,  // 6% engagement rate
					RetweetCount:    0,
					ReplyCount:      0,
					QuoteCount:      0,
					BookmarkCount:   0,
				},
				previous: &gotwtr.TweetMetrics{
					ImpressionCount: 1000,
					LikeCount:       50,  // 5% engagement rate
					RetweetCount:    0,
					ReplyCount:      0,
					QuoteCount:      0,
					BookmarkCount:   0,
				},
			},
			want: &gotwtr.AnalyticsComparison{
				Current:  6.0,
				Previous: 5.0,
				Change:   20.0, // ((6-5)/5)*100 = 20%
			},
		},
		{
			name: "negative growth",
			args: args{
				current: &gotwtr.TweetMetrics{
					ImpressionCount: 1000,
					LikeCount:       40,  // 4% engagement rate
					RetweetCount:    0,
					ReplyCount:      0,
					QuoteCount:      0,
					BookmarkCount:   0,
				},
				previous: &gotwtr.TweetMetrics{
					ImpressionCount: 1000,
					LikeCount:       50,  // 5% engagement rate
					RetweetCount:    0,
					ReplyCount:      0,
					QuoteCount:      0,
					BookmarkCount:   0,
				},
			},
			want: &gotwtr.AnalyticsComparison{
				Current:  4.0,
				Previous: 5.0,
				Change:   -20.0, // ((4-5)/5)*100 = -20%
			},
		},
		{
			name: "nil current metrics",
			args: args{
				current: nil,
				previous: &gotwtr.TweetMetrics{
					ImpressionCount: 1000,
					LikeCount:       50,
				},
			},
			want: nil,
		},
		{
			name: "nil previous metrics",
			args: args{
				current: &gotwtr.TweetMetrics{
					ImpressionCount: 1000,
					LikeCount:       50,
				},
				previous: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("key")
			got := c.CompareMetrics(tt.args.current, tt.args.previous)
			opt := cmp.Comparer(func(a, b float64) bool {
				return math.Abs(a-b) < 1e-9
			})
			if diff := cmp.Diff(tt.want, got, opt); diff != "" {
				t.Errorf("CompareMetrics() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}