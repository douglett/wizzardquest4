package gfx
import ray "github.com/gen2brain/raylib-go/raylib"
import "fmt"
import "os"
import "encoding/xml"
import "strings"
import "strconv"

type TiledMap struct {
	Xml       XMLMap
	Tw, Th    int
	Tsize     int
	Debug     bool
	Tileset   ray.Texture2D
}

type XMLMap struct {
	XMLName xml.Name       `xml:"map"`
	Width   int            `xml:"width,attr"`
	Height  int            `xml:"height,attr"`
	Layer   []XMLMapLayer  `xml:"layer"`
}

type XMLMapLayer struct {
	XMLName xml.Name `xml:"layer"`
	Name    string   `xml:"name,attr"`
	Data    string   `xml:"data"`
	IData   []int
}

func (tm *TiledMap) Load(fname string) error {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println("error", err)
		return err
	}
	defer file.Close()

	fmt.Println("decoding map:", fname)

	// get data from xml
	decoder := xml.NewDecoder(file)
	decoder.Decode(&tm.Xml)
	xml := &tm.Xml
	tm.Tw = xml.Width
	tm.Th = xml.Height

	// decode layer data
	for k := range xml.Layer {
		layer := &xml.Layer[k]
		data := strings.Split(layer.Data, ",")
		idata := []int{}
		for _, v := range data {
			i, _ := strconv.Atoi(strings.TrimSpace(v))
			idata = append(idata, i)
		}
		layer.IData = idata
		layer.Data = ""
	}
	
	return nil
}

func (tm *TiledMap) Zindex() int {
	return 0
}

func (tm *TiledMap) Paint(posx, posy int) {
	xml := &tm.Xml
	layer := &xml.Layer[0]
	coll := &xml.Layer[1]
	for y := range xml.Height {
		for x := range xml.Width {
			// show game tile
			tile := layer.IData[y * xml.Width + x]
			if tile > 0 {
				Blitt(tm.Tileset, tile - 1, (x * tm.Tsize) + posx, (y * tm.Tsize) + posy)
			}
			// show collision layer (optional)
			c := coll.IData[y * xml.Width + x]
			if tm.Debug && c > 0 {
				Rect((x * tm.Tsize) + posx, (y * tm.Tsize) + posy, tm.Tsize, tm.Tsize, ColorCollision)
			}
		}
	}
}

func (tm *TiledMap) Tile(x, y int) (int, bool) {
	xml := tm.Xml
	if x < 0 || y < 0 || x >= xml.Width || y >= xml.Height { return 0, true }
	t := xml.Layer[0].IData[y * xml.Width + x]
	c := xml.Layer[1].IData[y * xml.Width + x] > 0 // TODO: identify tile layer?
	return t, c
}

func (tm *TiledMap) Settile(x, y, tile int, col bool) {
	xml := tm.Xml
	if x < 0 || y < 0 || x >= xml.Width || y >= xml.Height { return }
	xml.Layer[0].IData[y * xml.Width + x] = tile
	lv := 0
	if col { lv = 1 } // go continues to annoy me...
	xml.Layer[1].IData[y * xml.Width + x] = lv // TODO: identify tile layer?
}
