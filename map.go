package main

import (
	"strconv"
)

type grid struct {
	tileMap  map[int]*tile
	minX     int
	minY     int
	maxX     int
	maxY     int
	tileSize int
}

func NewGrid(minX, minY, maxX, maxY int) *grid {
	gr := grid{}
	gr.tileMap = make(map[int]*tile)
	gr.minX = minX
	gr.minY = minY
	gr.maxX = maxX
	gr.maxY = maxY
	//gr.tileMap[0] = newTileHex(minX, minY)
	gr.tileSize = /*len(gr.tileMap[0].lines)*/ 16
	for y := gr.minY; y <= gr.maxY; y++ {
		for x := gr.minX; x <= gr.maxX; x++ {

			tl := newTileHex(x, y)
			id := hexToID(tl.hex)
			gr.tileMap[id] = tl
		}
	}
	//gr.tileSize = len(gr.tileMap[0].lines)
	return &gr
}

func getScreenCoords(t tile) hexCoords {
	return t.hex
}

func drawLine(s string) string {
	return s
}

func (gr *grid) tileByXY(x, y int) *tile {
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
	tilelines := 8
	totalSegments := (totalRows * totalCols * tilelines) + (totalCols * tilelines / 2)
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

func defineSegment(segment int, gr grid) string {
	totalCols := gr.maxX - gr.minX + 1
	col := (segment) % totalCols
	offset := false
	row := segment / totalCols / 8
	line := segment / totalCols % 8
	gridX := gr.minX + col
	str := "                "
	if absInt(gridX)%2 > 0 {
		offset = true
	}
	if offset {
		line = line - 4
		if line < 0 {
			line = line + 8
			row--
		}
	}
	gridY := gr.minY + row /*/8*/
	id := hexToID(hexCoords{gridX, gridY})
	if val, ok := gr.tileMap[id]; ok {
		str = val.lines[line]
		// } else {
		// 	str = "                "
	}
	if gridX == gr.maxX {
		str = str + " \n"
	}
	return str
}

func NewTile(x, y int) *tile {
	t := tile{}
	t.hex.col = x
	t.hex.row = y
	t.lines = Square(x, y)
	return &t
}

// func idForGrid(g grid, x, y int) int {
// 	return (g.maxX * y) + x
// }

func hexToID(hex hexCoords) int {
	cube := oddQToCube(hex)
	return spiralCubeToIDMAP[cube]
}

func Square(x, y int) []string {
	xCoord := convertCoord(x)
	yCoord := convertCoord(y)

	sqr := []string{
		"+--------------+",
		"|X" + xCoord + " Y" + yCoord + " Z" + yCoord + "|",
		"|X" + strconv.Itoa(x) + " Y" + strconv.Itoa(y) + "         |",
		"|              |",
		"|              |",
		"|              |",
		"|              |",
		"+--------------+",
	}
	return sqr
}

func (tl *tile) square() []string {
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

/*
  +-------------+
  |         ||  |
  |1234567890123|
  |             |+----------+
  |             |
  |             ||ABCDEFGHIJ|
  +-------------+|1234567890|
  +----------+|          |
  |          ||          |
  |          |+----------+
  |          |
  |          |
  +----------+



+--+    +--+
|  |    |  |
|  |+--+|  |
+--+|  |+--+
+--+|  |+--+
|  |+--+|  |
|  |+--+|  |
+--+|  |+--+
	|  |
	+--+
*/
