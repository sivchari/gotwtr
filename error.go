package gotwtr

type HTTPError struct {
	APIName string
	Status  string
	URL     string
}

func (e *HTTPError) Error() string {
	return e.APIName + ": " + e.Status + " " + e.URL
}

type APIResponseError struct {
	Title              string      `json:"title"`
	Detail             string      `json:"detail"`
	Type               string      `json:"type"`
	ResourceType       string      `json:"resource_type"`
	ResourceID         string      `json:"resource_id"`
	Parameter          string      `json:"parameter"`
	Parameters         Parameter   `json:"parameters"`
	Message            string      `json:"message"`
	Value              interface{} `json:"value"`
	Reason             string      `json:"reason,omitempty"`
	ClientID           string      `json:"client_id,omitempty"`
	RequiredEnrollment string      `json:"required_enrollment,omitempty"`
	RegistrationURL    string      `json:"registration_url,omitempty"`
	ConnectionIssue    string      `json:"connection_issue,omitempty"`
}

type Parameter struct {
	ID        []string `json:"id"`
	IDs       []string `json:"ids"`
	UserName  []string `json:"username"`
	UserNames []string `json:"usernames"`
}
