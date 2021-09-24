package field

type Media struct {
	MediaKey         string        `json:"media_key"`
	Type             string        `json:"type"`
	DurationMs       int           `json:"duration_ms,omitempty"`
	Height           int           `json:"height,omitempty"`
	NonPublicMetrics *MediaMetrics `json:"non_public_metrics,omitempty"`
	OrganicMetrics   *MediaMetrics `json:"organic_metrics,omitempty"`
	PreviewImageURL  string        `json:"preview_image_url,omitempty"`
	PromotedMetrics  *MediaMetrics `json:"promoted_metrics,omitempty"`
	PublicMetrics    *MediaMetrics `json:"public_metrics,omitempty"`
	Width            int           `json:"width,omitempty"`
	AltText          string        `json:"alt_text,omitempty"`
}

type MediaMetrics struct {
	Playback0Count   int `json:"playback_0_count,omitempty"`
	Playback25Count  int `json:"playback_25_count,omitempty"`
	Playback50Count  int `json:"playback_50_count,omitempty"`
	Playback75Count  int `json:"playback_75_count,omitempty"`
	Playback100Count int `json:"playback_100_count,omitempty"`
	ViewCount        int `json:"view_count,omitempty"`
}
