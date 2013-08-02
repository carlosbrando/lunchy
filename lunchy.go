package main

import (
	"log"
	"os"
)

func formatFlags(flags []string, c *Command) {
	for _, flag := range flags {
		switch flag {
		case "-F", "--force":
			c.force = true
		case "-v", "--verbose":
			c.verbose = true
		case "-w", "--write":
			c.write = true
		case "-l", "--long":
			c.long = true
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		showBanner()
		return
	}

	cmd := &Command{command: args[0]}
	if len(args) > 1 {
		cmd.pattern = args[1]
	}

	if len(args) > 2 {
		formatFlags(args[2:], cmd)
	}

	err := cmd.execute()
	if err != nil {
		log.Fatalf("Uh oh, I didn't expect this:\n%s\n", err)
	}
}
