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
