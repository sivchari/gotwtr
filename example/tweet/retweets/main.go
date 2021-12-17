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

	// post retweet
	p, err := client.PostRetweet(context.Background(), "user_id", "tweet_id")
	if err != nil {
		panic(err)
	}
	fmt.Println(p)

	// undo retweet
	d, err := client.UndoRetweet(context.Background(), "id", "source_tweet_id")
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
}
