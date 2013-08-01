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
    fmt.Println(agent.Name)
  }
  return nil
}
