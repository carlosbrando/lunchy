// Copyright 2013 Carlos Brando. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package agents

import (
	"fmt"
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

// FindOne returns only one agent when found.
// It returns nil if none were found.
// FindOne will print a list returns nil if multiple agents were found.
func FindOne(pattern string) (agent *Agent, err error) {
	agents, err := Find(pattern)
	if err != nil {
		return nil, err
	}

	switch len(agents) {
	case 0:
		fmt.Printf("No daemon was found matching '%s'.\n", pattern)
		return nil, nil
	case 1:
		return &agents[0], nil
	default:
		fmt.Printf("Multiple daemons found matching '%s'. You need to be more specific. Matches found are:\n", pattern)
		for _, agent := range agents {
			fmt.Println(agent.Name)
		}
		return nil, nil
	}
}
