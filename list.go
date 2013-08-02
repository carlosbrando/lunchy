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
