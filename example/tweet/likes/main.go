package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	// get users who liked the tweet that id is "id"
	users, err := client.LikesLookUpUsers(context.Background(), "id")
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	// get users who liked the tweet that id is "id" with option
	users, err = client.LikesLookUpUsers(context.Background(), "id", &gotwtr.LikesLookUpByTweetOpts{
		Expansions:  []gotwtr.Expansion{gotwtr.ExpansionPinnedTweetID},
		UserFields:  []gotwtr.UserField{gotwtr.UserFieldCreatedAt},
		TweetFields: []gotwtr.TweetField{gotwtr.TweetFieldCreatedAt},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
}
