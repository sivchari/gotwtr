package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func createOneToOneDM(ctx context.Context, c *client, participantID string, body *CreateOneToOneDMBody) (*CreateOneToOneDMResponse, error) {
	if participantID == "" {
		return nil, errors.New("create a one to one DM: participant id parameter is required")
	}
	ep := fmt.Sprintf(createOneToOneDMURL, participantID)
	j, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("create a one to one DM: can not marshal")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep, bytes.NewBuffer((j)))
	if err != nil {
		return nil, fmt.Errorf("create a one to one DM with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("create a one to one DM response: %w", err)
	}
	defer resp.Body.Close()

	var createOneToOneDM CreateOneToOneDMResponse
	if err := json.NewDecoder(resp.Body).Decode(&createOneToOneDM); err != nil {
		return nil, fmt.Errorf("create a one to one DM decode: %w", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return &createOneToOneDM, &HTTPError{
			APIName: "create one to one DM",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}
	return &createOneToOneDM, nil
}

func createNewGroupDM(ctx context.Context, c *client, conversationID string, body *CreateNewGroupDMBody) (*CreateNewGroupDMResponse, error) {
	if conversationID == "" {
		return nil, errors.New("create new group DM: conversation id parameter is required")
	}
	ep := fmt.Sprintf(createNewGroupDMURL, conversationID)
	j, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("create new group DM: can not marshal")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep, bytes.NewBuffer((j)))
	if err != nil {
		return nil, fmt.Errorf("create new group DM with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("create new group DM response: %w", err)
	}
	defer resp.Body.Close()

	var createNewGroupDM CreateNewGroupDMResponse
	if err := json.NewDecoder(resp.Body).Decode(&createNewGroupDM); err != nil {
		return nil, fmt.Errorf("create new group DM decode: %w", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return &createNewGroupDM, &HTTPError{
			APIName: "create new group DM",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}
	return &createNewGroupDM, nil
}

func postDM(ctx context.Context, c *client, body *PostDMBody) (*PostDMResponse, error) {
	j, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("post DM: can not marshal")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, postDMURL, bytes.NewBuffer((j)))
	if err != nil {
		return nil, fmt.Errorf("post DM with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("post DM response: %w", err)
	}
	defer resp.Body.Close()

	var postDM PostDMResponse
	if err := json.NewDecoder(resp.Body).Decode(&postDM); err != nil {
		return nil, fmt.Errorf("post DM decode: %w", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return &postDM, &HTTPError{
			APIName: "post DM",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}
	return &postDM, nil
}
