package main

import (
	"github.com/Galdoba/utils"
)

//Sector - содержит о себе всю информацию о вселенной.
type sector struct {
	name      string
	zone      []*zone
	zoneByHex map[hexCoords]*zone
	starByHex map[hexCoords]string
}

type zone struct {
	zoneID   int
	zoneType string
	zoneSize int
}

func NewSector() *sector {
	sect := &sector{}
	sect.zoneByHex = make(map[hexCoords]*zone)
	sect.starByHex = make(map[hexCoords]string)

	return sect
}

func (sect *sector) NewZone(id int, zoneType string, hex hexCoords) *zone {
	zone := &zone{}
	zone.zoneID = id
	zone.zoneType = zoneType
	zone.zoneSize = 1
	sect.zone = append(sect.zone, zone)
	sect.zoneByHex[hex] = zone

	return zone
}

func (zone *zone) expandZone(hex hexCoords) {
	zone.zoneSize++
	sect.zoneByHex[hex] = zone
}

func zoneUnknown() *zone {
	return &zone{0, "   UNCHARTED  ", 0}
}

func (sect *sector) setZones() {
	for _, val := range gr.tileMap {
		sect.zoneByHex[val.hex] = zoneUnknown()
	}

}

func (sctr *sector) getZone(hex hexCoords) string {
	return sctr.zoneByHex[hex].zoneType
}

func (sctr *sector) getStar(hex hexCoords) string {
	if sctr.starByHex[hex] != "" {
		return sctr.starByHex[hex]
	}
	return "              "
}

func randomNewZone() string {
	r1 := utils.RollDice("d6")
	if r1 == 6 {
		return "Weird Energy  "
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
		if sector.zoneByHex[neib] != zoneUnknown() {
			zL = append(zL, sector.zoneByHex[neib])
		}

	}
	return zL
}

func (sect *sector) addStarByHex(hex hexCoords, star string) {
	for len(star) < 14 {
		star = star + " "
	}
	sect.starByHex[hex] = star
}

/*
Sector
 Zone
  System
   World

*/
