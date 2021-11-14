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
		fmt.Println(data.Tweet)
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	}
}
