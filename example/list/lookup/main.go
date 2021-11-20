package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	// look up owned lists by id
	ls, err := client.LookUpOwnedListsByID(context.Background(), "id")
	if err != nil {
		panic(err)
	}
	for _, l := range ls.Lists {
		fmt.Println(l)
	}

	// look up list by ID
	l, err := client.LookUpListByID(context.Background(), "id")
	if err != nil {
		panic(err)
	}
	fmt.Println(*l.List)
}
