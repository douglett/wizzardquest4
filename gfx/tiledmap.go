package gfx
import ray "github.com/gen2brain/raylib-go/raylib"
import "fmt"
import "os"
import "encoding/xml"
import "strings"
import "strconv"

type TiledMap struct {
	xml            XMLMap
	tileset        ray.Texture2D
	tsize          int
	showCollision  bool
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
	decoder.Decode(&tm.xml)

	// decode layer data
	for k := range tm.xml.Layer {
		layer := &tm.xml.Layer[k]
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

func (tm *TiledMap) paint() {
	posx, posy := 0, 0
	layer := &tm.xml.Layer[0]
	coll := &tm.xml.Layer[1]
	for y := range tm.xml.Height {
		for x := range tm.xml.Width {
			// show game tile
			tile := layer.IData[y * tm.xml.Width + x]
			if tile > 0 {
				Screen.blitt(tm.tileset, tile - 1, (x * tm.tsize) + posx, (y * tm.tsize) + posy)
			}
			// show collision layer (optional)
			c := coll.IData[y * tm.xml.Width + x]
			if tm.showCollision && c > 0 {
				Screen.rect((x * tm.tsize) + posx, (y * tm.tsize) + posy, tm.tsize, tm.tsize, ColorCollision)
			}
		}
	}
}
