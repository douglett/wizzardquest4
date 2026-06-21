package main
import ray "github.com/gen2brain/raylib-go/raylib"
import "wizzardquest4/gfx"

type MapColors struct {
	x, y     int
	tw, th   int
	tsize    int
	colors   []ray.Color
	data     []int
}

func (mc *MapColors) Zindex() int {
	return 0
}

func (mc *MapColors) Paint(posx, posy int) {
	for y := range(mc.th) {
		for x := range(mc.tw) {
			c := mc.colors[mc.data[y * mc.tw + x]]
			gfx.Rect(posx + x*mc.tsize, posy + y*mc.tsize, mc.tsize, mc.tsize, c)
		}
	}
}
