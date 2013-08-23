package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/carlosbrando/lunchy/agents"
)

func (c *Command) edit() error {
	if c.pattern == "" {
		return errors.New("edit [name]")
	}

	agent, err := agents.FindOne(c.pattern)
	if err != nil {
		return err
	}

	if agent != nil {
		editor := os.Getenv("EDITOR")
		if editor == "" {
			return errors.New("EDITOR environment variable is not set")
		}

		cmd := exec.Command("sh", "-c", fmt.Sprintf("%s %s", editor, agent.Fullpath))
		err = cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
