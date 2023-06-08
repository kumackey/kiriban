# kiriban

## Usage example

```bash
$ go get github.com/kumackey/kiriban
```

### Determine if it is kiriban

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

### list up kiribans

```go:main.go
package main

import (
	"fmt"
	"github.com/kumackey/kiriban/pkg/kiriban"
)

func main() {
	c, _ := kiriban.NewChecker(kiriban.ExceptionalKiribanOption([]kiriban.ExceptionalKiriban{{1101, "birthday"}}))
	val := 100
	for val < 10000 {
		fmt.Println(val, c.JudgeKinds(val)[0].String())
		val = c.Next(val)
	}
}
```

```bash 
$ go run main.go
100 Trailing zeros
111 Repdigit
123 Consecutive
200 Trailing zeros
210 Consecutive
222 Repdigit
234 Consecutive
300 Trailing zeros
321 Consecutive
333 Repdigit
345 Consecutive
400 Trailing zeros
432 Consecutive
444 Repdigit
456 Consecutive
500 Trailing zeros
543 Consecutive
555 Repdigit
567 Consecutive
600 Trailing zeros
654 Consecutive
666 Repdigit
678 Consecutive
700 Trailing zeros
765 Consecutive
777 Repdigit
789 Consecutive
800 Trailing zeros
876 Consecutive
888 Repdigit
900 Trailing zeros
987 Consecutive
999 Repdigit
1000 Trailing zeros
1101 Exceptional kiriban: 1101, birthday
1111 Repdigit
1234 Consecutive
2000 Trailing zeros
2222 Repdigit
2345 Consecutive
3000 Trailing zeros
3210 Consecutive
3333 Repdigit
3456 Consecutive
4000 Trailing zeros
4321 Consecutive
4444 Repdigit
4567 Consecutive
5000 Trailing zeros
5432 Consecutive
5555 Repdigit
5678 Consecutive
6000 Trailing zeros
6543 Consecutive
6666 Repdigit
6789 Consecutive
7000 Trailing zeros
7654 Consecutive
7777 Repdigit
8000 Trailing zeros
8765 Consecutive
8888 Repdigit
9000 Trailing zeros
9876 Consecutive
9999 Repdigit
```


## What is "kiriban"?

"Kiriban" is an Internet term derived from the Japanese "kiri no ii ban go", which stands for "round number" or "nice
number."
It refers to the occurrence of a significant or noteworthy number in an online context, particularly when it appears as
an access count on a website.
The term has also been extended to refer to "round numbers" or "nice numbers" outside of websites.
