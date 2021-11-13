package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func following(ctx context.Context, c *client, id string, opt ...*FollowOption) (*FollowingResponse, error) {
	// check id
	if id == "" {
		return nil, errors.New("following by id: id parameter is required")
	}
	following := baseUserPath + "/" + id + "/following"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, following, nil)
	if err != nil {
		return nil, fmt.Errorf("following by id new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt FollowOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("following: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("following by id response: %w", err)
	}

	defer resp.Body.Close()

	var user FollowingResponse
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("following by id decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &user, &HTTPError{
			APIName: "following by id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &user, nil
}

func followers(ctx context.Context, c *client, id string, opt ...*FollowOption) (*FollowersResponse, error) {
	// check id
	if id == "" {
		return nil, errors.New("followers by id: id parameter is required")
	}
	followers := baseUserPath + "/" + id + "/followers"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, followers, nil)
	if err != nil {
		return nil, fmt.Errorf("followers by id new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt FollowOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("followers: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("followers by id response: %w", err)
	}

	defer resp.Body.Close()

	var user FollowersResponse
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("followers by id decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &user, &HTTPError{
			APIName: "followers by id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &user, nil
}

// tuid = target_user_id
func postFollowing(ctx context.Context, c *client, id string, tuid string, opt ...*FollowOption) (*PostFollowingResponse, error) {
	return nil, nil
}

// suid = source_user_id
// tuid = target_user_id
func undoFollowing(ctx context.Context, c *client, suid string, tuid string) (*UndoFollowingResponse, error) {
	return nil, nil
}
