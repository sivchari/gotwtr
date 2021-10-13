package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	// retweets lookup by ID
	t, err := client.RetweetsLookup(context.Background(), "id")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// post retweet by userID and tweetID
	r, err := client.PostRetweet(context.Background(), "user_id", "tweet_id")
	if err != nil {
		panic(err)
	}
	fmt.Println(r)

	// delete retweet by ID and sourceTweetID
	d, err := client.DeleteRetweet(context.Background(), "id", "source_tweet_id")
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
}
