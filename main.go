package main

import (
	"fmt"
	"os"
	"time"
)

const frequency = 5 * time.Minute

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <repro|check>\n", os.Args[0])
		os.Exit(1)
	}

	switch os.Args[1] {
	case "repro":
		for {
			fmt.Fprintf(os.Stderr, "===== starting test %s =====\n", time.Now())
			if err := repro(); err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			}
			fmt.Fprintf(os.Stderr, "===== test completed %s, waiting %s for next round =====\n", time.Now(), frequency)
			time.Sleep(frequency)
		}
	case "check":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "usage: %s check <journal>\n", os.Args[0])
			os.Exit(1)
		}
		if err := check(os.Args[2]); err != nil {
			panic(err)
		}
	}

}
