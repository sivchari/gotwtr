package gotwtr

import "net/http"

type ComplianceJobsOption struct {
	Type   ComplianceFieldType
	Status ComplianceFieldStatus
}

func (c *ComplianceJobsOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	q.Add("type", string(c.Type))
	if c.Status != "" {
		q.Add("status", string(c.Status))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

type CreateComplianceJobOption struct {
	Type      ComplianceFieldType `json:"type"`
	Name      string              `json:"name,omitempty"`
	Resumable bool                `json:"resumable,omitempty"`
}
