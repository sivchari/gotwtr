package gotwtr_test

import (
	"context"
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

func ExampleClient_RetrieveMultipleTweets() {
	client := gotwtr.New("key")
	ts, err := client.RetrieveMultipleTweets(context.Background(), []string{"tweet_id", "tweet_id2"})
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range ts.Tweets {
		log.Println(t)
	}
}

func ExampleClient_RetrieveSingleTweet() {
	client := gotwtr.New("key")
	t, err := client.RetrieveSingleTweet(context.Background(), "tweet_id")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*t.Tweet)
}

func ExampleClient_UserMentionTimeline() {
	client := gotwtr.New("key")
	tws, err := client.UserMentionTimeline(context.Background(), "user_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, tw := range tws.Tweets {
		log.Println(tw)
	}
}

func ExampleClient_UserTweetTimeline() {
	client := gotwtr.New("key")
	ts, err := client.UserTweetTimeline(context.Background(), "user_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range ts.Tweets {
		log.Println(t)
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
		log.Println("---")
		log.Println(t.Text)
	}

	log.Println("---meta---")
	log.Println(tsr.Meta)

}

func ExampleClient_CountRecentTweets() {
	client := gotwtr.New("key")
	ts, err := client.CountRecentTweets(context.Background(), "lakers")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ts.Meta.TotalTweetCount)
	for _, t := range ts.Counts {
		log.Println(t)
	}
}

func ExampleClient_CountAllTweets() {
	client := gotwtr.New("key")
	ts, err := client.CountAllTweets(context.Background(), "lakers")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ts.Meta.TotalTweetCount)
	for _, t := range ts.Counts {
		log.Println(t)
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
		log.Println(t)
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
		log.Println(t)
	}
}

func ExampleClient_ConnectToStream() {
	client := gotwtr.New("key")
	ch := make(chan gotwtr.ConnectToStreamResponse, 5)
	errCh := make(chan error)
	stream := client.ConnectToStream(context.Background(), ch, errCh)
	log.Println("streaming...")
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case data := <-ch:
				log.Println(data.Tweet)
			case err := <-errCh:
				log.Println(err)
			case <-ctx.Done():
				return
			}
		}
	}(ctx)
	time.Sleep(time.Second * 10)
	cancel()
	stream.Stop()
	log.Println("done")
}

func ExampleClient_VolumeStreams() {
	client := gotwtr.New("key")
	ch := make(chan gotwtr.VolumeStreamsResponse, 5)
	errCh := make(chan error)
	stream := client.VolumeStreams(context.Background(), ch, errCh)
	log.Println("streaming...")
	done := make(chan struct{})
	go func(done chan struct{}) {
		for {
			select {
			case data := <-ch:
				log.Println(data.Tweet)
			case err := <-errCh:
				log.Println(err)
			case <-done:
				return
			}
		}
	}(done)
	time.Sleep(time.Second * 10)
	close(done)
	stream.Stop()
	log.Println("done")
}

func ExampleClient_RetweetsLookup() {
	client := gotwtr.New("key")
	t, err := client.RetweetsLookup(context.Background(), "tweet_id")
	if err != nil {
		log.Println(err)
	}
	log.Println(t)
}

func ExampleClient_PostRetweet() {
	client := gotwtr.New("key")
	pr, err := client.PostRetweet(context.Background(), "user_id", "tweet_id")
	if err != nil {
		log.Println(err)
	}
	log.Println(pr)
}

func ExampleClient_UndoRetweet() {
	client := gotwtr.New("key")
	ur, err := client.UndoRetweet(context.Background(), "user_id", "source_tweet_id")
	if err != nil {
		log.Println(err)
	}
	log.Println(ur)
}

func ExampleClient_TweetsUserLiked() {
	client := gotwtr.New("key")
	tulr, err := client.TweetsUserLiked(context.Background(), "user_id")
	if err != nil {
		log.Println(err)
	}
	for _, tweet := range tulr.Tweets {
		log.Printf("id: %s, text: %s\n", tweet.ID, tweet.Text)
	}
}
func ExampleClient_UsersLikingTweet() {
	client := gotwtr.New("key")
	ultr, err := client.UsersLikingTweet(context.Background(), "tweet_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range ultr.Users {
		log.Printf("id: %s, name: %s\n", user.ID, user.UserName)
	}
}

func ExampleClient_PostUsersLikingTweet() {
	client := gotwtr.New("key")
	pult, err := client.PostUsersLikingTweet(context.Background(), "user_id", "tweet_id")
	if err != nil {
		log.Println(err)
	}
	log.Println(pult)
}

func ExampleClient_UndoUsersLikingTweet() {
	client := gotwtr.New("key")
	uult, err := client.UndoUsersLikingTweet(context.Background(), "user_id", "tweet_id")
	if err != nil {
		log.Println(err)
	}
	log.Println(uult)
}

