package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
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
		case filteredStreamRuleMaxLength < len(rule.Value):
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
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return &addOrDelete, &HTTPError{
			APIName: "add or delete",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &addOrDelete, nil
}


func (s *ConnectToStream) Stop() {
	close(s.done)
	s.wg.Wait()
}

func (s *ConnectToStream) retry(req *http.Request) {
	defer s.wg.Done()
	resp, err := s.client.Do(req)
	if err != nil {
		s.errCh <- err
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		s.errCh <- &HTTPError{
			APIName: "connect to stream",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
		fmt.Println(resp.Body)
	}
	dec := json.NewDecoder(resp.Body)

	for !stopped(s.done) {
		var connectToStream ConnectToStreamResponse
		err := dec.Decode(&connectToStream)
		if err != nil {
			if err == io.EOF {
				break
			}
			s.errCh <- err
		}
		s.ch <- connectToStream
	}
}

func connectToStream(ctx context.Context, c *client, ch chan<- ConnectToStreamResponse, errCh chan<- error, opt ...*ConnectToStreamOption) *ConnectToStream {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, filteredStream, nil)
	if err != nil {
		errCh <- fmt.Errorf("connect to stream new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var copt ConnectToStreamOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		copt = *opt[0]
	default:
		errCh<-errors.New("connect to stream: only one option is allowed")
	}
	copt.addQuery(req)

	s := &ConnectToStream{
		client: c.client,
		errCh:  errCh,
		ch:     ch,
		done:   make(chan struct{}),
		wg:     &sync.WaitGroup{},
	}

	s.wg.Add(1)
	go s.retry(req)
	return s
}