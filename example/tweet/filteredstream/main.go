package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	res, err := client.AddOrDeleteRules(context.Background(), &gotwtr.AddOrDeleteJSONBody{
		Add: []*gotwtr.Add{
			{
				Value: "puppy has:media",
				Tag:   "puppies with media",
			},
			{
				Value: "meme has:images",
			},
		},
		Delete: &gotwtr.Delete{IDs: []string{}},
	})
	if err != nil {
		panic(res)
	}
	// retrieve Stream rules
	ts, err := client.RetrieveStreamRules(context.Background())
	if err != nil {
		panic(err)
	}
	for _, t := range ts.Rules {
		fmt.Println(t)
	}
}
