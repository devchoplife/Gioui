package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		//Create new window
		w := app.NewWindow(
			app.Title("Egg Timer"),
			app.Size(unit.Dp(400), unit.Dp(600)),
		)

		// ops are for operations on the ui
		var ops op.Ops

		// startButton is a clickable widget
		var startButton widget.Clickable

		// the defines the material design style
		th := material.NewTheme(gofont.Collection())

		//Listen for events in the window
		for e := range w.Events() {
			// detect what type of event
			switch e := e.(type) {
			// this is sent when the application shoudl re-render
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				btn := material.Button(th, &startButton, "Start")
				btn.Layout(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()

	app.Main()
}
