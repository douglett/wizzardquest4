package main
import "wizzardquest4/gfx"
import "fmt"

// var tmap TiledMap

func main() {
	fmt.Println("starting WizzardQuest4!")
	gfx.Screen.Create()
	defer gfx.Screen.Destroy()

	for !gfx.Screen.ShouldQuit() {
		gfx.Screen.Begin()
		gfx.Screen.Flip()
	}
}