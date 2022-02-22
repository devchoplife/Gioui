package main

import (
	"image"
	"image/color"
	"log"
	"math"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
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
							// Draw a custom path, shaped like an egg
							var eggPath clip.Path
							op.Offset(f32.Pt(200, 150)).Add(gtx.Ops)
							eggPath.Begin(gtx.Ops)
							// Rotate from 0 to 360 degrees
							for deg := 0.0; deg <= 360; deg++ {

								// Egg math (really) at this brilliant site. Thanks!
								// https://observablehq.com/@toja/egg-curve
								// Convert degrees to radians
								rad := deg / 360 * 2 * math.Pi
								// Trig gives the distance in X and Y direction
								cosT := math.Cos(rad)
								sinT := math.Sin(rad)
								// Constants to define the eggshape
								a := 110.0
								b := 150.0
								d := 20.0
								// The x/y coordinates
								x := a * cosT
								y := -(math.Sqrt(b*b-d*d*cosT*cosT) + d*sinT) * sinT
								// Finally the point on the outline
								p := f32.Pt(float32(x), float32(y))
								// Draw the line to this point
								eggPath.LineTo(p)
							}
							// Close the path
							eggPath.Close()

							// Get hold of the actual clip
							eggArea := clip.Outline{Path: eggPath.End()}.Op()

							// Fill the shape
							// color := color.NRGBA{R: 255, G: 239, B: 174, A: 255}
							color := color.NRGBA{R: 255, G: uint8(239 * (1 - progress)), B: uint8(174 * (1 - progress)), A: 255}
							paint.FillShape(gtx.Ops, color, eggArea)

							d := image.Point{Y: 375}
							return layout.Dimensions{Size: d}
						},
					),
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
