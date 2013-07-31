package main

import (
  "os/user"
  "path/filepath"
  "strings"
)

func basename(f string) string {
  basename := strings.Split(filepath.Base(f), ".")
  return strings.Join(basename[:len(basename)-1], ".")
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

func plists(pattern string) []string {
  var list []string

  for _, dirname := range dirs() {
    files, err := filepath.Glob(dirname + "/*" + pattern + "*.plist")
    checkError(err)

    list = append(list, files...)
  }

  return list
}
