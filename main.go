package main
import ray "github.com/gen2brain/raylib-go/raylib"
import "wizzardquest4/gfx"
import "fmt"
import "slices"

// globals
var scene  = gfx.Container{}
var tmap   = gfx.TiledMap{ Tsize: 16, Debug: false }
var player = gfx.Sprite{ Id: "player", Tsize: 16, Tile: 2 }
var mobs   = []*gfx.Sprite{}
var moves  = 0

// start
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
	// scene
	gfx.Screen.Bgcolor = gfx.ColorOffBlack
	// map
	tmap.Tileset = ray.LoadTexture("assets/monotiles.png")
	tmap.Load("assets/level1.tmx")
	scene.Append(&tmap)
	// player sprite
	player.X, player.Y = 16, 16
	player.Tileset = ray.LoadTexture("assets/sprites.png")
	scene.Append(&player)
	// mob sprites
	tsprites := ray.LoadTexture("assets/sprites.png")
	for i := range(3) {
		m := gfx.Sprite{ Id: fmt.Sprintf("rslime%02d", i), Tsize: 16, X: 3*16, Y: (1+i*2)*16, Tile: 2 }
		m.Tileset = tsprites
		mobs = append(mobs, &m)
		scene.Append(&m)
	}

	// center screen
	scene.X = (gfx.Screen.Width - (tmap.Tw * tmap.Tsize)) / 2
	scene.Y = (gfx.Screen.Height - (tmap.Th * tmap.Tsize)) / 2
	fmt.Println(scene.X, scene.Y)

	for !gfx.ShouldQuit() {
		cmoves := moves
		// select direction
		switch {
			case ray.IsKeyDown(ray.KeyUp):
				movePlayer(0)
			case ray.IsKeyDown(ray.KeyRight):
				movePlayer(1)
			case ray.IsKeyDown(ray.KeyDown):
				movePlayer(2)
			case ray.IsKeyDown(ray.KeyLeft):
				movePlayer(3)
		}
		// after player move, move enemies
		if moves > cmoves {
			fmt.Println("moved")
			checkDoor()
		}
		// paint
		paintall()
	}
}

func paintall() {
	scene.Paint(0, 0)
	gfx.Text(fmt.Sprintf("moves: %d", moves), 1, 1, ray.Blue)
	gfx.Flip()
}

func dirToTiles(dir int) (int, int) {
	switch dir {
		case 0:  return  0, -1
		case 1:  return  1,  0
		case 2:  return  0,  1
		case 3:  return -1,  0
	}
	panic(fmt.Sprintf("unknown dir: %d", dir))
}

func movePlayer(dir int) {
	player.Tile = dir
	// check for collision
	t := tmap.Tsize
	tx, ty := player.X/t, player.Y/t
	dx, dy := dirToTiles(dir)
	_, c1 := tmap.Tile(tx + dx,   ty + dy)
	_, c2 := tmap.Tile(tx + dx*2, ty + dy*2)
	if c1 || c2 { return }
	// attack enemies
	collideEnemy(dir)
	// move player
	for i := 1; i <= t; i++ {
		player.Y = ty*t + i*dy*2
		player.X = tx*t + i*dx*2
		if i < t { paintall() }
	}
	moves++
}

func collideEnemy(dir int) *gfx.Sprite {
	t := tmap.Tsize
	tx, ty := player.X/16, player.Y/16
	dx, dy := dirToTiles(dir)
	for i, m := range(mobs) {
		if tx + dx*2 == m.X/t && ty + dy*2 == m.Y/t {
			fmt.Println("collide", m)
			mobs = slices.Delete(mobs, i, i+1)
			scene.Remove(m)
			return m
		}
	}
	return nil
}

func checkDoor() {
	if len(mobs) > 0 { return }
	tl := &tmap.Xml.Layer[0]
	cl := &tmap.Xml.Layer[1]
	for i, t := range(tl.IData) {
		if t == 16 {
			tl.IData[i] = 0
			cl.IData[i] = 0
		}
	}
}
