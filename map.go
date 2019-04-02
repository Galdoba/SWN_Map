package main

import (
	"strconv"
)

func main0() {
	gr := NewGrid(6, 7)
	for y := 0; y < gr.maxY; y++ {
		for x := 0; x < gr.maxX; x++ {
			tl := NewTile(1+x, 1+y)
			id := idForGrid(*gr, x, y)

			gr.tileMap[id] = tl
		}
	}
	gr.tileMap[5].content[4] = "|Test Info |"
	gr.tileMap[14].content[3] = "|   test2  |"
	drawGrid(*gr)
}

type grid struct {
	tileMap  map[int]*tile
	maxX     int
	maxY     int
	tileSize int
}

type screenCoords struct {
	xCoord int
	yCoord int
}

type tile struct {
	sc      screenCoords
	content []string
}

func NewGrid(maxX, maxY int) *grid {
	gr := grid{}
	gr.tileMap = make(map[int]*tile)
	gr.maxX = maxX
	gr.maxY = maxY
	gr.tileMap[0] = NewTile(0, 0)
	gr.tileSize = len(gr.tileMap[0].content)
	return &gr
}

func getScreenCoords(t tile) screenCoords {
	return t.sc
}

func drawLine(s string) string {
	return s
}

func drawGrid(gr grid) string {
	gridStr := ""
	for lineNum := 0; lineNum < gr.maxY*gr.tileSize; lineNum++ {
		line := ""
		if lineNum < 10 {
			line = line + "0"
		}
		line = line + strconv.Itoa(lineNum) + " | "
		var idStack []int
		tileContentNum := lineNum % gr.tileSize
		rowNum := lineNum / gr.tileSize
		prefix := ""
		if rowNum%2 > 0 {
			prefix = "      "
		}
		line = line + prefix
		for i := 0; i < gr.maxX; i++ {
			idStack = append(idStack, gr.maxX*(lineNum/gr.tileSize)+i)
		}

		for i := range idStack {
			line = line + gr.tileMap[idStack[i]].content[tileContentNum]
		}
		gridStr = gridStr + line + "\n"
	}
	return gridStr
}

func NewTile(x, y int) *tile {
	t := tile{}
	t.sc.xCoord = x
	t.sc.yCoord = y
	t.content = Square(x, y)
	return &t
}

func idForGrid(g grid, x, y int) int {
	return (g.maxX * y) + x
}

func Square(x, y int) []string {
	xCoord := convertCoord(x)
	yCoord := convertCoord(y)

	sqr := []string{
		"+----------+",
		"|          |",
		"|  " + xCoord + yCoord + "  |",
		"|          |",
		"|          |",
		"+----------+",
	}
	return sqr
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
		coord = " " + coord
	}
	return coord
}

/*
  +----------+
  |      ||  |
  |1234567890|
  |          |+----------+
  |          ||ABCDEFGHIJ|
  +----------+|1234567890|
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
