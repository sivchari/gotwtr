package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

func stopped(done <-chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func (s *VolumeStreams) Stop() {
	close(s.done)
	s.wg.Wait()
}

func (s *VolumeStreams) retry(req *http.Request) {
	defer s.wg.Done()
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
		return
	}
	dec := json.NewDecoder(resp.Body)
	for !stopped(s.done) {
		var v VolumeStreamsResponse
		err := dec.Decode(&v)
		if err != nil {
			if err == io.EOF {
				break
			}
			s.errCh <- err
		}
		s.ch <- v
	}
}

func volumeStreams(ctx context.Context, c *client, ch chan<- VolumeStreamsResponse, errCh chan<- error, opt ...*VolumeStreamsOption) *VolumeStreams {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, volumeStreamsURL, nil)
	if err != nil {
		errCh <- fmt.Errorf("sampled stream new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var vopt VolumeStreamsOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		vopt = *opt[0]
	default:
		errCh <- errors.New("sampled stream: only one option is allowed")
	}
	vopt.addQuery(req)

	vs := &VolumeStreams{
		client: c.client,
		errCh:  errCh,
		ch:     ch,
		done:   make(chan struct{}),
		wg:     &sync.WaitGroup{},
	}
	vs.wg.Add(1)
	go vs.retry(req)
	return vs
}
