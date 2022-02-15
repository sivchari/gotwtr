package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func complianceJobs(ctx context.Context, c *client, opt *ComplianceJobsOption) (*ComplianceJobsResponse, error) {
	if opt.Type == "" {
		return nil, errors.New("compliance jobs: type parameter is required")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, complianceJobsURL, nil)
	if err != nil {
		return nil, fmt.Errorf("compliance jobs new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	opt.addQuery(req)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("compliance jobs response: %w", err)
	}
	defer resp.Body.Close()

	var cj ComplianceJobsResponse
	if err := json.NewDecoder(resp.Body).Decode(&cj); err != nil {
		return nil, fmt.Errorf("compliance jobs: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &cj, &HTTPError{
			APIName: "compliance jobs",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}
	return &cj, nil
}
