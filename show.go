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

	// TODO: need to be case insensitive
	agents, err := agents.Find(c.pattern)
	if err != nil {
		return err
	}

	switch len(agents) {
	case 0:
		// do nothing
	case 1:
		data, err := ioutil.ReadFile(agents[0].Fullpath)
		if err != nil {
			return err
		}
		fmt.Println(string(data))
	default:
		fmt.Printf("Multiple daemons found matching '%s'. You need to be more specific. Matches found are:\n", c.pattern)
		for _, agent := range agents {
			fmt.Println(agent.Name)
		}
	}

	return nil
}
