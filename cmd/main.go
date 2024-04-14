package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/kumackey/kiriban/kiriban"
	"log"
	"os"
	"strconv"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	e := flag.String("e", "", "Event name")
	flag.Parse()

	_, err = toEventName(*e)
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

	d, err := kiriban.NewDeterminator(kiriban.EnableDigitBasedRoundDetermination())
	if err != nil {
		log.Fatalln(err)
	}

	if !d.IsKiriban(issueNumber) {
		fmt.Printf("#%d is not kiriban.\n", issueNumber)
		os.Exit(0)
	}

	fmt.Printf("#%d is kiriban!\n", issueNumber)

	ctx := context.Background()

	ic := newIssueCommenter(newGithubClient(ctx, cfg.githubToken), d)
	comment, err := ic.execute(ctx, cfg, issueNumber)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Commented: %s\n", comment.GetHTMLURL())
}
