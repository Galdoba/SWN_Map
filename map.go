package main

import (
	"strconv"

	"github.com/Galdoba/utils"
)

type grid struct {
	tileMap  map[int]*Tile
	minX     int
	minY     int
	maxX     int
	maxY     int
	TileSize int
}

//NewGrid - создает сеть
func NewGrid(minX, minY, maxX, maxY int) *grid {
	gr := grid{}
	gr.tileMap = make(map[int]*Tile)
	gr.minX = minX
	gr.minY = minY
	gr.maxX = maxX
	gr.maxY = maxY
	//gr.tileMap[0] = newTileHex(minX, minY)
	gr.TileSize = /*len(gr.tileMap[0].lines)*/ 16
	for y := gr.minY; y <= gr.maxY; y++ {
		for x := gr.minX; x <= gr.maxX; x++ {

			tl := newTileHex(x, y)
			id := hexToID(tl.hex)
			gr.tileMap[id] = tl
		}
	}

	//gr.TileSize = len(gr.tileMap[0].lines)
	return &gr
}

func getScreenCoords(t Tile) hexCoords {
	return t.hex
}

func drawLine(s string) string {
	return s
}

func (gr *grid) TileByXY(x, y int) *Tile {
	for _, val := range gr.tileMap {
		if val.hex.col == x && val.hex.row == y {
			return val
		}
	}
	return nil
}

func drawGrid(gr grid) string {
	gridStr := ""
	// var idList []int
	// for y := gr.minY; y <= gr.maxY; y++ {
	// 	for x := gr.minX; x <= gr.maxX; x++ {
	// 		id := hexToID(hexCoords{x, y})
	// 		idList = append(idList, id)
	// 		gridStr += "| " + strconv.Itoa(id) + "\n"
	// 	}
	// }
	tSegments := totalSegments(gr)
	for i := 0; i < tSegments; i++ {
		gridStr += defineSegment(i, gr)
	}
	return gridStr

}

func totalSegments(gr grid) int {
	totalRows := gr.maxY - gr.minY + 1
	totalCols := gr.maxX - gr.minX + 1
	Tilelines := 8
	totalSegments := (totalRows * totalCols * Tilelines) + (totalCols * Tilelines / 2)
	return totalSegments
}

func absInt(i int) int {
	if i > 0 {
		return i
	}
	if i < 0 {
		return i * -1
	}
	return 0
}

func offsetNeeded(col int) bool {
	if absInt(col)%2 > 0 {
		return true
	}
	return false
}

func defineSegment(segment int, gr grid) string {
	totalCols := gr.maxX - gr.minX + 1
	col := (segment) % totalCols
	//offset := false
	row := segment / totalCols / 8
	line := segment / totalCols % 8
	gridX := gr.minX + col
	str := "                "
	color := "white"
	// if absInt(gridX)%2 > 0 {
	// 	offset = true
	// }
	// if offset {
	// 	line = line - 4
	// 	if line < 0 {
	// 		line = line + 8
	// 		row--
	// 	}
	// }
	if offsetNeeded(gridX) {
		line = line - 4
		if line < 0 {
			line = line + 8
			row--
		}
	}
	gridY := gr.minY + row /*/8*/
	id := hexToID(hexCoords{gridX, gridY})
	if val, ok := gr.tileMap[id]; ok {
		color = zoneColor(val.LayerZone)
		str = val.lines[line]
		// if line == 2 {
		// 	str = "|" + layerInfoL(val.LayerZone) + "|"

		// }
		if line == 1 {
			str = "|" + utils.AsciiColor("white", hexCoordsStr(val.hex)) + utils.AsciiColor(color, "|")

		}

		if line == 3 {
			str = "|              |"
			if val.LayerStar != "" {
				str = "|      " + utils.AsciiColor("white", "\x23\x23") + utils.AsciiColor(color, "      |")
			}
		}
		if line == 4 {
			str = "|              |"
			if val.LayerStar != "" {
				str = "|      " + utils.AsciiColor("white", "\x23\x23") + utils.AsciiColor(color, "      |")
			}
		}

		// if line == 4 {
		// 	str = "|" + layerInfoL(val.LayerStar) + "|"

		// }

		// } else {
		// 	str = "                "
	}
	for len(str) < 14 {
		str = str + " "
	}
	str = utils.AsciiColor(color, str)
	if gridX == gr.maxX {
		str = str + " \n"
	}
	return str
}

