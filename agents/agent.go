// Copyright 2013 Carlos Brando. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package agents

import (
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
)

type Agent struct {
	Name     string
	Fullpath string
}

// basename returns only the basename of a fullpath filename.
func basename(f string) string {
	basename := strings.Split(filepath.Base(f), ".")

	if len(basename) > 1 {
		return strings.Join(basename[:len(basename)-1], ".")
	} else {
		return basename[0]
	}
}

// root returns true if program was called by a root user.
func root() bool {
	return syscall.Geteuid() == 0
}

// directories returns a list of all directories that contain agents.
func directories() ([]string, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	result := []string{"/Library/LaunchAgents", usr.HomeDir + "/Library/LaunchAgents"}

	if root() {
		result = append(result, "/Library/LaunchDaemons", "/System/Library/LaunchDaemons")
	}

	return result, nil
}

// Find returns only the agents that match with the pattern.
// If no pattern was assigned, returns all agents.
func Find(pattern string) ([]Agent, error) {
	var list []Agent

	dirs, err := directories()
	if err != nil {
		return nil, err
	}

	for _, dirname := range dirs {
		files, err := filepath.Glob(dirname + "/*.plist")
		if err != nil {
			return nil, err
		}

		for _, file := range files {
			matched, err := regexp.MatchString(strings.ToLower(pattern), strings.ToLower(file))
			if err != nil {
				return nil, err
			}

			if matched {
				agent := &Agent{Fullpath: file, Name: basename(file)}
				list = append(list, *agent)
			}
		}
	}

	return list, nil
}
