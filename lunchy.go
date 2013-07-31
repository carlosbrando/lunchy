package main

import "os"

var options Options

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
