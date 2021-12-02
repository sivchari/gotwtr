package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func userTweetTimeline(ctx context.Context, c *client, userID string, opt ...*UserTweetTimelineOption) (*UserTweetTimelineResponse, error) {
	if userID == "" {
		return nil, errors.New("user tweet timeline: id parameter is required")
	}
	ep := fmt.Sprintf(userTweetTimelineURL, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("user tweet timeline new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var uopt UserTweetTimelineOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		uopt = *opt[0]
	default:
		return nil, errors.New("user tweet timeline: only one option is allowed")
	}
	const (
		minimumMaxResults = 5
		maximumMaxResults = 100
		defaultMaxResults = 10
	)
	if uopt.MaxResults == 0 {
		uopt.MaxResults = defaultMaxResults
	}
	if uopt.MaxResults < minimumMaxResults || uopt.MaxResults > maximumMaxResults {
		return nil, fmt.Errorf("user tweet timeline: max results must be between %d and %d", minimumMaxResults, maximumMaxResults)
	}
	uopt.addQuery(req)

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

func userMentionTimeline(ctx context.Context, c *client, userID string, opt ...*UserMentionTimelineOption) (*UserMentionTimelineResponse, error) {
	if userID == "" {
		return nil, errors.New("user mention timeline: userID parameter is required")
	}
	ep := fmt.Sprintf(userMentionTimelineURL, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("user mention timeline new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var uopt UserMentionTimelineOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		uopt = *opt[0]
	default:
		return nil, errors.New("user mention timeline: only one option is allowed")
	}
	const (
		minimumMaxResults = 5
		maximumMaxResults = 100
		defaultMaxResults = 10
	)
	if uopt.MaxResults == 0 {
		uopt.MaxResults = defaultMaxResults
	}
	if uopt.MaxResults < minimumMaxResults || uopt.MaxResults > maximumMaxResults {
		return nil, fmt.Errorf("user mention timeline: max results must be between %d and %d", minimumMaxResults, maximumMaxResults)
	}
	uopt.addQuery(req)

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
