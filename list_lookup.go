package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func lookUpOwnedListsByID(ctx context.Context, c *client, id string, opt ...*ListLookUpOption) (*OwnedListsLookUpByIDResponse, error) {
	if id == "" {
		return nil, errors.New("owned lists lookup by id: id parameter is required")
	}
	lookUpOwnedListsByID := ownedListLookUp + "/" + id + "/owned_lists"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, lookUpOwnedListsByID, nil)
	if err != nil {
		return nil, fmt.Errorf("owned lists lookup by id new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt ListLookUpOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("owned lists lookup by id: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("owned lists lookup by id response: %w", err)
	}

	defer resp.Body.Close()

	var list OwnedListsLookUpByIDResponse
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, fmt.Errorf("owned lists lookup by id decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &list, &HTTPError{
			APIName: "owned lists lookup by id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &list, nil
}

func lookUpListByID(ctx context.Context, c *client, id string, opt ...*ListLookUpOption) (*ListLookUpByIDResponse, error) {
	if id == "" {
		return nil, errors.New("list lookup by id: id parameter is required")
	}
	listLookUpByID := listLookUp + "/" + id

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, listLookUpByID, nil)
	if err != nil {
		return nil, fmt.Errorf("list lookup new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt ListLookUpOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("list lookup by id: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("list lookup by id response: %w", err)
	}
	defer resp.Body.Close()

	var list ListLookUpByIDResponse
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, fmt.Errorf("list lookup by id decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &list, &HTTPError{
			APIName: "list lookup by id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &list, nil
}
