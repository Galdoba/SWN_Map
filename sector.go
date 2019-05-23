package main

import (
	"github.com/Galdoba/utils"
)

//Sector - содержит о себе всю информацию о вселенной.
type sector struct {
	name      string
	zoneByHex map[hexCoords]string
	starByHex map[hexCoords]string
}

func NewSector() *sector {
	sect := &sector{}
	sect.zoneByHex = make(map[hexCoords]string)
	return sect
}

func (gr *grid) putStars() {
	for _, val := range gr.tileMap {
		gr.sector.zoneByHex[val.hex] = "      unknown "
	}
	for _, val := range gr.tileMap {
		if gr.sector.zoneByHex[val.hex] == "      unknown " {
			r := utils.RollDice("d20")
			if r == 20 {
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
		return "Void"
	case 3:
		return "Dust Cloud"
	case 4:
		return "Plasma"
	}
	return "Error"
}

/*
Sector
 Zone
  System
   World

*/
