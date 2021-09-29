package gotwtr_test

import "net/http"

type roundTripFunc func(request *http.Request) *http.Response

func (rf roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return rf(request), nil
}

func mockHTTPClient(fn roundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
