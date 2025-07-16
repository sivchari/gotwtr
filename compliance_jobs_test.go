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

func Test_complianceJobs(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		copt   *gotwtr.ComplianceJobsOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.ComplianceJobsResponse
		wantErr bool
	}{
		{
			name: "200 ok",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `
                    {
                        "data": [
                            {
                                "type": "tweets",
                                "id": "1421185651106480129",
                                "resumable": false,
                                "upload_url": "https://storage.googleapis.com/twttr-tweet-compliance/1421185651106480129/submission/1202726487847104512_1421185651106480129?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=complianceapi-public-svc-acct%40twttr-compliance-public-prod.iam.gserviceaccount.com%2F20210730%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20210730T190718Z&X-Goog-Expires=900&X-Goog-SignedHeaders=content-type%3Bhost&X-Goog-Signature=5c197d6e2b54cd941006904d7f96a3ac4985f19b5b770c4b334d4defe495ccebda71650d8636fcfc7266b8e609c0de29255b6b46bf1ad883522fac78010a2936fdd46dd3afa1925a674311b51e1d6d19ab249aa51cc6d1afb65203847a1f998be41aff209d465d74d20b4b26898951035808afd5bd022445d0aeb7ffd8aa20486ee1b3e2ea3b6f9709dfd849fbdfacfb1542dca965d8473e6bfc9596df85fd1be716dd7ebbb4c6b995a0775472145bd778ec4175f2934f2823b21bba6604696301168e55d614098512ffee2bd1b0e363106fc6197e15d833c41bf83598dc4ce7e7f7e2edea0c07e3a55f815e1e28abeb5a24a3e5768fbaa70cf19e85c269d530",
                                "upload_expires_at": "2021-07-30T19:22:18.000Z",
                                "download_expires_at": "2021-08-06T19:07:18.000Z",
                                "download_url": "https://storage.googleapis.com/twttr-tweet-compliance/1421185651106480129/delivery/1202726487847104512_1421185651106480129?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=complianceapi-public-svc-acct%40twttr-compliance-public-prod.iam.gserviceaccount.com%2F20210805%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20210805T013511Z&X-Goog-Expires=604800&X-Goog-SignedHeaders=host&X-Goog-Signature=09de4feae68a6d4449eb7ce1f8f3551996552e7fba103005b3bd50ab318bb5215e4f5396ef29d17755deb6bf172b9d1dab61a04b249d39e87f6e2dbb31632b7e5f2d35f4f534e1f1522c9d7958b8745dd62471deb8d6992c80fd418628404f5f14eda3f557adf709403058910ea009e0c88ce81458ec9b915016a5c5901e2365b130db00b18fcb7da1b082e1a5c75f7bf7eeab8783675d1b6a56441ac6e9ffc972b1278a5853d2b94dda55e1a6e2068bc0ddd3cddc9213ec9cebb7cb5be931977bb28dda12c7c5e69d1f876b243f0f224076bf1b81149603319a2fc9cb82337bdbe05e7bbf184bcbdc17d43b3f5efbae72ea386d955ca10e702e00df31aabf32",
                                "created_at": "2021-07-30T19:07:18.000Z",
                                "status": "complete"
                            },
                            {
                                "type": "tweets",
                                "id": "1423095206576984067",
                                "resumable": false,
                                "upload_url": "https://storage.googleapis.com/twttr-tweet-compliance/1423095206576984067/submission/1202726487847104512_1423095206576984067?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=complianceapi-public-svc-acct%40twttr-compliance-public-prod.iam.gserviceaccount.com%2F20210805%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20210805T013511Z&X-Goog-Expires=900&X-Goog-SignedHeaders=content-type%3Bhost&X-Goog-Signature=ba08f588bea3873aa0465cf22015e583c2851a5ff14891d22430b1127288728f1aa303673e6895694e7017739871ff5ae59bbcde7d4ac7a14aaaafba98ad22ca818e99fb3ec7eaaf74b3ecfecbfb33711869b2e85d7666609276666ef4a8b396ae9616743a0cbd773962e5850f2942cd76be7373d608a140e041ca8492017d43fac9220fa145d0b2ecaf9f752d71fc8c4b81b67c5c22aa59ac87666f7d83714fdace72894d2911a3e36dd42028d0222e71054d6b28c8ef63d0f0000f228c8680bab9c8011b87d1a6c9a60e8cc9e8b6a83abf7c47a57772746c83b19849f5b4c938ccd0922990da5f2a81ff806edcb4667bb402fb1f1f6f5162768e0661648b21",
                                "upload_expires_at": "2021-08-05T01:50:11.000Z",
                                "download_expires_at": "2021-08-12T01:35:11.000Z",
                                "download_url": "https://storage.googleapis.com/twttr-tweet-compliance/1423095206576984067/delivery/1202726487847104512_1423095206576984067?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=complianceapi-public-svc-acct%40twttr-compliance-public-prod.iam.gserviceaccount.com%2F20210805%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20210805T013511Z&X-Goog-Expires=604800&X-Goog-SignedHeaders=host&X-Goog-Signature=09de4feae68a6d4449eb7ce1f8f3551996552e7fba103005b3bd50ab318bb5215e4f5396ef29d17755deb6bf172b9d1dab61a04b249d39e87f6e2dbb31632b7e5f2d35f4f534e1f1522c9d7958b8745dd62471deb8d6992c80fd418628404f5f14eda3f557adf709403058910ea009e0c88ce81458ec9b915016a5c5901e2365b130db00b18fcb7da1b082e1a5c75f7bf7eeab8783675d1b6a56441ac6e9ffc972b1278a5853d2b94dda55e1a6e2068bc0ddd3cddc9213ec9cebb7cb5be931977bb28dda12c7c5e69d1f876b243f0f224076bf1b81149603319a2fc9cb82337bdbe05e7bbf184bcbdc17d43b3f5efbae72ea386d955ca10e702e00df31aabf32",
                                "created_at": "2021-08-05T01:35:11.000Z",
                                "status": "expired"
                            }
                        ]
                    }`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				copt: &gotwtr.ComplianceJobsOption{
					Type: gotwtr.ComplianceFieldTypeTweets,
				},
			},
			want: &gotwtr.ComplianceJobsResponse{
				ComplianceJobsData: []*gotwtr.ComplianceJobsData{
					{
						Type:              "tweets",
						ID:                "1421185651106480129",
						Resumable:         false,
						UploadURL:         "https://storage.googleapis.com/twttr-tweet-compliance/1421185651106480129/submission/1202726487847104512_1421185651106480129?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=complianceapi-public-svc-acct%40twttr-compliance-public-prod.iam.gserviceaccount.com%2F20210730%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20210730T190718Z&X-Goog-Expires=900&X-Goog-SignedHeaders=content-type%3Bhost&X-Goog-Signature=5c197d6e2b54cd941006904d7f96a3ac4985f19b5b770c4b334d4defe495ccebda71650d8636fcfc7266b8e609c0de29255b6b46bf1ad883522fac78010a2936fdd46dd3afa1925a674311b51e1d6d19ab249aa51cc6d1afb65203847a1f998be41aff209d465d74d20b4b26898951035808afd5bd022445d0aeb7ffd8aa20486ee1b3e2ea3b6f9709dfd849fbdfacfb1542dca965d8473e6bfc9596df85fd1be716dd7ebbb4c6b995a0775472145bd778ec4175f2934f2823b21bba6604696301168e55d614098512ffee2bd1b0e363106fc6197e15d833c41bf83598dc4ce7e7f7e2edea0c07e3a55f815e1e28abeb5a24a3e5768fbaa70cf19e85c269d530",
						UploadExpiresAt:   "2021-07-30T19:22:18.000Z",
						DownloadExpiresAt: "2021-08-06T19:07:18.000Z",
						DownloadURL:       "https://storage.googleapis.com/twttr-tweet-compliance/1421185651106480129/delivery/1202726487847104512_1421185651106480129?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=complianceapi-public-svc-acct%40twttr-compliance-public-prod.iam.gserviceaccount.com%2F20210805%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20210805T013511Z&X-Goog-Expires=604800&X-Goog-SignedHeaders=host&X-Goog-Signature=09de4feae68a6d4449eb7ce1f8f3551996552e7fba103005b3bd50ab318bb5215e4f5396ef29d17755deb6bf172b9d1dab61a04b249d39e87f6e2dbb31632b7e5f2d35f4f534e1f1522c9d7958b8745dd62471deb8d6992c80fd418628404f5f14eda3f557adf709403058910ea009e0c88ce81458ec9b915016a5c5901e2365b130db00b18fcb7da1b082e1a5c75f7bf7eeab8783675d1b6a56441ac6e9ffc972b1278a5853d2b94dda55e1a6e2068bc0ddd3cddc9213ec9cebb7cb5be931977bb28dda12c7c5e69d1f876b243f0f224076bf1b81149603319a2fc9cb82337bdbe05e7bbf184bcbdc17d43b3f5efbae72ea386d955ca10e702e00df31aabf32",
						CreatedAt:         "2021-07-30T19:07:18.000Z",
						Status:            "complete",
					},
					{
						Type:              "tweets",
						ID:                "1423095206576984067",
						Resumable:         false,
						UploadURL:         "https://storage.googleapis.com/twttr-tweet-compliance/1423095206576984067/submission/1202726487847104512_1423095206576984067?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=complianceapi-public-svc-acct%40twttr-compliance-public-prod.iam.gserviceaccount.com%2F20210805%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20210805T013511Z&X-Goog-Expires=900&X-Goog-SignedHeaders=content-type%3Bhost&X-Goog-Signature=ba08f588bea3873aa0465cf22015e583c2851a5ff14891d22430b1127288728f1aa303673e6895694e7017739871ff5ae59bbcde7d4ac7a14aaaafba98ad22ca818e99fb3ec7eaaf74b3ecfecbfb33711869b2e85d7666609276666ef4a8b396ae9616743a0cbd773962e5850f2942cd76be7373d608a140e041ca8492017d43fac9220fa145d0b2ecaf9f752d71fc8c4b81b67c5c22aa59ac87666f7d83714fdace72894d2911a3e36dd42028d0222e71054d6b28c8ef63d0f0000f228c8680bab9c8011b87d1a6c9a60e8cc9e8b6a83abf7c47a57772746c83b19849f5b4c938ccd0922990da5f2a81ff806edcb4667bb402fb1f1f6f5162768e0661648b21",
						UploadExpiresAt:   "2021-08-05T01:50:11.000Z",
						DownloadExpiresAt: "2021-08-12T01:35:11.000Z",
						DownloadURL:       "https://storage.googleapis.com/twttr-tweet-compliance/1423095206576984067/delivery/1202726487847104512_1423095206576984067?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=complianceapi-public-svc-acct%40twttr-compliance-public-prod.iam.gserviceaccount.com%2F20210805%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20210805T013511Z&X-Goog-Expires=604800&X-Goog-SignedHeaders=host&X-Goog-Signature=09de4feae68a6d4449eb7ce1f8f3551996552e7fba103005b3bd50ab318bb5215e4f5396ef29d17755deb6bf172b9d1dab61a04b249d39e87f6e2dbb31632b7e5f2d35f4f534e1f1522c9d7958b8745dd62471deb8d6992c80fd418628404f5f14eda3f557adf709403058910ea009e0c88ce81458ec9b915016a5c5901e2365b130db00b18fcb7da1b082e1a5c75f7bf7eeab8783675d1b6a56441ac6e9ffc972b1278a5853d2b94dda55e1a6e2068bc0ddd3cddc9213ec9cebb7cb5be931977bb28dda12c7c5e69d1f876b243f0f224076bf1b81149603319a2fc9cb82337bdbe05e7bbf184bcbdc17d43b3f5efbae72ea386d955ca10e702e00df31aabf32",
						CreatedAt:         "2021-08-05T01:35:11.000Z",
						Status:            "expired",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "200 ok no compliance jobs",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				copt: &gotwtr.ComplianceJobsOption{
					Type: gotwtr.ComplianceFieldTypeTweets,
				},
			},
			want:    &gotwtr.ComplianceJobsResponse{},
			wantErr: false,
		},
		{
			name: "no type parameter",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					return nil
				}),
				copt: &gotwtr.ComplianceJobsOption{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for i, tt := range tests {
		i := i
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.ComplianceJobs(tt.args.ctx, tt.args.copt)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComplianceJobs() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ComplianceJobs() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}

		})
	}
}

