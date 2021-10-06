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
	Reason             string      `json:"reason"`
	ClientID           string      `json:"client_id"`
	RequiredEnrollment string      `json:"required_enrollment"`
	RegistrationURL    string      `json:"registration_url"`
}

type Parameter struct {
	IDs []string `json:"ids"`
}
