package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func stopped(done <-chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func (s *StreamResponse) Stop() {
	close(s.done)
	s.wg.Wait()
}

func (s *StreamResponse) retry(req *http.Request) {
	defer s.wg.Done()
	for !stopped(s.done) {
		resp, err := s.client.Do(req)
		if err != nil {
			s.errCh <- err
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			s.errCh <- &HTTPError{
				APIName: "sampled stream",
				Status:  resp.Status,
				URL:     req.URL.String(),
			}
		}
		var res SampledStreamResponse
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			s.errCh <- err
		}
		s.ch <- res
		/*
			App rate limit: 50 requests per 15-minute window
			FYI https://developer.twitter.com/en/docs/twitter-api/tweets/volume-streams/api-reference/get-tweets-sample-stream
		*/
		time.Sleep(time.Second * 18)
	}
}

func sampledStream(ctx context.Context, c *client, ch chan<- SampledStreamResponse, errCh chan<- error, opt ...*SampledStreamOpts) *StreamResponse {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, sampleStream, nil)
	if err != nil {
		errCh <- fmt.Errorf("sampled stream new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var sopt SampledStreamOpts
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		sopt = *opt[0]
	default:
		errCh <- errors.New("sampled stream: only one option is allowed")
	}
	sopt.addQuery(req)

	s := &StreamResponse{
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
