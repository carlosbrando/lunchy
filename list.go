package main

import "fmt"

func (c *Command) list() {
	for _, filename := range plists(c.pattern) {
		fmt.Println(basename(filename))
	}
}
