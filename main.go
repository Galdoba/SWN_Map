package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jroimartin/gocui"
)

var counter int
var ticker int
var tickerGo bool
var appErr error
var runStart time.Time
var gr *grid
var mapCellX int
var mapCellY int
var mapCellXLast int
var mapCellYLast int

func main() {
	initGrids()
	runStart = time.Now()
	counter = 1
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	bindKeys(g)

	tile1 := newTileHex(0, 0)
	tile2 := newTileHex(0, 0)
	tile3 := newTileHex(0, 0)

	fmt.Println(spiralCubeToIDMAP[tile1.cube], "is ID for tile 1")
	fmt.Println(spiralCubeToIDMAP[tile2.cube], "is ID for tile 2")
	fmt.Println(spiralCubeToIDMAP[tile3.cube], "is ID for tile 3")

	// fmt.Println(cubeRing(tile1.cube, 2))
	// fmt.Println(cubeSpiral(tile1.cube, 0))
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Spiral with", i, "radius has", len(cubeSpiral(tile1.cube, i)), "hexes and has id =")
	// }

	//minX, minY, maxX, maxY := hexRectangleDimentions(tile1.hex, tile2.hex, tile3.hex)

	//gr = NewGrid(minX, minY, maxX, maxY)
	gr = NewGrid(hexRectangleDimentions(newTileHex(0, 0).hex))
	//Tile("06","02")

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			g.Update(layout)

			if tickerGo {
				ticker = ticker + counter
			}

		}
	}()

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
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
		s := time.Since(runStart).Round(time.Millisecond)
		pureSeconds := float64(time.Millisecond) + 567
		fmt.Fprintf(v, "Program working: %s\n Sec: %d\n", s, pureSeconds/1000)
		fmt.Fprintf(v, "%d, %d\n", ticker, counter)
		fmt.Fprintf(v, "rume 'm' = %d", string(rune(109)))
		fmt.Fprintf(v, "\n"+strconv.Itoa(mapCellX)+" mX"+"   "+strconv.Itoa(mapCellY)+" mY")
		fmt.Fprintf(v, "\nTile Clicked: ")
	case "Info":
		v.Clear()

		fmt.Fprintf(v, drawGrid(*gr))
	}

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
