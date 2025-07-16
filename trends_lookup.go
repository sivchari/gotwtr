package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func trendsByWOEID(ctx context.Context, c *client, woeid string) (*TrendsByWOEIDResponse, error) {
	if woeid == "" {
		return nil, errors.New("trends by woeid: woeid parameter is required")
	}

	ep := fmt.Sprintf(trendsByWOEIDURL, woeid)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("trends by woeid new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("trends by woeid response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	var tr TrendsByWOEIDResponse
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return nil, fmt.Errorf("trends by woeid decode: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &tr, &HTTPError{
			APIName: "trends by woeid",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &tr, nil
}