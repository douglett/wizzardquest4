package gfx
import ray "github.com/gen2brain/raylib-go/raylib"

type Shape struct {
	Paintable
	Color  ray.Color
	W, H   int
}

func (sh *Shape) Paint() {
	Rect(sh.X, sh.Y, sh.W, sh.H, sh.Color)
}
