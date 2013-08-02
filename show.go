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

	agents, err := agents.Find(c.pattern)
	if err != nil {
		return err
	}

	if len(agents) > 1 {
		return errors.New("Multiple daemons found matching.")
		// TODO: puts "Multiple daemons found matching '#{name}'. You need to be more specific. Matches found are:\n#{files.keys.join("\n")}"
	} else {
		data, err := ioutil.ReadFile(agents[0].Fullpath)
		if err != nil {
			return err
		}
		fmt.Println(string(data))
	}

	return nil
}
