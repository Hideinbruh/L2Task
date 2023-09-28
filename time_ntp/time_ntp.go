package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func main() {
	timeNTP, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting time\n")
		os.Exit(1)
	}
	fmt.Printf("Текущее время: %s\n", timeNTP)
}
