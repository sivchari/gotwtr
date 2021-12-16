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
	client := gotwtr.New("key")
	tws, err := client.UserMentionTimeline(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, tw := range tws.Tweets {
		fmt.Println(tw)
	}
}
func ExampleClient_UserTweetTimeline() {
	client := gotwtr.New("key")
	ts, err := client.UserTweetTimeline(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range ts.Tweets {
		fmt.Println(t)
	}
}
func ExampleClient_SearchRecentTweets() {
	client := gotwtr.New("key")
	tsr, err := client.SearchRecentTweets(context.Background(), "go", &gotwtr.SearchTweetsOption{
		TweetFields: []gotwtr.TweetField{
			gotwtr.TweetFieldAuthorID,
			gotwtr.TweetFieldAttachments,
		},
		MaxResults: 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range tsr.Tweets {
		fmt.Println("---")
		fmt.Println(t.Text)
	}

	fmt.Println("---meta---")
	fmt.Println(tsr.Meta)

}
func ExampleClient_CountsRecentTweet() {
	client := gotwtr.New("key")
	ts, err := client.CountsRecentTweet(context.Background(), "from:TwitterDev")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ts.Meta.TotalTweetCount)
	for _, t := range ts.Counts {
		fmt.Println(t)
	}
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
