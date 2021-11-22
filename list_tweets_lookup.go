package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func lookUpListsTweetsByID(ctx context.Context, c *client, id string, opt ...*ListLookUpOption) (*ListsTweetsLookUpByIDResponse, error) {
	if id == "" {
		return nil, errors.New("lists tweets lookup by id: id parameter is required")
	}
	lookUpListsTweetsByID := listsTweetsLookUp + "/" + id + "/tweets"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, lookUpListsTweetsByID, nil)
	if err != nil {
		return nil, fmt.Errorf("lists tweets lookup by id new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt ListLookUpOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("lists tweets lookup by id: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("lists tweets lookup by id response: %w", err)
	}

	defer resp.Body.Close()

	var list ListsTweetsLookUpByIDResponse
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, fmt.Errorf("lists tweets lookup by id decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &list, &HTTPError{
			APIName: "lists tweets lookup by id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &list, nil
}
