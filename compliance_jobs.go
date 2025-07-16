package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func complianceJobs(ctx context.Context, c *client, opt ...*ComplianceJobsOption) (*ComplianceJobsResponse, error) {
	var copt ComplianceJobsOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		copt = *opt[0]
	default:
		return nil, errors.New("compliance jobs: only one option is allowed")
	}

	if copt.Type == "" {
		return nil, errors.New("compliance jobs: type parameter is required")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, complianceJobsURL, nil)
	if err != nil {
		return nil, fmt.Errorf("compliance jobs new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	copt.addQuery(req)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("compliance jobs response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

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

func complianceJob(ctx context.Context, c *client, cjID int) (*ComplianceJobResponse, error) {
	ep := fmt.Sprintf(complianceJobURL, cjID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("compliance job new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("compliance job response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cj ComplianceJobResponse
	if err := json.NewDecoder(resp.Body).Decode(&cj); err != nil {
		return nil, fmt.Errorf("compliance job: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &cj, &HTTPError{
			APIName: "compliance job",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}
	return &cj, nil
}

func createComplianceJob(ctx context.Context, c *client, opt ...*CreateComplianceJobOption) (*CreateComplianceJobResponse, error) {
	var copt CreateComplianceJobOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		copt = *opt[0]
	default:
		return nil, errors.New("create compliance job: only one option is allowed")
	}
	if copt.Type == "" {
		return nil, errors.New("create compliance job: type parameter is required")
	}
	j, err := json.Marshal(copt)
	if err != nil {
		return nil, fmt.Errorf("create compliance job: can not marshal: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, createComplianceJobURL, bytes.NewBuffer(j))
	if err != nil {
		return nil, fmt.Errorf("create compliance job new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("create compliance job response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cresponse CreateComplianceJobResponse
	if err := json.NewDecoder(resp.Body).Decode(&cresponse); err != nil {
		return nil, fmt.Errorf("create compliance job: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &cresponse, &HTTPError{
			APIName: "create compliance job",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}
	return &cresponse, nil
}
