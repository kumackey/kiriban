package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/kumackey/kiriban/kiriban"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	t := flag.String("t", "", "Event name")
	udks := flag.String("u", "", "User-defined kiribans")
	flag.Parse()

	_, err = toEventName(*t)
	if err != nil {
		log.Fatalln(err)
	}

	eks, err := toExceptionalKiribans(*udks)
	if err != nil {
		log.Fatalln(err)
	}

	args := flag.Args()
	if len(args) != 1 {
		log.Fatalln("Invalid arguments")
	}

	issueNumber, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(err)
	}

	c, err := kiriban.NewChecker(kiriban.EnableDigitBasedRoundDetermination(), kiriban.ExceptionalKiribanOption(eks))
	if err != nil {
		log.Fatalln(err)
	}

	if !c.IsKiriban(issueNumber) {
		fmt.Printf("#%c is not kiriban.\n", issueNumber)
		os.Exit(0)
	}

	fmt.Printf("#%c is kiriban!\n", issueNumber)

	ctx := context.Background()

	ic := newIssueCommenter(newGithubClient(ctx, cfg.githubToken), c)
	comment, err := ic.execute(ctx, cfg, issueNumber)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Commented: %s\n", comment.GetHTMLURL())
}

func toExceptionalKiribans(e string) ([]kiriban.ExceptionalKiriban, error) {
	if e == "" {
		return nil, nil
	}

	e = strings.ReplaceAll(e, " ", "")
	parts := strings.Split(e, ",")
	kiribans := make([]kiriban.ExceptionalKiriban, 0, len(parts))

	for _, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid exceptional kiriban value: %v", err)
		}

		kiribans = append(kiribans, kiriban.ExceptionalKiriban{Value: value})
	}

	return kiribans, nil
}
