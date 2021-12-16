package gotwtr_test

import (
	"context"
	"fmt"
	"log"
	"time"

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
func ExampleClient_AddOrDeleteRules_add() {
	client := gotwtr.New("key")
	_, err := client.AddOrDeleteRules(context.Background(), &gotwtr.AddOrDeleteJSONBody{
		Add: []*gotwtr.AddRule{
			{
				Value: "puppy has:media",
				Tag:   "puppies with media",
			},
			{
				Value: "meme has:images",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleClient_AddOrDeleteRules_delete() {
	client := gotwtr.New("key")
	// retrieve Stream rules
	ts, err := client.RetrieveStreamRules(context.Background())
	if err != nil {
		panic(err)
	}
	var ids []string
	for _, t := range ts.Rules {
		fmt.Println(t)
		ids = append(ids, t.ID)
	}

	// delete Stream rules
	_, err = client.AddOrDeleteRules(context.Background(), &gotwtr.AddOrDeleteJSONBody{
		Delete: &gotwtr.DeleteRule{
			IDs: ids,
		},
	})
	if err != nil {
		panic(err)
	}
}
func ExampleClient_RetrieveStreamRules() {
	client := gotwtr.New("key")
	ts, err := client.RetrieveStreamRules(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range ts.Rules {
		fmt.Println(t)
	}
}
func ExampleClient_ConnectToStream() {
	client := gotwtr.New("key")
	ch := make(chan gotwtr.ConnectToStreamResponse, 5)
	errCh := make(chan error)
	stream := client.ConnectToStream(context.Background(), ch, errCh)
	fmt.Println("streaming...")
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case data := <-ch:
				fmt.Println(data.Tweet)
			case err := <-errCh:
				fmt.Println(err)
			case <-ctx.Done():
				break
			}
		}
	}(ctx)
	time.Sleep(time.Second * 10)
	cancel()
	stream.Stop()
	fmt.Println("done")
}
func ExampleClient_VolumeStreams() {

}
func ExampleClient_RetweetsLookup() {

}
func ExampleClient_TweetsUserLiked() {

}
func ExampleClient_UsersLikingTweet() {

}
