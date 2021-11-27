package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func discoverSpaces(ctx context.Context, c *client, userIDs []string, opt ...*DiscoverSpacesOption) (*DiscoverSpacesResponse, error) {
	switch {
	case len(userIDs) == 0:
		return nil, errors.New("discover spaces: ids parameter is required")
	case len(userIDs) > discoverSpacesMaxIDs:
		return nil, errors.New("discover spaces: ids parameter must be less than or equal to 100")
	default:
	}
	ep := discoverSpacesURL
	for i, uid := range userIDs {
		if i+1 < len(userIDs) {
			ep += fmt.Sprintf("%s,", uid)
		} else {
			ep += uid
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("discover spaces: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var dopt DiscoverSpacesOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		dopt = *opt[0]
	default:
		return nil, errors.New("discover spaces: only one option is allowed")
	}
	dopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("discover spaces: %w", err)
	}
	defer resp.Body.Close()

	var dsr DiscoverSpacesResponse

	if err := json.NewDecoder(resp.Body).Decode(&dsr); err != nil {
		return nil, fmt.Errorf("discover spaces: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &dsr, &HTTPError{
			APIName: "discover spaces",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &dsr, nil
}
