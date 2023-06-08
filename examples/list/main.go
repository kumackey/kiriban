package main

import (
	"fmt"
	"github.com/kumackey/kiriban/pkg/kiriban"
)

func main() {
	c, _ := kiriban.NewChecker(kiriban.ExceptionalKiribanOption([]kiriban.ExceptionalKiriban{{Value: 1101, Reason: "birthday"}}))
	val := 100
	for val < 10000 {
		fmt.Println(val, c.JudgeKinds(val)[0].String())
		val = c.Next(val)
	}
}
