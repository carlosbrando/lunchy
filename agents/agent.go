package agents

import (
  "os/user"
  "path/filepath"
  "strings"
)

type Agent struct {
  Name     string
  Fullpath string
}

func basename(f string) string {
  basename := strings.Split(filepath.Base(f), ".")

  if len(basename) > 1 {
    return strings.Join(basename[:len(basename)-1], ".")
  } else {
    return basename[0]
  }
}

func directories() ([]string, error) {
  usr, err := user.Current()
  if err != nil {
    return nil, err
  }

  return []string{"/Library/LaunchAgents", usr.HomeDir + "/Library/LaunchAgents"}, nil

  // TODO: add root option
  // if root {
  //   result = append(result, "/Library/LaunchDaemons", "/System/Library/LaunchDaemons")
  // }
}

func Find(pattern string) ([]Agent, error) {
  var list []Agent

  dirs, err := directories()
  if err != nil {
    return nil, err
  }

  for _, dirname := range dirs {
    files, err := filepath.Glob(dirname + "/*" + pattern + "*.plist")
    if err != nil {
      return nil, err
    }

    for _, file := range files {
      agent := &Agent{Fullpath: file, Name: basename(file)}
      list = append(list, *agent)
    }
  }

  return list, nil
}
