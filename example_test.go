package gotwtr_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sivchari/gotwtr"
)

func ExampleClient_GenerateAppOnlyBearerToken() {
	c := gotwtr.New(
		"key",
		gotwtr.WithConsumerSecret("sec"),
	)
	b, err := c.GenerateAppOnlyBearerToken(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if !b {
		log.Fatal("failed to generate bearer token")
	}
}

// func ExampleClient_InvalidatingBearerToken() {
// 	client := gotwtr.New("key")
// 	b, err := client.InvalidatingBearerToken(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if !b {
// 		log.Fatal("failed to invalidate bearer token")
// 	}
// }

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
				return
			}
		}
	}(ctx)
	time.Sleep(time.Second * 10)
	cancel()
	stream.Stop()
	fmt.Println("done")
}

func ExampleClient_VolumeStreams() {
	client := gotwtr.New("key")
	ch := make(chan gotwtr.VolumeStreamsResponse, 5)
	errCh := make(chan error)
	stream := client.VolumeStreams(context.Background(), ch, errCh)
	fmt.Println("streaming...")
	done := make(chan struct{})
	go func(done chan struct{}) {
		for {
			select {
			case data := <-ch:
				fmt.Println(data.Tweet)
			case err := <-errCh:
				fmt.Println(err)
			case <-done:
				return
			}
		}
	}(done)
	time.Sleep(time.Second * 10)
	close(done)
	stream.Stop()
	fmt.Println("done")
}

func ExampleClient_RetweetsLookup() {
	client := gotwtr.New("key")
	t, err := client.RetweetsLookup(context.Background(), "id")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(t)
}

func ExampleClient_PostRetweet() {
	client := gotwtr.New("key")
	pr, err := client.PostRetweet(context.Background(), "uid", "tid")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(pr)
}

func ExampleClient_UndoRetweet() {
	client := gotwtr.New("key")
	ur, err := client.UndoRetweet(context.Background(), "id", "stid")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ur)
}

func ExampleClient_TweetsUserLiked_noOption() {
	client := gotwtr.New("key")
	tulr, err := client.TweetsUserLiked(context.Background(), "user_id")
	if err != nil {
		log.Println(err)
	}
	for _, tweet := range tulr.Tweets {
		fmt.Printf("id: %s, text: %s\n", tweet.ID, tweet.Text)
	}
}

func ExampleClient_TweetsUserLiked_option() {
	client := gotwtr.New("key")
	tulr, err := client.TweetsUserLiked(context.Background(), "user_id", &gotwtr.TweetsUserLikedOpts{
		TweetFields: []gotwtr.TweetField{gotwtr.TweetFieldCreatedAt, gotwtr.TweetFieldSource},
		MaxResults:  10,
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, tweet := range tulr.Tweets {
		fmt.Printf("id: %s, text: %s, created_at: %s, source: %s\n", tweet.ID, tweet.Text, tweet.CreatedAt, tweet.Source)
	}
}

func ExampleClient_UsersLikingTweet_noOption() {
	client := gotwtr.New("key")
	ultr, err := client.UsersLikingTweet(context.Background(), "tweet_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range ultr.Users {
		fmt.Printf("id: %s, name: %s\n", user.ID, user.UserName)
	}
}

func ExampleClient_UsersLikingTweet_option() {
	client := gotwtr.New("key")
	ultr, err := client.UsersLikingTweet(context.Background(), "tweet_id", &gotwtr.UsersLikingTweetOption{
		Expansions:  []gotwtr.Expansion{gotwtr.ExpansionPinnedTweetID},
		UserFields:  []gotwtr.UserField{gotwtr.UserFieldCreatedAt},
		TweetFields: []gotwtr.TweetField{gotwtr.TweetFieldCreatedAt},
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range ultr.Users {
		fmt.Printf("id: %s, name: %s, created_at: %v\n", user.ID, user.UserName, user.CreatedAt)
	}
	if ultr.Includes != nil {
		for _, tweet := range ultr.Includes.Tweets {
			fmt.Printf("tweet_id: %s, created_at: %v\n", tweet.ID, tweet.CreatedAt)
		}
	}
}
