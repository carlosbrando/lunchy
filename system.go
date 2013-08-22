// Copyright 2013 Carlos Brando. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/carlosbrando/lunchy/agents"
)

func (c *Command) launchctl() error {
	agent, err := agents.FindOne(c.pattern)
	if err != nil {
		return err
	}

	if agent != nil {
		switch c.command {
		case "start":
			fmt.Println("starting", agent.Name)
		case "stop":
			fmt.Println("stopping", agent.Name)
		case "restart":
			c.clone("stop").launchctl()
			c.clone("start").launchctl()
		}

		runCmd("launchctl", c.args(agent.Fullpath)...)
	}

	return nil
}

func (c *Command) args(more ...string) []string {
	var args []string

	if c.force {
		args = append(args, "-F")
	}

	if c.write {
		args = append(args, "-w")
	}

	switch c.command {
	case "start":
		args = append([]string{"load"}, args...)
	case "stop":
		args = append([]string{"unload"}, args...)
	}

	return append(args, more...)
}