func zoneColor(zoneLayer string) string {
	switch zoneLayer {
	case "Weird Energy":
		return "green"
	case "Nebula":
		return "cyan"
	case "Void":
		return "blue"
	case "Dust Cloud":
		return "yellow"
	case "Plasma":
		return "magenta"
	}
	return "white"
}

func (gr *grid) tileByClick(mX, mY int) int {
	col := mX / 16
	hexX := gr.minX + col
	if offsetNeeded(hexX) {
		mY = mY - 4
		if mY < 0 {
			return -1
		}
	}
	row := mY / 8
	hexY := gr.minY + row
	hexCrd := setHexCoords(hexX, hexY)
	id := hexToID(hexCrd)
	return id
}

// func NewTile(x, y int) *Tile {
// 	t := Tile{}
// 	t.hex.col = x
// 	t.hex.row = y
// 	t.cube = oddQToCube(t.hex)
// 	t.ID = hexToID(t.hex)
// 	t.lines = newSquare()
// 	return &t
// }

// func NewTileByID(id int) *Tile {
// 	t := Tile{}
// 	t.ID = id
// 	t.hex = hexFromID(id)
// 	t.cube = cubeFromID(id)
// 	t.lines = newSquare()

// 	return &t
// }

//Tile -
type Tile struct {
	hex       hexCoords
	cube      cubeCoords
	ID        int
	LayerHex  string
	LayerZone string
	LayerStar string
	lines     []string
}

func newTileHex(col, row int) *Tile {
	tile := &Tile{}
	tile.hex = setHexCoords(col, row)
	tile.cube = oddQToCube(tile.hex)
	tile.ID = spiralCubeToIDMAP[tile.cube]
	tile.LayerStar = "NO DATA"
	tile.LayerZone = "NO DATA"
	tile.lines = []string{
		"+--------------+",
		"|" + hexCoordsStr(tile.hex) + "|",
		"|              |",
		"|              |",
		"|              |",
		"|              |",
		"|              |",
		"+--------------+",
	}
	//	fmt.Println("Create:" + strconv.Itoa(col) + " " + strconv.Itoa(row))
	return tile
}

func hexToID(hex hexCoords) int {
	cube := oddQToCube(hex)
	return spiralCubeToIDMAP[cube]
}

//Square -
func Square(x, y int) []string {
	xCoord := convertCoord(x)
	yCoord := convertCoord(y)

	sqr := []string{
		"+--------------+",
		"|X" + xCoord + " Y" + yCoord + " Z" + yCoord + "|",
		"|X" + strconv.Itoa(x) + " Y" + strconv.Itoa(y) + "         |",
		"|              |",
		"| " + convertCoord(hexToID(hexCoords{x, y})) + "           |",
		"|              |",
		"|              |",
		"+--------------+",
	}
	return sqr
}

func newSquare() []string {
	sqr := []string{
		"+--------------+",
		"|              |",
		"|              |",
		"|              |",
		"|              |",
		"|              |",
		"|              |",
		"+--------------+",
	}
	return sqr
}

func (tl *Tile) square() []string {
	return tl.lines
}

func convertCoord(i int) string {
	neg := false
	if i < 0 {
		i = -i
		neg = true
	}
	coord := strconv.Itoa(i)
	if i < 10 && i > -1 {
		coord = "0" + coord
	}
	if neg {
		coord = "-" + coord
	} else {
		coord = "+" + coord
	}
	return coord
}
