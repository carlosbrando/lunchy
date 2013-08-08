// Copyright 2013 Carlos Brando. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/carlosbrando/lunchy/agents"
)

func (c *Command) show() error {
	if c.pattern == "" {
		return errors.New("show [name]")
	}

	agent, err := agents.FindOne(c.pattern)
	if err != nil {
		return err
	}

	if agent != nil {
		data, err := ioutil.ReadFile(agent.Fullpath)
		if err != nil {
			return err
		}
		fmt.Println(string(data))
	}

	return nil
}
