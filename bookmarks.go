package gotwtr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//https://developer.twitter.com/en/docs/twitter-api/tweets/bookmarks/api-reference/get-users-id-bookmarks
func lookupUserBookmarks(ctx context.Context, c *client, userID string, opt ...*LookupUserBookmarksOption) (*LookupUserBookmarksResponse, error) {
	if userID == "" {
		return nil, errors.New("lookup user bookmarks: user id parameter is required")
	}
	ep := fmt.Sprintf(lookupUserBookmarksURL, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("lookup user bookmarks new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var lopt LookupUserBookmarksOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		lopt = *opt[0]
	default:
		return nil, errors.New("lookup user bookmarks: only one option is allowed")
	}
	lopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("lookup user bookmarks response: %w", err)
	}
	defer resp.Body.Close()

	var lookupUserBookmarks LookupUserBookmarksResponse
	if err := json.NewDecoder(resp.Body).Decode(&lookupUserBookmarks); err != nil {
		return nil, fmt.Errorf("lookup user bookmarks decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &lookupUserBookmarks, &HTTPError{
			APIName: "lookup user bookmarks",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &lookupUserBookmarks, nil
}

//https://developer.twitter.com/en/docs/twitter-api/tweets/bookmarks/api-reference/post-users-id-bookmarks
func bookmarkTweet(ctx context.Context, c *client, userID string, body *BookmarkTweetBody) (*BookmarkTweetResponse, error) {
	if userID == "" {
		return nil, errors.New("bookmark tweet: user id parameter is required")
	}
	if body.TweetID == "" {
		return nil, errors.New("bookmark tweet: tweet id parameter is required")
	}
	ep := fmt.Sprintf(bookmarkTweetURL, userID)

	j, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("bookmark tweet : can not marshal")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep, bytes.NewBuffer(j))
	if err != nil {
		return nil, fmt.Errorf("bookmark tweet new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("bookmark tweet response: %w", err)
	}
	defer resp.Body.Close()

	var bookmarkTweet BookmarkTweetResponse
	if err := json.NewDecoder(resp.Body).Decode(&bookmarkTweet); err != nil {
		return nil, fmt.Errorf("bookmark tweet decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &bookmarkTweet, &HTTPError{
			APIName: "bookmark tweet",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &bookmarkTweet, nil
}

//https://developer.twitter.com/en/docs/twitter-api/tweets/bookmarks/api-reference/delete-users-id-bookmarks-tweet_id
func removeBookmarkOfTweet(ctx context.Context, c *client, userID string, tweetID string) (*RemoveBookmarkOfTweetResponse, error) {
	if userID == "" {
		return nil, errors.New("remove bookmark of tweet: user id parameter is required")
	}
	if tweetID == "" {
		return nil, errors.New("remove bookmark of tweet: tweet id parameter is required")
	}
	ep := fmt.Sprintf(removeBookmarkOfTweetURL, userID, tweetID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("remove bookmark of tweet new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("remove bookmark of tweet response: %w", err)
	}
	defer resp.Body.Close()

	var removeBookmarkOfTweet RemoveBookmarkOfTweetResponse
	if err := json.NewDecoder(resp.Body).Decode(&removeBookmarkOfTweet); err != nil {
		return nil, fmt.Errorf("remove bookmark of tweet decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &removeBookmarkOfTweet, &HTTPError{
			APIName: "remove bookmark of tweet",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &removeBookmarkOfTweet, nil
}
