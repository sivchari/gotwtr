package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("keuy")
	lmr, err := client.ListsSpecifiedUser(context.Background(), "84839422")
	if err != nil {
		panic(err)
	}
	for _, lm := range lmr.Lists {
		fmt.Println(lm)
	}
}
