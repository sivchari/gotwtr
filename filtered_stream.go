package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type RetrieveStreamRulesOption struct {
	IDs []string
}

type FilteredRule struct {
	ID    string `json:"id"`
	Value string `json:"value"`
	Tag   string `json:"tag,omitempty"`
}

type RetrieveStreamRulesResponse struct {
	Rules  []*FilteredRule          `json:"data"`
	Meta   *RetrieveStreamRulesMeta `json:"meta"`
	Errors []*APIResponseError      `json:"errors,omitempty"`
}

type RetrieveStreamRulesMeta struct {
	Sent string // TODO: Is it number ?
}

func (t *RetrieveStreamRulesOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(t.IDs) > 0 {
		q.Add("ids", strings.Join(t.IDs, ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

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

type AddOrDeleteRulesOption struct {
	DryRun bool // If it is true, test a the syntax of your rule without submitting it
}

type AddOrDeleteRulesResponse struct {
	Rules  []*FilteredRule       `json:"data"`
	Meta   *AddOrDeleteRulesMeta `json:"meta"`
	Errors []*APIResponseError   `json:"errors,omitempty"`
}

type AddOrDeleteRulesMeta struct {
	Sent    string                  `json:"sent"` // TODO: Is it number ?
	Summary *AddOrDeleteMetaSummary `json:"summary"`
}

type AddOrDeleteMetaSummary struct {
	Created    int `json:"created"`
	NotCreated int `json:"not_created"`
	Deleted    int `json:"deleted"`
	NotDeleted int `json:"not_deleted"`
	Valid      int `json:"valid"`
	Invalid    int `json:"invalid"`
}

type AddOrDeleteJSONBody struct {
	Add    []*Add  `json:"add,omitempty"`
	Delete *Delete `json:"delete,omitempty"`
}

type Add struct {
	Value string `json:"value"`
	Tag   string `json:"tag,omitempty"`
}

type Delete struct {
	IDs []string `json:"ids"`
}

func (t *AddOrDeleteRulesOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if t.DryRun {
		q.Add("dry_run", strconv.FormatBool(t.DryRun))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
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
