package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	// user tweet timeline
	ts, err := client.UserTweetTimeline(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range ts.Tweets {
		fmt.Println(t)
	}

	// user mension timeline
	tws, err := client.UserMensionTimeline(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, tw := range tws.Tweets {
		fmt.Println(tw)
	}
}
