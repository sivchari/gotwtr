package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	// look up
	ts, err := client.LookUpTweets(context.Background(), []string{"id", "id2"})
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range ts.Tweets {
		fmt.Println(t)
	}

	// look up by ID
	t, err := client.LookUpTweetByID(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*t.Tweet)
}
