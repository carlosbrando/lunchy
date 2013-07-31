package main

import "os"

type Options struct {
  command string
  pattern string
}

var options Options

func execute() {
  switch options.command {
  case "list", "ls":
    list()
  default:
    showBanner()
  }
}

func main() {
  args := os.Args[1:]
  if len(args) < 1 {
    showBanner()
    return
  }

  options = Options{command: args[0]}
  if len(args) > 1 {
    options.pattern = args[1]
  }

  execute()
}
