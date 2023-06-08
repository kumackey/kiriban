# kiriban

## Usage example

```bash
$ go get github.com/kumackey/kiriban
```

```go:main.go
package main

import (
	"fmt"
	"github.com/kumackey/kiriban/pkg/kiriban"
)

func main() {
	c, _ := kiriban.NewChecker()

	v1 := c.IsKiriban(10000)
	fmt.Printf("10000 is kiriban? -> %t\n", v1)
	// Output: 10000 is kiriban? -> true

	v2 := c.IsKiriban(11111)
	fmt.Printf("11111 is kiriban? -> %t\n", v2)
	// Output: 11111 is kiriban? -> true

	v3 := c.IsKiriban(12345)
	fmt.Printf("12345 is kiriban? -> %t\n", v3)
	// Output: 12345 is kiriban? -> true

	v4 := c.IsKiriban(10001)
	fmt.Printf("10001 is kiriban? -> %t\n", v4)
	// Output: 10001 is kiriban? -> false

	v5 := c.Next(10001)
	fmt.Printf("next kiriban of 10001 -> %d\n", v5)
	// Output: next kiriban of 10001 -> 11111
}
```

See below for other usage examples.

https://pkg.go.dev/github.com/kumackey/kiriban/pkg/kiriban#pkg-examples

## What is "kiriban"?

"Kiriban" is an Internet term derived from the Japanese "kiri no ii ban go", which stands for "round number" or "nice
number."
It refers to the occurrence of a significant or noteworthy number in an online context, particularly when it appears as
an access count on a website.
The term has also been extended to refer to "round numbers" or "nice numbers" outside of websites.
