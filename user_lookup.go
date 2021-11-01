package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func lookUpUsers(ctx context.Context, c *client, ids []string, opt ...*UserLookUpOption) (*UserLookUpResponse, error) {
	// check ids
	switch {
	case len(ids) == 0:
		return nil, errors.New("user lookup: ids parameter is required")
	case len(ids) > userLookUpMaxIDs:
		return nil, errors.New("user lookup: ids parameter must be less than or equal to 100")
	default:
	}
	lookUpUsers := baseUserPath + "?ids="
	// join ids to url
	for i, id := range ids {
		if i+1 < len(ids) {
			lookUpUsers += fmt.Sprintf("%s,", id)
		} else {
			lookUpUsers += id
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, lookUpUsers, nil)
	if err != nil {
		return nil, fmt.Errorf("user lookup new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt UserLookUpOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("user lookup: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("user lookup response: %w", err)
	}
	defer resp.Body.Close()

	var user UserLookUpResponse
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("user lookup decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &user, &HTTPError{
			APIName: "user lookup",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &user, nil
}

func lookUpUserByID(ctx context.Context, c *client, id string, opt ...*UserLookUpOption) (*UserLookUpByIDResponse, error) {
	if id == "" {
		return nil, errors.New("user lookup by id: id parameter is required")
	}
	lookUpUserByID := baseUserPath + "/" + id

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, lookUpUserByID, nil)
	if err != nil {
		return nil, fmt.Errorf("user lookup by id new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt UserLookUpOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("user lookup: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("user lookup by id response: %w", err)
	}

	defer resp.Body.Close()

	var user UserLookUpByIDResponse
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("user lookup by id decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &user, &HTTPError{
			APIName: "user lookup by id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &user, nil
}

// lookup single user by username
func lookUpUserByUserName(ctx context.Context, c *client, name string, opt ...*UserLookUpOption) (*UserLookUpByUserNameResponse, error) {
	if name == "" {
		return nil, errors.New("user lookup by username: name parameter is required")
	}
	lookUpUserByUserName := baseUserPath + "/by/username/" + name

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, lookUpUserByUserName, nil)
	if err != nil {
		return nil, fmt.Errorf("user lookup by username new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt UserLookUpOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("user lookup by username: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("user lookup by username response: %w", err)
	}

	defer resp.Body.Close()

	var user UserLookUpByUserNameResponse
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("user lookup by username decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &user, &HTTPError{
			APIName: "user lookup by username",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &user, nil
}

// lookup multiple users by user names
func lookUpUsersByUserNames(ctx context.Context, c *client, names []string, opt ...*UserLookUpOption) (*UsersLookUpByUserNamesResponse, error) {
	// check names
	switch {
	case len(names) == 0:
		return nil, errors.New("users lookup: names parameter is required")
	case len(names) > userLookUpMaxIDs:
		return nil, errors.New("users lookup: names parameter must be less than or equal to 100")
	default:
	}
	lookUpUsersByUserName := baseUserPath + "/by?usernames="
	// join ids to url
	for i, name := range names {
		if i+1 < len(names) {
			lookUpUsersByUserName += fmt.Sprintf("%s,", name)
		} else {
			lookUpUsersByUserName += name
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, lookUpUsersByUserName, nil)
	if err != nil {
		return nil, fmt.Errorf("users lookup by usernames new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt UserLookUpOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("users lookup by usernames: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("users lookup by usernames response: %w", err)
	}

	defer resp.Body.Close()

	var users UsersLookUpByUserNamesResponse
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, fmt.Errorf("users lookup by usernames decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &users, &HTTPError{
			APIName: "users lookup by usernames",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &users, nil
}
