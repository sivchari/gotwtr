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

func Test_uploadMedia(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx       context.Context
		client    *http.Client
		media     io.Reader
		mediaType string
		opt       []*gotwtr.MediaUploadOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.MediaUploadResponse
		wantErr bool
	}{
		{
			name: "200 ok upload successful",
			args: args{
				ctx:       context.Background(),
				media:     strings.NewReader("fake image data"),
				mediaType: gotwtr.MediaTypeJPEG,
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"media_id": "1234567890123456789",
						"media_key": "2_1234567890123456789",
						"expires_after_secs": 86400
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				opt: []*gotwtr.MediaUploadOption{
					{
						MediaCategory: gotwtr.MediaCategoryTweetImage,
						AltText:       "A sample image",
					},
				},
			},
			want: &gotwtr.MediaUploadResponse{
				MediaID:         "1234567890123456789",
				MediaKey:        "2_1234567890123456789",
				ExpiresAfterSecs: 86400,
			},
			wantErr: false,
		},
		{
			name: "nil media reader",
			args: args{
				ctx:       context.Background(),
				media:     nil,
				mediaType: gotwtr.MediaTypeJPEG,
				client:    http.DefaultClient,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty media type",
			args: args{
				ctx:       context.Background(),
				media:     strings.NewReader("data"),
				mediaType: "",
				client:    http.DefaultClient,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.UploadMedia(tt.args.ctx, tt.args.media, tt.args.mediaType, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.UploadMedia() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.UploadMedia() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func Test_initializeChunkedUpload(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		req    *gotwtr.MediaUploadInitRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.MediaUploadResponse
		wantErr bool
	}{
		{
			name: "201 created chunked upload initialized",
			args: args{
				ctx: context.Background(),
				req: &gotwtr.MediaUploadInitRequest{
					Command:       "INIT",
					MediaType:     gotwtr.MediaTypeMP4,
					TotalBytes:    10485760, // 10MB
					MediaCategory: gotwtr.MediaCategoryTweetVideo,
				},
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"media_id": "9876543210987654321",
						"expires_after_secs": 86400
					}`
					return &http.Response{
						StatusCode: http.StatusCreated,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
			},
			want: &gotwtr.MediaUploadResponse{
				MediaID:         "9876543210987654321",
				ExpiresAfterSecs: 86400,
			},
			wantErr: false,
		},
		{
			name: "nil request",
			args: args{
				ctx:    context.Background(),
				req:    nil,
				client: http.DefaultClient,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty media type",
			args: args{
				ctx: context.Background(),
				req: &gotwtr.MediaUploadInitRequest{
					Command:    "INIT",
					MediaType:  "",
					TotalBytes: 1000,
				},
				client: http.DefaultClient,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.InitializeChunkedUpload(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.InitializeChunkedUpload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.InitializeChunkedUpload() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func Test_appendChunkedUpload(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		req    *gotwtr.MediaUploadAppendRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "204 no content chunk uploaded successfully",
			args: args{
				ctx: context.Background(),
				req: &gotwtr.MediaUploadAppendRequest{
					Command:      "APPEND",
					MediaID:      "9876543210987654321",
					SegmentIndex: 0,
					Media:        []byte("fake video chunk data"),
				},
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					return &http.Response{
						StatusCode: http.StatusNoContent,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader("")),
					}
				}),
			},
			wantErr: false,
		},
		{
			name: "nil request",
			args: args{
				ctx:    context.Background(),
				req:    nil,
				client: http.DefaultClient,
			},
			wantErr: true,
		},
		{
			name: "empty media_id",
			args: args{
				ctx: context.Background(),
				req: &gotwtr.MediaUploadAppendRequest{
					Command:      "APPEND",
					MediaID:      "",
					SegmentIndex: 0,
					Media:        []byte("data"),
				},
				client: http.DefaultClient,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("key", gotwtr.WithHTTPClient(tt.args.client))
			err := c.AppendChunkedUpload(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.AppendChunkedUpload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_finalizeChunkedUpload(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		req    *gotwtr.MediaUploadFinalizeRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.MediaUploadResponse
		wantErr bool
	}{
		{
			name: "200 ok finalize successful with processing",
			args: args{
				ctx: context.Background(),
				req: &gotwtr.MediaUploadFinalizeRequest{
					Command: "FINALIZE",
					MediaID: "9876543210987654321",
				},
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"media_id": "9876543210987654321",
						"media_key": "2_9876543210987654321",
						"processing_info": {
							"state": "pending",
							"check_after_secs": 5
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
			},
			want: &gotwtr.MediaUploadResponse{
				MediaID:  "9876543210987654321",
				MediaKey: "2_9876543210987654321",
				ProcessingInfo: &gotwtr.MediaProcessingInfo{
					State:          "pending",
					CheckAfterSecs: 5,
				},
			},
			wantErr: false,
		},
		{
			name: "nil request",
			args: args{
				ctx:    context.Background(),
				req:    nil,
				client: http.DefaultClient,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.FinalizeChunkedUpload(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.FinalizeChunkedUpload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.FinalizeChunkedUpload() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func Test_checkUploadStatus(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		req    *gotwtr.MediaUploadStatusRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.MediaUploadResponse
		wantErr bool
	}{
		{
			name: "200 ok processing complete",
			args: args{
				ctx: context.Background(),
				req: &gotwtr.MediaUploadStatusRequest{
					Command: "STATUS",
					MediaID: "9876543210987654321",
				},
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"media_id": "9876543210987654321",
						"media_key": "2_9876543210987654321",
						"processing_info": {
							"state": "succeeded",
							"progress_percent": 100
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
			},
			want: &gotwtr.MediaUploadResponse{
				MediaID:  "9876543210987654321",
				MediaKey: "2_9876543210987654321",
				ProcessingInfo: &gotwtr.MediaProcessingInfo{
					State:           "succeeded",
					ProgressPercent: 100,
				},
			},
			wantErr: false,
		},
		{
			name: "nil request",
			args: args{
				ctx:    context.Background(),
				req:    nil,
				client: http.DefaultClient,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.CheckUploadStatus(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CheckUploadStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.CheckUploadStatus() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}