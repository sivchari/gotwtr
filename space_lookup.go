package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

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
