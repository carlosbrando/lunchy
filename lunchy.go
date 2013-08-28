// Copyright 2013 Carlos Brando. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		printBanner()
		return
	}

	cmd := &Command{command: args[0]}
	cmd.parseFlags(args[1:])

	err := cmd.exec()
	if err != nil {
		fmt.Println(err)
	}
}