func ExampleClient_RetrieveMultipleUsersWithIDs() {
	client := gotwtr.New("key")
	// look up users
	us, err := client.RetrieveMultipleUsersWithIDs(context.Background(), []string{"user_id", "user_id2"})
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range us.Users {
		log.Println(u)
	}
}

func ExampleClient_RetrieveSingleUserWithID() {
	client := gotwtr.New("key")
	u, err := client.RetrieveSingleUserWithID(context.Background(), "user_id")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(u)
}

func ExampleClient_RetrieveMultipleUsersWithUserNames() {
	client := gotwtr.New("key")
	uns, err := client.RetrieveMultipleUsersWithUserNames(context.Background(), []string{"username", "username2"})
	if err != nil {
		log.Fatal(err)
	}
	for _, un := range uns.Users {
		log.Println(un)
	}
}

func ExampleClient_RetrieveSingleUserWithUserName() {
	client := gotwtr.New("key")
	un, err := client.RetrieveSingleUserWithUserName(context.Background(), "username")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(un)
}

func ExampleClient_Following() {
	client := gotwtr.New("key")
	f, err := client.Following(context.Background(), "user_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range f.Users {
		log.Println(user)
	}
}

func ExampleClient_Followers() {
	client := gotwtr.New("key")
	f, err := client.Followers(context.Background(), "user_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range f.Users {
		log.Println(user)
	}
}

func ExampleClient_PostFollowing() {
	client := gotwtr.New("key")
	pf, err := client.PostFollowing(context.Background(), "user_id", "target_user_id")
	if err != nil {
		panic(err)
	}
	log.Println(pf)
}

func ExampleClient_UndoFollowing() {
	client := gotwtr.New("key")
	uf, err := client.UndoFollowing(context.Background(), "source_user_id", "target_user_id")
	if err != nil {
		panic(err)
	}
	log.Println(uf)
}

func ExampleClient_Blocking() {
	client := gotwtr.New("key")
	b, err := client.Blocking(context.Background(), "user_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range b.Users {
		log.Println(user)
	}
}

func ExampleClient_PostBlocking() {
	client := gotwtr.New("key")
	pb, err := client.PostBlocking(context.Background(), "user_id", "target_user_id")
	if err != nil {
		log.Println(err)
	}
	log.Println(pb)
}

func ExampleClient_UndoBlocking() {
	client := gotwtr.New("key")
	ub, err := client.UndoBlocking(context.Background(), "source_user_id", "target_user_id")
	if err != nil {
		log.Println(err)
	}
	log.Println(ub)
}

func ExampleClient_Muting() {
	client := gotwtr.New("key")
	m, err := client.Muting(context.Background(), "user_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range m.Users {
		log.Println(user)
	}
}

func ExampleClient_PostMuting() {
	client := gotwtr.New("key")
	pm, err := client.PostMuting(context.Background(), "user_id", "target_user_id")
	if err != nil {
		panic(err)
	}
	log.Println(pm)
}

func ExampleClient_UndoMuting() {
	client := gotwtr.New("key")
	um, err := client.UndoMuting(context.Background(), "source_user_id", "target_user_id")
	if err != nil {
		panic(err)
	}
	log.Println(um)
}

func ExampleClient_LookUpSpace() {
	client := gotwtr.New("key")
	s, err := client.LookUpSpace(context.Background(), "space_id")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s)
}

func ExampleClient_LookUpSpaces() {
	client := gotwtr.New("key")
	ss, err := client.LookUpSpaces(context.Background(), []string{"space_id", "space_id2"})
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range ss.Spaces {
		log.Println(s)
	}
}

func ExampleClient_UsersPurchasedSpaceTicket() {
	client := gotwtr.New("key")
	tickets, err := client.UsersPurchasedSpaceTicket(context.Background(), "space_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range tickets.Users {
		log.Println(user)
	}
}

func ExampleClient_DiscoverSpaces() {
	client := gotwtr.New("key")
	discover, err := client.DiscoverSpaces(context.Background(), []string{"user_id", "user_id2"})
	if err != nil {
		log.Fatal(err)
	}
	for _, space := range discover.Spaces {
		log.Println(space)
	}
}

func ExampleClient_SearchSpaces() {
	client := gotwtr.New("key")
	spaces, err := client.SearchSpaces(context.Background(), "query")
	if err != nil {
		log.Fatal(err)
	}
	for _, space := range spaces.Spaces {
		log.Println(space)
	}
}

func ExampleClient_LookUpList() {
	client := gotwtr.New("key")
	l, err := client.LookUpList(context.Background(), "list_id")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(l)
}

func ExampleClient_LookUpAllListsOwned() {
	client := gotwtr.New("key")
	lists, err := client.LookUpAllListsOwned(context.Background(), "user_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, list := range lists.Lists {
		log.Println(list)
	}
}

