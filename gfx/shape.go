package gfx
import ray "github.com/gen2brain/raylib-go/raylib"

type Shape struct {
	X, Y, Z, W, H  int
	Color          ray.Color
}

func (sh *Shape) Zindex() int {
	return sh.Z
}

func (sh *Shape) Paint(x, y int) {
	Rect(sh.X + y, sh.Y + y, sh.W, sh.H, sh.Color)
}
