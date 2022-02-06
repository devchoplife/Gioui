package main

import (
	"gioui.org/app"
	"gioui.org/unit"
)

func main() {
	go func() {
		//Create new window
		w := app.NewWindow(
			app.Title("Egg Timer"),
			app.Size(unit.Dp(400), unit.Dp(600)),
		)

		//Listen for events in the window
		for range w.Events() {
		}
	}()

	app.Main()
}
