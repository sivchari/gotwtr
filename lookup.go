package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func lookUp(ctx context.Context, c *client, ids []string, opt ...*TweetOption) (*TweetLookUpResponse, error) {
	// check ids
	switch {
	case len(ids) == 0:
		return nil, errors.New("tweet lookup: ids parameter is required")
	case len(ids) > tweetLookUpMaxIDs:
		return nil, errors.New("tweet lookup: ids parameter must be less than or equal to 100")
	default:
	}
	tweetLookUp := tweetLookUp + "?ids="
	// join ids to url
	for i, id := range ids {
		if i+1 < len(ids) {
			tweetLookUp += fmt.Sprintf("%s,", id)
		} else {
			tweetLookUp += id
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, tweetLookUp, nil)
	if err != nil {
		return nil, fmt.Errorf("tweet lookup new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt TweetOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("tweet lookup: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("tweet lookup response: %w", err)
	}
	defer resp.Body.Close()

	var tweet TweetLookUpResponse
	if err := json.NewDecoder(resp.Body).Decode(&tweet); err != nil {
		return nil, fmt.Errorf("tweet lookup decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &tweet, &HTTPError{
			APIName: "tweet lookup",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &tweet, nil
}

func lookUpByID(ctx context.Context, c *client, id string, opt ...*TweetOption) (*TweetLookUpByIDResponse, error) {
	if id == "" {
		return nil, errors.New("tweet lookup by id: id parameter is required")
	}
	tweetLookUpByID := tweetLookUp + "/" + id
	fmt.Println(tweetLookUpByID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, tweetLookUpByID, nil)
	if err != nil {
		return nil, fmt.Errorf("tweet lookup by id new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("tweet lookup by id response: %w", err)
	}

	defer resp.Body.Close()

	var tweet TweetLookUpByIDResponse
	if err := json.NewDecoder(resp.Body).Decode(&tweet); err != nil {
		return nil, fmt.Errorf("tweet lookup by id decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &tweet, &HTTPError{
			APIName: "tweet lookup by id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &tweet, nil
}
