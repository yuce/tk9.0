package main

import "fmt"
import . "modernc.org/tk9.0"
import _ "modernc.org/tk9.0/themes/azure"

func main() {
	ActivateTheme("azure light")
	sty := Opts{Padx("2m"), Pady("2m"), Ipadx("1m"), Ipady("1m")}
	lbl := TLabel()
	Grid(lbl)
	Grid(TExit(), sty)
	lbl.Configure(Txt(fmt.Sprint(GridContent(App))))
	App.SetResizable(false, false)
	App.Wait()
}
