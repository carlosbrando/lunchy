// Copyright 2013 Carlos Brando. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package main

import (
	"fmt"
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
		fmt.Println(err)
	}
}
