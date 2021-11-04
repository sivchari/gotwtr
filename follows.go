package gotwtr

import (
	"bytes"
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
	if id == "" {
		return nil, errors.New("post following by id: id parameter is required")
	}
	postFollowingPath := baseUserPath + "/" + id + "/following"

	if tuid == "" {
		return nil, errors.New("post following by target_user_id: target_user_id parameter is required")
	}
	body := &FollowingBody{
		TargetUserID: tuid,
	}
	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("post following: can not marshal")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, postFollowingPath, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("post following new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("OAuth %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("post folllowing response: %w", err)
	}
	defer resp.Body.Close()

	var postFollowing PostFollowingResponse
	if err := json.NewDecoder(resp.Body).Decode(&postFollowing); err != nil {
		return nil, fmt.Errorf("post following decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &postFollowing, &HTTPError{
			APIName: "post following",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &postFollowing, nil
}

// suid = source_user_id
// tuid = target_user_id
func deleteFollowing(ctx context.Context, c *client, suid string, tuid string) (*DeleteFollowingResponse, error) {
	if suid == "" {
		return nil, errors.New("delete following by source_user_id: source_user_id parameter is required")
	}
	if tuid == "" {
		return nil, errors.New("delete following by target_user_id: target_user_id parameter is required")
	}
	deleteFollowingPath := baseUserPath + "/" + suid + "/following/" + tuid

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, deleteFollowingPath, nil)
	if err != nil {
		return nil, fmt.Errorf("delete following new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("OAuth %s", c.bearerToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("delete following response: %w", err)
	}
	defer resp.Body.Close()

	var deleteFollowing DeleteFollowingResponse
	if err := json.NewDecoder(resp.Body).Decode(&deleteFollowing); err != nil {
		return nil, fmt.Errorf("delete following decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &deleteFollowing, &HTTPError{
			APIName: "delete following",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &deleteFollowing, nil
}
