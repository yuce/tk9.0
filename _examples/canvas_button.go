// Example: Making a button in a frame within canvas responsive.
//
// See also https://g.co/gemini/share/6ae2c7d481cd

package main

import . "modernc.org/tk9.0"

func main() {
	canvas := Canvas(Background(Red), Width(300), Height(200))
	Grid(canvas, Sticky(NEWS))
	frame := TFrame()
	button := frame.Button(Txt("Click me"), Command(func() { MessageBox(Detail("Clicked"), Type("ok")) }))
	Grid(button)
	canvas.CreateLine(0, 0, 100, 80, Fill(Blue))
	canvas.CreateText(150, 150, Txt("Canvas Text"))
	canvas.CreateWindow(100, 80, ItemWindow(frame.Window), Anchor("nw"))
	App.Wait()
}
