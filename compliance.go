package gotwtr

type ComplianceFieldType string

const (
	ComplianceFieldTypeTweets ComplianceFieldType = "tweets"
	ComplianceFieldTypeUsers  ComplianceFieldType = "users"
)

type ComplianceFieldStatus string

const (
	ComplianseFieldStatusCreated    ComplianceFieldStatus = "created"
	ComplianseFieldStatusInProgress ComplianceFieldStatus = "in_progress"
	ComplianseFieldStatusFailed     ComplianceFieldStatus = "failed"
	ComplianseFieldStatusCompletae  ComplianceFieldStatus = "complete"
)

type ComplianceJobsResponse struct {
	ComplianceJobsData []*ComplianceJobsData `json:"data"`
	Errors             []*APIResponseError   `json:"errors"`
}

type ComplianceJobsData struct {
	ID                string `json:"id"`
	CreatedAt         string `json:"created_at"`
	Type              string `json:"type"`
	Name              string `json:"name"`
	UploadURL         string `json:"upload_url"`
	UploadExpiresAt   string `json:"upload_expires_at"`
	DownloadURL       string `json:"download_url"`
	DownloadExpiresAt string `json:"download_expires_at"`
	Status            string `json:"status"`
	Resumable         bool   `json:"resumable"`
	Error             string `json:"error,omitempty"`
}
