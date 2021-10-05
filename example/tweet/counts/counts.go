package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	ts, err := client.CountsRecentTweet(context.Background(), "from:TwitterDev")
	if err != nil {
		log.Fatal(err)
	}
	println(ts.Meta.TotalTweetCount)
	for _, t := range ts.Counts {
		fmt.Println(t)
	}
}
