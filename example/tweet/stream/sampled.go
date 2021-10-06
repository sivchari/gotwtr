package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	// sampled stream
	tsr, err := client.SampledStream(context.Background(), &gotwtr.SampledStreamOpts)
	if err != nil {
		panic(err)
	}
	for _, t := range tsr.Tweets {
		fmt.Println(t)
	}
}
