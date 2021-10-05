package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func userTweetTimeline(ctx context.Context, c *client, id string, opt ...*TweetOption) (*UserTweetTimelineResponse, error) {
	// check id
	if len(id) == 0 {
		return nil, errors.New("user tweet timeline: id parameter is required")
	}
	userTweetTimeline := timeline + "?id=" + id + "/tweets"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, userTweetTimeline, nil)
	if err != nil {
		return nil, fmt.Errorf("user tweet timeline new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt TweetOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("user tweet timeline: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("user tweet timeline response: %w", err)
	}
	defer resp.Body.Close()

	var timelines UserTweetTimelineResponse
	if err := json.NewDecoder(resp.Body).Decode(&timelines); err != nil {
		return nil, fmt.Errorf("user tweet timeline decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &timelines, &HTTPError{
			APIName: "user tweet timeline",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &timelines, nil
}

func userMensionTimeline(ctx context.Context, c *client, id string, opt ...*TweetOption) (*UserMensionTimelineResponse, error) {
	// check id
	if len(id) == 0 {
		return nil, errors.New("user mension timeline: id parameter is required")
	}
	userMensionTimeline := timeline + "?id=" + id + "/mensions"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, userMensionTimeline, nil)
	if err != nil {
		return nil, fmt.Errorf("user mension timeline new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt TweetOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("user mension timeline: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("user mension timeline response: %w", err)
	}
	defer resp.Body.Close()

	var timelines UserMensionTimelineResponse
	if err := json.NewDecoder(resp.Body).Decode(&timelines); err != nil {
		return nil, fmt.Errorf("user mension timeline decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &timelines, &HTTPError{
			APIName: "user mension timeline",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &timelines, nil
}
