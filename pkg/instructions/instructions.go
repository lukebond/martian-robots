package instructions

import (
  "bufio"
  "errors"
  "fmt"
  "os"
  "strconv"
  "strings"

  "github.com/lukebond/martian-robots/pkg/grid"
  "github.com/lukebond/martian-robots/pkg/robot"
)

type Instructions struct {
  MaxX        int
  MaxY        int
  Robots      []robot.Robot
  Sequences   []string
}

func LoadInstructionsFromFile(filename string) (*Instructions, error) {
  file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

  i := Instructions{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

  // grid dimensions
  if !scanner.Scan() {
    return nil, errors.New("Failed to read grid dimensions from file")
  }
  fields := strings.Fields(scanner.Text())
  i.MaxX, err = strconv.Atoi(fields[0])
  if err != nil {
    return nil, errors.New("int-to-string conversion error reading grid dimensions")
  }
  i.MaxY, err = strconv.Atoi(fields[1])
  if err != nil {
    return nil, errors.New("int-to-string conversion error reading grid dimensions")
  }

  // robots
	for {
    // position & orientation
    if !scanner.Scan() {
      break
    }
    fields = strings.Fields(scanner.Text())
    r := robot.Robot{}
    r.X, err = strconv.Atoi(fields[0])
    if err != nil {
      return nil, errors.New("int-to-string conversion error reading robot position")
    }
    r.Y, err = strconv.Atoi(fields[1])
    if err != nil {
      return nil, errors.New("int-to-string conversion error reading robot position")
    }
    r.Orientation = fields[2]
    i.Robots = append(i.Robots, r)

    // sequence of instructions
    if !scanner.Scan() {
      return nil, errors.New("Failed to sequence of robot instructions from file")
    }
    i.Sequences = append(i.Sequences, scanner.Text())
	}

	file.Close()
  return &i, nil
}

func (i *Instructions) ProcessInstructions() {
  g := grid.NewGrid(i.MaxX + 1, i.MaxY + 1)
  for r := 0; r < len(i.Robots); r++ {
    ProcessSequence(&i.Robots[r], i.Sequences[r], g)
  }
}

func ProcessSequence(r *robot.Robot, seq string, g *grid.Grid) string {
  var lost = false
  for s := 0; s < len(seq); s++ {
    if lost {
      break
    }
    instruction := seq[s]
    switch instruction {
    case 'L':
      r.RotateLeft()
    case 'R':
      r.RotateRight()
    case 'F':
      lastX := r.X
      lastY := r.Y
      r.Forward()
      if r.X < 0 || r.X >= g.Width || r.Y < 0 || r.Y >= g.Height {
        fmt.Printf("Robot lost at %d,%d. Last position: %d,%d\n", r.X, r.Y, lastX, lastY)
        g.SetScent(lastX, lastY)
        r.X = lastX
        r.Y = lastY
        lost = true
      }
    default:
      fmt.Printf("Ignoring instruction '%v'", instruction)
    }
  }
  var output = fmt.Sprintf("%d %d %v", r.X, r.Y, r.Orientation)
  if lost {
    output = output + " LOST"
  }
  return output
}
