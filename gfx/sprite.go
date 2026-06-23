package gfx
import ray "github.com/gen2brain/raylib-go/raylib"

type Sprite struct {
	X, Y, Z   int
	Tile      int
	Tsize     int
	Tileset   ray.Texture2D
}

func (spr *Sprite) Zindex() int {
	return spr.Z
}

func (spr *Sprite) Paint(x, y int) {
	Blitt(spr.Tileset, spr.Tile, spr.X + x, spr.Y + y)
}
