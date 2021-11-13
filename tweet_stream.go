package gotwtr

import (
	"net/http"
	"sync"
)

type SampledStreamResponse struct {
	Tweets   []*Tweet            `json:"data"`
	Includes *TweetIncludes      `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
}

type StreamResponse struct {
	client *http.Client
	errCh  chan<- error
	ch     chan<- SampledStreamResponse
	done   chan struct{}
	wg     *sync.WaitGroup
}
