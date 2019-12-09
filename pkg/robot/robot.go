package robot

import (
  "errors"
)

type Robot struct {
	X           int
	Y           int
	Orientation string
}

func (r *Robot) RotateLeft() error {
	switch r.Orientation {
	case "N":
    r.Orientation = "W"
	case "S":
    r.Orientation = "E"
	case "E":
    r.Orientation = "N"
	case "W":
    r.Orientation = "S"
	default:
    return errors.New("Invalid orientation: " + r.Orientation)
	}
  return nil
}

func (r *Robot) RotateRight() error {
	switch r.Orientation {
	case "N":
    r.Orientation = "E"
	case "S":
    r.Orientation = "W"
	case "E":
    r.Orientation = "S"
	case "W":
    r.Orientation = "N"
	default:
    return errors.New("Invalid orientation: " + r.Orientation)
  }
  return nil
}

func (r *Robot) Forward() error {
	switch r.Orientation {
	case "N":
    r.Y++
	case "S":
    r.Y--
	case "E":
    r.X++
	case "W":
    r.X--
	default:
    return errors.New("Invalid orientation: " + r.Orientation)
  }
  return nil
}
