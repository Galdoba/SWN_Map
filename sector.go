package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Galdoba/utils"
)

//Sector - содержит о себе всю информацию о вселенной.

// func (gr *grid) setZones() {
// 	for _, val := range gr.tileMap {
// 		sect.zoneByHex[val.hex] = zoneUnknown()
// 	}
// 	for _, val := range gr.tileMap {
// 		bZone := borderZone(val.hex)
// 		fmt.Println(bZone, val.hex)
// 		if bZone == zoneClearSpace() || bZone == zoneUnknown() {
// 			//sect.zoneByHex[val.hex] = scanA2()
// 			sect.scanA(val.hex)
// 		} else {
// 			sect.scanB(val.hex, bZone)
// 		}
// 		sect.scanC(val.hex)
// 		//appendToSectorFile("Added hex:"+val.hex.String()+"\n", "string2\n")
// 	}
// }

func randomNewZone() string {
	r1 := utils.RollDice("d6")
	if r1 == 6 {
		return "Weird Energy"
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

// func scanA2() string {
// 	r := utils.RollDice("d20")
// 	if r == 20 {
// 		return newNaturalZone()
// 	}
// 	return zoneClearSpace().zoneType
// }

// func scanB(hex hexCoords, nZone *zone) {
// 	adjZones := adjustedZones(hex, nZone)
// 	dice := "d0"
// 	tn := 0
// 	switch nZone.zoneType {
// 	case "Nebula        ":
// 		dice = "d20"
// 		tn = 20
// 	case "Void          ":
// 		dice = "d20"
// 		tn = 20
// 	case "Dust Cloud    ":
// 		dice = "d12"
// 		tn = 12
// 	case "Plasma        ":
// 		dice = "d10"
// 		tn = 10
// 	case "Weird Energy  ":
// 		dice = "d8"
// 		tn = 8
// 	}
// 	zRoll := utils.RollDice(dice, nZone.zoneSize-adjZones)
// 	if zRoll > tn {

// 		sect.zoneByHex[hex] = zoneClearSpace()
// 	} else {
// 		fmt.Println(hex, dice, zRoll, tn)
// 		//panic(0)
// 		nZone.expandZone(hex)
// 	}
// }

// func (sect *sector) scanC(hex hexCoords) {
// 	zone := sect.zoneByHex[hex]
// 	pMod := 0
// 	if zone.zoneType == "Nebula        " {
// 		pMod++
// 	}
// 	if zone.zoneType == "Void          " {
// 		pMod--
// 	}
// 	pRoll := utils.RollDice("d6", pMod)
// 	if pRoll < 4 {
// 		sect.addStarByHex(hex, "              ")
// 	}
// 	if pRoll == 4 {
// 		tRoll := utils.RollDice("d8", -4)
// 		if tRoll < 1 {
// 			tRoll = 0
// 		}
// 		star := "T-000" + strconv.Itoa(tRoll)
// 		sect.addStarByHex(hex, star+"        ")
// 	}
// 	if pRoll > 4 {
// 		sect.addStarByHex(hex, "Star          ")
// 	}
// }

func appendToSectorFile(s ...string) {
	f, err := os.OpenFile("Sector.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	for i := range s {
		_, err = f.Write([]byte(s[i]))
		if err != nil {
			log.Fatal(err)
		}
	}
	f.Close()
}

func (gr *grid) scanSector() {
	for _, v := range gr.tileMap {
		time.Sleep(1)
		if !borderZonesExists(v) {
			v.LayerZone = scanA()
		} else {
			//v.LayerZone = "Check Expand"
			//gr.tileMap[k] = v
			//zone := adjustedZone(v)
			//size := v.zoneSize()
			v.LayerZone = scanB(adjustedZone(v))
		}
		v.LayerStar = scanC(v)

	}
}

func neighborTiles(tile *Tile) []*Tile {
	var neighbors []*Tile
	for i := 0; i < 6; i++ {
		nID := spiralCubeToIDMAP[cubeNeighbor(tile.cube, i)]
		if nTile, ok := gr.tileMap[nID]; ok { //проверяем только те тайлы которые находятся на карте
			neighbors = append(neighbors, nTile)
		}
	}
	return neighbors
}

func borderZonesExists(tile *Tile) bool {
	nTile := neighborTiles(tile)
	for i := range nTile {
		if nTile[i].isZone() {
			return true
		}
	}
	return false
}

func scanA() string {
	r := utils.RollDice("d20")
	if r != 20 {
		return ""
	}
	return newNaturalZone()
}

func scanB(zone string) string {
	r := utils.RollDice("d6")
	if r > 3 {
		return ""
	}
	return zone
}

func scanC(tile *Tile) string {
	pMod := 0
	if tile.LayerZone == "Nebula" {
		pMod++
	}
	if tile.LayerZone == "Void" {
		pMod--
	}
	pRoll := utils.RollDice("d6", pMod)
	if pRoll < 4 {
		return ""
	}
	if pRoll == 4 {
		tRoll := utils.RollDice("d8", -4)
		if tRoll < 1 {
			tRoll = 0
		}
		return "T-000" + strconv.Itoa(tRoll)

	}
	if pRoll > 4 {
		syst := NewStarSystem(utils.RollDice("d100"))
		return syst.getStarClasses() + "-" + syst.getStarCode()
	}
	return ""
}

func (tile *Tile) isZone() bool {
	if tile.LayerZone != "NO DATA" && tile.LayerZone != "" {
		return true
	}
	return false
}

func adjustedZone(tile *Tile) string {
	nTiles := neighborTiles(tile)
	var zones []string
	for i := range nTiles {
		if !nTiles[i].isZone() {
			continue
		}
		zones = appendIfNew(zones, nTiles[i].LayerZone)
	}
	return utils.RandomFromList(zones)
}

func (tile *Tile) zoneSize() int {
	if !tile.isZone() {
		return 0
	}
	var ids []string
	ids = appendIfNew(ids, strconv.Itoa(tile.ID))
	size := 1

	for i := range ids {
		id, _ := strconv.Atoi(ids[i])
		nTiles := neighborTiles(gr.tileMap[id])
		for j := range nTiles {
			if nTiles[j].LayerZone == tile.LayerZone {
				size++
				ids = appendIfNew(ids, strconv.Itoa(tile.ID))
			}
		}
	}
	fmt.Println(tile.hex, tile.LayerZone, size)
	return size
}

func appendIfNew(slice []string, elem string) []string {
	for i := range slice {
		if slice[i] == elem {
			return slice
		}
	}
	slice = append(slice, elem)
	return slice
}
