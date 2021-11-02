package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func lookUpSpaces(ctx context.Context, c *client, ids []string, opt ...*SpaceLookUpOption) (*SpaceLookUpResponse, error) {
	// check ids
	switch {
	case len(ids) == 0:
		return nil, errors.New("space lookup: ids parameter is required")
	case len(ids) > spaceLookUpMaxIDs:
		return nil, errors.New("space lookup: ids parameter must be less than or equal to 100")
	default:
	}

	spaceLookUp := spaceLookUp + "?ids="
	// join ids to url
	for i, id := range ids {
		if i+1 < len(ids) {
			spaceLookUp += fmt.Sprintf("%s,", id)
		} else {
			spaceLookUp += id
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, spaceLookUp, nil)
	if err != nil {
		return nil, fmt.Errorf("space lookup new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var sopt SpaceLookUpOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		sopt = *opt[0]
	default:
		return nil, errors.New("space lookup: only one option is allowed")
	}
	sopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("space lookup response: %w", err)
	}
	defer resp.Body.Close()

	var slr SpaceLookUpResponse

	if err := json.NewDecoder(resp.Body).Decode(&slr); err != nil {
		return nil, fmt.Errorf("space lookup: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &slr, &HTTPError{
			APIName: "space lookup",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &slr, nil
}

func lookUpSpaceByID(ctx context.Context, c *client, id string, opt ...*SpaceLookUpOption) (*SpaceLookUpByIDResponse, error) {
	if id == "" {
		return nil, errors.New("space lookup by id: id parameter is required")
	}
	spaceLookUpByID := spaceLookUp + "/" + id

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, spaceLookUpByID, nil)
	if err != nil {
		return nil, fmt.Errorf("space lookup by id: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var sopt SpaceLookUpOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		sopt = *opt[0]
	default:
		return nil, errors.New("space lookup by id: only one option is allowed")
	}
	sopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("space lookup by id: %w", err)
	}

	defer resp.Body.Close()

	var slr SpaceLookUpByIDResponse
	if err := json.NewDecoder(resp.Body).Decode(&slr); err != nil {
		return nil, fmt.Errorf("space lookup by id: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &slr, &HTTPError{
			APIName: "space lookup by id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &slr, nil
}
