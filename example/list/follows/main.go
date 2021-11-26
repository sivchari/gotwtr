package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	// look up list followers by id
	us, err := client.LookUpListFollowersByID(context.Background(), "id")
	if err != nil {
		panic(err)
	}
	for _, u := range us.Users {
		fmt.Println(u)
	}

	// look up lists user following by ID
	ls, err := client.LookUpListsUserFollowingByID(context.Background(), "id")
	if err != nil {
		panic(err)
	}
	for _, l := range ls.Lists {
		fmt.Println(l)
	}
}