func Test_complianceJob(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		cjID   int
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.ComplianceJobResponse
		wantErr bool
	}{
		{
			name: "200 ok",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `
                    {
						"data": {
							"download_expires_at": "2021-08-12T01:35:11.000Z",
							"download_url": "https://storage.googleapis.com/twttr-tweet-compliance/1423095206576984067/delivery/1202726487847104512_1423095206576984067?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=complianceapi-public-svc-acct%40twttr-compliance-public-prod.iam.gserviceaccount.com%2F20210805%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20210805T013511Z&X-Goog-Expires=604800&X-Goog-SignedHeaders=host&X-Goog-Signature=09de4feae68a6d4449eb7ce1f8f3551996552e7fba103005b3bd50ab318bb5215e4f5396ef29d17755deb6bf172b9d1dab61a04b249d39e87f6e2dbb31632b7e5f2d35f4f534e1f1522c9d7958b8745dd62471deb8d6992c80fd418628404f5f14eda3f557adf709403058910ea009e0c88ce81458ec9b915016a5c5901e2365b130db00b18fcb7da1b082e1a5c75f7bf7eeab8783675d1b6a56441ac6e9ffc972b1278a5853d2b94dda55e1a6e2068bc0ddd3cddc9213ec9cebb7cb5be931977bb28dda12c7c5e69d1f876b243f0f224076bf1b81149603319a2fc9cb82337bdbe05e7bbf184bcbdc17d43b3f5efbae72ea386d955ca10e702e00df31aabf32",
							"resumable": false,
							"upload_expires_at": "2021-08-05T01:50:11.000Z",
							"created_at": "2021-08-05T01:35:11.000Z",
							"upload_url": "https://storage.googleapis.com/twttr-tweet-compliance/1423095206576984067/submission/1202726487847104512_1423095206576984067?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=complianceapi-public-svc-acct%40twttr-compliance-public-prod.iam.gserviceaccount.com%2F20210805%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20210805T013511Z&X-Goog-Expires=900&X-Goog-SignedHeaders=content-type%3Bhost&X-Goog-Signature=ba08f588bea3873aa0465cf22015e583c2851a5ff14891d22430b1127288728f1aa303673e6895694e7017739871ff5ae59bbcde7d4ac7a14aaaafba98ad22ca818e99fb3ec7eaaf74b3ecfecbfb33711869b2e85d7666609276666ef4a8b396ae9616743a0cbd773962e5850f2942cd76be7373d608a140e041ca8492017d43fac9220fa145d0b2ecaf9f752d71fc8c4b81b67c5c22aa59ac87666f7d83714fdace72894d2911a3e36dd42028d0222e71054d6b28c8ef63d0f0000f228c8680bab9c8011b87d1a6c9a60e8cc9e8b6a83abf7c47a57772746c83b19849f5b4c938ccd0922990da5f2a81ff806edcb4667bb402fb1f1f6f5162768e0661648b21",
							"id": "1423095206576984067",
							"type": "tweets",
							"status": "expired"
						}
                    }`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				cjID: 1423095206576984067,
			},
			want: &gotwtr.ComplianceJobResponse{
				ComplianceJobData: &gotwtr.ComplianceJobData{
					DownloadExpiresAt: "2021-08-12T01:35:11.000Z",
					DownloadURL:       "https://storage.googleapis.com/twttr-tweet-compliance/1423095206576984067/delivery/1202726487847104512_1423095206576984067?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=complianceapi-public-svc-acct%40twttr-compliance-public-prod.iam.gserviceaccount.com%2F20210805%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20210805T013511Z&X-Goog-Expires=604800&X-Goog-SignedHeaders=host&X-Goog-Signature=09de4feae68a6d4449eb7ce1f8f3551996552e7fba103005b3bd50ab318bb5215e4f5396ef29d17755deb6bf172b9d1dab61a04b249d39e87f6e2dbb31632b7e5f2d35f4f534e1f1522c9d7958b8745dd62471deb8d6992c80fd418628404f5f14eda3f557adf709403058910ea009e0c88ce81458ec9b915016a5c5901e2365b130db00b18fcb7da1b082e1a5c75f7bf7eeab8783675d1b6a56441ac6e9ffc972b1278a5853d2b94dda55e1a6e2068bc0ddd3cddc9213ec9cebb7cb5be931977bb28dda12c7c5e69d1f876b243f0f224076bf1b81149603319a2fc9cb82337bdbe05e7bbf184bcbdc17d43b3f5efbae72ea386d955ca10e702e00df31aabf32",
					Resumable:         false,
					UploadExpiresAt:   "2021-08-05T01:50:11.000Z",
					CreatedAt:         "2021-08-05T01:35:11.000Z",
					UploadURL:         "https://storage.googleapis.com/twttr-tweet-compliance/1423095206576984067/submission/1202726487847104512_1423095206576984067?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=complianceapi-public-svc-acct%40twttr-compliance-public-prod.iam.gserviceaccount.com%2F20210805%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20210805T013511Z&X-Goog-Expires=900&X-Goog-SignedHeaders=content-type%3Bhost&X-Goog-Signature=ba08f588bea3873aa0465cf22015e583c2851a5ff14891d22430b1127288728f1aa303673e6895694e7017739871ff5ae59bbcde7d4ac7a14aaaafba98ad22ca818e99fb3ec7eaaf74b3ecfecbfb33711869b2e85d7666609276666ef4a8b396ae9616743a0cbd773962e5850f2942cd76be7373d608a140e041ca8492017d43fac9220fa145d0b2ecaf9f752d71fc8c4b81b67c5c22aa59ac87666f7d83714fdace72894d2911a3e36dd42028d0222e71054d6b28c8ef63d0f0000f228c8680bab9c8011b87d1a6c9a60e8cc9e8b6a83abf7c47a57772746c83b19849f5b4c938ccd0922990da5f2a81ff806edcb4667bb402fb1f1f6f5162768e0661648b21",
					ID:                "1423095206576984067",
					Type:              "tweets",
					Status:            "expired",
				},
			},
			wantErr: false,
		},
		{
			name: "no compliance jobs",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `
						{
							"errors": [
								{
									"value": "111111111111",
									"detail": "Could not find compliance_job with id: [111111111111].",
									"title": "Not Found Error",
									"resource_type": "compliance_job",
									"parameter": "id",
									"resource_id": "111111111111",
									"type": "https://api.twitter.com/2/problems/resource-not-found"
								}
							]
						}
					`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				cjID: 111111111111,
			},
			want: &gotwtr.ComplianceJobResponse{
				Errors: []*gotwtr.APIResponseError{
					{
						Value:        "111111111111",
						Detail:       "Could not find compliance_job with id: [111111111111].",
						Title:        "Not Found Error",
						ResourceType: "compliance_job",
						Parameter:    "id",
						ResourceID:   "111111111111",
						Type:         "https://api.twitter.com/2/problems/resource-not-found",
					},
				},
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		i := i
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.ComplianceJob(tt.args.ctx, tt.args.cjID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComplianceJob() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ComplianceJob() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}

		})
	}
}

