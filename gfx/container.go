package gfx

type Container struct {
	X, Y  int
	Children  []Paintable
}

func (con * Container) Append(p Paintable) {
	con.Children = append(con.Children, p)
}

func (con * Container) Paint(x, y int) {
	for _, p := range(con.Children) {
		p.Paint(con.X + x, con.Y + y)
	}
}
