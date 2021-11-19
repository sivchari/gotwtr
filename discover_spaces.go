package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func discoverSpacesByUserIDs(ctx context.Context, c *client, ids []string, opt ...*DiscoverSpacesOption) (*DiscoverSpacesByUserIDsResponse, error) {
	// check ids
	switch {
	case len(ids) == 0:
		return nil, errors.New("discover spaces: ids parameter is required")
	case len(ids) > discoverSpacesMaxIDs:
		return nil, errors.New("discover spaces: ids parameter must be less than or equal to 100")
	default:
	}

	spaceDiscover := spaceDiscover + "?user_ids="
	// join ids to url
	for i, id := range ids {
		if i+1 < len(ids) {
			spaceDiscover += fmt.Sprintf("%s,", id)
		} else {
			spaceDiscover += id
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, spaceDiscover, nil)
	if err != nil {
		return nil, fmt.Errorf("discover spaces new request with ctx: %w", err)
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

	var dsr DiscoverSpacesByUserIDsResponse

	if err := json.NewDecoder(resp.Body).Decode(&dsr); err != nil {
		return nil, fmt.Errorf("discover spaces: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			APIName: "discover spaces",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &dsr, nil
}
