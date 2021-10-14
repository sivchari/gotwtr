package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func searchSpaces(ctx context.Context, c *client, query string, opt ...*SearchSpacesOption) (*SearchSpacesResponse, error) {
	switch {
	case len(query) == 0:
		return nil, errors.New("search spaces: query is empty")
	default:
	}

	searchSpaces := spacesSearch + "?query=" + query

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, searchSpaces, nil)
	if err != nil {
		return nil, fmt.Errorf("search spaces: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt SearchSpacesOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("search spaces: too many options")
	}
	topt.addQuery(req)
	fmt.Println(req.URL.String())

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
