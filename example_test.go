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
		log.Fatal(err)
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
		log.Fatal(err)
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

func ExampleClient_RetrieveMultipleUsersWithIDs() {
	client := gotwtr.New("key")
	us, err := client.RetrieveMultipleUsersWithIDs(context.Background(), []string{"id", "id2"})
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range us.Users {
		fmt.Println(u)
	}
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

func ExampleClient_RetrieveSingleUserWithID() {
	client := gotwtr.New("key")
	u, err := client.RetrieveSingleUserWithID(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*u.User)
}

func ExampleClient_RetrieveSingleUserWithUserName() {
	client := gotwtr.New("key")
	un, err := client.RetrieveSingleUserWithUserName(context.Background(), "username")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*un.User)
}

func ExampleClient_Following() {
	client := gotwtr.New("key")
	followingUsers, err := client.Following(context.Background(), "id", &gotwtr.FollowOption{
		MaxResults: 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range followingUsers.Users {
		fmt.Println(u)
	}
}

func ExampleClient_Followers() {
	client := gotwtr.New("key")
	followerUsers, err := client.Followers(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range followerUsers.Users {
		fmt.Println(u)
	}
}

func ExampleClient_LookUpAllListsOwned() {
	client := gotwtr.New("key")
	ls, err := client.LookUpAllListsOwned(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range ls.Lists {
		fmt.Println(l)
	}
}

func ExampleClient_LookUpList() {
	client := gotwtr.New("key")
	l, err := client.LookUpList(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*l.List)
}

func ExampleClient_LookUpListFollowers() {
	client := gotwtr.New("key")
	us, err := client.LookUpListFollowers(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range us.Users {
		fmt.Println(u)
	}
}

func ExampleClient_LookUpAllListsUserFollows() {
	client := gotwtr.New("key")
	ls, err := client.LookUpAllListsUserFollows(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range ls.Lists {
		fmt.Println(l)
	}
}

func ExampleClient_ListsSpecifiedUser() {
	client := gotwtr.New("key")
	lmr, err := client.ListsSpecifiedUser(context.Background(), "84839422")
	if err != nil {
		log.Fatal(err)
	}
	for _, lm := range lmr.Lists {
		fmt.Println(lm)
	}
}

func ExampleClient_ListMembers() {
	client := gotwtr.New("key")
	lms, err := client.ListMembers(context.Background(), "listid")
	if err != nil {
		log.Fatal(err)
	}
	for _, lm := range lms.Users {
		fmt.Println(lm)
	}
}
func ExampleClient_LookUpListTweets() {
	client := gotwtr.New("key")
	ts, err := client.LookUpListTweets(context.Background(), "id")
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range ts.Tweets {
		fmt.Println(t)
	}
}

func ExampleClient_DiscoverSpaces_option() {
	client := gotwtr.New("key")
	dsr, err := client.DiscoverSpaces(context.Background(), []string{"id"}, &gotwtr.DiscoverSpacesOption{
		Expansions: []gotwtr.Expansion{
			gotwtr.ExpansionHostIDs,
			gotwtr.ExpansionCreatorID,
			gotwtr.ExpansionInvitedUserIDs,
			gotwtr.ExpansionSpeakerIDs,
		},
		TopicFields: []gotwtr.TopicField{
			gotwtr.TopicFieldName,
			gotwtr.TopicFieldID,
			gotwtr.TopicFieldDescription,
		},
		UserFields: []gotwtr.UserField{
			gotwtr.UserFieldCreatedAt,
			gotwtr.UserFieldDescription,
			gotwtr.UserFieldEntities,
			gotwtr.UserFieldID,
			gotwtr.UserFieldLocation,
			gotwtr.UserFieldName,
			gotwtr.UserFieldPinnedTweetID,
			gotwtr.UserFieldProfileImageURL,
			gotwtr.UserFieldProtected,
			gotwtr.UserFieldPublicMetrics,
			gotwtr.UserFieldURL,
			gotwtr.UserFieldUserName,
			gotwtr.UserFieldVerified,
			gotwtr.UserFieldWithHeld,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("---")
	for _, t := range dsr.Spaces {
		fmt.Println(t)
	}
	fmt.Println("---")
	for _, t := range dsr.Includes.Topics {
		fmt.Println(t)
	}
	fmt.Println("---")
	for _, t := range dsr.Includes.Users {
		fmt.Println(t)
	}
}

func ExampleClient_DiscoverSpaces_noOption() {
	client := gotwtr.New("key")
	dsr, err := client.DiscoverSpaces(context.Background(), []string{"id"})
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range dsr.Spaces {
		fmt.Println(t)
	}
}

func ExampleClient_LookUpSpace() {
	client := gotwtr.New("key")
	s, err := client.LookUpSpace(context.Background(), "spaceid")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s.Space)
}

func ExampleClient_LookUpSpaces() {
	client := gotwtr.New("key")
	ss, err := client.LookUpSpaces(context.Background(), []string{
		"spaceid1",
		"spaceid2",
	})
	if err != nil {
		log.Fatal(err)
	}
	for i, s := range ss.Spaces {
		fmt.Printf("index: %d, val: %v\n", i, s)
	}
}

func ExampleClient_UsersPurchasedSpaceTicket() {
	client := gotwtr.New("key")
	str, err := client.UsersPurchasedSpaceTicket(context.Background(), "spaceid")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}

func ExampleClient_SearchSpaces_option() {
	client := gotwtr.New("key")
	ssr, err := client.SearchSpaces(context.Background(), "hello", &gotwtr.SearchSpacesOption{
		SpaceFields: []gotwtr.SpaceField{
			gotwtr.SpaceFieldHostIDs,
			gotwtr.SpaceFieldCreatedAt,
			gotwtr.SpaceFieldCreatorID,
			gotwtr.SpaceFieldID,
			gotwtr.SpaceFieldLanguage,
			gotwtr.SpaceFieldInvittedUserIDs,
			gotwtr.SpaceFieldParticipantCount,
			gotwtr.SpaceFieldSpeakerIDs,
			gotwtr.SpaceFieldStartedAt,
			gotwtr.SpaceFieldState,
			gotwtr.SpaceFieldTitle,
			gotwtr.SpaceFieldUpdatedAt,
			gotwtr.SpaceFieldScheduledStart,
			gotwtr.SpaceFieldIsTicketed,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range ssr.Spaces {
		fmt.Println(s)
	}
	fmt.Println(ssr.Meta.ResultCount)
}

func ExampleClient_SearchSpaces_noOption() {
	client := gotwtr.New("key")
	ssr, err := client.SearchSpaces(context.Background(), "hello")
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range ssr.Spaces {
		fmt.Println(s)
	}
}
