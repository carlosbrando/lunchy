package main

import (
  "log"
  "os"
)

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

func checkError(err error) {
  if err != nil {
    log.Fatalf("Uh oh, I didn't expect this:\n%s\n", err)
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
