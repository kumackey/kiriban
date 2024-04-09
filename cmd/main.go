package main

import (
	"fmt"
	"os"
)

func main() {
	events := os.Getenv("INPUT_EVENTS")
	fmt.Println(events)
}
