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
	for i := range(3) {
		m := gfx.Sprite{ Id: fmt.Sprintf("rslime%02d", i), Tsize: 16, X: 3*16, Y: (1+i*2)*16, Tile: 2 }
		m.Tileset = ray.LoadTexture("assets/sprites.png")
		mobs = append(mobs, &m)
		scene.Append(&m)
	}

	// center screen
	scene.X = (gfx.Screen.Width - (tmap.Tw * tmap.Tsize)) / 2
	scene.Y = (gfx.Screen.Height - (tmap.Th * tmap.Tsize)) / 2
	fmt.Println(scene.X, scene.Y)

	dir := -1
	walk := 0
	moves := 0

	for !gfx.ShouldQuit() {
		// select direction
		if dir == -1 {
			tx, ty := player.X/16, player.Y/16
			switch {
				case ray.IsKeyDown(ray.KeyUp):
					player.Tile = 0
					_, c1 := tmap.Tile(tx, ty-1)
					_, c2 := tmap.Tile(tx, ty-2)
					collideEnemy(0)
					if !c1 && !c2 { dir = 0; moves++ }
				case ray.IsKeyDown(ray.KeyRight):
					player.Tile = 1
					_, c1 := tmap.Tile(tx+1, ty)
					_, c2 := tmap.Tile(tx+2, ty)
					collideEnemy(1)
					if !c1 && !c2 { dir = 1; moves++ }
				case ray.IsKeyDown(ray.KeyDown):
					player.Tile = 2
					_, c1 := tmap.Tile(tx, ty+1)
					_, c2 := tmap.Tile(tx, ty+2)
					collideEnemy(2)
					if !c1 && !c2 { dir = 2; moves++ }
				case ray.IsKeyDown(ray.KeyLeft):
					player.Tile = 3
					_, c1 := tmap.Tile(tx-1, ty)
					_, c2 := tmap.Tile(tx-2, ty)
					collideEnemy(3)
					if !c1 && !c2 { dir = 3; moves++ }
			}
		}
		// player walk animation
		if dir > -1 {
			switch dir {
				case 0:  player.Y -= 2
				case 1:  player.X += 2
				case 2:  player.Y += 2
				case 3:  player.X -= 2
			}
			walk++
			if walk >= 16 { dir = -1; walk = 0 }
		}

		// paint
		if dir > -1 { player.Tile = dir }
		scene.Paint(0, 0)
		gfx.Text(fmt.Sprintf("moves: %d", moves), 1, 1, ray.Blue)
		gfx.Flip()
	}
}

func collideEnemy(dir int) *gfx.Sprite {
	t := tmap.Tsize
	dx, dy := 0, 0
	switch dir {
		case 0:  dy = -2
		case 1:  dx = 2
		case 2:  dy = 2
		case 3:  dx = -2
	}
	for i, m := range(mobs) {
		if player.X/t + dx == m.X/t && player.Y/t + dy == m.Y/t {
			fmt.Println("collide", m)
			mobs = slices.Delete(mobs, i, i+1)
			scene.Remove(m)
			return m
		}
	}
	return nil
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
