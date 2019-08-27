package main

import (
	"log"
	"os"
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
		if utils.RandomBool() {
			v.LayerStar = "scanC(v)	"
		}
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

func appendIfNew(slice []string, elem string) []string {
	for i := range slice {
		if slice[i] == elem {
			return slice
		}
	}
	slice = append(slice, elem)
	return slice
}
