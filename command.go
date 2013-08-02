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
	default:
		showBanner()
	}

	return nil
}
