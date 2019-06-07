package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Galdoba/utils"

	"github.com/jroimartin/gocui"
)

var counter int
var ticker int
var tickerGo bool
var appErr error
var runStart time.Time
var gr *grid

func main() {
	seed := utils.RandomSeed()
	fmt.Println(seed)
	opt, _ := utils.TakeOptions("Define Sector:\n", "New Sector", "Load Sector")
	minX := 0
	minY := 0
	maxX := 0
	maxY := 0
	if opt == 1 {
		minX = utils.InputInt("Set Grid.minX")
		minY = utils.InputInt("Set Grid.minY")
		maxX = utils.InputInt("Set Grid.maxX")
		maxY = utils.InputInt("Set Grid.maxY")
	}

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

	gr = NewGrid(minX, minY, maxX, maxY)

	fmt.Println(drawGrid(*gr))
	gr.scanSector()
	// for _, v := range gr.tileMap {
	// 	if v.isZone() {
	// 		v.zoneSize()
	// 	}
	// }

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
		fmt.Fprintf(v, "\n"+strconv.Itoa(mapCellXLast)+" mXl"+"   "+strconv.Itoa(mapCellYLast)+" mYl")
		fmt.Fprintf(v, "\nTile Clicked: \n")
		hexID := gr.tileByClick(mapCellX, mapCellY)
		lastHexID := gr.tileByClick(mapCellXLast, mapCellYLast)
		var lines []string
		var data string
		if val, ok := gr.tileMap[hexID]; ok {
			lines = val.lines
			data = val.toString()
		} else {
			lines = append(lines, " ")
		}

		for i := range lines {
			fmt.Fprintf(v, lines[i]+"\n")
		}
		fmt.Fprintf(v, data)
		if hexID != lastHexID {
			return
		}
		if val, ok := gr.tileMap[lastHexID]; ok {
			data2 := val.toString()
			fmt.Fprintf(v, data2)
		}
	case "Info":
		v.Clear()

		fmt.Fprintf(v, drawGrid(*gr))
	}

}

/*

Navigation Mode
+--------------+
|Sector: V01H01| - coordinates
|        Nebula| - space conditions
|Star:  G2134a3| - star profile
|*P  *St    *Mf| - main world profile
|      Ansa Tau| - main world name
|player is here| - player marker
+--------------+

Trade Mode
+--------------+
|Sector: V01H01| - coordinates
|      Ansa Tau| - main world name
|P9/L10/F5     | - population profile
|      L-123456| - regulation profile
|              | - main world name
|P    F    Pr  | - player/Factor marker
+--------------+
 Sector: V01H01

{
	map: "Sect"
	rows: 15
	cols: 10
	hex: ...hexes

}






*/
