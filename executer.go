package main

type Options struct {
  command string
  pattern string
}

func execute() {
  switch options.command {
  case "list", "ls":
    list()
  default:
    showBanner()
  }
}
