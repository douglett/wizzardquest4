package main
import ray "github.com/gen2brain/raylib-go/raylib"
import "wizzardquest4/gfx"
import "fmt"

// var tmap TiledMap

func main() {
	fmt.Println("starting WizzardQuest4!")
	gfx.Create()
	defer gfx.Destroy()

	// test1()
	// test2()
	test3()
}

// tiled map test
func test3() {
	gfx.Screen.Bgcolor = gfx.ColorOffBlack
	scene := gfx.Container{}
	tmap := gfx.TiledMap{ Tsize: 16, Debug: false }
	tmap.Tileset = ray.LoadTexture("assets/monotiles.png")
	tmap.Load("assets/level1.tmx")
	scene.Append(&tmap)

	scene.X = (gfx.Screen.Width - (tmap.Tw * tmap.Tsize)) / 2
	scene.Y = (gfx.Screen.Height - (tmap.Th * tmap.Tsize)) / 2
	fmt.Println(scene.X, scene.Y)

	for !gfx.ShouldQuit() {
		scene.Paint(0, 0)
		gfx.Text("tiledmap test", 1, 1, ray.White)
		gfx.Flip()
	}
}

// color map test
func test2() {
	scene := gfx.Container{}

	gmap := MapColors{
		tsize: 16,
		colors: []ray.Color{
			ray.Black,
			ray.Red,
			ray.Blue,
		},
		tw: 5,
		th: 5,
		data: []int {
			1, 1, 1, 1, 1,
			1, 0, 0, 0, 1,
			1, 2, 0, 0, 1,
			1, 0, 0, 0, 1,
			1, 1, 1, 1, 1,
		},
	}

	scene.Append(&gmap)

	for !gfx.ShouldQuit() {
		scene.Paint(0, 0)
		gfx.Text("mapcolors test", 10, 10, ray.White)
		gfx.Flip()
	}
}

// basic drawing test
func test1() {
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

	x := 0

	for !gfx.ShouldQuit() {
		scene.Paint(0, 0)

		gfx.Text("hello worldq", 10, 10, ray.White)
		gfx.Text("hello worldq", 10, 18, ray.Pink)
		gfx.Text13("hello worldq", 10, 30, ray.White)
		gfx.Text13("hello worldq", 10, 42, ray.Pink)

		x = (x + 1) % (60 * 10)
		gfx.Text(fmt.Sprintf("%d", x/10), 10, 60, ray.Blue)

		gfx.Flip()
		// scene.X++
		// gfx.Screen.Offx++
	}
}
