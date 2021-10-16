package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func retrieveStreamRules(ctx context.Context, c *client, opt ...*RetrieveStreamRulesOption) (*RetrieveStreamRulesResponse, error) {
	filteredStreamRules := filteredStream + "/rules"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, filteredStreamRules, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve stream rules new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt RetrieveStreamRulesOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("retrieve stream rules: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("retrieve stream rules: %w", err)
	}
	defer resp.Body.Close()

	var tweet RetrieveStreamRulesResponse
	if err := json.NewDecoder(resp.Body).Decode(&tweet); err != nil {
		return nil, fmt.Errorf("retrieve stream rules decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &tweet, &HTTPError{
			APIName: "retrieve stream rules",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &tweet, nil
}

func addOrDeleteRules(ctx context.Context, c *client, body *AddOrDeleteJSONBody, opt ...*AddOrDeleteRulesOption) (*AddOrDeleteRulesResponse, error) {
	filteredStreamRules := filteredStream + "/rules"

	switch {
	case len(body.Add) == 0 && len(body.Delete.IDs) == 0:
		return nil, errors.New("add or delete rules: add or delete.ids are required")
	default:
	}

	for _, rule := range body.Add {
		switch {
		case len(rule.Value) == 0:
			return nil, errors.New("add or delete rules: add is required")
		case 512 < len(rule.Value):
			return nil, errors.New("add or delete rules: delete is required")
		default:
		}
	}

	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("add or delete rules : can not marshal")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, filteredStreamRules, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("add or delete rules new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "application/json")

	var topt AddOrDeleteRulesOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("add or delete rules: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("add or delete rules: %w", err)
	}
	defer resp.Body.Close()

	var addOrDelete AddOrDeleteRulesResponse
	if err := json.NewDecoder(resp.Body).Decode(&addOrDelete); err != nil {
		return nil, fmt.Errorf("add or delete rules decode: %w", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return &addOrDelete, &HTTPError{
			APIName: "stream rules",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &addOrDelete, nil
}

func connectToStream(ctx context.Context, c *client, opt ...*ConnectToStreamOption) (*ConnectToStreamResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, filteredStream, nil)
	if err != nil {
		return nil, fmt.Errorf("connect to stream new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt ConnectToStreamOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("connect to stream: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("connect to stream: %w", err)
	}
	defer resp.Body.Close()

	var connectToStream ConnectToStreamResponse
	if err := json.NewDecoder(resp.Body).Decode(&connectToStream); err != nil {
		return nil, fmt.Errorf("connect to stream decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &connectToStream, &HTTPError{
			APIName: "connect to stream",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &connectToStream, nil
}
