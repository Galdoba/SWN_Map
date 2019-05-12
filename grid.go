package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Galdoba/utils"
)

const (
	directionN  = 0
	directionNE = 1
	directionSE = 2
	directionS  = 3
	directionSW = 4
	directionNW = 5
)

//Hex -
type tile struct {
	hex   hexCoords
	cube  cubeCoords
	lines []string
}

func newTileHex(col, row int) *tile {
	tile := &tile{}
	tile.hex = setHexCoords(col, row)
	tile.cube = oddQToCube(tile.hex)
	tile.lines = []string{
		"+--------------+",
		"|" + strconv.Itoa(tile.hex.col) + " " + strconv.Itoa(tile.hex.row) + "           |",
		"|              |",
		"|" + cubeCoordsStr(tile.cube) + "|",
		"|              |",
		"|              |",
		"|              |",
		"+--------------+",
	}
	return tile
}

type cubeCoords struct {
	x int
	y int
	z int
}

func cubeCoordsStr(cube cubeCoords) string {
	fmt.Println(cube)
	xStr := coordNumToStr("X", cube.x)
	yStr := coordNumToStr("Y", cube.y)
	zStr := coordNumToStr("Z", cube.z)
	output := xStr + " " + yStr + " " + zStr
	return output
}

func coordNumToStr(coordName string, x int) string {
	xStr := coordName
	if x < 0 {
		xStr += "-"
		x = x * -1
	} else {
		xStr += " "
	}
	fmt.Println("1:", xStr)
	if x < 10 && x > -10 {
		xStr += "0"
		xStr += strconv.Itoa(x)
	} else {
		xStr += strconv.Itoa(x)
	}
	return xStr
}

func setCubeCoords(x, y, z int) cubeCoords {
	cube := cubeCoords{}
	cube.x = x
	cube.y = y
	cube.z = z
	return cube
}

func oddQToCube(hex hexCoords) cubeCoords {
	x := hex.col
	z := hex.row - (hex.col-(hex.col&1))/2
	y := -x - z
	return setCubeCoords(x, y, z)
}

type hexCoords struct {
	col int
	row int
}

func setHexCoords(c, r int) hexCoords {
	offCrds := hexCoords{}
	offCrds.col = c
	offCrds.row = r
	return offCrds
}

func cubeToHex(cube cubeCoords) hexCoords {
	col := cube.x
	row := cube.z + (cube.x-(cube.x&1))/2
	return setHexCoords(col, row)
}

var hexDirections [][]hexCoords

func initGrids() {
	// hexDirections = [][]hexCoords{
	// 	{hexCoords{1, 0}, hexCoords{1, -1}, hexCoords{0, -1}, hexCoords{-1, -1}, hexCoords{-1, 0}, hexCoords{0, 1}},
	// 	{hexCoords{1, 1}, hexCoords{1, 0}, hexCoords{0, -1}, hexCoords{-1, 0}, hexCoords{-1, 1}, hexCoords{0, 1}},
	// }
	hexDirections = [][]hexCoords{
		{hexCoords{0, -1}, hexCoords{1, -1}, hexCoords{1, 0}, hexCoords{0, 1}, hexCoords{-1, 0}, hexCoords{-1, -1}},
		{hexCoords{0, -1}, hexCoords{1, 0}, hexCoords{1, 1}, hexCoords{0, 1}, hexCoords{-1, 1}, hexCoords{-1, 0}},
	}

}

func info(t tile) {
	fmt.Println("hex coords:", t.hex)
	fmt.Println("cube coords:", t.cube)
}

func drawTile(t tile) {
	for i := range t.lines {
		fmt.Println(t.lines[i])
	}
}

func hexNeighbor(hex hexCoords, direction int) hexCoords {
	parity := hex.col & 1
	dir := hexDirections[parity][direction]
	return setHexCoords(hex.col+dir.col, hex.row+dir.row)
}

func cubeNeighbor(cube cubeCoords, direction int) cubeCoords {
	hex := cubeToHex(cube)
	parity := hex.col & 1
	dir := hexDirections[parity][direction]
	hexN := setHexCoords(hex.col+dir.col, hex.row+dir.row)
	return oddQToCube(hexN)
}

func cubeDistance(cubeA, cubeB cubeCoords) int {
	return int(math.Abs(float64(cubeA.x-cubeB.x)) + math.Abs(float64(cubeA.y-cubeB.y)) + math.Abs(float64(cubeA.z-cubeB.z))/2)
}

func hexDistance(hexA, hexB hexCoords) int {
	cubeA := oddQToCube(hexA)
	cubeB := oddQToCube(hexB)
	return cubeDistance(cubeA, cubeB)
}

//подвопросом
func cubeLine(cubeA, cubeB cubeCoords) []cubeCoords {
	dist := cubeDistance(cubeA, cubeB)
	var line []cubeCoords
	line = append(line, cubeA)
	localDist := dist
	//fmt.Println(localDist, "localDist")
	tick := 0
	for localDist > 0 {
		localDist = cubeDistance(cubeA, cubeB)
		tick++
		//	fmt.Println(tick)
		//	fmt.Println(localDist)
		if tick > 1000 {
			return line
		}
		//line = append(line, cubeA) //доделать
		for dir := 0; dir < 6; dir++ {
			cubeN := cubeNeighbor(cubeA, dir)
			//fmt.Println(cubeDistance(cubeN, cubeB), "cube", cubeN)
			if cubeDistance(cubeN, cubeB) < localDist {
				//	fmt.Println(cubeDistance(cubeN, cubeB), "cube", cubeN, "Pick IT!!")
				line = append(line, cubeN) //доделать
				cubeA = cubeN
				break
			}
		}
	}
	for i := range line {
		fmt.Println(line[i])
	}
	return line
}

func hexRectangleDimentions(hex1, hex2 hexCoords) (int, int, int, int) {
	maxX := utils.Max(hex1.col, hex2.col)
	minX := utils.Min(hex1.col, hex2.col)
	//xDim := maxX - minX + 1
	maxY := utils.Max(hex1.row, hex2.row)
	minY := utils.Min(hex1.row, hex2.row)
	//yDim := maxY - minY + 1
	return minX, minY, maxX, maxY
}
