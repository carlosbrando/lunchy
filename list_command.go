package main

import (
  "fmt"
)

type ListCommand struct {
  command *Command
}

func (l *ListCommand) execute() {
  for _, filename := range plists(l.command.pattern) {
    fmt.Println(basename(filename))
  }
}
