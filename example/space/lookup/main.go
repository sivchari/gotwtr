package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")

	// look up by ID
	s, err := client.LookUpSpaceByID(context.Background(), "spaceid")
	if err != nil {
		panic(err)
	}
	fmt.Println(s.Space)
}
