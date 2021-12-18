package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New(gotwtr.WithBearerToken("key"))
	// look up lists tweets by id
	ts, err := client.LookUpListTweets(context.Background(), "id")
	if err != nil {
		panic(err)
	}
	for _, t := range ts.Tweets {
		fmt.Println(t)
	}
}
