package game
import "fmt"
import ray "github.com/gen2brain/raylib-go/raylib"

// useful colors
var ColorCollision = ray.Color{255, 0, 0, 100}
var ColorBlack     = ray.Color{16, 8, 32, 255}
var ColorWhite     = ray.Color{255, 255, 255, 255}

// game state
var screen = Screen{
	width: 160, height: 160, zoom: 4, tsize: 16,
	winname: "WizzardQuest4",
	bgcolor: ColorBlack,
}
var tmap TiledMap

// begin
func Start() {
	fmt.Println("starting WizzardQuest4!")
	screen.create()
	defer screen.destroy()

	for !ray.WindowShouldClose() {
		screen.begin()
		screen.flip()
	}
}
