package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	tsr, err := client.SearchRecentTweets(context.Background(), "go", &gotwtr.TweetSearchOption{
		TweetFields: []gotwtr.TweetField{
			gotwtr.TweetFieldAuthorID,
			gotwtr.TweetFieldAttachments,
		},
		MaxResults: 10,
	})
	if err != nil {
		panic(err)
	}
	for _, t := range tsr.Tweets {
		fmt.Println("---")
		fmt.Println(t.Text)
	}

	fmt.Println("---meta---")
	fmt.Println(tsr.Meta)
}
