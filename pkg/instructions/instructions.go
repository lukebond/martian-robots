package instructions

import (
  "bufio"
  "errors"
  "os"
  "strconv"
  "strings"

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
