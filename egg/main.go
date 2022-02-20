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
		win := app.NewWindow(
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
		for e := range win.Events() {
			// detect what type of event it is
			switch e := e.(type) {
			// this is sent when the application should re-render
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				// Using the flex layout
				layout.Flex{
					// Vertical alignment, top to bottom
					Axis: layout.Vertical,
					// Empty space is left at the start i.e. at the top
					Spacing: layout.SpaceStart,
				}.Layout(gtx,
					// We insert tow rigid elements
					// First, a button
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(th, &startButton, "Start")
							return btn.Layout(gtx)
						},
					),

					// then an empty spacer
					layout.Rigid(
						// The height of the spacer is 25 Device independent pixels
						layout.Spacer{Height: unit.Dp(25)}.Layout,
					),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()

	app.Main()
}
