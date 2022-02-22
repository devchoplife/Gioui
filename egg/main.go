package main

import (
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions

var incrementProgress chan float32

// the draw function handles the layout
func draw(w *app.Window) error {
	var boiling bool
	var progress float32

	// ops are the operations from the UI
	var ops op.Ops

	// startButton is a clickable widget
	var startButton widget.Clickable

	// the defines the material design style
	th := material.NewTheme(gofont.Collection())

	//Listen for events in the window
	for {
		select {
		case e := <-w.Events():
			// detect what type of event it is
			switch e := e.(type) {
			// this is sent when the application should re-render
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				// Using the flex layout
				if startButton.Clicked() {
					boiling = !boiling
				}
				layout.Flex{
					// Vertical alignment, top to bottom
					Axis: layout.Vertical,
					// Empty space is left at the start i.e. at the top
					Spacing: layout.SpaceStart,
				}.Layout(gtx,
					// We insert rigid elements
					// Insert the progress bar
					layout.Rigid(
						func(gtx C) D {
							bar := material.ProgressBar(th, (progress)) // Progress is used here
							return bar.Layout(gtx)
						},
					),
					// First, a button
					layout.Rigid(
						func(gtx C) D {
							margins := layout.Inset{
								Top:    unit.Dp(25),
								Bottom: unit.Dp(25),
								Right:  unit.Dp(35),
								Left:   unit.Dp(35),
							}

							var text string
							if !boiling {
								text = "Start"
							} else {
								text = "Stop"
							}

							btn := material.Button(th, &startButton, text)
							return margins.Layout(gtx,
								func(gtx C) D {
									return btn.Layout(gtx)
								},
							)
						},
					),

					// then an empty spacer
					layout.Rigid(
						// The height of the spacer is 25 Device independent pixels
						layout.Spacer{Height: unit.Dp(25)}.Layout,
					),
				)
				e.Frame(gtx.Ops)

				// THis is sent when the application is closed
			case system.DestroyEvent:
				return e.Err
			}

		case p := <-incrementProgress:
			if boiling && progress < 1 {
				progress += p
				w.Invalidate()
			}
		}
	}
}

func main() {
	incrementProgress = make(chan float32)

	go func() {
		for {
			time.Sleep(time.Second / 25)
			incrementProgress <- 0.004
		}
	}()

	go func() {
		//Create new window
		win := app.NewWindow(
			app.Title("Egg Timer"),
			app.Size(unit.Dp(400), unit.Dp(600)),
		)

		if err := draw(win); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()

	app.Main()
}
