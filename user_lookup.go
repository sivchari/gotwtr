package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func retriveMultipleUsersWithIDs(ctx context.Context, c *client, userIDs []string, opt ...*RetriveUserOption) (*UsersResponse, error) {
	switch {
	case len(userIDs) == 0:
		return nil, errors.New("retrive multiple users with ids: ids parameter is required")
	case len(userIDs) > userLookUpMaxIDs:
		return nil, errors.New("retrive multiple users with ids: ids parameter must be less than or equal to 100")
	default:
	}
	ep := retriveMultipleUsersWithIDsURL
	for i, uid := range userIDs {
		if i+1 < len(userIDs) {
			ep += fmt.Sprintf("%s,", uid)
		} else {
			ep += uid
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("retrive multiple users with ids new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var ropt RetriveUserOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		ropt = *opt[0]
	default:
		return nil, errors.New("retrive multiple users with ids: only one option is allowed")
	}
	ropt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("retrive multiple users with ids response: %w", err)
	}
	defer resp.Body.Close()

	var ur UsersResponse
	if err := json.NewDecoder(resp.Body).Decode(&ur); err != nil {
		return nil, fmt.Errorf("retrive multiple users with ids decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &ur, &HTTPError{
			APIName: "user lookup",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &ur, nil
}

func retriveSingleUserWithID(ctx context.Context, c *client, userID string, opt ...*RetriveUserOption) (*UserResponse, error) {
	if userID == "" {
		return nil, errors.New("retrive single user with id: user id is required")
	}
	ep := fmt.Sprintf(retriveSingleUserWithIDURL, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("retrive single user with id new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var ropt RetriveUserOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		ropt = *opt[0]
	default:
		return nil, errors.New("retrive single user with id: only one option is allowed")
	}
	ropt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("retrive single user with id response: %w", err)
	}

	defer resp.Body.Close()

	var ur UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&ur); err != nil {
		return nil, fmt.Errorf("user lookup by id decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &ur, &HTTPError{
			APIName: "retrive single user with id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &ur, nil
}

func retriveMultipleUsersWithUserNames(ctx context.Context, c *client, userNames []string, opt ...*RetriveUserOption) (*UsersResponse, error) {
	switch {
	case len(userNames) == 0:
		return nil, errors.New("retrive multiple users with user names: user names parameter is required")
	case len(userNames) > userLookUpMaxIDs:
		return nil, errors.New("retrive multiple users with user names: user names parameter must be less than or equal to 100")
	default:
	}
	ep := retriveMultipleUsersWithUserNamesURL
	for i, un := range userNames {
		if i+1 < len(un) {
			ep += fmt.Sprintf("%s,", un)
		} else {
			ep += un
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("retrive multiple users with user names new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var ropt RetriveUserOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		ropt = *opt[0]
	default:
		return nil, errors.New("retrive multiple users with user names: only one option is allowed")
	}
	ropt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("retrive multiple users with user names response: %w", err)
	}

	defer resp.Body.Close()

	var ur UsersResponse
	if err := json.NewDecoder(resp.Body).Decode(&ur); err != nil {
		return nil, fmt.Errorf("users lookup by usernames decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &ur, &HTTPError{
			APIName: "users lookup by usernames",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &ur, nil
}

func retriveSingleUserWithUserName(ctx context.Context, c *client, userName string, opt ...*RetriveUserOption) (*UserResponse, error) {
	if userName == "" {
		return nil, errors.New("retrive single user with user name: user name is required")
	}
	ep := fmt.Sprintf(retriveSingleUserWithUserNameURL, userName)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("retrive single user with user name new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var ropt RetriveUserOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		ropt = *opt[0]
	default:
		return nil, errors.New("retrive single user with user name: only one option is allowed")
	}
	ropt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("retrive single user with user name response: %w", err)
	}

	defer resp.Body.Close()

	var ur UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&ur); err != nil {
		return nil, fmt.Errorf("user lookup by username decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &ur, &HTTPError{
			APIName: "user lookup by username",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &ur, nil
}
