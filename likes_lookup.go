package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func likesLookUpUsers(ctx context.Context, c *client, id string, opt ...*LikesLookUpByTweetOpts) (*LikesLookUpByTweetResponse, error) {
	// check id
	if len(id) == 0 {
		return nil, errors.New("likes look up by tweet: id parameter is required")
	}
	LikesLookUpByTweet := tweetLookUp + "/" + id + "/liking_users"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, LikesLookUpByTweet, nil)
	if err != nil {
		return nil, fmt.Errorf("likes look up by tweet new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt LikesLookUpByTweetOpts
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("likes look up by tweet: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("likes look up by tweet response: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("likes look up by tweet response status: %d", http.StatusNotFound)
	}

	var usersWhoLiked LikesLookUpByTweetResponse
	if err := json.NewDecoder(resp.Body).Decode(&usersWhoLiked); err != nil {
		return nil, fmt.Errorf("likes look up by tweet decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &usersWhoLiked, &HTTPError{
			APIName: "likes look up by tweet",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &usersWhoLiked, nil
}
