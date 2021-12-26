package main

import "gioui.org/app"

func main() {
	go func() {
		//Create new window
		w := app.NewWindow()

		//Listen for events in the window
		for range w.Events() {
		}
	}()

	app.Main()
}
