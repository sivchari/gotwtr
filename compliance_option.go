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
}
