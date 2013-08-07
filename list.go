// Copyright 2013 Carlos Brando. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/carlosbrando/lunchy/agents"
)

func (c *Command) list() error {
	agents, err := agents.Find(c.pattern)
	if err != nil {
		return err
	}

	for _, agent := range agents {
		if c.long {
			fmt.Println(agent.Fullpath)
		} else {
			fmt.Println(agent.Name)
		}
	}
	return nil
}
