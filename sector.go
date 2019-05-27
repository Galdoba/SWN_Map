package main

import (
	"github.com/Galdoba/utils"
)

//Sector - содержит о себе всю информацию о вселенной.
type sector struct {
	name       string
	zone       []*zone
	zoneByHex  map[hexCoords]string
	zoneByHex0 map[hexCoords]*zone
	starByHex  map[hexCoords]string
}

type zone struct {
	zoneID   int
	zoneType string
	zoneSize int
}

func NewSector() *sector {
	sect := &sector{}
	sect.zoneByHex = make(map[hexCoords]string)
	sect.zoneByHex0 = make(map[hexCoords]*zone)
	return sect
}

func zoneUnknown() *zone {
	return &zone{0, "      Unknown ", 0}
}

func (gr *grid) setZones() {
	totalZones := 0
	for _, val := range gr.tileMap {
		gr.sector.zoneByHex[val.hex] = "      unknown "
		gr.sector.zoneByHex0[val.hex] = zoneUnknown()
	}

	for _, val := range gr.tileMap {
		if gr.sector.zoneByHex[val.hex] == "      unknown " {
			zl := gr.sector.borderZonesList(oddQToCube(val.hex))
			if len(zl) < 1 {
				//step A
			}
			r := utils.RollDice("d20")
			if r == 20 {
				totalZones++
				gr.sector.zoneByHex[val.hex] = newZone()

			} else {
				gr.sector.zoneByHex[val.hex] = "Normal Space  "
			}
		}
	}
}

func (sctr *sector) getStar(hex hexCoords) string {
	if sctr.zoneByHex[hex] != "" {
		return sctr.zoneByHex[hex]
	}
	return "              "
}

func newZone() string {
	r1 := utils.RollDice("d6")
	if r1 == 6 {
		return "Weird Energy "
	}
	return newNaturalZone()
}

func newNaturalZone() string {
	r := utils.RollDice("d4")
	switch r {
	case 1:
		return "Nebula"
	case 2:
		return "Void          "
	case 3:
		return "Dust Cloud    "
	case 4:
		return "Plasma        "
	}
	return "Error"
}

func (sector *sector) borderZonesList(cube cubeCoords) []*zone {
	var zL []*zone
	for i := 0; i < 6; i++ {
		neib := cubeToHex(cubeNeighbor(cube, i))
		if sector.zoneByHex0[neib] != zoneUnknown() {
			zL = append(zL, sector.zoneByHex0[neib])
		}

	}
	return zL
}

/*
Sector
 Zone
  System
   World

*/
