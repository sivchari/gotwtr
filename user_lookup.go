package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func retrieveMultipleUsersWithIDs(ctx context.Context, c *client, userIDs []string, opt ...*RetrieveUserOption) (*UsersResponse, error) {
	switch {
	case len(userIDs) == 0:
		return nil, errors.New("retrieve multiple users with ids: ids parameter is required")
	case len(userIDs) > userLookUpMaxIDs:
		return nil, errors.New("retrieve multiple users with ids: ids parameter must be less than or equal to 100")
	default:
	}
	ep := retrieveMultipleUsersWithIDsURL + strings.Join(userIDs, ",")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve multiple users with ids new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var ropt RetrieveUserOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		ropt = *opt[0]
	default:
		return nil, errors.New("retrieve multiple users with ids: only one option is allowed")
	}
	ropt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("retrieve multiple users with ids response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var ur UsersResponse
	if err := json.NewDecoder(resp.Body).Decode(&ur); err != nil {
		return nil, fmt.Errorf("retrieve multiple users with ids decode: %w", err)
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

func retrieveSingleUserWithID(ctx context.Context, c *client, userID string, opt ...*RetrieveUserOption) (*UserResponse, error) {
	if userID == "" {
		return nil, errors.New("retrieve single user with id: user id is required")
	}
	ep := fmt.Sprintf(retrieveSingleUserWithIDURL, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve single user with id new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var ropt RetrieveUserOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		ropt = *opt[0]
	default:
		return nil, errors.New("retrieve single user with id: only one option is allowed")
	}
	ropt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("retrieve single user with id response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	var ur UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&ur); err != nil {
		return nil, fmt.Errorf("retrieve single user with id decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &ur, &HTTPError{
			APIName: "retrieve single user with id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &ur, nil
}

func retrieveMultipleUsersWithUserNames(ctx context.Context, c *client, userNames []string, opt ...*RetrieveUserOption) (*UsersResponse, error) {
	switch {
	case len(userNames) == 0:
		return nil, errors.New("retrieve multiple users with user names: user names parameter is required")
	case len(userNames) > userLookUpMaxIDs:
		return nil, errors.New("retrieve multiple users with user names: user names parameter must be less than or equal to 100")
	default:
	}
	ep := retrieveMultipleUsersWithUserNamesURL + strings.Join(userNames, ",")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve multiple users with user names new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var ropt RetrieveUserOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		ropt = *opt[0]
	default:
		return nil, errors.New("retrieve multiple users with user names: only one option is allowed")
	}
	ropt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("retrieve multiple users with user names response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	var ur UsersResponse
	if err := json.NewDecoder(resp.Body).Decode(&ur); err != nil {
		return nil, fmt.Errorf("retrieve multiple users with user names decode: %w", err)
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

func retrieveSingleUserWithUserName(ctx context.Context, c *client, userName string, opt ...*RetrieveUserOption) (*UserResponse, error) {
	if userName == "" {
		return nil, errors.New("retrieve single user with user name: user name is required")
	}
	ep := fmt.Sprintf(retrieveSingleUserWithUserNameURL, userName)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve single user with user name new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var ropt RetrieveUserOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		ropt = *opt[0]
	default:
		return nil, errors.New("retrieve single user with user name: only one option is allowed")
	}
	ropt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("retrieve single user with user name response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	var ur UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&ur); err != nil {
		return nil, fmt.Errorf("retrieve single user with user name decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &ur, &HTTPError{
			APIName: "retrieve single user with user name",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &ur, nil
}

func searchUsers(ctx context.Context, c *client, query string, opt ...*SearchUsersOption) (*SearchUsersResponse, error) {
	if len(query) == 0 {
		return nil, errors.New("search users: query parameter is required")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, searchUsersURL, nil)
	if err != nil {
		return nil, fmt.Errorf("search users new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	q := req.URL.Query()
	q.Add("query", query)
	req.URL.RawQuery = q.Encode()

	var sopt SearchUsersOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		sopt = *opt[0]
	default:
		return nil, errors.New("search users: only one option is allowed")
	}
	sopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("search users response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var sur SearchUsersResponse
	if err := json.NewDecoder(resp.Body).Decode(&sur); err != nil {
		return nil, fmt.Errorf("search users decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &sur, &HTTPError{
			APIName: "search users",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &sur, nil
}
