package gotwtr_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/sivchari/gotwtr"
)

func Test_sampledStream(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		opt    []*gotwtr.VolumeStreamsOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.VolumeStreamsResponse
		wantErr bool
	}{
		{
			name: "200 ok default payload",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					body := `{
						"data": {
							"id": "1067094924124872705",
							"text": "Just getting started with Twitter APIs? Find out what you need in order to build an app. Watch this video! https://t.co/Hg8nkfoizN"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				opt: []*gotwtr.VolumeStreamsOption{},
			},
			want: &gotwtr.VolumeStreamsResponse{
				Tweet: &gotwtr.Tweet{
					ID:   "1067094924124872705",
					Text: "Just getting started with Twitter APIs? Find out what you need in order to build an app. Watch this video! https://t.co/Hg8nkfoizN",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ch := make(chan gotwtr.VolumeStreamsResponse)
			errCh := make(chan error)
			c := gotwtr.New("key", gotwtr.WithHTTPClient(tt.args.client))
			c.VolumeStreams(tt.args.ctx, ch, errCh, tt.args.opt...)
			select {
			case got := <-ch:
				if diff := cmp.Diff(tt.want, &got); diff != "" {
					t.Errorf("client.VolumeStreams() mismatch (-want +got):\n%s", diff)
					return
				}
			case err := <-errCh:
				if (err != nil) != tt.wantErr {
					t.Errorf("client.VolumeStreams() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}

func Test_sampledStream10(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		opt    []*gotwtr.VolumeStreamsOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.VolumeStreamsResponse
		wantErr bool
	}{
		{
			name: "200 ok default payload",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					body := `{
						"data": {
							"id": "1067094924124872705",
							"text": "Just getting started with Twitter APIs? Find out what you need in order to build an app. Watch this video! https://t.co/Hg8nkfoizN"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				opt: []*gotwtr.VolumeStreamsOption{},
			},
			want: &gotwtr.VolumeStreamsResponse{
				Tweet: &gotwtr.Tweet{
					ID:   "1067094924124872705",
					Text: "Just getting started with Twitter APIs? Find out what you need in order to build an app. Watch this video! https://t.co/Hg8nkfoizN",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ch := make(chan gotwtr.VolumeStreamsResponse)
			errCh := make(chan error)
			c := gotwtr.New("key", gotwtr.WithHTTPClient(tt.args.client))
			c.VolumeStreams10(tt.args.ctx, ch, errCh, tt.args.opt...)
			select {
			case got := <-ch:
				if diff := cmp.Diff(tt.want, &got); diff != "" {
					t.Errorf("client.VolumeStreams10() mismatch (-want +got):\n%s", diff)
					return
				}
			case err := <-errCh:
				if (err != nil) != tt.wantErr {
					t.Errorf("client.VolumeStreams10() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}