func ExampleClient_CreateNewList() {
	client := gotwtr.New("key")
	l, err := client.CreateNewList(context.Background(), &gotwtr.CreateNewListBody{
		Name: "test v2 create list",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(l)
}

func ExampleClient_DeleteList() {
	client := gotwtr.New("key")
	l, err := client.DeleteList(context.Background(), "list_id")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(l)
}

func ExampleClient_UpdateMetaDataForList() {
	client := gotwtr.New("key")
	l, err := client.UpdateMetaDataForList(context.Background(), "list_id")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(l)
}

func ExampleClient_LookupUserBookmarks() {
	client := gotwtr.New("key")
	l, err := client.LookupUserBookmarks(context.Background(), "userID")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(l)
}

func ExampleClient_BookmarkTweet() {
	client := gotwtr.New("key")
	l, err := client.BookmarkTweet(context.Background(), "user_id", &gotwtr.BookmarkTweetBody{
		TweetID: "2022",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(l)
}

func ExampleClient_RemoveBookmarkOfTweet() {
	client := gotwtr.New("key")
	l, err := client.RemoveBookmarkOfTweet(context.Background(), "user_id", "tweet_id")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(l)
}

func ExampleClient_LookUpListTweets() {
	client := gotwtr.New("key")
	lt, err := client.LookUpListTweets(context.Background(), "list_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, tweet := range lt.Tweets {
		log.Println(tweet)
	}
}

func ExampleClient_ListMembers() {
	client := gotwtr.New("key")
	members, err := client.ListMembers(context.Background(), "list_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, member := range members.Users {
		log.Println(member)
	}
}

func ExampleClient_ListsSpecifiedUser() {
	client := gotwtr.New("key")
	lists, err := client.ListsSpecifiedUser(context.Background(), "user_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, list := range lists.Lists {
		log.Println(list)
	}
}

func ExampleClient_PostListMembers() {
	client := gotwtr.New("key")
	plm, err := client.PostListMembers(context.Background(), "list_id", "user_id")
	if err != nil {
		log.Println(err)
	}
	log.Println(plm)
}

func ExampleClient_UndoListMembers() {
	client := gotwtr.New("key")
	ulm, err := client.UndoListMembers(context.Background(), "list_id", "user_id")
	if err != nil {
		log.Println(err)
	}
	log.Println(ulm)
}

func ExampleClient_ListFollowers() {
	client := gotwtr.New("key")
	followers, err := client.ListFollowers(context.Background(), "list_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range followers.Users {
		log.Println(user)
	}
}

func ExampleClient_AllListsUserFollows() {
	client := gotwtr.New("key")
	lists, err := client.AllListsUserFollows(context.Background(), "user_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, list := range lists.Lists {
		log.Println(list)
	}
}

func ExampleClient_PostListFollows() {
	client := gotwtr.New("key")
	plf, err := client.PostListFollows(context.Background(), "list_id", "user_id")
	if err != nil {
		log.Println(err)
	}
	log.Println(plf)
}

func ExampleClient_UndoListFollows() {
	client := gotwtr.New("key")
	ulf, err := client.UndoListFollows(context.Background(), "list_id", "user_id")
	if err != nil {
		log.Println(err)
	}
	log.Println(ulf)
}

func ExampleClient_PinnedLists() {
	client := gotwtr.New("key")
	pl, err := client.PinnedLists(context.Background(), "user_id")
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range pl.Lists {
		log.Println(l)
	}
}

func ExampleClient_PostPinnedLists() {
	client := gotwtr.New("key")
	ppl, err := client.PostPinnedLists(context.Background(), "list_id", "user_id")
	if err != nil {
		log.Println(err)
	}
	log.Println(ppl)
}

func ExampleClient_UndoPinnedLists() {
	client := gotwtr.New("key")
	upl, err := client.UndoPinnedLists(context.Background(), "list_id", "user_id")
	if err != nil {
		log.Println(err)
	}
	log.Println(upl)
}

func ExampleClient_ComplianceJobs() {
	client := gotwtr.New("key")
	cj, err := client.ComplianceJobs(context.Background(), &gotwtr.ComplianceJobsOption{
		Type: gotwtr.ComplianceFieldTypeTweets,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cj)
}

func ExampleClient_ComplianceJob() {
	client := gotwtr.New("key")
	cj, err := client.ComplianceJob(context.Background(), 1382081613278814209)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cj)
	for _, e := range cj.Errors {
		log.Println(e)
	}
}

func ExampleClient_PostTweet() {
	client := gotwtr.New("key")
	t, err := client.PostTweet(context.Background(), &gotwtr.PostTweetOption{
		Text: "Hello World",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(t)
}

func ExampleClient_DeleteTweet() {
	client := gotwtr.New("key")
	_, err := client.DeleteTweet(context.Background(), "tweet_id")
	if err != nil {
		log.Fatal(err)
	}
}
