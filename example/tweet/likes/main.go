package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	// get users who liked the tweet that id is "tweet_id"
	ultr, err := client.UsersLikingTweet(context.Background(), "tweet_id")
	if err != nil {
		panic(err)
	}
	fmt.Println("------no option------")
	for _, user := range ultr.Users {
		fmt.Printf("id: %s, name: %s\n", user.ID, user.UserName)
	}

	// get users who liked the tweet that id is "tweet_id" with option
	ultr, err = client.UsersLikingTweet(context.Background(), "tweet_id", &gotwtr.UsersLikingTweetOption{
		Expansions:  []gotwtr.Expansion{gotwtr.ExpansionPinnedTweetID},
		UserFields:  []gotwtr.UserField{gotwtr.UserFieldCreatedAt},
		TweetFields: []gotwtr.TweetField{gotwtr.TweetFieldCreatedAt},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("------with option------")
	for _, user := range ultr.Users {
		fmt.Printf("id: %s, name: %s, created_at: %v\n", user.ID, user.UserName, user.CreatedAt)
	}
	if ultr.Includes != nil {
		for _, tweet := range ultr.Includes.Tweets {
			fmt.Printf("tweet_id: %s, created_at: %v\n", tweet.ID, tweet.CreatedAt)
		}
	}

	// get tweets which the user liked with no potion
	tulr, err := client.TweetsUserLiked(context.Background(), "user_id")
	if err != nil {
		panic(err)
	}
	fmt.Println("------no option------")
	for _, tweet := range tulr.Tweets {
		fmt.Printf("id: %s, text: %s\n", tweet.ID, tweet.Text)
	}

	// get tweets which the user liked with potion
	tulr, err = client.TweetsUserLiked(context.Background(), "user_id", &gotwtr.TweetsUserLikedOpts{
		TweetFields: []gotwtr.TweetField{gotwtr.TweetFieldCreatedAt, gotwtr.TweetFieldSource},
		MaxResults:  10,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println("------with option------")
	for _, tweet := range tulr.Tweets {
		fmt.Printf("id: %s, text: %s, created_at: %s, source: %s\n", tweet.ID, tweet.Text, tweet.CreatedAt, tweet.Source)
	}
}
