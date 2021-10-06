package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type FilteredStreamOption struct {
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

func (t *FilteredStreamOption) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(t.IDs) > 0 {
		q.Add("ids", strings.Join(t.IDs, ","))
	}
	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

func retrieveStreamRules(ctx context.Context, c *client, opt ...*FilteredStreamOption) (*RetrieveStreamRulesResponse, error) {
	filteredStreamRules := filteredStream + "/rules"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, filteredStreamRules, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve stream rules new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt FilteredStreamOption
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
