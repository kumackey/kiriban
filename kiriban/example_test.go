package kiriban_test

import (
	"fmt"
	"github.com/kumackey/kiriban/kiriban"
)

func ExampleChecker_IsKiriban() {
	c, _ := kiriban.NewChecker()

	v1 := c.IsKiriban(10000)
	fmt.Printf("10000 is kiriban? -> %t\n", v1)

	v2 := c.IsKiriban(10001)
	fmt.Printf("10001 is kiriban? -> %t\n", v2)

	// Output:
	// 10000 is kiriban? -> true
	// 10001 is kiriban? -> false
}

func ExampleChecker_JudgeKinds() {
	c, _ := kiriban.NewChecker()

	v3 := c.JudgeKinds(100000)
	fmt.Printf("100000 is %s\n", v3[0])
	// Output:
	// 100000 is Round number
}

func ExampleChecker_Next() {
	c, _ := kiriban.NewChecker(kiriban.ExceptionalKiribanOption([]kiriban.ExceptionalKiriban{
		// Any kiriban can be added.
		{Value: 1101, Reason: "birthday"},
	}))

	val := 100
	for val < 10000 {
		fmt.Println(val, c.JudgeKinds(val)[0].String())
		val = c.Next(val)
	}

	// Output:
	// 100 Round number
	// 111 Repdigit
	// 123 Consecutive
	// 200 Round number
	// 210 Consecutive
	// 222 Repdigit
	// 234 Consecutive
	// 300 Round number
	// 321 Consecutive
	// 333 Repdigit
	// 345 Consecutive
	// 400 Round number
	// 432 Consecutive
	// 444 Repdigit
	// 456 Consecutive
	// 500 Round number
	// 543 Consecutive
	// 555 Repdigit
	// 567 Consecutive
	// 600 Round number
	// 654 Consecutive
	// 666 Repdigit
	// 678 Consecutive
	// 700 Round number
	// 765 Consecutive
	// 777 Repdigit
	// 789 Consecutive
	// 800 Round number
	// 876 Consecutive
	// 888 Repdigit
	// 900 Round number
	// 987 Consecutive
	// 999 Repdigit
	// 1000 Round number
	// 1101 Exceptional kiriban: 1101, birthday
	// 1111 Repdigit
	// 1234 Consecutive
	// 2000 Round number
	// 2222 Repdigit
	// 2345 Consecutive
	// 3000 Round number
	// 3210 Consecutive
	// 3333 Repdigit
	// 3456 Consecutive
	// 4000 Round number
	// 4321 Consecutive
	// 4444 Repdigit
	// 4567 Consecutive
	// 5000 Round number
	// 5432 Consecutive
	// 5555 Repdigit
	// 5678 Consecutive
	// 6000 Round number
	// 6543 Consecutive
	// 6666 Repdigit
	// 6789 Consecutive
	// 7000 Round number
	// 7654 Consecutive
	// 7777 Repdigit
	// 8000 Round number
	// 8765 Consecutive
	// 8888 Repdigit
	// 9000 Round number
	// 9876 Consecutive
	// 9999 Repdigit
}
