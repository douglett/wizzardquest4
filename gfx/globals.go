package gfx
import ray "github.com/gen2brain/raylib-go/raylib"

// useful colors
var ColorCollision = ray.Color{255, 0, 0, 100}
var ColorOffBlack  = ray.Color{16, 8, 32, 255}
var ColorWhite     = ray.Color{255, 255, 255, 255}
var ColorBlack     = ray.Color{0, 0, 0, 255}

// base drawable type
type Paintable struct {
	X, Y, Z  int
	Name     string
}

// game state
var Screen = ScreenType{
	width: 160, height: 160, zoom: 4, tsize: 16,
	winname: "WizzardQuest4",
	bgcolor: ColorBlack,
}
