package main

import (
	"fmt"
	"strconv"

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

var zoneUncharted zone
var zoneClear zone

func initSector() {
	zoneUncharted = zone{0, "   UNCHARTED  ", 0}
	zoneClear = zone{0, "              ", 0}
}

func NewSector() *sector {
	sect := &sector{}
	sect.zoneByHex = make(map[hexCoords]*zone)
	sect.starByHex = make(map[hexCoords]string)

	return sect
}

func NewZone(id int, zoneType string, hex hexCoords) *zone {
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
	//return &zone{0, "   UNCHARTED  ", 0}
	return &zoneUncharted
}

func zoneClearSpace() *zone {
	//return &zone{0, "cl            ", 0}
	return &zoneClear
}

func (sect *sector) setZones() {
	for _, val := range gr.tileMap {
		sect.zoneByHex[val.hex] = zoneUnknown()
	}
	for _, val := range gr.tileMap {
		bZone := borderZone(val.hex)
		fmt.Println(bZone, val.hex)
		if bZone == zoneClearSpace() || bZone == zoneUnknown() {
			sect.scanA(val.hex)
		} else {
			sect.scanB(val.hex, bZone)
		}
		sect.scanC(val.hex)
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
		return "Nebula        "
	case 2:
		return "Void          "
	case 3:
		return "Dust Cloud    "
	case 4:
		return "Plasma        "
	}
	return "Error"
}

func borderZone(hex hexCoords) *zone {
	cube := oddQToCube(hex)
	for i := 0; i < 6; i++ {
		neib := cubeToHex(cubeNeighbor(cube, i))
		if sect.zoneByHex[neib] != zoneUnknown() && sect.zoneByHex[neib] != zoneClearSpace() {
			if sect.zoneByHex[neib] == nil {

				return zoneClearSpace() //когда проверяется хекс вне пределов карты он не имеет зоны - принудительно присваиваем ему "чистый космос"
			}
			return sect.zoneByHex[neib]
		}
	}

	return zoneClearSpace()
}

func (sect *sector) addStarByHex(hex hexCoords, star string) {
	for len(star) < 14 {
		star = star + " "
	}
	sect.starByHex[hex] = star
}

func (sect *sector) scanA(hex hexCoords) {
	r := utils.RollDice("d20")
	sect.zoneByHex[hex] = zoneClearSpace()
	if r == 20 {
		sect.zoneByHex[hex] = NewZone(len(sect.zone)+1, newNaturalZone(), hex)
	}
}

func (sect *sector) scanB(hex hexCoords, nZone *zone) {
	adjZones := adjustedZones(hex, nZone)
	dice := "d0"
	tn := 0
	switch nZone.zoneType {
	case "Nebula        ":
		dice = "d20"
		tn = 20
	case "Void          ":
		dice = "d20"
		tn = 20
	case "Dust Cloud    ":
		dice = "d12"
		tn = 12
	case "Plasma        ":
		dice = "d10"
		tn = 10
	case "Weird Energy  ":
		dice = "d8"
		tn = 8
	}
	zRoll := utils.RollDice(dice, nZone.zoneSize-adjZones)
	if zRoll > tn {

		sect.zoneByHex[hex] = zoneClearSpace()
	} else {
		fmt.Println(hex, dice, zRoll, tn)
		//panic(0)
		nZone.expandZone(hex)
	}
}

func (sect *sector) scanC(hex hexCoords) {
	zone := sect.zoneByHex[hex]
	pMod := 0
	if zone.zoneType == "Nebula        " {
		pMod++
	}
	if zone.zoneType == "Void          " {
		pMod--
	}
	pRoll := utils.RollDice("d6", pMod)
	if pRoll < 4 {
		sect.addStarByHex(hex, "              ")
	}
	if pRoll == 4 {
		tRoll := utils.RollDice("d8", -4)
		if tRoll < 1 {
			tRoll = 0
		}
		star := "T-000" + strconv.Itoa(tRoll)
		sect.addStarByHex(hex, star+"        ")
	}
	if pRoll > 4 {
		sect.addStarByHex(hex, "Star          ")
	}
}

func adjustedZones(hex hexCoords, zone *zone) int {
	adjZones := 0
	cube := oddQToCube(hex)
	for i := 0; i < 6; i++ {
		neib := cubeToHex(cubeNeighbor(cube, i))
		if sect.zoneByHex[neib] == zone {
			adjZones++
		}
	}

	return adjZones
}

func equals(zoneA, zoneB *zone) bool {
	if zoneA.zoneID != zoneB.zoneID {
		return false
	}
	if zoneA.zoneSize != zoneB.zoneSize {
		return false
	}
	if zoneA.zoneType != zoneB.zoneType {
		return false
	}
	return true
}

/*
Sector
 Zone
  System
   World

*/
