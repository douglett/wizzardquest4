package main
import "wizzardquest4/gfx"
import "fmt"
import ray "github.com/gen2brain/raylib-go/raylib"

// var tmap TiledMap

func main() {
	fmt.Println("starting WizzardQuest4!")
	gfx.Create()
	defer gfx.Destroy()

	scene := gfx.Container{ X: 0, Y: 0 }

	r := gfx.Shape{
		X: 10, Y: 10,
		W: 20, H: 20,
		Color: gfx.ColorWhite,
	}

	scene.Append(&r)
	r.Color.G = 0
	r.Color.B = 0
	
	scene.Append(&gfx.Shape{
		X: 50, Y: 20,
		W: 20, H: 20,
		Color: ray.Color{0, 0, 255, 255},
	})

	for !gfx.ShouldQuit() {
		scene.Paint(0, 0)
		gfx.Qbtext("hello worldq", 10, 10, ray.White)
		gfx.Qbtext("hello worldq", 10, 18, ray.Pink)
		gfx.Flip()
		// scene.X++
		// gfx.Screen.Offx++
	}
}