// Copyright 2013 Carlos Brando. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/carlosbrando/lunchy/agents"
)

func (c *Command) stop() error {
	agent, err := agents.FindOne(c.pattern)
	if err != nil {
		return err
	}

	if agent != nil {
		fmt.Println("stopping", agent.Name)
		execProcess("launchctl", "unload", agent.Fullpath)
	}

	return nil
}
