package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")

	// lookup followings
	followingUsers, err := client.Following(context.Background(), "id", &gotwtr.FollowOption{
		MaxResults: 10,
	})
	if err != nil {
		panic(err)
	}
	for _, u := range followingUsers.Users {
		fmt.Println(u)
	}

	// lookup followers
	followerUsers, err := client.Followers(context.Background(), "id")
	if err != nil {
		panic(err)
	}
	for _, u := range followerUsers.Users {
		fmt.Println(u)
	}

	// post following
	p, err := client.PostFollowing(context.Background(), "id", "target_user_id")
	if err != nil {
		panic(err)
	}
	fmt.Println(p)

	// undo following
	d, err := client.UndoFollowing(context.Background(), "source_user_id", "target_user_id")
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
}
