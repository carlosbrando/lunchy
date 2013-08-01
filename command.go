package main

type Command struct {
  command string
  pattern string
}

func (c *Command) execute() error {
  switch c.command {
  case "list", "ls":
    c.list()
  default:
    showBanner()
  }

  return nil
}
