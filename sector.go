package main

import (
	"github.com/Galdoba/utils"
)

//Sector - содержит о себе всю информацию о вселенной.
type sector struct {
	name      string
	zoneByHex map[hexCoords]string
}

func NewSector() *sector {
	sect := &sector{}
	sect.zoneByHex = make(map[hexCoords]string)
	return sect
}

func (sctr *sector) putStars() {
	for _, val := range gr.tileMap {
		sctr.zoneByHex[val.hex] = "              "
	}
	for _, val := range gr.tileMap {
		if sctr.zoneByHex[val.hex] == "              " {
			r := utils.RollDice("d6")
			if r > 3 {
				sctr.zoneByHex[val.hex] = "      SS      "
			}
		}
	}
}

func (sctr *sector) getStar(hex hexCoords) string {
	return sctr.zoneByHex[hex]
}

/*
Sector
 Zone
  System
   World

*/
