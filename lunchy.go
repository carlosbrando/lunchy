package main

import (
  "fmt"
  "log"
  "lunchy/cmd"
  "os"
  "os/user"
  "path/filepath"
  "strings"
)

var pattern string

var longFlag bool

func checkError(err error) {
  if err != nil {
    log.Fatalf("Uh oh, I didn't expect this:\n%s\n", err)
  }
}

func execute(command string) {
  switch command {
  case "ls":
    ls()
  default:
    showBanner()
  }
}

func dirs() []string {
  usr, err := user.Current()
  checkError(err)

  return []string{"/Library/LaunchAgents", usr.HomeDir + "/Library/LaunchAgents"}

  // TODO: add root option
  // if root {
  //   result = append(result, "/Library/LaunchDaemons", "/System/Library/LaunchDaemons")
  // }
}

func plists() []string {
  var list []string

  for _, dirname := range dirs() {
    files, err := filepath.Glob(dirname + "/*" + pattern + "*.plist")
    checkError(err)

    list = append(list, files...)
  }

  return list
}

func basename(f string) string {
  basename := strings.Split(filepath.Base(f), ".")
  return strings.Join(basename[:len(basename)-1], ".")
}

func ls() {
  for _, filename := range plists() {
    fmt.Println(basename(filename))
  }
}

func init() {

}

func main() {
  args := os.Args[1:]
  if len(args) < 1 {
    showBanner()
    return
  }

  command := args[0]
  if len(args) > 1 {
    pattern = args[1]
  }

  execute(command)

  cmd.List()
}
