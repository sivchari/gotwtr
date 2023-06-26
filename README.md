# gotwtr

[![Go Reference](https://pkg.go.dev/badge/github.com/sivchari/gotwtr.svg)](https://pkg.go.dev/github.com/sivchari/gotwtr)
[![reviewdog](https://github.com/sivchari/gotwtr/actions/workflows/lint.yml/badge.svg)](https://github.com/sivchari/gotwtr/actions/workflows/lint.yml)
[![test](https://github.com/sivchari/gotwtr/actions/workflows/test.yml/badge.svg)](https://github.com/sivchari/gotwtr/actions/workflows/test.yml)

gotwtr is a Go client library for the Twitter v2 API.

## Note

We covers only Twitter v2 API supported by OAuth 2.0 Bearer Token.

We will had worked on it, when new one is be handled OAuth 2.0.

## Installation

```console
go get github.com/sivchari/gotwtr
```

## Documentation

Please see [GoDoc](https://pkg.go.dev/github.com/sivchari/gotwtr)

## Example

### Tweet lookup
```go
package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("YOUR_TWITTER_BEARER_TOKEN")
	// look up multiple tweets
	ts, err := client.RetrieveMultipleTweets(context.Background(), []string{"id", "id2"})
	if err != nil {
		panic(err)
	}
	for _, t := range ts.Tweets {
		fmt.Println(t)
	}

	// look up single tweet
	t, err := client.RetrieveSingleTweet(context.Background(), "id")
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Tweet)
}
```

If you wanna more example, please see examples dir.

These are covered all code gotwtr provides Twitter v2 API interface.

## Contributing

We are welcome to contribute to this project.

Fork and make a Pull Request, or create an Issue if you see any problem or any enhancement, feature request.
