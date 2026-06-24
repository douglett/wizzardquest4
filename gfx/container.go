package gfx
import "slices"

type Container struct {
	X, Y, Z   int
	Children  []Paintable
}

func (con *Container) Append(p Paintable) {
	con.Children = append(con.Children, p)
}

func (con *Container) Remove(p Paintable) {
	for i, c := range(con.Children) {
		if c == p {	
			con.Children = slices.Delete(con.Children, i, i+1)
		}
	}
}

func (con *Container) Zindex() int {
	return con.Z
}

func (con *Container) Paint(x, y int) {
	for _, p := range(con.Children) {
		p.Paint(con.X + x, con.Y + y)
	}
}
