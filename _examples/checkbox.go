package main

import "fmt"
import . "modernc.org/tk9.0"
import _ "modernc.org/tk9.0/themes/azure"

func main() {
	ActivateTheme("azure light")
	menuBar := Menu()
	fileMenu := menuBar.Menu()
	bindVar := Variable("")
	checkBox := fileMenu.AddCheckbutton(Lbl("Check"))
	fileMenu.EntryConfigure(checkBox, bindVar)
	menuBar.AddCascade(Lbl("File"), Underline(0), Mnu(fileMenu))
	style := Opts{Padx("2m"), Pady("2m"), Ipadx("1m"), Ipady("1m")}
	eventLabel := TLabel(Txt("Event=\"\""))
	Bind(fileMenu, "<<MenuSelect>>", Command(func() {
		eventLabel.Configure(Txt(fmt.Sprintf("Event=%q", bindVar.Get())))
	}))
	Grid(eventLabel, style)
	readLabel := TLabel(Txt("Read=\"\""))
	Grid(TButton(Txt("Read variable"), Command(func() {
		readLabel.Configure(Txt(fmt.Sprintf("Read=%q", bindVar.Get())))
	})), style)
	Grid(readLabel, style)
	Grid(TExit(), style)
	App.SetResizable(false, false)
	App.Configure(Mnu(menuBar)).Wait()
}
