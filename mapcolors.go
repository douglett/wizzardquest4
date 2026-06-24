package main
import ray "github.com/gen2brain/raylib-go/raylib"
import "wizzardquest4/gfx"

type MapColors struct {
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

// gmap := MapColors{
// 	tsize: 16,
// 	colors: []ray.Color{
// 		ray.Black,
// 		ray.Red,
// 		ray.Blue,
// 	},
// 	tw: 5,
// 	th: 5,
// 	data: []int {
// 		1, 1, 1, 1, 1,
// 		1, 0, 0, 0, 1,
// 		1, 2, 0, 0, 1,
// 		1, 0, 0, 0, 1,
// 		1, 1, 1, 1, 1,
// 	},
// }
