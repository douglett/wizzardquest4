package main
import "wizzardquest4/gfx"
import "fmt"

// var tmap TiledMap

func main() {
	fmt.Println("starting WizzardQuest4!")
	gfx.Create()
	defer gfx.Destroy()

	for !gfx.ShouldQuit() {
		gfx.Flip()
	}
}