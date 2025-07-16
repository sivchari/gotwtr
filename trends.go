package gotwtr

// Trend represents a trending topic
type Trend struct {
	TrendName  string `json:"trend_name"`
	TweetCount int    `json:"tweet_count,omitempty"`
}

// TrendsByWOEIDResponse represents the response for trends by WOEID lookup
type TrendsByWOEIDResponse struct {
	Trends []*Trend            `json:"data,omitempty"`
	Errors []*APIResponseError `json:"errors,omitempty"`
	Title  string              `json:"title,omitempty"`
	Detail string              `json:"detail,omitempty"`
	Type   string              `json:"type,omitempty"`
}