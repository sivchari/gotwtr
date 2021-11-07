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
}
