package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New(gotwtr.WithBearerToken("key"))
	lms, err := client.ListMembers(context.Background(), "listid")
	if err != nil {
		panic(err)
	}
	for _, lm := range lms.Users {
		fmt.Println(lm)
	}
}
