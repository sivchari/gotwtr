package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

func uploadMedia(ctx context.Context, c *client, media io.Reader, mediaType string, opt ...*MediaUploadOption) (*MediaUploadResponse, error) {
	if media == nil {
		return nil, errors.New("upload media: media parameter is required")
	}
	if mediaType == "" {
		return nil, errors.New("upload media: mediaType parameter is required")
	}

	// Read media data
	mediaData, err := io.ReadAll(media)
	if err != nil {
		return nil, fmt.Errorf("upload media: failed to read media data: %w", err)
	}

	// Create multipart form
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add media file
	part, err := writer.CreateFormFile("media", "media")
	if err != nil {
		return nil, fmt.Errorf("upload media: failed to create form file: %w", err)
	}
	if _, err := part.Write(mediaData); err != nil {
		return nil, fmt.Errorf("upload media: failed to write media data: %w", err)
	}

	// Add optional fields
	var mopt MediaUploadOption
	if len(opt) > 0 {
		mopt = *opt[0]
	}

	if mopt.MediaCategory != "" {
		if err := writer.WriteField("media_category", mopt.MediaCategory); err != nil {
			return nil, fmt.Errorf("upload media: failed to write media_category: %w", err)
		}
	}

	if mopt.AltText != "" {
		if err := writer.WriteField("alt_text", mopt.AltText); err != nil {
			return nil, fmt.Errorf("upload media: failed to write alt_text: %w", err)
		}
	}

	for i, owner := range mopt.AdditionalOwners {
		fieldName := fmt.Sprintf("additional_owners[%d]", i)
		if err := writer.WriteField(fieldName, owner); err != nil {
			return nil, fmt.Errorf("upload media: failed to write additional_owners: %w", err)
		}
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("upload media: failed to close multipart writer: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, mediaUploadURL, &buf)
	if err != nil {
		return nil, fmt.Errorf("upload media new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("upload media response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	var mur MediaUploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&mur); err != nil {
		return nil, fmt.Errorf("upload media decode: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return &mur, &HTTPError{
			APIName: "upload media",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &mur, nil
}

func initializeChunkedUpload(ctx context.Context, c *client, req *MediaUploadInitRequest) (*MediaUploadResponse, error) {
	if req == nil {
		return nil, errors.New("initialize chunked upload: request parameter is required")
	}
	if req.MediaType == "" {
		return nil, errors.New("initialize chunked upload: media_type is required")
	}
	if req.TotalBytes <= 0 {
		return nil, errors.New("initialize chunked upload: total_bytes must be greater than 0")
	}

	// Prepare form data
	data := url.Values{}
	data.Set("command", "INIT")
	data.Set("media_type", req.MediaType)
	data.Set("total_bytes", strconv.FormatInt(req.TotalBytes, 10))

	if req.MediaCategory != "" {
		data.Set("media_category", req.MediaCategory)
	}

	for i, owner := range req.AdditionalOwners {
		data.Set(fmt.Sprintf("additional_owners[%d]", i), owner)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, mediaUploadURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("initialize chunked upload new request with ctx: %w", err)
	}

	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("initialize chunked upload response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	var mur MediaUploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&mur); err != nil {
		return nil, fmt.Errorf("initialize chunked upload decode: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return &mur, &HTTPError{
			APIName: "initialize chunked upload",
			Status:  resp.Status,
			URL:     httpReq.URL.String(),
		}
	}

	return &mur, nil
}

func appendChunkedUpload(ctx context.Context, c *client, req *MediaUploadAppendRequest) error {
	if req == nil {
		return errors.New("append chunked upload: request parameter is required")
	}
	if req.MediaID == "" {
		return errors.New("append chunked upload: media_id is required")
	}
	if req.Media == nil || len(req.Media) == 0 {
		return errors.New("append chunked upload: media data is required")
	}

	// Create multipart form
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add command
	if err := writer.WriteField("command", "APPEND"); err != nil {
		return fmt.Errorf("append chunked upload: failed to write command: %w", err)
	}

	// Add media_id
	if err := writer.WriteField("media_id", req.MediaID); err != nil {
		return fmt.Errorf("append chunked upload: failed to write media_id: %w", err)
	}

	// Add segment_index
	if err := writer.WriteField("segment_index", strconv.Itoa(req.SegmentIndex)); err != nil {
		return fmt.Errorf("append chunked upload: failed to write segment_index: %w", err)
	}

	// Add media chunk
	part, err := writer.CreateFormFile("media", "chunk")
	if err != nil {
		return fmt.Errorf("append chunked upload: failed to create form file: %w", err)
	}
	if _, err := part.Write(req.Media); err != nil {
		return fmt.Errorf("append chunked upload: failed to write media chunk: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("append chunked upload: failed to close multipart writer: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, mediaUploadURL, &buf)
	if err != nil {
		return fmt.Errorf("append chunked upload new request with ctx: %w", err)
	}

	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("append chunked upload response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		// Read error response
		var mur MediaUploadResponse
		if decodeErr := json.NewDecoder(resp.Body).Decode(&mur); decodeErr == nil && len(mur.Errors) > 0 {
			return &HTTPError{
				APIName: "append chunked upload",
				Status:  resp.Status,
				URL:     httpReq.URL.String(),
			}
		}
		return &HTTPError{
			APIName: "append chunked upload",
			Status:  resp.Status,
			URL:     httpReq.URL.String(),
		}
	}

	return nil
}

func finalizeChunkedUpload(ctx context.Context, c *client, req *MediaUploadFinalizeRequest) (*MediaUploadResponse, error) {
	if req == nil {
		return nil, errors.New("finalize chunked upload: request parameter is required")
	}
	if req.MediaID == "" {
		return nil, errors.New("finalize chunked upload: media_id is required")
	}

	// Prepare form data
	data := url.Values{}
	data.Set("command", "FINALIZE")
	data.Set("media_id", req.MediaID)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, mediaUploadURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("finalize chunked upload new request with ctx: %w", err)
	}

	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("finalize chunked upload response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	var mur MediaUploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&mur); err != nil {
		return nil, fmt.Errorf("finalize chunked upload decode: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return &mur, &HTTPError{
			APIName: "finalize chunked upload",
			Status:  resp.Status,
			URL:     httpReq.URL.String(),
		}
	}

	return &mur, nil
}

func checkUploadStatus(ctx context.Context, c *client, req *MediaUploadStatusRequest) (*MediaUploadResponse, error) {
	if req == nil {
		return nil, errors.New("check upload status: request parameter is required")
	}
	if req.MediaID == "" {
		return nil, errors.New("check upload status: media_id is required")
	}

	// Prepare query parameters
	params := url.Values{}
	params.Set("command", "STATUS")
	params.Set("media_id", req.MediaID)

	// Create URL with query parameters
	reqURL := fmt.Sprintf("%s?%s", mediaUploadURL, params.Encode())

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("check upload status new request with ctx: %w", err)
	}

	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("check upload status response: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	var mur MediaUploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&mur); err != nil {
		return nil, fmt.Errorf("check upload status decode: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &mur, &HTTPError{
			APIName: "check upload status",
			Status:  resp.Status,
			URL:     httpReq.URL.String(),
		}
	}

	return &mur, nil
}