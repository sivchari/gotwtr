package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func usersLikingTweet(ctx context.Context, c *client, tweetID string, opt ...*UsersLikingTweetOption) (*UsersLikingTweetResponse, error) {
	if tweetID == "" {
		return nil, errors.New("users liking tweet: tweet id parameter is required")
	}
	ep := fmt.Sprintf(usersLikingTweetURL, tweetID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("users liking tweet new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var uopt UsersLikingTweetOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		uopt = *opt[0]
	default:
		return nil, errors.New("users liking tweet: only one option is allowed")
	}
	uopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("users liking tweet: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var ultr UsersLikingTweetResponse
	if err := json.NewDecoder(resp.Body).Decode(&ultr); err != nil {
		return nil, fmt.Errorf("users liking tweet: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &ultr, &HTTPError{
			APIName: "users liking tweet",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &ultr, nil
}

func tweetsUserLiked(ctx context.Context, c *client, userID string, opt ...*TweetsUserLikedOption) (*TweetsUserLikedResponse, error) {
	if userID == "" {
		return nil, errors.New("tweets user liked: user id parameter is required")
	}
	ep := fmt.Sprintf(tweetsUserLikedURL, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("tweets user liked new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt TweetsUserLikedOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("tweets user liked: only one option is allowed")
	}
	const (
		minimumMaxResults = 10
		maximumMaxResults = 100
		defaultMaxResults = 100
	)
	if topt.MaxResults == 0 {
		topt.MaxResults = defaultMaxResults
	}
	if topt.MaxResults < minimumMaxResults || topt.MaxResults > maximumMaxResults {
		return nil, fmt.Errorf("tweets user liked: max results must be between %d and %d", minimumMaxResults, maximumMaxResults)
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("tweets user liked: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var tulr TweetsUserLikedResponse
	if err := json.NewDecoder(resp.Body).Decode(&tulr); err != nil {
		return nil, fmt.Errorf("tweets user liked: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &tulr, &HTTPError{
			APIName: "tweets user liked",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &tulr, nil
}

func postUsersLikingTweet(ctx context.Context, c *client, userID string, tweetID string) (*PostUsersLikingTweetResponse, error) {
	if userID == "" {
		return nil, errors.New("post users liking tweet: userID parameter is required")
	}
	ep := fmt.Sprintf(postUsersLikingTweetURL, userID)

	if tweetID == "" {
		return nil, errors.New("post users liking tweet: tweetID parameter is required")
	}
	body := &UsersLikingBody{
		TweetID: tweetID,
	}
	j, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("post users liking tweet: can not marshal")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep, bytes.NewBuffer(j))
	if err != nil {
		return nil, fmt.Errorf("post users liking tweet new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("post users liking tweet response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var postUsersLikingTweet PostUsersLikingTweetResponse
	if err := json.NewDecoder(resp.Body).Decode(&postUsersLikingTweet); err != nil {
		return nil, fmt.Errorf("post users liking tweet decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &postUsersLikingTweet, &HTTPError{
			APIName: "post users liking tweet",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &postUsersLikingTweet, nil
}

func undoUsersLikingTweet(ctx context.Context, c *client, userID string, tweetID string) (*UndoUsersLikingTweetResponse, error) {
	if userID == "" {
		return nil, errors.New("undo users liking tweet: userID parameter is required")
	}
	if tweetID == "" {
		return nil, errors.New("undo users liking tweet: tweetID parameter is required")
	}
	ep := fmt.Sprintf(undoUsersLikingTweetURL, userID, tweetID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("undo users liking tweet new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("undo users liking tweet response: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var undoUsersLikingTweet UndoUsersLikingTweetResponse
	if err := json.NewDecoder(resp.Body).Decode(&undoUsersLikingTweet); err != nil {
		return nil, fmt.Errorf("undo users liking tweet decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &undoUsersLikingTweet, &HTTPError{
			APIName: "undo users liking tweet",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &undoUsersLikingTweet, nil
}
