package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New(gotwtr.WithBearerToken("key"))
	// look up users
	us, err := client.RetrieveMultipleUsersWithIDs(context.Background(), []string{"id", "id2"})
	if err != nil {
		panic(err)
	}
	for _, u := range us.Users {
		fmt.Println(u)
	}

	// look up user by ID
	u, err := client.RetrieveSingleUserWithID(context.Background(), "id")
	if err != nil {
		panic(err)
	}
	fmt.Println(*u.User)

	// look up users by username
	uns, err := client.RetrieveMultipleUsersWithUserNames(context.Background(), []string{"username", "username2"})
	if err != nil {
		panic(err)
	}
	for _, un := range uns.Users {
		fmt.Println(un)
	}

	// look up user by username
	un, err := client.RetrieveSingleUserWithUserName(context.Background(), "username")
	if err != nil {
		panic(err)
	}
	fmt.Println(*un.User)
}
