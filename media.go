package gotwtr

type MediaField string

/*
	Media field will only return
	if you've also included the expansions=attachments.media_keys query parameter in your request.
*/
const (
	MediaFieldDurationMS       MediaField = "duration_ms"
	MediaFieldHeight           MediaField = "height"
	MediaFieldMediaKey         MediaField = "media_key"
	MediaFieldPreviewImageURL  MediaField = "preview_image_url"
	MediaFieldType             MediaField = "type"
	MediaFieldURL              MediaField = "url"
	MediaFieldWidth            MediaField = "width"
	MediaFieldPublicMetrics    MediaField = "public_metrics"
	MediaFieldNonPublicMetrics MediaField = "non_public_metrics"
	MediaFieldOrganicMetrics   MediaField = "organic_metrics"
	MediaFieldPromotedMetrics  MediaField = "promoted_metrics"
	MediaFieldAltText          MediaField = "alt_text"
)

type Media struct {
	MediaKey         string        `json:"media_key"`
	Type             string        `json:"type"`
	URL              string        `json:"url,omitempty"`
	DurationMs       int           `json:"duration_ms,omitempty"`
	Height           int           `json:"height,omitempty"`
	NonPublicMetrics *MediaMetrics `json:"non_public_metrics,omitempty"`
	OrganicMetrics   *MediaMetrics `json:"organic_metrics,omitempty"`
	PreviewImageURL  string        `json:"preview_image_url,omitempty"`
	PromotedMetrics  *MediaMetrics `json:"promoted_metrics,omitempty"`
	PublicMetrics    *MediaMetrics `json:"public_metrics,omitempty"`
	Width            int           `json:"width,omitempty"`
	AltText          string        `json:"alt_text,omitempty"`
	MediaIDs         []string      `json:"media_ids,omitempty"`
	TaggedUserIDs    []string      `json:"tagged_user_ids,omitempty"`
}

type MediaMetrics struct {
	Playback0Count   int `json:"playback_0_count,omitempty"`
	Playback25Count  int `json:"playback_25_count,omitempty"`
	Playback50Count  int `json:"playback_50_count,omitempty"`
	Playback75Count  int `json:"playback_75_count,omitempty"`
	Playback100Count int `json:"playback_100_count,omitempty"`
	ViewCount        int `json:"view_count,omitempty"`
}

func mediaFieldsToString(mfs []MediaField) []string {
	slice := make([]string, len(mfs))
	for i, mf := range mfs {
		slice[i] = string(mf)
	}
	return slice
}
