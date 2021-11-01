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
	Title        string      `json:"title"`
	Detail       string      `json:"detail"`
	Type         string      `json:"type"`
	ResourceType string      `json:"resource_type"`
	ResourceID   string      `json:"resource_id"`
	Parameter    string      `json:"parameter"`
	Parameters   Parameter   `json:"parameters"`
	Message      string      `json:"message"`
	Value        interface{} `json:"value"`
}

type Parameter struct {
	ID  []string `json:"id"`
	IDs []string `json:"ids"`
	UserName []string `json:"username"`
	UserNames []string `json:"usernames"`
}
