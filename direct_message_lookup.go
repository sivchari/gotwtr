package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func lookUpAllOneToOneDM(ctx context.Context, c *client, participantID string, opt ...*DirectMessageOption) (*LookUpAllOneToOneDMResponse, error) {
	if participantID == "" {
		return nil, errors.New("lookup all one to one DM: participant id parameter is required")
	}
	ep := fmt.Sprintf(lookUpAllOneToOneDMURL, participantID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("lookup all one to one DM with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var dmopt DirectMessageOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		dmopt = *opt[0]
	default:
		return nil, errors.New("lookup all one to one DM: only one option is allowed")
	}
	dmopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("lookup all one to one DM response: %w", err)
	}
	defer resp.Body.Close()

	var lookUpAllOneToOneDM LookUpAllOneToOneDMResponse
	if err := json.NewDecoder(resp.Body).Decode(&lookUpAllOneToOneDM); err != nil {
		return nil, fmt.Errorf("lookup all DM decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &lookUpAllOneToOneDM, &HTTPError{
			APIName: "lookup all one to one DM",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}
	return &lookUpAllOneToOneDM, nil
}

func lookUpDM(ctx context.Context, c *client, dmConversationID string, opt ...*DirectMessageOption) (*LookUpDMResponse, error) {
	if dmConversationID == "" {
		return nil, errors.New("lookup DM: dm conversation id parameter is required")
	}
	ep := fmt.Sprintf(lookUpDMURL, dmConversationID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("lookup DM with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var dmopt DirectMessageOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		dmopt = *opt[0]
	default:
		return nil, errors.New("lookup DM: only one option is allowed")
	}
	dmopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("lookup DM response: %w", err)
	}
	defer resp.Body.Close()

	var lookUpDM LookUpDMResponse
	if err := json.NewDecoder(resp.Body).Decode(&lookUpDM); err != nil {
		return nil, fmt.Errorf("lookup DM decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &lookUpDM, &HTTPError{
			APIName: "lookup DM",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}
	return &lookUpDM, nil
}

func lookUpAllDM(ctx context.Context, c *client, opt ...*DirectMessageOption) (*LookUpAllDMResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, lookUpAllDMURL, nil)
	if err != nil {
		return nil, fmt.Errorf("lookup all DM with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var dmopt DirectMessageOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		dmopt = *opt[0]
	default:
		return nil, errors.New("lookup all DM: only one option is allowed")
	}
	dmopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("lookup all DM response: %w", err)
	}
	defer resp.Body.Close()

	var lookUpAllDM LookUpAllDMResponse
	if err := json.NewDecoder(resp.Body).Decode(&lookUpAllDM); err != nil {
		return nil, fmt.Errorf("lookup all DM decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &lookUpAllDM, &HTTPError{
			APIName: "lookup all DM",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}
	return &lookUpAllDM, nil
}
