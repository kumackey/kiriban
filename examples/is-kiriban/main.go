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

	v2 := c.IsKiriban(10001)
	fmt.Printf("10001 is kiriban? -> %t\n", v2)
	// Output: 10001 is kiriban? -> false

	v3 := c.JudgeKinds(100000)
	fmt.Printf("100000 is %s\n", v3[0])
	// Output: 100000 is Trailing zeros

	v4 := c.Next(100000)
	fmt.Printf("next kiriban of 100000 is %d\n", v4)
	// Output: next kiriban of 100000 is 111111
}
