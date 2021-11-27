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
	ch := make(chan gotwtr.VolumeStreamsResponse, 5)
	errCh := make(chan error)
	stream := client.VolumeStreams(context.Background(), ch, errCh)
	fmt.Println("streaming...")
	done := make(chan struct{})
	go func(done chan struct{}) {
		for {
			select {
			case data := <-ch:
				fmt.Println(data.Tweet)
			case err := <-errCh:
				fmt.Println(err)
			case <-done:
				break
			}
		}
	}(done)
	time.Sleep(time.Second * 10)
	close(done)
	stream.Stop()
	fmt.Println("done")
}
