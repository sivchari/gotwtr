package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func userTweetTimeline(ctx context.Context, c *client, id string, opt ...*UserTweetTimelineOpts) (*UserTweetTimelineResponse, error) {
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

	var topt UserTweetTimelineOpts
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

func userMentionTimeline(ctx context.Context, c *client, id string, opt ...*UserMentionTimelineOpts) (*UserMentionTimelineResponse, error) {
	// check id
	if len(id) == 0 {
		return nil, errors.New("user mention timeline: id parameter is required")
	}
	userMentionTimeline := timeline + "?id=" + id + "/mentions"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, userMentionTimeline, nil)
	if err != nil {
		return nil, fmt.Errorf("user mention timeline new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt UserMentionTimelineOpts
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("user mention timeline: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("user mention timeline response: %w", err)
	}
	defer resp.Body.Close()

	var timelines UserMentionTimelineResponse
	if err := json.NewDecoder(resp.Body).Decode(&timelines); err != nil {
		return nil, fmt.Errorf("user mention timeline decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &timelines, &HTTPError{
			APIName: "user mention timeline",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &timelines, nil
}
