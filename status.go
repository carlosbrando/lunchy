// Copyright 2013 Carlos Brando. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/carlosbrando/lunchy/agents"
)

var headerPrinted bool = false

// printHeader prints the header if it was not printed before.
func printHeader(header string) {
	if !headerPrinted {
		fmt.Println(header)
		headerPrinted = true
	}
}

func (c *Command) status() error {
	cmd := exec.Command("launchctl", "list")

	out, err := cmd.Output()
	if err != nil {
		return err
	}

	output := strings.Split(string(out), "\n")

	agents, err := agents.Find(c.pattern)
	if err != nil {
		return err
	}

	for _, line := range output {
		for _, agent := range agents {
			matched, err := regexp.MatchString(agent.Name, line)
			if err != nil {
				return err
			}

			if matched {
				printHeader(output[0])
				fmt.Println(line)
			}
		}
	}

	return nil
}
