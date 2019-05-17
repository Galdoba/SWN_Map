package main

import (
	"fmt"
	"strconv"
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
	id    int
	lines []string
}

func newTileHex(col, row int) *tile {
	tile := &tile{}
	tile.hex = setHexCoords(col, row)
	tile.cube = oddQToCube(tile.hex)
	tile.id = spiralCubeToIDMAP[tile.cube]
	tile.lines = []string{
		"+--------------+",
		"|              |",
		"|" + hexCoordsStr(tile.hex) + "|",
		"|" + cubeCoordsStr(tile.cube) + "|",
		"|              |",
		"|              |",
		"|              |",
		"+--------------+",
	}

	//	fmt.Println("Create:" + strconv.Itoa(col) + " " + strconv.Itoa(row))
	return tile
}

func hexCoordsStr(hex hexCoords) string {
	//fmt.Println(cube)
	xStr := coordNumToStr("X", hex.col)
	yStr := coordNumToStr("Y", hex.row)

	output := xStr + " " + yStr + "     "
	return output
}

type cubeCoords struct {
	x int
	y int
	z int
}

func cubeCoordsStr(cube cubeCoords) string {
	//fmt.Println(cube)
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
	//fmt.Println("1:", xStr)
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
var cubeDirections []cubeCoords
var spiralCubeToIDMAP map[cubeCoords]int

func initGrids() {
	// hexDirections = [][]hexCoords{
	// 	{hexCoords{1, 0}, hexCoords{1, -1}, hexCoords{0, -1}, hexCoords{-1, -1}, hexCoords{-1, 0}, hexCoords{0, 1}},
	// 	{hexCoords{1, 1}, hexCoords{1, 0}, hexCoords{0, -1}, hexCoords{-1, 0}, hexCoords{-1, 1}, hexCoords{0, 1}},
	// }
	hexDirections = [][]hexCoords{
		{hexCoords{0, -1}, hexCoords{1, -1}, hexCoords{1, 0}, hexCoords{0, 1}, hexCoords{-1, 0}, hexCoords{-1, -1}},
		{hexCoords{0, -1}, hexCoords{1, 0}, hexCoords{1, 1}, hexCoords{0, 1}, hexCoords{-1, 1}, hexCoords{-1, 0}},
	}
	cubeDirections = []cubeCoords{
		cubeCoords{0, 1, -1}, cubeCoords{1, 0, -1}, cubeCoords{1, -1, 0}, cubeCoords{0, -1, 1}, cubeCoords{-1, 0, 1}, cubeCoords{-1, 1, 0},
	}
	grandSpiral := cubeSpiral(cubeCoords{0, 0, 0}, 200)
	idMAP := make(map[cubeCoords]int)
	for i := range grandSpiral {
		idMAP[grandSpiral[i]] = i
	}
	spiralCubeToIDMAP = idMAP
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
	// hex := cubeToHex(cube)
	// parity := hex.col & 1
	// dir := hexDirections[parity][direction]
	// hexN := setHexCoords(hex.col+dir.col, hex.row+dir.row)
	// return oddQToCube(hexN)
	cubeN := cubeCoords{cube.x + cubeDirections[direction].x, cube.y + cubeDirections[direction].y, cube.z + cubeDirections[direction].z}
	return cubeN
}

func cubeDistance(cubeA, cubeB cubeCoords) int {
	//return int(math.Abs(float64(cubeA.x-cubeB.x)) + math.Abs(float64(cubeA.y-cubeB.y)) + math.Abs(float64(cubeA.z-cubeB.z))/2)
	xDif := cubeA.x - cubeB.x
	if xDif < 0 {
		xDif = xDif * -1
	}
	yDif := cubeA.y - cubeB.y
	if yDif < 0 {
		yDif = yDif * -1
	}
	zDif := cubeA.z - cubeB.z
	if zDif < 0 {
		zDif = zDif * -1
	}
	difArr := []int{xDif, yDif, zDif}
	return maxFromIntArray(difArr)
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

func hexRectangleDimentions(hex ...hexCoords) (int, int, int, int) {
	var rowArray []int
	var colArray []int
	for i := range hex {
		rowArray = append(rowArray, hex[i].row)
		colArray = append(colArray, hex[i].col)
	}
	minX := minFromIntArray(colArray)
	minY := minFromIntArray(rowArray)
	maxX := maxFromIntArray(colArray)
	maxY := maxFromIntArray(rowArray)

	return minX, minY, maxX, maxY
}

func minFromIntArray(slice []int) int {
	min := slice[0]
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func maxFromIntArray(slice []int) int {
	min := slice[0]
	for i := range slice {
		if slice[i] > min {
			min = slice[i]
		}
	}
	return min
}

func cubeRing(center cubeCoords, radius int) (ring []cubeCoords) {
	//двигаемся на север пока не удалимся на radius от center
	var start cubeCoords
	for cubeDistance(center, start) < radius {
		start = cubeNeighbor(start, directionN)
	}
	//запоминаем точку старта
	//проверяем соседей:
	//каждый встреченный сосед находящийся на radius от center отправляется в ring
	ring = append(ring, start)
	done := false
	for !done {
		for i := 0; i < 6; i++ {
			ringAplicant := cubeNeighbor(ring[len(ring)-1], i)
			if ringAplicant == start {
				return ring
			}
			if cubeDistance(center, ringAplicant) == radius {
				if !coordsInRing(ringAplicant, ring) {
					ring = append(ring, ringAplicant)
					break
				}
			}
		}
	}
	//если новый сосед равен старту - возвращаем ring
	return ring
}

func cubeSpiral(center cubeCoords, radius int) (spiral []cubeCoords) {
	spiral = append(spiral, center)
	for i := 1; i < radius+1; i++ {
		spiral = append(spiral, cubeRing(center, i)...)
	}
	return spiral
}

func coordsInRing(cube cubeCoords, ring []cubeCoords) bool {
	for i := range ring {
		if cube == ring[i] {
			return true
		}
	}
	return false
}
