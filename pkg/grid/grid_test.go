package grid_test

import (
	"testing"

	"github.com/lukebond/martian-robots/pkg/grid"
)

func TestGridCreate(t *testing.T) {
	g := grid.NewGrid(10, 10)
	if g.Width != 10 || g.Height != 10 {
		t.Errorf("Grid initialiser sets the wrong sizes")
	}
}

func TestGridCreateScentInit(t *testing.T) {
	g := grid.NewGrid(10, 10)
	if g.CheckScent(0, 0) {
		t.Errorf("Grid initialiser scent is set on create")
	}
}

func TestGridSetScent(t *testing.T) {
	g := grid.NewGrid(10, 10)
	g.SetScent(5, 5)
	if !g.CheckScent(5, 5) {
		t.Errorf("Setting grid scent failed")
	}
}
