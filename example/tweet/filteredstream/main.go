package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("AAAAAAAAAAAAAAAAAAAAAIhYUQEAAAAAV8vbjMIPhhXKPVI2Ea8n5RAR%2BbM%3DNh7spVQNQRA9vOHxkeogX2BPNbHSEEGhwnnLKuWgcQ7uFPuVoI")
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
