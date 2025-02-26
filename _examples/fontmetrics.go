package main

import (
	"fmt"

	. "modernc.org/tk9.0"
	_ "modernc.org/tk9.0/themes/azure"
)

func main() {
	for _, family := range []string{
		DefaultFont, TextFont, FixedFont, MenuFont, HeadingFont,
		CaptionFont, SmallCaptionFont, IconFont, TooltipFont,
	} {
		font := NewFont(Family(family))
		fmt.Printf("%20s : ascent=%s, descent=%s, linespace=%s, fixed=%s\n",
			family,
			font.Metrics(FontMetricAscent),
			font.Metrics(FontMetricDescent),
			font.Metrics(FontMetricLinespace),
			font.Metrics(FontMetricFixed),
		)
	}
}
