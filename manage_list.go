package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/post-lists
func createNewList(ctx context.Context, c *client, listName string, opt ...*CreateNewListOption) (*CreateNewListResponse, error) {
	if listName == "" {
		return nil, errors.New("create new list: tweet id parameter is required")
	}
	ep := fmt.Sprintf(listURL, listName)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("create new list new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var copt CreateNewListOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		copt = *opt[0]
	default:
		return nil, errors.New("create new list: only one option is allowed")
	}
	copt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("create new list response: %w", err)
	}
	defer resp.Body.Close()

	var createNewList CreateNewListResponse
	if err := json.NewDecoder(resp.Body).Decode(&createNewList); err != nil {
		return nil, fmt.Errorf("create new list decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &createNewList, &HTTPError{
			APIName: "create new list",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &createNewList, nil
}

//https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/delete-lists-id
func deleteList(ctx context.Context, c *client, listID string) (*DeleteListResponse, error) {
	if listID == "" {
		return nil, errors.New("delete list: list id parameter is required")
	}
	ep := fmt.Sprintf(lookUpListURL, listID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("delete list new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("delete list response: %w", err)
	}
	defer resp.Body.Close()

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

//https://developer.twitter.com/en/docs/twitter-api/lists/manage-lists/api-reference/put-lists-id
func updateMetaDataForList(ctx context.Context, c *client, listName string, opt ...*UpdateMetaDataForListOption) (*UpdateMetaDataForListResponse, error) {
	if listName == "" {
		return nil, errors.New("update meta data for list: tweet id parameter is required")
	}
	ep := fmt.Sprintf(lookUpListURL, listName)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("update meta data for list new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var copt UpdateMetaDataForListOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		copt = *opt[0]
	default:
		return nil, errors.New("update meta data for list: only one option is allowed")
	}
	copt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("update meta data for list response: %w", err)
	}
	defer resp.Body.Close()

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
