package gotwtr

// MediaUploadResponse represents the response from media upload operations
type MediaUploadResponse struct {
	MediaID         string               `json:"media_id"`
	MediaKey        string               `json:"media_key,omitempty"`
	ExpiresAfterSecs int                 `json:"expires_after_secs,omitempty"`
	ProcessingInfo  *MediaProcessingInfo `json:"processing_info,omitempty"`
	Errors          []*APIResponseError  `json:"errors,omitempty"`
	Title           string               `json:"title,omitempty"`
	Detail          string               `json:"detail,omitempty"`
	Type            string               `json:"type,omitempty"`
}

// MediaProcessingInfo represents media processing status
type MediaProcessingInfo struct {
	State          string                 `json:"state"`          // pending, in_progress, failed, succeeded
	CheckAfterSecs int                    `json:"check_after_secs,omitempty"`
	ProgressPercent int                   `json:"progress_percent,omitempty"`
	Error          *MediaProcessingError  `json:"error,omitempty"`
}

// MediaProcessingError represents errors during media processing
type MediaProcessingError struct {
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

// MediaUploadInitRequest represents the INIT command for chunked upload
type MediaUploadInitRequest struct {
	Command         string   `json:"command"`          // "INIT"
	MediaType       string   `json:"media_type"`       // MIME type (e.g., "image/jpeg", "video/mp4")
	TotalBytes      int64    `json:"total_bytes"`
	MediaCategory   string   `json:"media_category,omitempty"`   // e.g., "tweet_image", "tweet_video", "dm_image"
	AdditionalOwners []string `json:"additional_owners,omitempty"`
}

// MediaUploadAppendRequest represents the APPEND command for chunked upload
type MediaUploadAppendRequest struct {
	Command      string `json:"command"`       // "APPEND"
	MediaID      string `json:"media_id"`
	SegmentIndex int    `json:"segment_index"`
	Media        []byte `json:"media"`         // File chunk data
}

// MediaUploadFinalizeRequest represents the FINALIZE command for chunked upload
type MediaUploadFinalizeRequest struct {
	Command string `json:"command"`  // "FINALIZE"
	MediaID string `json:"media_id"`
}

// MediaUploadStatusRequest represents the STATUS command to check upload progress
type MediaUploadStatusRequest struct {
	Command string `json:"command"`  // "STATUS"
	MediaID string `json:"media_id"`
}