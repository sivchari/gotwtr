package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")

	followingUsers, err := client.Following(context.Background(), "id")
	if err != nil {
		panic(err)
	}
	for _, u := range followingUsers.Users {
		fmt.Println(u)
	}

	followerUsers, err := client.Followers(context.Background(), "id")
	if err != nil {
		panic(err)
	}
	for _, u := range followerUsers.Users {
		fmt.Println(u)
	}
}
