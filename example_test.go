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

func ExampleClient_TweetsUserLiked() {
	client := gotwtr.New("key")
	tulr, err := client.TweetsUserLiked(context.Background(), "user_id")
	if err != nil {
		log.Println(err)
	}
	for _, tweet := range tulr.Tweets {
		fmt.Printf("id: %s, text: %s\n", tweet.ID, tweet.Text)
	}
}
func ExampleClient_UsersLikingTweet() {
	client := gotwtr.New("key")
	ultr, err := client.UsersLikingTweet(context.Background(), "tweet_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range ultr.Users {
		fmt.Printf("id: %s, name: %s\n", user.ID, user.UserName)
	}
}

func ExampleClient_RetrieveMultipleUsersWithIDs() {
	client := gotwtr.New("key")
	// look up users
	us, err := client.RetrieveMultipleUsersWithIDs(context.Background(), []string{"id", "id2"})
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range us.Users {
		fmt.Println(u)
	}
}

func ExampleClient_RetrieveSingleUserWithID() {
	client := gotwtr.New("key")
	u, err := client.RetrieveSingleUserWithID(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}

func ExampleClient_RetrieveMultipleUsersWithUserNames() {
	client := gotwtr.New("key")
	uns, err := client.RetrieveMultipleUsersWithUserNames(context.Background(), []string{"username", "username2"})
	if err != nil {
		log.Fatal(err)
	}
	for _, un := range uns.Users {
		fmt.Println(un)
	}
}

func ExampleClient_RetrieveSingleUserWithUserName() {
	client := gotwtr.New("key")
	un, err := client.RetrieveSingleUserWithUserName(context.Background(), "username")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(un)
}

func ExampleClient_Following() {
	client := gotwtr.New("key")
	f, err := client.Following(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range f.Users {
		fmt.Println(user)
	}
}

func ExampleClient_Followers() {
	client := gotwtr.New("key")
	f, err := client.Followers(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range f.Users {
		fmt.Println(user)
	}
}

func ExampleClient_PostFollowing() {
	client := gotwtr.New("key")
	pf, err := client.PostFollowing(context.Background(), "user_id", "target_user_id")
	if err != nil {
		panic(err)
	}
	fmt.Println(pf)
}

func ExampleClient_UndoFollowing() {
	client := gotwtr.New("key")
	uf, err := client.UndoFollowing(context.Background(), "source_user_id", "target_user_id")
	if err != nil {
		panic(err)
	}
	fmt.Println(uf)
}

func ExampleClient_Muting() {
	client := gotwtr.New("key")
	m, err := client.Muting(context.Background(), "user_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range m.Users {
		fmt.Println(user)
	}
}

func ExampleClient_PostMuting() {
	client := gotwtr.New("key")
	pm, err := client.PostMuting(context.Background(), "user_id", "target_user_id")
	if err != nil {
		panic(err)
	}
	fmt.Println(pm)
}

func ExampleClient_UndoMuting() {
	client := gotwtr.New("key")
	um, err := client.UndoMuting(context.Background(), "source_user_id", "target_user_id")
	if err != nil {
		panic(err)
	}
	fmt.Println(um)
}

func ExampleClient_LookUpSpace() {
	client := gotwtr.New("key")
	s, err := client.LookUpSpace(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}

func ExampleClient_LookUpSpaces() {
	client := gotwtr.New("key")
	ss, err := client.LookUpSpaces(context.Background(), []string{"id", "id2"})
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range ss.Spaces {
		fmt.Println(s)
	}
}

func ExampleClient_UsersPurchasedSpaceTicket() {
	client := gotwtr.New("key")
	tickets, err := client.UsersPurchasedSpaceTicket(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range tickets.Users {
		fmt.Println(user)
	}
}

func ExampleClient_DiscoverSpaces() {
	client := gotwtr.New("key")
	discover, err := client.DiscoverSpaces(context.Background(), []string{"id", "id2"})
	if err != nil {
		log.Fatal(err)
	}
	for _, space := range discover.Spaces {
		fmt.Println(space)
	}
}

func ExampleClient_SearchSpaces() {
	client := gotwtr.New("key")
	spaces, err := client.SearchSpaces(context.Background(), "query")
	if err != nil {
		log.Fatal(err)
	}
	for _, space := range spaces.Spaces {
		fmt.Println(space)
	}
}

func ExampleClient_LookUpList() {
	client := gotwtr.New("key")
	l, err := client.LookUpList(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(l)
}

func ExampleClient_LookUpAllListsOwned() {
	client := gotwtr.New("key")
	lists, err := client.LookUpAllListsOwned(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, list := range lists.Lists {
		fmt.Println(list)
	}
}

func ExampleClient_LookUpListTweets() {
	client := gotwtr.New("key")
	lt, err := client.LookUpListTweets(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, tweet := range lt.Tweets {
		fmt.Println(tweet)
	}
}

func ExampleClient_ListMembers() {
	client := gotwtr.New("key")
	members, err := client.ListMembers(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, member := range members.Users {
		fmt.Println(member)
	}
}

func ExampleClient_ListsSpecifiedUser() {
	client := gotwtr.New("key")
	lists, err := client.ListsSpecifiedUser(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, list := range lists.Lists {
		fmt.Println(list)
	}
}

func ExampleClient_LookUpListFollowers() {
	client := gotwtr.New("key")
	followers, err := client.LookUpListFollowers(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range followers.Users {
		fmt.Println(user)
	}
}

func ExampleClient_LookUpAllListsUserFollows() {
	client := gotwtr.New("key")
	lists, err := client.LookUpAllListsUserFollows(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, list := range lists.Lists {
		fmt.Println(list)
	}
}
