package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func listMembers(ctx context.Context, c *client, listid string, opt ...*ListMembersOption) (*ListMembersResponse, error) {
	if listid == "" {
		return nil, errors.New("look up list members: id parameter is required")
	}
	lm := fmt.Sprintf(listMembersURL, listid)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, lm, nil)
	if err != nil {
		return nil, fmt.Errorf("look up list members: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt ListMembersOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("look up list members: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("look up list members: %w", err)
	}

	defer resp.Body.Close()

	var lmr ListMembersResponse
	if err := json.NewDecoder(resp.Body).Decode(&lmr); err != nil {
		return nil, fmt.Errorf("look up list members: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &lmr, &HTTPError{
			APIName: "owned lists lookup by id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &lmr, nil
}
