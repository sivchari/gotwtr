package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	// look up by ID
	s, err := client.LookUpSpace(context.Background(), "spaceid")
	if err != nil {
		panic(err)
	}
	fmt.Println(s.Space)

	ss, err := client.LookUpSpaces(context.Background(), []string{
		"spaceid1",
		"spaceid2",
	})
	if err != nil {
		panic(err)
	}

	for i, s := range ss.Spaces {
		fmt.Printf("index: %d, val: %v\n", i, s)
	}

	str, err := client.UsersPurchasedSpaceTicket(context.Background(), "spaceid")
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
