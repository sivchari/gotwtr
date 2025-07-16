package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func searchPostsEligibleForNotes(ctx context.Context, c *client, opt ...*SearchPostsEligibleForNotesOption) (*SearchPostsEligibleForNotesResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, searchPostsEligibleForNotesURL, nil)
	if err != nil {
		return nil, fmt.Errorf("search posts eligible for notes new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var sopt SearchPostsEligibleForNotesOption
	switch len(opt) {
	case 0:
		// Set default test_mode = true
		sopt.TestMode = true
	case 1:
		sopt = *opt[0]
		sopt.TestMode = true // Ensure test_mode is always true
	default:
		return nil, errors.New("search posts eligible for notes: only one option is allowed")
	}
	sopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("search posts eligible for notes response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	var spr SearchPostsEligibleForNotesResponse
	if err := json.NewDecoder(resp.Body).Decode(&spr); err != nil {
		return nil, fmt.Errorf("search posts eligible for notes decode: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &spr, &HTTPError{
			APIName: "search posts eligible for notes",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &spr, nil
}

func searchNotesWritten(ctx context.Context, c *client, opt ...*SearchNotesWrittenOption) (*SearchNotesWrittenResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, searchNotesWrittenURL, nil)
	if err != nil {
		return nil, fmt.Errorf("search notes written new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var sopt SearchNotesWrittenOption
	switch len(opt) {
	case 0:
		// Set default test_mode = true
		sopt.TestMode = true
	case 1:
		sopt = *opt[0]
		sopt.TestMode = true // Ensure test_mode is always true
	default:
		return nil, errors.New("search notes written: only one option is allowed")
	}
	sopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("search notes written response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	var snr SearchNotesWrittenResponse
	if err := json.NewDecoder(resp.Body).Decode(&snr); err != nil {
		return nil, fmt.Errorf("search notes written decode: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &snr, &HTTPError{
			APIName: "search notes written",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &snr, nil
}

func createCommunityNote(ctx context.Context, c *client, body *CreateCommunityNoteBody) (*CreateCommunityNoteResponse, error) {
	if body == nil {
		return nil, errors.New("create community note: body parameter is required")
	}
	if body.PostID == "" {
		return nil, errors.New("create community note: post_id is required")
	}
	if body.Info == nil {
		return nil, errors.New("create community note: info is required")
	}

	// Ensure test_mode is always true
	body.TestMode = true

	j, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("create community note json marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, createCommunityNoteURL, bytes.NewReader(j))
	if err != nil {
		return nil, fmt.Errorf("create community note new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("create community note response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	var cnr CreateCommunityNoteResponse
	if err := json.NewDecoder(resp.Body).Decode(&cnr); err != nil {
		return nil, fmt.Errorf("create community note decode: %w", err)
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return &cnr, &HTTPError{
			APIName: "create community note",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &cnr, nil
}