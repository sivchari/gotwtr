package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func searchSpaces(ctx context.Context, c *client, searchTerm string, opt ...*SearchSpacesOption) (*SearchSpacesResponse, error) {
	if searchTerm == "" {
		return nil, errors.New("search spaces: searchTerm parameter is required")
	}
	ep := fmt.Sprintf(searchSpacesURL, searchTerm)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("search spaces new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var sopt SearchSpacesOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		sopt = *opt[0]
	default:
		return nil, errors.New("search spaces: too many options")
	}
	sopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("search spaces: %w", err)
	}
	defer resp.Body.Close()

	var ssr SearchSpacesResponse
	if err := json.NewDecoder(resp.Body).Decode(&ssr); err != nil {
		return nil, fmt.Errorf("search spaces: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			APIName: "search spaces",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &ssr, nil
}
