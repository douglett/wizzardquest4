package main
import "wizzardquest4/gfx"
import "fmt"

// var tmap TiledMap

func main() {
	fmt.Println("starting WizzardQuest4!")
	gfx.Create()
	defer gfx.Destroy()

	r := gfx.Shape{
		Paintable: gfx.Paintable{ X: 10, Y: 10 },
		W: 20, H: 20,
		Color: gfx.ColorWhite,
	}

	for !gfx.ShouldQuit() {
		r.Paint()
		gfx.Flip()
	}
}