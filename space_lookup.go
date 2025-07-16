package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func lookUpSpace(ctx context.Context, c *client, spaceID string, opt ...*SpaceOption) (*SpaceResponse, error) {
	if spaceID == "" {
		return nil, errors.New("look up space: spaceID parameter is required")
	}
	ep := fmt.Sprintf(spaceURL, spaceID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("look up space new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var sopt SpaceOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		sopt = *opt[0]
	default:
		return nil, errors.New("look up space: only one option is allowed")
	}
	sopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("look up space response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	var sr SpaceResponse
	if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		return nil, fmt.Errorf("space lookup: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &sr, &HTTPError{
			APIName: "space lookup by id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &sr, nil
}

func lookUpSpaces(ctx context.Context, c *client, spaceIDs []string, opt ...*SpaceOption) (*SpacesResponse, error) {
	switch {
	case len(spaceIDs) == 0:
		return nil, errors.New("look up spaces: spaceIDs parameter is required")
	case len(spaceIDs) > spaceLookUpMaxIDs:
		return nil, fmt.Errorf("look up spaces: spaceIDs parameter must be less than %d", spaceLookUpMaxIDs)
	default:
	}
	ep := spacesURL
	for i, sid := range spaceIDs {
		if i+1 < len(spaceIDs) {
			ep += fmt.Sprintf("%s,", sid)
		} else {
			ep += sid
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("look up spaces new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var sopt SpaceOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		sopt = *opt[0]
	default:
		return nil, errors.New("look up spaces: only one option is allowed")
	}
	sopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("look up spaces response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var sr SpacesResponse

	if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		return nil, fmt.Errorf("look up spaces: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &sr, &HTTPError{
			APIName: "look up spaces",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &sr, nil
}

func usersPurchasedSpaceTicket(ctx context.Context, c *client, spaceID string, opt ...*UsersPurchasedSpaceTicketOption) (*UsersPurchasedSpaceTicketResponse, error) {
	if spaceID == "" {
		return nil, errors.New("users purchased space ticket: spaceID parameter is required")
	}
	ep := fmt.Sprintf(usersPurchasedSpaceTicketURL, spaceID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("users purchased space ticket new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var uopt UsersPurchasedSpaceTicketOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		uopt = *opt[0]
	default:
		return nil, errors.New("users purchased space ticket: only one option is allowed")
	}
	uopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("users purchased space ticket response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	var upstr UsersPurchasedSpaceTicketResponse
	if err := json.NewDecoder(resp.Body).Decode(&upstr); err != nil {
		return nil, fmt.Errorf("users purchased space ticket: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &upstr, &HTTPError{
			APIName: "users purchased space ticket",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &upstr, nil
}

func spacesTweets(ctx context.Context, c *client, spaceID string, opt ...*SpacesTweetsOption) (*SpacesTweetsResponse, error) {
	if spaceID == "" {
		return nil, errors.New("spaces tweets: spaceID parameter is required")
	}
	ep := fmt.Sprintf(spacesTweetsURL, spaceID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("spaces tweets new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var sopt SpacesTweetsOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		sopt = *opt[0]
	default:
		return nil, errors.New("spaces tweets: only one option is allowed")
	}
	sopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("spaces tweets response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	var str SpacesTweetsResponse
	if err := json.NewDecoder(resp.Body).Decode(&str); err != nil {
		return nil, fmt.Errorf("spaces tweets: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &str, &HTTPError{
			APIName: "spaces tweets",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &str, nil
}
