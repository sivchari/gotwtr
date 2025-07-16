package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func createNewList(ctx context.Context, c *client, body *CreateNewListBody) (*CreateNewListResponse, error) {
	if body.Name == "" {
		return nil, errors.New("create new list: name parameter is required")
	}
	if len(body.Name) > 25 {
		return nil, errors.New("create new list: name parameter must be less than 25 characters")
	}
	j, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("create new list : can not marshal")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, createNewListURL, bytes.NewBuffer(j))
	if err != nil {
		return nil, fmt.Errorf("create new list new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("create new list response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var createNewList CreateNewListResponse
	if err := json.NewDecoder(resp.Body).Decode(&createNewList); err != nil {
		return nil, fmt.Errorf("create new list decode: %w", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return &createNewList, &HTTPError{
			APIName: "create new list",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &createNewList, nil
}

func deleteList(ctx context.Context, c *client, listID string) (*DeleteListResponse, error) {
	if listID == "" {
		return nil, errors.New("delete list: list id parameter is required")
	}
	ep := fmt.Sprintf(deleteListURL, listID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("delete list new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("delete list response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var deleteList DeleteListResponse
	if err := json.NewDecoder(resp.Body).Decode(&deleteList); err != nil {
		return nil, fmt.Errorf("delete list decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &deleteList, &HTTPError{
			APIName: "delete list",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &deleteList, nil
}

func updateMetaDataForList(ctx context.Context, c *client, listID string, body ...*UpdateMetaDataForListBody) (*UpdateMetaDataForListResponse, error) {
	if listID == "" {
		return nil, errors.New("update meta data for list: list id parameter is required")
	}
	ep := fmt.Sprintf(updateMetaDataForListURL, listID)

	var ubody UpdateMetaDataForListBody
	switch len(body) {
	case 0:
		// do nothing
	case 1:
		ubody = *body[0]
	default:
		return nil, errors.New("update meta data for list: only one option is allowed")
	}

	j, err := json.Marshal(ubody)
	if err != nil {
		return nil, errors.New("update meta data for list : can not marshal")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, ep, bytes.NewBuffer(j))
	if err != nil {
		return nil, fmt.Errorf("update meta data for list new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("update meta data for list response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var updateMetaDataForList UpdateMetaDataForListResponse
	if err := json.NewDecoder(resp.Body).Decode(&updateMetaDataForList); err != nil {
		return nil, fmt.Errorf("update meta data for list decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &updateMetaDataForList, &HTTPError{
			APIName: "update meta data for list",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &updateMetaDataForList, nil
}
