package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
)

var counter int
var ticker int
var tickerGo bool
var appErr error
var runStart time.Time
var grid0 *grid
var mapCellX int
var mapCellY int
var mapCellXLast int
var mapCellYLast int

func main() {
	runStart = time.Now()
	counter = 1
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	bindKeys(g)

	tile1 := newTileHex(-1, -1)
	tile2 := newTileHex(3, 2)

	minX, minY, maxX, maxY := hexRectangleDimentions(tile1.hex, tile2.hex)

	gr := NewGrid(minX, minY, maxX, maxY)
	for y := minY; y < gr.maxY; y++ {
		for x := minX; x < gr.maxX; x++ {
			tl := NewTile(x, y)
			id := idForGrid(*gr, x, y)

			gr.tileMap[id] = tl
		}
	}
	grid0 = gr

	//Tile("06","02")

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			g.Execute(layout)

			if tickerGo {
				ticker = ticker + counter
			}

		}
	}()

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func tileStrings(x, y int) []string {
	tileID := idForGrid(*grid0, x, y)
	sqr := grid0.tileMap[tileID].content
	return sqr
}

func tileByID(id int) *Tile {
	return grid0.tileMap[id]
}

//Создает и отрисовывает все окна - к этому моменту программа должна иметь
//представление что где и в каком окне должно быть.
//Запускается при каждом обновлении экрана
//TODO: прощупать стоит ли хранить содержимое окна где-либо вне его.
func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	v1, v1Err := newPanelInfo(g, "Size", 0, 0, maxX/4, maxY-3)
	if v1Err != nil && v1Err != gocui.ErrUnknownView {
		return v1Err
	}
	v2, err := newPanelInfo(g, "Info", maxX/4+1, 0, (maxX/4)*3-2, maxY-maxY/4)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	fillPanel(v1)
	fillPanel(v2)
	return nil
}

func newPanelInfo(g *gocui.Gui, panelName string, pX, pY, pW, pH int) (*gocui.View, error) {

	v, err := g.SetView(panelName, pX, pY, pX+pW, pY+pH)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}
	v.Title = panelName
	if panelName == "Info" {

	}
	return v, nil
}

func fillPanel(v *gocui.View) {
	switch v.Title {
	case "Size":
		v.Clear()

		t := time.Now().Format("2006-Jan-02 15:04:05")
		ts := runStart.Format("2006-Jan-02 15:04:05")
		fmt.Fprintf(v, "Current Real Time: %s \n", t)
		fmt.Fprintf(v, "RunStart: %s\n", ts)
		//s := time.Since(runStart).Round(time.Millisecond)
		s := 567.2
		pureSeconds := float64(time.Millisecond) + 567

		fmt.Fprintf(v, "Program working: %s\n Sec: %d\n", s, pureSeconds/1000)
		fmt.Fprintf(v, "%d, %d\n", ticker, counter)
		fmt.Fprintf(v, "rume 'm' = %d", string(rune(109)))
		fmt.Fprintf(v, "\n Random Roll: ", strconv.Itoa(roll1dX(counter, 0)), "//////////")
		fmt.Fprintf(v, "\n"+strconv.Itoa(mapCellX)+" mX"+"   "+strconv.Itoa(mapCellY)+" mY")
		fmt.Fprintf(v, "\nTile Clicked: ")
		allStr := drawGrid(*grid0)
		lines := strings.Split(allStr, "\n")
		bytesAr := []byte(lines[mapCellY])

		fmt.Fprintf(v, "\nLine: "+string(bytesAr[mapCellX]))
		tileID := mapCellsToID(mapCellX, mapCellY)
		mapXCoords, mapYCoords := mapCoordinates(mapCellX, mapCellY)
		sqr := tileStrings(mapXCoords, mapYCoords)
		for i := range sqr {
			fmt.Fprintf(v, "\n"+sqr[i])
			//fmt.Fprintf(v, "\ntickerGo is active")
		}
		fmt.Fprintf(v, "\n Tile ID: "+strconv.Itoa(tileID))
	case "Info":
		v.Clear()

		fmt.Fprintf(v, drawGrid(*grid0))
	}

}

func mapCellsToID(mapCellX, mapCellY int) int {
	mapXCoords, mapYCoords := mapCoordinates(mapCellX, mapCellY)
	tileID := idForGrid(*grid0, mapXCoords, mapYCoords)
	return tileID
}

func mapCoordinates(mapCellX, mapCellY int) (mapXCoords int, mapYCoords int) {
	mapYCoords = (0 + mapCellY) / 6
	offset := mapYCoords % 2
	if offset == 1 {
		offset = 6
	}
	mapXCoords = ((mapCellX + offset - 4) / 16)
	return mapXCoords, mapYCoords
}

/*

+--------------+
|X-00 Y+00 Z-00|
|        Nebula|
|Star:  G2134a3|
|*P  *St    *Mf|
|      Ansa Tau|
|player is here|
+--------------+
*/
