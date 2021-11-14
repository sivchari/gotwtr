package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	// sampled stream
	ch := make(chan gotwtr.SampledStreamResponse, 5)
	errCh := make(chan error)
	stream := client.SampledStream(context.Background(), ch, errCh)
	fmt.Println("streaming...")
	var i int
	for ; i < 5; i++ {
		select {
		case data := <-ch:
			fmt.Println(data)
		case err := <-errCh:
			if err != nil {
				panic(err)
			}
		default:
			time.Sleep(time.Second * 10)
		}
	}
	stream.Stop()
	fmt.Println("done")
}
