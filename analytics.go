package gotwtr

import "time"

// SpaceAnalytics represents extended analytics for Spaces
type SpaceAnalytics struct {
	LiveListenerCount    int    `json:"live_listener_count,omitempty"`
	TotalListenerCount   int    `json:"total_listener_count,omitempty"`
	SpeakerRequestCount  int    `json:"speaker_request_count,omitempty"`
	RecordingDownloads   int    `json:"recording_downloads,omitempty"`
	ShareCount           int    `json:"share_count,omitempty"`
	Duration             string `json:"duration,omitempty"`
}

// ListAnalytics represents analytics for Lists
type ListAnalytics struct {
	ViewCount        int     `json:"view_count,omitempty"`
	MemberGrowthRate float64 `json:"member_growth_rate,omitempty"`
	TweetCount       int     `json:"tweet_count,omitempty"`
	EngagementRate   float64 `json:"engagement_rate,omitempty"`
}

// MediaAnalytics represents analytics for Media objects
type MediaAnalytics struct {
	ViewCount      int `json:"view_count,omitempty"`
	PlayCount      int `json:"play_count,omitempty"`
	DownloadCount  int `json:"download_count,omitempty"`
	ShareCount     int `json:"share_count,omitempty"`
	CompletionRate int `json:"completion_rate,omitempty"` // Percentage for videos
}

// EngagementMetrics provides calculated engagement metrics
type EngagementMetrics struct {
	EngagementRate    float64 `json:"engagement_rate"`
	LikeRate         float64 `json:"like_rate"`
	RetweetRate      float64 `json:"retweet_rate"`
	ReplyRate        float64 `json:"reply_rate"`
	QuoteRate        float64 `json:"quote_rate"`
	BookmarkRate     float64 `json:"bookmark_rate"`
	ClickThroughRate float64 `json:"click_through_rate"`
}

// AnalyticsTimeframe represents time-based analytics filters
type AnalyticsTimeframe struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Granularity string  `json:"granularity"` // day, hour, minute
}

// AnalyticsComparison represents comparative analytics data
type AnalyticsComparison struct {
	Current  interface{} `json:"current"`
	Previous interface{} `json:"previous"`
	Change   float64     `json:"change_percentage"`
}

// TweetAnalyticsSummary provides aggregated tweet analytics
type TweetAnalyticsSummary struct {
	TotalTweets      int                `json:"total_tweets"`
	TotalImpressions int                `json:"total_impressions"`
	TotalEngagements int                `json:"total_engagements"`
	AverageMetrics   *TweetMetrics      `json:"average_metrics"`
	TopPerforming    []*Tweet          `json:"top_performing"`
	Timeframe        *AnalyticsTimeframe `json:"timeframe"`
}

// UserAnalyticsSummary provides aggregated user analytics
type UserAnalyticsSummary struct {
	FollowerGrowth   *AnalyticsComparison `json:"follower_growth"`
	EngagementTrends *EngagementMetrics   `json:"engagement_trends"`
	TopTweets        []*Tweet            `json:"top_tweets"`
	Timeframe        *AnalyticsTimeframe  `json:"timeframe"`
}

// Analytics utility functions

// CalculateEngagementRate calculates engagement rate from public metrics
func CalculateEngagementRate(metrics *TweetMetrics) float64 {
	if metrics == nil || metrics.ImpressionCount == 0 {
		return 0.0
	}
	
	totalEngagements := metrics.LikeCount + metrics.RetweetCount + 
					   metrics.ReplyCount + metrics.QuoteCount + metrics.BookmarkCount
	
	return (float64(totalEngagements) / float64(metrics.ImpressionCount)) * 100
}

// CalculateClickThroughRate calculates CTR from metrics
func CalculateClickThroughRate(metrics *TweetMetrics) float64 {
	if metrics == nil || metrics.ImpressionCount == 0 {
		return 0.0
	}
	
	totalClicks := metrics.URLLinkClicks + metrics.UserProfileClicks
	
	return (float64(totalClicks) / float64(metrics.ImpressionCount)) * 100
}

// CalculateEngagementMetrics calculates detailed engagement metrics
func CalculateEngagementMetrics(metrics *TweetMetrics) *EngagementMetrics {
	if metrics == nil || metrics.ImpressionCount == 0 {
		return &EngagementMetrics{}
	}
	
	impressions := float64(metrics.ImpressionCount)
	
	return &EngagementMetrics{
		EngagementRate:   CalculateEngagementRate(metrics),
		LikeRate:         (float64(metrics.LikeCount) / impressions) * 100,
		RetweetRate:      (float64(metrics.RetweetCount) / impressions) * 100,
		ReplyRate:        (float64(metrics.ReplyCount) / impressions) * 100,
		QuoteRate:        (float64(metrics.QuoteCount) / impressions) * 100,
		BookmarkRate:     (float64(metrics.BookmarkCount) / impressions) * 100,
		ClickThroughRate: CalculateClickThroughRate(metrics),
	}
}

// CompareMetrics compares two sets of metrics and returns change percentage
func CompareMetrics(current, previous *TweetMetrics) *AnalyticsComparison {
	if current == nil || previous == nil {
		return nil
	}
	
	currentEngagement := CalculateEngagementRate(current)
	previousEngagement := CalculateEngagementRate(previous)
	
	var change float64
	if previousEngagement != 0 {
		change = ((currentEngagement - previousEngagement) / previousEngagement) * 100
	}
	
	return &AnalyticsComparison{
		Current:  currentEngagement,
		Previous: previousEngagement,
		Change:   change,
	}
}

// GetTopPerformingTweets sorts tweets by engagement rate
func GetTopPerformingTweets(tweets []*Tweet, limit int) []*Tweet {
	if len(tweets) == 0 {
		return []*Tweet{}
	}
	
	// Create a copy to avoid modifying the original slice
	sortedTweets := make([]*Tweet, len(tweets))
	copy(sortedTweets, tweets)
	
	// Simple sort by engagement rate (this could be optimized with proper sorting)
	for i := 0; i < len(sortedTweets)-1; i++ {
		for j := i + 1; j < len(sortedTweets); j++ {
			rate1 := CalculateEngagementRate(sortedTweets[i].PublicMetrics)
			rate2 := CalculateEngagementRate(sortedTweets[j].PublicMetrics)
			if rate2 > rate1 {
				sortedTweets[i], sortedTweets[j] = sortedTweets[j], sortedTweets[i]
			}
		}
	}
	
	if limit > 0 && limit < len(sortedTweets) {
		return sortedTweets[:limit]
	}
	
	return sortedTweets
}