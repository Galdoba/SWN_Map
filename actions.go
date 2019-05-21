package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func bindKeys(g *gocui.Gui) {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, actionQuit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", rune(113) /*q*/, gocui.ModNone, actionIncreaseCounter); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", rune(109) /*m*/, gocui.ModNone, actionToggleTicker); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("Info", gocui.MouseMiddle, gocui.ModNone, actionButtonClick); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("Info", gocui.MouseLeft, gocui.ModNone, actionSelectClick); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone, actionMoveRight); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone, actionMoveLeft); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, actionMoveUp); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, actionMoveDown); err != nil {
		log.Panicln(err)
	}

}

func actionQuit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

//////

func actionIncreaseCounter(g *gocui.Gui, v *gocui.View) error {
	counter++
	if counter > 50 {
		return gocui.ErrQuit
	}
	//toggleTicker()
	return nil
}

func actionToggleTicker(g *gocui.Gui, v *gocui.View) error {
	toggleTicker()
	return nil
}

func toggleTicker() {
	if tickerGo {
		tickerGo = false
	} else {
		tickerGo = true
	}

}

//////

func actionButtonClick(g *gocui.Gui, v *gocui.View) error {
	err := executeClick(v)
	hexID := gr.tileByClick(mapCellX, mapCellY)
	lines := gr.tileMap[hexID].lines
	view, _ := g.View("Size")
	fmt.Fprintf(view, "ping")
	for i := range lines {
		fmt.Fprintf(view, lines[i])
	}
	return err
}

func actionSelectClick(g *gocui.Gui, v *gocui.View) error {
	err := executeClick(v)

	return err
}

func actionChangeColor(g *gocui.Gui, v *gocui.View) error {
	view, err := g.View(v.Name())
	bg := view.BgColor
	bg++
	if bg > 7 {
		bg = 0
	}
	view.BgColor = bg
	return err
}

func executeClick(v *gocui.View) error {
	switch v.Title {
	case "Info":
		mapCellXLast = mapCellX
		mapCellYLast = mapCellY
		oX, oY := v.Origin()
		clX, clY := v.Cursor()
		mapCellX = oX + clX
		mapCellY = oY + clY
		toggleTicker()
		fillPanel(v)

	default:

	}

	return nil
}

func actionMoveRight(g *gocui.Gui, v *gocui.View) error {
	view, err := g.View("Info")
	if err != nil {
		return err
	}
	voX, voY := view.Origin()
	view.SetOrigin(voX+2, voY)

	return err
}

func actionMoveLeft(g *gocui.Gui, v *gocui.View) error {
	view, err := g.View("Info")
	if err != nil {
		return err
	}
	voX, voY := view.Origin()
	view.SetOrigin(voX-2, voY)

	return err
}

func actionMoveUp(g *gocui.Gui, v *gocui.View) error {
	view, err := g.View("Info")
	if err != nil {
		return err
	}
	voX, voY := view.Origin()
	view.SetOrigin(voX, voY-1)

	return err
}

func actionMoveDown(g *gocui.Gui, v *gocui.View) error {
	view, err := g.View("Info")
	if err != nil {
		return err
	}
	voX, voY := view.Origin()
	view.SetOrigin(voX, voY+1)

	return err
}
