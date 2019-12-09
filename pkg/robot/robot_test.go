package robot_test

import (
	"testing"

	"github.com/lukebond/martian-robots/pkg/robot"
)

func TestRotateLeftE(t *testing.T) {
	r := robot.Robot{X: 0, Y: 0, Orientation: "E"}
	r.RotateLeft()
	if r.Orientation != "N" {
		t.Errorf("On rotate left from East, orientation should be North")
	}
}

func TestRotateLeftN(t *testing.T) {
	r := robot.Robot{X: 0, Y: 0, Orientation: "N"}
	r.RotateLeft()
	if r.Orientation != "W" {
		t.Errorf("On rotate left from North, orientation should be West")
	}
}

func TestRotateLeftW(t *testing.T) {
	r := robot.Robot{X: 0, Y: 0, Orientation: "W"}
	r.RotateLeft()
	if r.Orientation != "S" {
		t.Errorf("On rotate left from West, orientation should be South")
	}
}

func TestRotateLeftS(t *testing.T) {
	r := robot.Robot{X: 0, Y: 0, Orientation: "S"}
	r.RotateLeft()
	if r.Orientation != "E" {
		t.Errorf("On rotate left from South, orientation should be East")
	}
}

func TestRotateRightE(t *testing.T) {
	r := robot.Robot{X: 0, Y: 0, Orientation: "E"}
	r.RotateRight()
	if r.Orientation != "S" {
		t.Errorf("On rotate right from East, orientation should be South")
	}
}

func TestRotateRightN(t *testing.T) {
	r := robot.Robot{X: 0, Y: 0, Orientation: "N"}
	r.RotateRight()
	if r.Orientation != "E" {
		t.Errorf("On rotate right from North, orientation should be East")
	}
}

func TestRotateRightW(t *testing.T) {
	r := robot.Robot{X: 0, Y: 0, Orientation: "W"}
	r.RotateRight()
	if r.Orientation != "N" {
		t.Errorf("On rotate right from West, orientation should be North")
	}
}

func TestRotateRightS(t *testing.T) {
	r := robot.Robot{X: 0, Y: 0, Orientation: "S"}
	r.RotateRight()
	if r.Orientation != "W" {
		t.Errorf("On rotate right from South, orientation should be West")
	}
}

func TestForwardE(t *testing.T) {
	r := robot.Robot{X: 1, Y: 1, Orientation: "E"}
	r.Forward()
	if r.X != 2 {
		t.Errorf("On move forward from orientation East, X should increment by 1")
	}
}

func TestForwardN(t *testing.T) {
	r := robot.Robot{X: 1, Y: 1, Orientation: "N"}
	r.Forward()
	if r.Y != 2 {
		t.Errorf("On move forward from orientation North, Y should increment by 1")
	}
}

func TestForwardW(t *testing.T) {
	r := robot.Robot{X: 1, Y: 1, Orientation: "W"}
	r.Forward()
	if r.X != 0 {
		t.Errorf("On move forward from orientation West, X should decrement by 1")
	}
}

func TestForwardS(t *testing.T) {
	r := robot.Robot{X: 1, Y: 1, Orientation: "S"}
	r.Forward()
	if r.Y != 0 {
		t.Errorf("On move forward from orientation South, Y should decrement by 1")
	}
}
