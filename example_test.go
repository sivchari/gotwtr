package gotwtr_test

import (
	"context"
	"fmt"
	"log"

	"github.com/sivchari/gotwtr"
)

func ExampleClient_RetrieveMultipleTweets() {
	client := gotwtr.New("key")
	ts, err := client.RetrieveMultipleTweets(context.Background(), []string{"id", "id2"})
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range ts.Tweets {
		fmt.Println(t)
	}
}
func ExampleClient_RetrieveSingleTweet() {
	client := gotwtr.New("key")
	t, err := client.RetrieveSingleTweet(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*t.Tweet)
}
func ExampleClient_UserMentionTimeline() {

}
func ExampleClient_UserTweetTimeline() {

}
func ExampleClient_SearchRecentTweets() {

}
func ExampleClient_CountsRecentTweet() {

}
func ExampleClient_AddOrDeleteRules() {

}
func ExampleClient_RetrieveStreamRules() {

}
func ExampleClient_ConnectToStream() {

}
func ExampleClient_VolumeStreams() {

}
func ExampleClient_RetweetsLookup() {

}
func ExampleClient_TweetsUserLiked() {

}
func ExampleClient_UsersLikingTweet() {

}
