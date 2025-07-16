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

func Test_generateAppOnlyBearerToken(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx            context.Context
		client         *http.Client
		consumerKey    string
		consumerSecret string
	}
	tests := []struct {
		name string
		args args
		want struct {
			BearerToken    string
			ConsumerKey    string
			ConsumerSecret string
		}
		wantErr bool
	}{
		{
			name: "200 ok",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"token_type": "bearer",
						"access_token": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				consumerKey:    "consumerKey",
				consumerSecret: "consumerSecret",
			},
			want: struct {
				BearerToken    string
				ConsumerKey    string
				ConsumerSecret string
			}{
				BearerToken:    "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
				ConsumerKey:    "consumerKey",
				ConsumerSecret: "consumerSecret",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New(
				"key",
				gotwtr.WithHTTPClient(tt.args.client),
				gotwtr.WithConsumerKey(tt.args.consumerKey),
				gotwtr.WithConsumerSecret(tt.args.consumerSecret),
			)
			b, err := c.GenerateAppOnlyBearerToken(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GenerateAppOnlyBearerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !b {
				t.Errorf("client.GenerateAppOnlyBearerToken() = %v, want %v", b, true)
				return
			}
			nowc := c.ExportClient()
			cstate := struct {
				BearerToken    string
				ConsumerKey    string
				ConsumerSecret string
			}{
				BearerToken:    nowc["bearerToken"],
				ConsumerKey:    nowc["consumerKey"],
				ConsumerSecret: nowc["consumerSecret"],
			}
			if diff := cmp.Diff(tt.want, cstate); diff != "" {
				t.Errorf("client.GenerateAppOnlyBearerToken() diff = %v", diff)
				return
			}
		})
	}
}

func Test_InvalidateToken(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx         context.Context
		client      *http.Client
		bearerToken string
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.InvalidateTokenResponse
		wantErr bool
	}{
		{
			name: "200 ok",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"access_token": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				bearerToken: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			},
			want: &gotwtr.InvalidateTokenResponse{
				AccessToken: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			},
			wantErr: false,
		},
		{
			name: "no bearer token",
			args: args{
				ctx:         context.Background(),
				client:      mockHTTPClient(func(request *http.Request) *http.Response { return nil }),
				bearerToken: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New(
				tt.args.bearerToken,
				gotwtr.WithHTTPClient(tt.args.client),
				gotwtr.WithConsumerKey("consumerKey"),
				gotwtr.WithConsumerSecret("consumerSecret"),
			)
			got, err := c.InvalidateToken(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.InvalidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.InvalidateToken() diff = %v", diff)
				return
			}
		})
	}
}
