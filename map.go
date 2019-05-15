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
	gr.tileMap[0] = newTileHex(minX, minY)
	gr.tileSize = /*len(gr.tileMap[0].lines)*/ 16
	idNum := 0
	for y := gr.minY; y < gr.maxY; y++ {
		for x := gr.minX; x < gr.maxX; x++ {
			idNum++
			tl := newTileHex(x, y)
			id := /*idForGrid(*gr, x, y)*/ idNum
			gr.tileMap[id] = tl
		}
	}
	gr.tileSize = len(gr.tileMap[0].lines)
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
	for y := gr.minY; y <= gr.maxY; y++ {
		for x := gr.minX; x <= gr.maxX; x++ {
			//tile := newTileHex(x, y)
			//for i := range tile.lines {
			gridStr += "draw Tile(" + strconv.Itoa(x) + ";" + strconv.Itoa(y) + ") cubeCoords: " + /*cubeCoordsStr(oddQToCube(hexCoords{x, y})) + */ "\n"
			

		}
	}
	return gridStr

	// for lineNum := 0; lineNum < gr.maxY*gr.tileSize; lineNum++ {
	// 	line := ""
	// 	if lineNum < 10 {
	// 		line = line + "0"
	// 	}
	// 	line = line + strconv.Itoa(lineNum) + " | "
	// 	var idStack []int
	// 	tileContentNum := lineNum % gr.tileSize
	// 	rowNum := lineNum / gr.tileSize
	// 	prefix := ""
	// 	if rowNum%2 > 0 {
	// 		prefix = "        " //Offset For Drawing
	// 	}
	// 	line = line + prefix
	// 	for i := 0; i < gr.maxX; i++ {
	// 		idStack = append(idStack, gr.maxX*(lineNum/gr.tileSize)+i)
	// 	}

	// 	for i := range idStack {
	// 		line = line + gr.tileMap[idStack[i]].lines[tileContentNum]
	// 	}
	// 	gridStr = gridStr + line + "\n"
	// }
	return gridStr
}

func NewTile(x, y int) *tile {
	t := tile{}
	t.hex.col = x
	t.hex.row = y
	t.lines = Square(x, y)
	return &t
}

func idForGrid(g grid, x, y int) int {
	return (g.maxX * y) + x
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
