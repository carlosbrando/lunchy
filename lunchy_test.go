package main

import (
  "testing"
)

func TestDirs(t *testing.T) {
  dirs := dirs()

  if len(dirs) != 2 {
    t.Error("erro")
  }
}
