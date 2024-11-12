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

func ExampleChecker_KiribanKind() {
	c, _ := kiriban.NewChecker()

	v3, _ := c.KiribanKind(100000)
	fmt.Printf("100000 is %s\n", v3)
	// Output:
	// 100000 is Round
}

func ExampleChecker_Next() {
	c, _ := kiriban.NewChecker(kiriban.ExceptionalKiribanOption([]kiriban.ExceptionalKiriban{
		{Value: 1101, Reason: "birthday"},
	}))

	val := 0
	for val < 10000 {
		val = c.Next(val)
		kind, _ := c.KiribanKind(val)
		fmt.Println(val, kind.String())
	}

	// Output:
	// 10 Round
	// 11 Repdigit
	// 20 Round
	// 22 Repdigit
	// 30 Round
	// 33 Repdigit
	// 40 Round
	// 44 Repdigit
	// 50 Round
	// 55 Repdigit
	// 60 Round
	// 66 Repdigit
	// 70 Round
	// 77 Repdigit
	// 80 Round
	// 88 Repdigit
	// 90 Round
	// 99 Repdigit
	// 100 Round
	// 111 Repdigit
	// 123 Consecutive
	// 200 Round
	// 210 Consecutive
	// 222 Repdigit
	// 234 Consecutive
	// 300 Round
	// 321 Consecutive
	// 333 Repdigit
	// 345 Consecutive
	// 400 Round
	// 432 Consecutive
	// 444 Repdigit
	// 456 Consecutive
	// 500 Round
	// 543 Consecutive
	// 555 Repdigit
	// 567 Consecutive
	// 600 Round
	// 654 Consecutive
	// 666 Repdigit
	// 678 Consecutive
	// 700 Round
	// 765 Consecutive
	// 777 Repdigit
	// 789 Consecutive
	// 800 Round
	// 876 Consecutive
	// 888 Repdigit
	// 900 Round
	// 987 Consecutive
	// 999 Repdigit
	// 1000 Round
	// 1101 Exceptional kiriban: 1101, birthday
	// 1111 Repdigit
	// 1234 Consecutive
	// 2000 Round
	// 2222 Repdigit
	// 2345 Consecutive
	// 3000 Round
	// 3210 Consecutive
	// 3333 Repdigit
	// 3456 Consecutive
	// 4000 Round
	// 4321 Consecutive
	// 4444 Repdigit
	// 4567 Consecutive
	// 5000 Round
	// 5432 Consecutive
	// 5555 Repdigit
	// 5678 Consecutive
	// 6000 Round
	// 6543 Consecutive
	// 6666 Repdigit
	// 6789 Consecutive
	// 7000 Round
	// 7654 Consecutive
	// 7777 Repdigit
	// 8000 Round
	// 8765 Consecutive
	// 8888 Repdigit
	// 9000 Round
	// 9876 Consecutive
	// 9999 Repdigit
	// 10000 Round
}

func ExampleKiriban_Iter() {
	opt := kiriban.ExceptionalKiribanOption([]kiriban.ExceptionalKiriban{
		{Value: 1101, Reason: "birthday"},
	})
	for i := range kiriban.Iter(opt) {
		if i > 10000 {
			break
		}

		fmt.Println(i)
	}

	// Output:
	// 10
	// 11
	// 20
	// 22
	// 30
	// 33
	// 40
	// 44
	// 50
	// 55
	// 60
	// 66
	// 70
	// 77
	// 80
	// 88
	// 90
	// 99
	// 100
	// 111
	// 123
	// 200
	// 210
	// 222
	// 234
	// 300
	// 321
	// 333
	// 345
	// 400
	// 432
	// 444
	// 456
	// 500
	// 543
	// 555
	// 567
	// 600
	// 654
	// 666
	// 678
	// 700
	// 765
	// 777
	// 789
	// 800
	// 876
	// 888
	// 900
	// 987
	// 999
	// 1000
	// 1101
	// 1111
	// 1234
	// 2000
	// 2222
	// 2345
	// 3000
	// 3210
	// 3333
	// 3456
	// 4000
	// 4321
	// 4444
	// 4567
	// 5000
	// 5432
	// 5555
	// 5678
	// 6000
	// 6543
	// 6666
	// 6789
	// 7000
	// 7654
	// 7777
	// 8000
	// 8765
	// 8888
	// 9000
	// 9876
	// 9999
	// 10000
}
