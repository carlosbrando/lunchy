// Copyright 2013 Carlos Brando. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package main

type Command struct {
	command string
	pattern string
	long    bool
	force   bool
	verbose bool
	write   bool
}

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
	default:
		showBanner()
	}

	return nil
}
