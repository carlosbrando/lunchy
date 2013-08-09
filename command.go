// Copyright 2013 Carlos Brando. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package main

import (
	"os"
	"os/exec"
	"syscall"
)

type Command struct {
	command string
	pattern string
	long    bool
	force   bool
	verbose bool
	write   bool
}

// execute verifies what command to run.
func (c *Command) execute() error {
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
	case "start", "stop":
		if err := c.launchctl(); err != nil {
			return err
		}
	default:
		showBanner()
	}

	return nil
}

func runCmd(name string, arg ...string) (output string, err error) {
	out, err := exec.Command(name, arg...).Output()
	if err == nil {
		output = string(out)
	}
	return
}

// execProcess executes a process replacing the current Go process.
func execProcess(process string, args ...string) error {
	binary, err := exec.LookPath(process)
	if err != nil {
		return err
	}

	err = syscall.Exec(binary, append([]string{process}, args...), os.Environ())
	if err != nil {
		return err
	}

	return nil
}
