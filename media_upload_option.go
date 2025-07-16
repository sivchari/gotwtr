package gotwtr

// MediaUploadOption represents options for simple media upload
type MediaUploadOption struct {
	MediaCategory    string   `json:"media_category,omitempty"`    // e.g., "tweet_image", "tweet_video", "dm_image", "dm_video"
	AdditionalOwners []string `json:"additional_owners,omitempty"` // Additional user IDs with upload permissions
	AltText          string   `json:"alt_text,omitempty"`          // Alt text for accessibility
}

// MediaCategory constants for different media types
const (
	MediaCategoryTweetImage = "tweet_image"
	MediaCategoryTweetVideo = "tweet_video"
	MediaCategoryTweetGIF   = "tweet_gif"
	MediaCategoryDMImage    = "dm_image"
	MediaCategoryDMVideo    = "dm_video"
	MediaCategoryDMGIF      = "dm_gif"
)

// Common MIME types for media upload
const (
	MediaTypeJPEG = "image/jpeg"
	MediaTypePNG  = "image/png"
	MediaTypeGIF  = "image/gif"
	MediaTypeWebP = "image/webp"
	MediaTypeMP4  = "video/mp4"
	MediaTypeMOV  = "video/quicktime"
	MediaTypeAVI  = "video/x-msvideo"
	MediaTypeWEBM = "video/webm"
)

// ChunkedUploadThreshold defines the file size threshold for using chunked upload (5MB)
const ChunkedUploadThreshold = 5 * 1024 * 1024

// MaxChunkSize defines the maximum size of each upload chunk (5MB)
const MaxChunkSize = 5 * 1024 * 1024