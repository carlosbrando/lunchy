package main

import (
  "log"
  "os"
)

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

  cmd := Command{command: args[0]}
  if len(args) > 1 {
    cmd.pattern = args[1]
  }

  cmd.Execute()
}
