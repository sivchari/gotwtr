package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	// look up
	ts, err := client.RetrieveMultipleTweets(context.Background(), []string{"id", "id2"})
	if err != nil {
		panic(err)
	}
	for _, t := range ts.Tweets {
		fmt.Println(t)
	}

	// look up by ID
	t, err := client.RetrieveSingleTweet(context.Background(), "id")
	if err != nil {
		panic(err)
	}
	fmt.Println(*t.Tweet)
}
