// Example using TextWidget.InsertML and html.Preview
//
// Inspired by https://gitlab.com/cznic/tk9.0/-/issues/68
package main

import (
	"io"
	"strings"

	"modernc.org/htmlview"
	. "modernc.org/tk9.0"
	_ "modernc.org/tk9.0/themes/azure"
)

type reader string

// Implements html.Config
func (r reader) URL2Reader(string) (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader(string(r))), nil
}

func main() {
	ActivateTheme("azure light")
	text := Text()
	text.TagConfigure("red", Foreground(Red))
	text.TagConfigure("yellow", Foreground(Yellow))

	text.InsertML("<red>Red.</red><yellow>NOT yellow</yellow><br>")
	// InsertML does not handle white space the same way HTML does.
	text.InsertML("<red>Red.</red> <yellow>Yellow</yellow><br>")
	text.InsertML("<red>Red.</red><yellow> Yellow</yellow><br>")

	// Sometimes html.Preview might be a better choice than InsertML. Note the
	// different white space handling and the necessity to not have whitespace
	// between the `<br>` and `<red>` tags, otherwise there will be a space
	// character rendered at the line beginning.
	pv := html.NewPreview(App, reader(`
<style>
	red {color: red}
	yellow {color: yellow}
</style>
<red>Red.</red><yellow>NOT yellow</yellow>
<br><red>Red.</red> <yellow>Yellow</yellow>
<br><red>Red.</red><yellow> Yellow</yellow><br>`))
	pv.SetWindowLocation("", nil, true)

	sty := Opts{Padx("2m"), Pady("2m"), Ipadx("1m"), Ipady("1m")}
	Grid(text, sty)
	Grid(pv, sty)
	Grid(TExit(), sty)
	App.SetResizable(false, false)
	App.Wait()
}
