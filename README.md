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

	v2 := c.IsKiriban(10001)
	fmt.Printf("10001 is kiriban? -> %t\n", v2)

	v3 := c.JudgeKinds(100000)
	fmt.Printf("100000 is %s\n", v3[0])

	v4 := c.Next(100000)
	fmt.Printf("next kiriban of 100000 is %d\n", v4)
}
```

```bash
$ go run main.go
10000 is kiriban? -> true
10001 is kiriban? -> false
100000 is Trailing zeros
next kiriban of 100000 is 111111
```

## What is "kiriban"?

"Kiriban" is an Internet term derived from the Japanese "kiri no ii ban go", which stands for "round number" or "nice
number."
It refers to the occurrence of a significant or noteworthy number in an online context, particularly when it appears as
an access count on a website.
The term has also been extended to refer to "round numbers" or "nice numbers" outside of websites.
