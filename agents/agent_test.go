package agents

import (
  "testing"
)

func TestDirectories(t *testing.T) {
  dirs, err := directories()
  if err != nil {
    t.Fatal(err)
  }

  if len(dirs) != 2 {
    t.Fatal("should returns exact two directories")
  }
}

func TestBasename(t *testing.T) {
  files := []string{"filename.ext", "dir/filename.ext", "filename", "dir/filename"}

  for _, f := range files {
    base := basename(f)
    if base != "filename" {
      t.Fatalf("basename returned the wrong value: '%s' for '%s'", base, f)
    }
  }
}

func TestFind(t *testing.T) {
  agents, err := Find("")
  if err != nil {
    t.Fatal(err)
  }

  if len(agents) == 0 {
    t.Fatal("should return more than one agent")
  }
}
