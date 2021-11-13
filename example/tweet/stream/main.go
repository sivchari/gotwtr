package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	// sampled stream
	ch := make(chan gotwtr.SampledStreamResponse)
	errCh := make(chan error)
	client.SampledStream(context.Background(), ch, errCh)
	select {
	case data := <-ch:
		for _, d := range data.Tweets {
			fmt.Println(d)
		}
	case err:= <-errCh:
		if err != nil {
			panic(err)
		}
	}
}
