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

func Test_CreateNewList(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		body   *gotwtr.CreateNewListBody
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.CreateNewListResponse
		wantErr bool
	}{
		{
			name: "200 Create a list",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": {
								"id": "1441162269824405510",
								"name": "test v2 create list"
                        }
                    }`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				body: &gotwtr.CreateNewListBody{},
			},
			want: &gotwtr.CreateNewListResponse{
				Data: &gotwtr.CreateNewListData{
					ID:   "1441162269824405510",
					Name: "test v2 create list",
				},
			},
			wantErr: false,
		},
		{
			name: "200 Create a list Option",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
						"data": {
								"id": "1441162269824405511",
								"name": "name-for-new-list"
                        }
                    }`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				body: &gotwtr.CreateNewListBody{
					Name:        "name-for-new-list",
					Description: "description-of-list",
					Private:     false,
				},
			},
			want: &gotwtr.CreateNewListResponse{
				Data: &gotwtr.CreateNewListData{
					ID:   "1441162269824405511",
					Name: "name-for-new-list",
				},
			},
			wantErr: false,
		},
		/*
		        {
		            name: "404 not found",
		            args: args{
		                ctx: context.Background(),
		                client: mockHTTPClient(func(req *http.Request) *http.Response {
		                    body := `{
								"errors":[
									{
										"parameters":{
											"id":[
												"111111111111111111111111"
											]
										},
										"message":"The id query parameter value [111111111111111111111111] is not valid"
									}
								],
								"title":"Invalid Request",
								"detail":"One or more parameters to your request was invalid.",
								"type":"https://api.twitter.com/2/problems/invalid-request"
							}`
		                    return &http.Response{
		                        StatusCode: http.StatusNotFound,
		                        Body:       io.NopCloser(strings.NewReader(body)),
		                    }
		                }),
		                body: &gotwtr.CreateNewListBody{},
		            },
		            want: &gotwtr.CreateNewListResponse{
		                Data: nil,
		                Errors: []*gotwtr.APIResponseError{
		                    {
		                        Parameters: gotwtr.Parameter{
		                            ID: []string{"111111111111111111111111"},
		                        },
		                        Message: "The id query parameter value [111111111111111111111111] is not valid",
		                    },
		                },
		                Meta:   nil,
		                Title:  "Invalid Request",
		                Detail: "One or more parameters to your request was invalid.",
		                Type:   "https://api.twitter.com/2/problems/invalid-request",
		            },
		            wantErr: true,
		        },
		*/
	}
	for i, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.CreateNewList(tt.args.ctx, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CreateNewList() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.CreateNewList() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}

func Test_DeleteList(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		listID string
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.DeleteListResponse
		wantErr bool
	}{
		{
			name: "200 Delete a list",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
                      "data": {
                        "deleted": true
                       }
                    }`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				listID: "1441162269824405510",
			},
			want: &gotwtr.DeleteListResponse{
				Data: &gotwtr.DeleteListData{
					Deleted: true,
				},
			},
			wantErr: false,
		},
		/*
		        {
		            name: "404 not found",
		            args: args{
		                ctx: context.Background(),
		                client: mockHTTPClient(func(req *http.Request) *http.Response {
		                    body := `{
								"errors":[
									{
										"parameters":{
											"id":[
												"111111111111111111111111"
											]
										},
										"message":"The id query parameter value [111111111111111111111111] is not valid"
									}
								],
								"title":"Invalid Request",
								"detail":"One or more parameters to your request was invalid.",
								"type":"https://api.twitter.com/2/problems/invalid-request"
							}`
		                    return &http.Response{
		                        StatusCode: http.StatusNotFound,
		                        Body:       io.NopCloser(strings.NewReader(body)),
		                    }
		                }),
		                listID: "",
		            },
		            want: &gotwtr.DeleteListResponse{
		                Data: nil,
		                Errors: []*gotwtr.APIResponseError{
		                    {
		                        Parameters: gotwtr.Parameter{
		                            ID: []string{"111111111111111111111111"},
		                        },
		                        Message: "The id query parameter value [111111111111111111111111] is not valid",
		                    },
		                },
		                Meta:   nil,
		                Title:  "Invalid Request",
		                Detail: "One or more parameters to your request was invalid.",
		                Type:   "https://api.twitter.com/2/problems/invalid-request",
		            },
		            wantErr: true,
		        },
		*/
	}
	for i, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.DeleteList(tt.args.ctx, tt.args.listID)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.DeleteList() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.DeleteList() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}

func Test_UpdateMetaDataForList(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		listID string
		body   []*gotwtr.UpdateMetaDataForListBody
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.UpdateMetaDataForListResponse
		wantErr bool
	}{
		{
			name: "200 Update a list",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
                        "data": {
                            "updated": true
                        }
                    }`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				listID: "1441163524802158595",
				body:   nil,
			},
			want: &gotwtr.UpdateMetaDataForListResponse{
				Data: &gotwtr.UpdateMetaDataForListData{
					Updated: true,
				},
			},
			wantErr: false,
		},
		{
			name: "200 Update a list option",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(req *http.Request) *http.Response {
					body := `{
                        "data": {
                            "updated": true
                        }
                    }`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				listID: "1441163524802158595",
				body: []*gotwtr.UpdateMetaDataForListBody{
					{
						Name: "test v2 update list",
					},
				},
			},
			want: &gotwtr.UpdateMetaDataForListResponse{
				Data: &gotwtr.UpdateMetaDataForListData{
					Updated: true,
				},
			},
			wantErr: false,
		},
		/*
		           {
		               name: "404 not found",
		               args: args{
		                   ctx: context.Background(),
		                   client: mockHTTPClient(func(req *http.Request) *http.Response {
		                       body := `{
		   						"errors":[
		   							{
		   								"parameters":{
		   									"id":[
		   										"111111111111111111111111"
		   									]
		   								},
		   								"message":"The id query parameter value [111111111111111111111111] is not valid"
		   							}
		   						],
		   						"title":"Invalid Request",
		   						"detail":"One or more parameters to your request was invalid.",
		   						"type":"https://api.twitter.com/2/problems/invalid-request"
		   					}`
		                       return &http.Response{
		                           StatusCode: http.StatusNotFound,
		                           Body:       io.NopCloser(strings.NewReader(body)),
		                       }
		                   }),
		                   listID: "",
		               },
		               want: &gotwtr.DeleteListResponse{
		                   Data: nil,
		                   Errors: []*gotwtr.APIResponseError{
		                       {
		                           Parameters: gotwtr.Parameter{
		                               ID: []string{"111111111111111111111111"},
		                           },
		                           Message: "The id query parameter value [111111111111111111111111] is not valid",
		                       },
		                   },
		                   Meta:   nil,
		                   Title:  "Invalid Request",
		                   Detail: "One or more parameters to your request was invalid.",
		                   Type:   "https://api.twitter.com/2/problems/invalid-request",
		               },
		               wantErr: true,
		           },
		*/
	}
	for i, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.UpdateMetaDataForList(tt.args.ctx, tt.args.listID, tt.args.body...)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.UpdateMetaDataForList() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("client.UpdateMetaDataForList() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}
