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
	totalRows := gr.maxY - gr.minY + 1
	totalCols := gr.maxX - gr.minX + 1
	tilelines := 8
	totalSegments := (totalRows * totalCols * tilelines) + (totalCols * tilelines / 2)
	//panic("stop")
	gridStr += "Tile ID =" + strconv.Itoa(totalCols) + strconv.Itoa(totalRows) + strconv.Itoa(tilelines) + " " + strconv.Itoa(totalSegments)
	for y := gr.minY; y <= gr.maxY; y++ {
		for x := gr.minX; x <= gr.maxX; x++ {
			// gridStr += "draw Tile(" + strconv.Itoa(x) + ";" + strconv.Itoa(y) + ") cubeCoords: " + /*cubeCoordsStr(oddQToCube(hexCoords{x, y})) + */ "\n"
			// id := hexToID(hexCoords{x, y})
			// var sqrLN []string
			// if val, ok := gr.tileMap[id]; ok {
			// 	sqrLN = val.lines
			// }
			// gridStr += "Tile ID =" + strconv.Itoa(id) + "\n"
			// for i := range sqrLN {
			// 	gridStr += sqrLN[i] + "\n"
			// }

		}
	}
	return gridStr

}

func defineSegment(segment int, gr grid) string {
	/*
		псевдокод:
		segment / totalCol = row
		segment % totalCol = col
		col%2 = offset (if 0 = false)
		if offset {
			line = line - 4
			if line < 0 {
				row = row - 1
				line = line + 8
				if row < 0 {
					segment = BLANK
				}
			}
		}
		return gr.hex(col - gr.minX, row - gr.minY).line[line]
	*/
	offset := false
	if segment%2 == 1 {
		offset = true
	}
	totalCols := gr.maxX - gr.minX + 1
	line := segment % totalCols

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
