package instructions_test

import (
  "fmt"
  "testing"

  "github.com/lukebond/martian-robots/pkg/instructions"
)

func TestLoadInstructions(t *testing.T) {
  i, err := instructions.LoadInstructionsFromFile("sample-input.txt")
  fmt.Printf("%v\n", i)
  fmt.Printf("%v\n", err)
  if err != nil {
    t.Errorf("Error while loading instructions")
  }
  if i.MaxX != 5 || i.MaxY != 3 {
    t.Errorf("Loading instructions resulted in wrong grid dimensions")
  }
  if len(i.Robots) != 3 {
    t.Errorf("Loading instructions resulted in wrong number of robots")
  }
  if i.Robots[2].X != 0 || i.Robots[2].Y != 3 || i.Robots[2].Orientation != "W" {
    t.Errorf("Loading instructions resulted in incorrect robot position & orientation")
  }
}
