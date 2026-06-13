package gfx
import ray "github.com/gen2brain/raylib-go/raylib"

// useful colors
var ColorCollision = ray.Color{255, 0, 0, 100}
var ColorOffBlack  = ray.Color{16, 8, 32, 255}
var ColorWhite     = ray.Color{255, 255, 255, 255}
var ColorBlack     = ray.Color{0, 0, 0, 255}

type Paintable interface {
	Paint  (x, y int)
	Zindex () int
}

// game state
var Screen = ScreenType{
	Width: 160, Height: 160, Zoom: 4, Tsize: 16,
	Winname: "WizzardQuest4",
}
