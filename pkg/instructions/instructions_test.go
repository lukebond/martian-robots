package instructions_test

import (
  "fmt"
  "testing"

  "github.com/lukebond/martian-robots/pkg/grid"
  "github.com/lukebond/martian-robots/pkg/instructions"
  "github.com/lukebond/martian-robots/pkg/robot"
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

func TestProcessSequence1(t *testing.T) {
  g := grid.NewGrid(6, 4)
  r := robot.Robot{X: 1, Y: 1, Orientation: "E"}
  output := instructions.ProcessSequence(&r, "RFRFRFRF", g)
  if output != "1 1 E" {
    t.Errorf("Test sequence output '%v' didn't match sample data", output)
  }
}

func TestProcessSequence2(t *testing.T) {
  g := grid.NewGrid(6, 4)
  r := robot.Robot{X: 3, Y: 2, Orientation: "N"}
  output := instructions.ProcessSequence(&r, "FRRFLLFFRRFLL", g)
  if output != "3 3 N LOST" {
    t.Errorf("Test sequence output '%v' didn't match sample data", output)
  }
}

func TestProcessSequence3(t *testing.T) {
  g := grid.NewGrid(6, 4)
  r := robot.Robot{X: 0, Y: 3, Orientation: "W"}
  output := instructions.ProcessSequence(&r, "LLFFFLFLFL", g)
  if output != "2 3 S" {
    t.Errorf("Test sequence output '%v' didn't match sample data", output)
  }
}
