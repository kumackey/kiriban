package kiriban_test

import (
	"fmt"
	"github.com/kumackey/kiriban/kiriban"
)

func ExampleDeterminator_IsKiriban() {
	d, _ := kiriban.NewDeterminator()

	v1 := d.IsKiriban(10000)
	fmt.Printf("10000 is kiriban? -> %t\n", v1)

	v2 := d.IsKiriban(10001)
	fmt.Printf("10001 is kiriban? -> %t\n", v2)

	// Output:
	// 10000 is kiriban? -> true
	// 10001 is kiriban? -> false
}

func ExampleDeterminator_KiribanKinds() {
	d, _ := kiriban.NewDeterminator()

	v3 := d.KiribanKinds(100000)
	fmt.Printf("100000 is %s\n", v3[0])
	// Output:
	// 100000 is Round
}

func ExampleDeterminator_Next() {
	d, _ := kiriban.NewDeterminator(kiriban.ExceptionalKiribanOption([]kiriban.ExceptionalKiriban{
		// Any kiriban can be added.
		{Value: 1101, Reason: "birthday"},
	}))

	val := 0
	for val < 10000 {
		val = d.Next(val)
		fmt.Println(val, d.KiribanKinds(val)[0].String())
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
