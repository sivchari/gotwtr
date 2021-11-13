package main

import (
	"context"
	"fmt"
	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	_, err := client.AddOrDeleteRules(context.Background(), &gotwtr.AddOrDeleteJSONBody{
		Add: []*gotwtr.AddRule{
			{
				Value: "puppy has:media",
				Tag:   "puppies with media",
			},
			{
				Value: "meme has:images",
			},
		},
	})
	if err != nil {
		panic(err)
	}

	responses, err := client.ConnectToStream(context.Background(), 10)
	if err != nil {
		panic(err)
	}
	if responses.Error != nil {
		fmt.Println(responses.Error)
		return
	}
	for _, resp := range responses.Chunks {
		fmt.Println(resp.Tweet.Text)
	}

	// retrieve Stream rules
	ts, err := client.RetrieveStreamRules(context.Background())
	if err != nil {
		panic(err)
	}
	var ids []string
	for _, t := range ts.Rules {
		fmt.Println(t)
		ids = append(ids, t.ID)
	}

	// delete Stream rules
	_, err = client.AddOrDeleteRules(context.Background(), &gotwtr.AddOrDeleteJSONBody{
		Delete: &gotwtr.DeleteRule{
			IDs: ids,
		},
	})
	if err != nil {
		panic(err)
	}
}
