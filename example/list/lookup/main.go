package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New(gotwtr.WithBearerToken("key"))
	// look up owned lists by id
	ls, err := client.LookUpAllListsOwned(context.Background(), "id")
	if err != nil {
		panic(err)
	}
	for _, l := range ls.Lists {
		fmt.Println(l)
	}

	// look up list by ID
	l, err := client.LookUpList(context.Background(), "id")
	if err != nil {
		panic(err)
	}
	fmt.Println(*l.List)
}
