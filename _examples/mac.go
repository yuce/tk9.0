package main

import (
	. "modernc.org/tk9.0"
	_ "modernc.org/tk9.0/themes/azure"
)

var quit = false

func main() {
	t1 := Label(Txt("I haven't been hidden"))
	MacOnHide(func() {
		t1.Configure(Txt("I was hidden!"))
	})

	t2 := Label(Txt("I haven't been shown"))
	MacOnShow(func() {
		t2.Configure(Txt("I was shown!"))
	})

	t3 := Label(Txt("I haven't received a filename via Open With..."))
	MacOpenDocument(func(e string) {
		t2.Configure(Txt(e))
	})

	t4 := Label(Txt("I haven't been reopened"))
	MacReopenApplication(func() {
		t4.Configure(Txt("I was clicked on in the dock!"))
	})

	t5 := Label(Txt("I haven't had a Quit sent yet"))
	MacQuit(func() {
		t5.Configure(Txt("Oh, you tried to quit! Do it again to actually exit."))
		if quit {
			Destroy(App)
		} else {
			quit = true
		}
	})

	Grid(t1)
	Grid(t2)
	Grid(t3)
	Grid(t4)
	Grid(t5)
	ActivateTheme("azure light")
	App.SetResizable(false, false)
	App.Wait()
}
