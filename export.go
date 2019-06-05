package main

import "strconv"

func tileToExportMap(t Tile) (string, string) {
	key := strconv.Itoa(t.ID)
	val := ""
	val += hexCoordsStr(t.hex) + "|"
	val += t.lines[2] + "|"
	val += t.lines[3] + "|"
	val += t.lines[4]

	return key, val
}
