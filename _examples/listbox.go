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

	Grid(myList, Row(0), Column(0), Sticky(NEWS), Columnspan(3), style)
	Grid(label, Row(1), Column(0), Columnspan(3), style)
	Grid(TButton(Txt("Insert"), Command(func() {
		myList.Insert("end", time.Now().Format(time.RFC3339Nano))
	})), Row(2), Column(0), Columnspan(2), style)
	Grid(TButton(Txt("Delete"), Command(func() {
		myList.Delete("end")
	})), Row(3), Column(0), Columnspan(2), style)
	Grid(TExit(), Row(4), Column(0), Columnspan(3), style)

	// these two buttons demonstrate how to simulate the synthetic
	// <<ListboxSelect>> event programmatically.
	Grid(TButton(Txt("⬆️"), Command(func() {
		selected := myList.Curselection()
		if len(selected) == 0 {
			myList.SelectionSet("end")
		} else if selected[0] > 0 {
			// if the listbox were in "extended" mode,
			// and more than one element were selected,
			// this branch would not fire, and the selection
			// would grow upward in the list.
			if len(selected) == 1 {
				myList.SelectionClear("0", "end")
			}
			myList.SelectionSet(selected[0] - 1)
			myList.See(selected[0] - 1) // ensure the selection stays visible
		}
		label.Configure(Txt(myList.Get(myList.Curselection()[0])))
	})), Row(2), Column(2), Columnspan(1), style)

	Grid(TButton(Txt("⬇️"), Command(func() {
		selected := myList.Curselection()
		if len(selected) == 0 {
			myList.SelectionSet(0)
		} else if selected[0] < myList.Size()-1 {
			if len(selected) == 1 {
				myList.SelectionClear("0", "end")
			}
			myList.SelectionSet(selected[0] + 1)
			myList.See(selected[0] + 1)
		}
		label.Configure(Txt(myList.Get(myList.Curselection()[0])))
	})), Row(3), Column(2), Columnspan(1), style)

	myList.Insert(0, "Choice1", "Choice2", "Choice3")
	ActivateTheme("azure light")
	App.SetResizable(false, false)
	App.Wait()
}
