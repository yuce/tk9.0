// Adapted from https://www.tutorialspoint.com/tcl-tk/tk_listbox_widget.htm
package main

import (
	"time"

	. "modernc.org/tk9.0"
	_ "modernc.org/tk9.0/themes/azure"
)

func main() {
	myList := Listbox()
	label := TLabel(Txt("No choice selected"))
	Bind(myList, "<<ListboxSelect>>", Command(func() {
		label.Configure(Txt(myList.Get(myList.Curselection()[0])))
	}))
	style := Opts{Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m")}
	Grid(myList, Row(0), Column(0), Sticky(NEWS), Columnspan(2), style)
	Grid(label, Row(1), Column(0), Columnspan(2), style)
	Grid(TButton(Txt("Insert"), Command(func() {
		myList.Insert("end", time.Now().Format(time.RFC3339Nano))
	})), Row(2), Column(0), Columnspan(2), style)
	Grid(TButton(Txt("Delete"), Command(func() {
		myList.Delete("end")
	})), Row(3), Column(0), Columnspan(2), style)
	Grid(TExit(), Row(4), Column(0), Columnspan(2), style)
	myList.Insert(0, "Choice1", "Choice2", "Choice3")
	ActivateTheme("azure light")
	App.SetResizable(false, false)
	App.Wait()
}
