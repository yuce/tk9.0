package main

import . "modernc.org/tk9.0"
import . "modernc.org/tk9.0/extensions/eval"
import _ "modernc.org/tk9.0/themes/azure"

func main() {
	ActivateTheme("azure light")
	InitializeExtension("eval")
	Eval(`
proc nextFrame {image {index 0}} {
    if {[catch {tmpimg configure -format "gif -index $index"} stderr]} {
        set nextIndex 0
        set time 1
    } else {
        set nextIndex [expr {$index + 1}]
        set metadata [tmpimg cget -metadata]
        if {    [dict exists $metadata "disposal method"]
                && [dict get $metadata "disposal method"] eq "do not dispose"
        } {
            $image copy tmpimg -compositingrule overlay
        } else {
            $image copy tmpimg -compositingrule set
        }
        if {[dict exists $metadata "delay time"]} {
            set time [expr {[dict get $metadata "delay time"]*10}]
        } else {
            set time 1
        }
    }
    after $time nextFrame $image $nextIndex
}

set img [image create photo -file "hexahedron.gif"]
# Create a helper image
image create photo tmpimg -file [$img cget -file]

label .w -image $img
grid .w

nextFrame $img
`)
	sty := Opts{Padx("2m"), Pady("2m"), Ipadx("1m"), Ipady("1m")}
	Grid(TExit(), sty)
	App.SetResizable(false, false)
	App.Wait()
}