func Test_createComplianceJob(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		opt    []*gotwtr.CreateComplianceJobOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.CreateComplianceJobResponse
		wantErr bool
	}{
		{
			name: "200 ok",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					data := `{
						"data": {
							"resumable": false,
							"type": "tweets",
							"download_expires_at": "2021-08-13T17:04:26.000Z",
							"created_at": "2021-08-06T17:04:26.000Z",
							"upload_url": "https://storage.googleapis.com/twttr-tweet-compliance/1423691444842209280/submission/1202726487847104512_1423691444842209280?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=complianceapi-public-svc-acct%40twttr-compliance-public-prod.iam.gserviceaccount.com%2F20210806%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20210806T170426Z&X-Goog-Expires=900&X-Goog-SignedHeaders=content-type%3Bhost&X-Goog-Signature=b057f9d07f80c7d4d73dc1184e2801e495054f61ff48ee1f6a74b7125bb62b02f3afa66a445e57634a0e78a546308632d845e28e092c10ae90cadb8655b0daafcee9ec8c7d95d2099437117db61d6789c8334e55ce5ee7b76c0d7dd383bd270d0c5a266ebbe0aa51365b332fe2c04942937526102871faa72e9255d2a8683d2dadcbd5ece0de18144a6dc74a6a53cdd4e5bb98261032047bf7d085be44a0126300aa3bb94d0657e532b538303ff217e20aaacbf638393addb6d7705966f1e5334040f150d930b857593e3e365381c0cf6e6ac4a24584c762adc75a27b769333e9a299dc16f4d771661d7aecc44d583bea1ff5f99fe9d08c87e55865610efdde2",
							"download_url": "https://storage.googleapis.com/twttr-tweet-compliance/1423691444842209280/delivery/1202726487847104512_1423691444842209280?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=complianceapi-public-svc-acct%40twttr-compliance-public-prod.iam.gserviceaccount.com%2F20210806%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20210806T170426Z&X-Goog-Expires=604800&X-Goog-SignedHeaders=host&X-Goog-Signature=840f364a22c259458b8705f8c96781db09c31a0bd844974f2a88b0a45f3455bc8e2f6ecc6c349abc8311a0040278114af3771ff0de3de6fee4230761573a44244aa5bb2d763829680d3d6bfee2a01538021fbb2f7b9d718e376945aea6355bf861618b968db597027eec317efaf434702d940ba805299ebfdae7af7f028a5ea89e74dd990920e0e879036cc0e2044228195356f0aa63ab89bfef5d6ede2fbf1789c2fe1a3e73dc58236775409e15f49acf72f5f91585a8ad0e5b073e5d6197cf8437aab82358ac9b0df81b5cdb2d6864f8d6e9725587ab92b5dbfc2d3968a5ee796d3940fd1594933f5a9653191dcfbbd63a8ccd02a56c2ef17764000591739d",
							"id": "1423691444842209280",
							"status": "created",
							"upload_expires_at": "2021-08-06T17:19:26.000Z"
						}
					}`
					return &http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(strings.NewReader(data)),
					}
				}),
				opt: []*gotwtr.CreateComplianceJobOption{
					{
						Type: gotwtr.ComplianceFieldTypeTweets,
					},
				},
			},
			want: &gotwtr.CreateComplianceJobResponse{
				CreateComplianceJobData: &gotwtr.CreateComplianceJobData{
					Resumable:         false,
					Type:              "tweets",
					DownloadExpiresAt: "2021-08-13T17:04:26.000Z",
					CreatedAt:         "2021-08-06T17:04:26.000Z",
					UploadURL:         "https://storage.googleapis.com/twttr-tweet-compliance/1423691444842209280/submission/1202726487847104512_1423691444842209280?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=complianceapi-public-svc-acct%40twttr-compliance-public-prod.iam.gserviceaccount.com%2F20210806%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20210806T170426Z&X-Goog-Expires=900&X-Goog-SignedHeaders=content-type%3Bhost&X-Goog-Signature=b057f9d07f80c7d4d73dc1184e2801e495054f61ff48ee1f6a74b7125bb62b02f3afa66a445e57634a0e78a546308632d845e28e092c10ae90cadb8655b0daafcee9ec8c7d95d2099437117db61d6789c8334e55ce5ee7b76c0d7dd383bd270d0c5a266ebbe0aa51365b332fe2c04942937526102871faa72e9255d2a8683d2dadcbd5ece0de18144a6dc74a6a53cdd4e5bb98261032047bf7d085be44a0126300aa3bb94d0657e532b538303ff217e20aaacbf638393addb6d7705966f1e5334040f150d930b857593e3e365381c0cf6e6ac4a24584c762adc75a27b769333e9a299dc16f4d771661d7aecc44d583bea1ff5f99fe9d08c87e55865610efdde2",
					DownloadURL:       "https://storage.googleapis.com/twttr-tweet-compliance/1423691444842209280/delivery/1202726487847104512_1423691444842209280?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=complianceapi-public-svc-acct%40twttr-compliance-public-prod.iam.gserviceaccount.com%2F20210806%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20210806T170426Z&X-Goog-Expires=604800&X-Goog-SignedHeaders=host&X-Goog-Signature=840f364a22c259458b8705f8c96781db09c31a0bd844974f2a88b0a45f3455bc8e2f6ecc6c349abc8311a0040278114af3771ff0de3de6fee4230761573a44244aa5bb2d763829680d3d6bfee2a01538021fbb2f7b9d718e376945aea6355bf861618b968db597027eec317efaf434702d940ba805299ebfdae7af7f028a5ea89e74dd990920e0e879036cc0e2044228195356f0aa63ab89bfef5d6ede2fbf1789c2fe1a3e73dc58236775409e15f49acf72f5f91585a8ad0e5b073e5d6197cf8437aab82358ac9b0df81b5cdb2d6864f8d6e9725587ab92b5dbfc2d3968a5ee796d3940fd1594933f5a9653191dcfbbd63a8ccd02a56c2ef17764000591739d",
					ID:                "1423691444842209280",
					Status:            "created",
					UploadExpiresAt:   "2021-08-06T17:19:26.000Z",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.CreateComplianceJob(tt.args.ctx, tt.args.opt...)
			if (err != nil) != tt.wantErr {
				t.Errorf("createComplianceJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("createComplianceJob() mismatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
