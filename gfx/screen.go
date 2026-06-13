package gfx
import ray "github.com/gen2brain/raylib-go/raylib"
import "fmt"
import "math"

type ScreenType struct {
	Width, Height, Zoom, Tsize int 
	Winname     string
	Bgcolor     ray.Color
	Offx, Offy  int
	camera      ray.Camera2D
	// tileset     ray.Texture2D
	// sound       ray.Sound
}

// use screen singleton from globals
var scr = &Screen

func Create() error {
	// defaults
	if scr.Width <= 0     { scr.Width = 640 }
	if scr.Height <= 0    { scr.Height = 480 }
	if scr.Zoom <= 0      { scr.Zoom = 1 }
	if scr.Tsize <= 0     { scr.Tsize = 16 }
	if scr.Bgcolor.A == 0 { scr.Bgcolor = ray.Black }
	if scr.Winname == ""  { scr.Winname = "Screen" }
	scr.camera.Zoom = float32(scr.Zoom)
	// init raylib
	// ray.SetTraceLogLevel(ray.LogInfo)
	ray.SetTraceLogLevel(ray.LogWarning)
	ray.InitWindow(int32(scr.Width * scr.Zoom), int32(scr.Height * scr.Zoom), scr.Winname)
	ray.InitAudioDevice()
	ray.SetTargetFPS(60)
	// ok
	fmt.Println("Screen initialized:", scr.Width, scr.Height)
	Begin()
	return nil
}

func Destroy() {
	ray.CloseAudioDevice()
	ray.CloseWindow()
	fmt.Println("Screen destroyed")
}

func ShouldQuit() bool {
	return ray.WindowShouldClose()
}

func Begin() {
	ray.BeginDrawing()
	ray.BeginMode2D(scr.camera)
	ray.ClearBackground(scr.Bgcolor)
}

func Flip() {
	// show framerate
	fps := fmt.Sprintf("%d", ray.GetFPS())
	fontw := int32(10)
	txtw := ray.MeasureText(fps, fontw)
	ray.DrawText(fps, int32(scr.Width) - (txtw + 2), 1, fontw, ray.Green)
	// flip
	ray.EndMode2D()
	ray.EndDrawing()
	// begin drawing mode for next frame
	Begin()
}

// blit sub-pixel positions
func Blitf(tex ray.Texture2D, x, y float64) {
	Blit(tex, int(math.Round(x)), int(math.Round(y)))
}
func Blittf(tex ray.Texture2D, tile int, x, y float64) {
	Blitt(tex, tile, int(math.Round(x)), int(math.Round(y)))
}

// texture blitting
func Blit(tex ray.Texture2D, x, y int) {
	ray.DrawTexture(tex, int32(x + scr.Offx), int32(y + scr.Offy), ray.White)
}

// blit texture as tileset
func Blitt(tex ray.Texture2D, tile, x, y int) {
	tx := tile % (int(tex.Width) / scr.Tsize)
	ty := tile / (int(tex.Width) / scr.Tsize)
	src := ray.Rectangle{
		float32(tx * scr.Tsize), float32(ty * scr.Tsize),
		float32(scr.Tsize), float32(scr.Tsize),
	}
	dst := ray.Vector2{ float32(x + scr.Offx), float32(y + scr.Offy) }
	ray.DrawTextureRec(tex, src, dst, ray.White)
}

func Rect(x, y, w, h int, color ray.Color) {
	ray.DrawRectangle(int32(x + scr.Offx), int32(y + scr.Offy), int32(w), int32(h), color)
}

func Text(s string, x, y int) {
	ray.DrawText(s, int32(x + scr.Offx), int32(y + scr.Offy), 10, ray.White)
}
