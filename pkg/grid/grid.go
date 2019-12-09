package grid

type Grid struct {
  Width   int
  Height  int
  scents  [][]bool
}

func NewGrid(width int, height int) *Grid {
  g := Grid{width, height, make([][]bool, height)}
  for i := 0; i < height; i++ {
		g.scents[i] = make([]bool, width)
	}
  return &g
}

func (g *Grid) CheckScent(x int, y int) bool {
  return g.scents[y][x]
}

func (g *Grid) SetScent(x int, y int) {
  g.scents[y][x] = true
}
