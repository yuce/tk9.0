package main

import (
	"fmt"
	"strconv"

	. "modernc.org/tk9.0"
)

func parseGeometry(s string) (width int, height int, x int, y int) {
	coll := []rune{}
	seen := false
	for i, r := range s {
		switch r {
		case 'x':
			width, _ = strconv.Atoi(string(coll))
			coll = []rune{}
		case '+':
			if !seen {
				height, _ = strconv.Atoi(string(coll))
				coll = []rune{}
				seen = true
			} else {
				x, _ = strconv.Atoi(string(coll))
				y, _ = strconv.Atoi(s[i:])
				return
			}
		default:
			coll = append(coll, r)
		}
	}
	return
}

var counter = 0

func main() {
	otherWindow := Toplevel()
	otherWindow.Window.Configure(Padx("4m"), Pady("3m"))
	WmWithdraw(otherWindow.Window)

	Pack(otherWindow.TButton(Txt("Hide"), Command(func() {
		WmWithdraw(otherWindow.Window)
	})))
	Pack(App.TButton(Txt("Show"), Command(func() {
		h, w, x, y := parseGeometry(WmGeometry(App))
		switch counter % 4 {
		case 0:
			x = x + w*2
			y = y - h
		case 1:
			x = x - w*2
			y = y - h
		case 2:
			x = x - w*2
			y = y + h
		case 3:
			x = x + w*2
			y = y + h
		}
		WmGeometry(otherWindow.Window, fmt.Sprintf("+%d+%d", x, y))
		WmDeiconify(otherWindow.Window)
		counter++
	})))
	Bind(otherWindow, "<Destroy>", Command(func(e *Event) {
		// closed by user click on <x>
		Destroy(App)
	}))

	App.Wait()
}
