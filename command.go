// Copyright 2013 Carlos Brando. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package main

import "os/exec"

type Command struct {
	command string
	pattern string
	long    bool
	force   bool
	verbose bool
	write   bool
}

// execute verifies what command to run.
func (c *Command) exec() error {
	switch c.command {
	case "list", "ls":
		if err := c.list(); err != nil {
			return err
		}
	case "show":
		if err := c.show(); err != nil {
			return err
		}
	case "status":
		if err := c.status(); err != nil {
			return err
		}
	case "start", "stop", "restart":
		if err := c.launchctl(); err != nil {
			return err
		}
	case "edit":
		if err := c.edit(); err != nil {
			return err
		}
	default:
		printBanner()
	}

	return nil
}

func (c *Command) clone(command string) *Command {
	return &Command{
		command: command,
		pattern: c.pattern,
		long:    c.long,
		force:   c.force,
		verbose: c.verbose,
		write:   c.write,
	}
}

func runCmd(name string, args ...string) (output string, err error) {
	binary, err := exec.LookPath(name)
	if err != nil {
		return
	}

	out, err := exec.Command(binary, args...).Output()
	if err == nil {
		output = string(out)
	}
	return
}
