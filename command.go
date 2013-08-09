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
	case "start":
		if err := c.start(); err != nil {
			return err
		}
	default:
		showBanner()
	}

	return nil
}

// runCommand execute a command and returns the output as a string.
func runCommand(name string, arg ...string) (string, error) {
	output, err := exec.Command(name, arg...).Output()

	if err == nil {
		return string(output), nil
	} else {
		return "", err
	}
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
