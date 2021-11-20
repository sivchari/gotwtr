package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func lookUpUsersWhoLiked(ctx context.Context, c *client, tweetID string, opt ...*LookUpUsersWhoLikedOpts) (*LookUpUsersWhoLikedResponse, error) {
	// check id
	if len(tweetID) == 0 {
		return nil, errors.New("look up users who liked: tweetID parameter is required")
	}
	likesLookUpUsersURL := likesLookUpUsers + "/" + tweetID + "/liking_users"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, likesLookUpUsersURL, nil)
	if err != nil {
		return nil, fmt.Errorf("look up users who liked: new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt LookUpUsersWhoLikedOpts
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("look up users who liked: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("look up users who liked response: %w", err)
	}
	defer resp.Body.Close()

	var usersWhoLiked LookUpUsersWhoLikedResponse
	if err := json.NewDecoder(resp.Body).Decode(&usersWhoLiked); err != nil {
		return nil, fmt.Errorf("look up users who liked decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &usersWhoLiked, &HTTPError{
			APIName: "look up users who liked",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &usersWhoLiked, nil
}
