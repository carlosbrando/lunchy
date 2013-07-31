package main

type Command struct {
  command string
  pattern string
}

func (c *Command) Execute() error {
  switch c.command {
  case "list", "ls":
    cmd := ListCommand{command: c}
    cmd.execute()
  default:
    showBanner()
  }

  return nil
}
