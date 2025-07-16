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

func Test_searchPostsEligibleForNotes(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		opt    []*gotwtr.SearchPostsEligibleForNotesOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.SearchPostsEligibleForNotesResponse
		wantErr bool
	}{
		{
			name: "200 ok the request was successful",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"data": [
							{
								"id": "1234567890",
								"text": "This is a sample post eligible for Community Notes",
								"author_id": "987654321"
							}
						],
						"meta": {
							"result_count": 1
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				opt: []*gotwtr.SearchPostsEligibleForNotesOption{
					{
						MaxResults: 10,
					},
				},
			},
			want: &gotwtr.SearchPostsEligibleForNotesResponse{
				Data: []*gotwtr.Post{
					{
						ID:       "1234567890",
						Text:     "This is a sample post eligible for Community Notes",
						AuthorID: "987654321",
					},
				},
				Meta: &gotwtr.SearchNotesMeta{
					ResultCount: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "404 not found",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"title": "Not Found Error",
						"detail": "The requested resource was not found.",
						"type": "https://api.x.com/2/problems/resource-not-found"
					}`
					return &http.Response{
						StatusCode: http.StatusNotFound,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
			},
			want: &gotwtr.SearchPostsEligibleForNotesResponse{
				Title:  "Not Found Error",
				Detail: "The requested resource was not found.",
				Type:   "https://api.x.com/2/problems/resource-not-found",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.SearchPostsEligibleForNotes(tt.args.ctx, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.SearchPostsEligibleForNotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.SearchPostsEligibleForNotes() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func Test_searchNotesWritten(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		opt    []*gotwtr.SearchNotesWrittenOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.SearchNotesWrittenResponse
		wantErr bool
	}{
		{
			name: "200 ok the request was successful",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"data": [
							{
								"id": "note123",
								"text": "This claim needs more context...",
								"author_id": "user123",
								"post_id": "1234567890",
								"classification": "needs_more_context"
							}
						],
						"meta": {
							"result_count": 1
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				opt: []*gotwtr.SearchNotesWrittenOption{
					{
						MaxResults: 10,
					},
				},
			},
			want: &gotwtr.SearchNotesWrittenResponse{
				Data: []*gotwtr.CommunityNote{
					{
						ID:             "note123",
						Text:           "This claim needs more context...",
						AuthorID:       "user123",
						PostID:         "1234567890",
						Classification: "needs_more_context",
					},
				},
				Meta: &gotwtr.SearchNotesMeta{
					ResultCount: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.SearchNotesWritten(tt.args.ctx, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.SearchNotesWritten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.SearchNotesWritten() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}

func Test_createCommunityNote(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		body   *gotwtr.CreateCommunityNoteBody
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.CreateCommunityNoteResponse
		wantErr bool
	}{
		{
			name: "201 created successfully",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"data": {
							"id": "note456",
							"text": "This post needs additional context to be understood correctly.",
							"author_id": "user123",
							"post_id": "1234567890",
							"classification": "needs_more_context"
						}
					}`
					return &http.Response{
						StatusCode: http.StatusCreated,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				body: &gotwtr.CreateCommunityNoteBody{
					TestMode: true,
					PostID:   "1234567890",
					Info: &gotwtr.CommunityNoteInfoRequest{
						Text:           "This post needs additional context to be understood correctly.",
						Classification: "needs_more_context",
					},
				},
			},
			want: &gotwtr.CreateCommunityNoteResponse{
				Data: &gotwtr.CommunityNote{
					ID:             "note456",
					Text:           "This post needs additional context to be understood correctly.",
					AuthorID:       "user123",
					PostID:         "1234567890",
					Classification: "needs_more_context",
				},
			},
			wantErr: false,
		},
		{
			name: "missing body",
			args: args{
				ctx:    context.Background(),
				client: http.DefaultClient,
				body:   nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing post_id",
			args: args{
				ctx:    context.Background(),
				client: http.DefaultClient,
				body: &gotwtr.CreateCommunityNoteBody{
					TestMode: true,
					PostID:   "",
					Info: &gotwtr.CommunityNoteInfoRequest{
						Text:           "This post needs additional context.",
						Classification: "needs_more_context",
					},
				},
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
			got, err := c.CreateCommunityNote(tt.args.ctx, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CreateCommunityNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.CreateCommunityNote() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}