package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func sampledStream(ctx context.Context, c *client, opt ...*SampledStreamOpts) (*SampledStreamResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, sampleStream, nil)
	if err != nil {
		return nil, fmt.Errorf("sampled stream new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt SampledStreamOpts
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("sampled stream: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sampled stream response: %w", err)
	}
	defer resp.Body.Close()

	var sampledStream SampledStreamResponse
	if err := json.NewDecoder(resp.Body).Decode(&sampledStream); err != nil {
		return nil, fmt.Errorf("sampled stream decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &sampledStream, &HTTPError{
			APIName: "sampled stream",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &sampledStream, nil
}
