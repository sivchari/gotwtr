package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func lookUpListFollowersByID(ctx context.Context, c *client, id string, opt ...*ListLookUpOption) (*ListFollowersLookUpByIDResponse, error) {
	if id == "" {
		return nil, errors.New("list followers lookup by id: id parameter is required")
	}
	lookUpListFollowersByID := listLookUp + "/" + id + "/followers"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, lookUpListFollowersByID, nil)
	if err != nil {
		return nil, fmt.Errorf("list followers lookup by id new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt ListLookUpOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("list followers lookup by id: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("list followers lookup by id response: %w", err)
	}

	defer resp.Body.Close()

	var list ListFollowersLookUpByIDResponse
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, fmt.Errorf("list followers lookup by id decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &list, &HTTPError{
			APIName: "list followers lookup by id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &list, nil
}

func lookUpListsUserFollowingByID(ctx context.Context, c *client, id string, opt ...*ListLookUpOption) (*ListsUserFollowingLookUpByIDResponse, error) {
	if id == "" {
		return nil, errors.New("lists user following lookup by id: id parameter is required")
	}
	listsUserFollowingLookUpByID := followingLookUp + "/" + id + "/followed_lists"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, listsUserFollowingLookUpByID, nil)
	if err != nil {
		return nil, fmt.Errorf("lists user following lookup new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt ListLookUpOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("lists user following lookup by id: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("lists user following lookup by id response: %w", err)
	}
	defer resp.Body.Close()

	var list ListsUserFollowingLookUpByIDResponse
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, fmt.Errorf("lists user following lookup by id decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &list, &HTTPError{
			APIName: "lists user following lookup by id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &list, nil
}
