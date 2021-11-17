package main

import (
	"context"
	"fmt"
	"github.com/sivchari/gotwtr"
	"time"
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

	// connect to  stream
	ch := make(chan gotwtr.ConnectToStreamResponse, 5)
	errCh := make(chan error)
	stream := client.ConnectToStream(context.Background(), ch, errCh)
	fmt.Println("streaming...")
	ctx, cancel:= context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case data := <-ch:
				fmt.Println(data.Tweet)
			case err := <-errCh:
				fmt.Println(err)
			case <-ctx.Done():
				break
			}
		}
	}(ctx)
	time.Sleep(time.Second * 10)
	cancel()
	stream.Stop()
	fmt.Println("done")

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
