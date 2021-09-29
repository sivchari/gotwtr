package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	ts, err := client.LookUpTweets(context.Background(), []string{"tweet ids"})
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range ts.Tweets {
		fmt.Println(t)
	}
}
