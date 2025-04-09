# tk9.0: The CGo-free, cross platform GUI toolkit for Go

[![LiberaPay](https://liberapay.com/assets/widgets/donate.svg)](https://liberapay.com/jnml/donate)
[![receives](https://img.shields.io/liberapay/receives/jnml.svg?logo=liberapay)](https://liberapay.com/jnml/donate)
[![patrons](https://img.shields.io/liberapay/patrons/jnml.svg?logo=liberapay)](https://liberapay.com/jnml/donate)

![photo](_examples/photo.png "photo")

Using Go embedded images (_examples/photo.go).

     1	package main
     2	
     3	import _ "embed"
     4	import . "modernc.org/tk9.0"
     5	import _ "modernc.org/tk9.0/themes/azure"
     6	
     7	//go:embed gopher.png
     8	var gopher []byte
     9	
    10	func main() {
    11		ActivateTheme("azure light")
    12		Pack(Label(Image(NewPhoto(Data(gopher)))),
    13			TExit(),
    14			Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
    15		App.Center().Wait()
    16	}

![menu](_examples/menu.png "menu")

Cascading menus (_examples/menu.go)

     1	package main
     2	
     3	import (
     4		"fmt"
     5		. "modernc.org/tk9.0"
     6		_ "modernc.org/tk9.0/themes/azure"
     7		"runtime"
     8	)
     9	
    10	func main() {
    11		menubar := Menu()
    12	
    13		fileMenu := menubar.Menu()
    14		fileMenu.AddCommand(Lbl("New"), Underline(0), Accelerator("Ctrl+N"))
    15		fileMenu.AddCommand(Lbl("Open..."), Underline(0), Accelerator("Ctrl+O"), Command(func() { GetOpenFile() }))
    16		Bind(App, "<Control-o>", Command(func() { fileMenu.Invoke(1) }))
    17		fileMenu.AddCommand(Lbl("Save"), Underline(0), Accelerator("Ctrl+S"))
    18		fileMenu.AddCommand(Lbl("Save As..."), Underline(5))
    19		fileMenu.AddCommand(Lbl("Close"), Underline(0), Accelerator("Crtl+W"))
    20		fileMenu.AddSeparator()
    21		fileMenu.AddCommand(Lbl("Exit"), Underline(1), Accelerator("Ctrl+Q"), ExitHandler())
    22		Bind(App, "<Control-q>", Command(func() { fileMenu.Invoke(6) }))
    23		menubar.AddCascade(Lbl("File"), Underline(0), Mnu(fileMenu))
    24	
    25		editMenu := menubar.Menu()
    26		editMenu.AddCommand(Lbl("Undo"))
    27		editMenu.AddSeparator()
    28		editMenu.AddCommand(Lbl("Cut"))
    29		editMenu.AddCommand(Lbl("Copy"))
    30		editMenu.AddCommand(Lbl("Paste"))
    31		editMenu.AddCommand(Lbl("Delete"))
    32		editMenu.AddCommand(Lbl("Select All"))
    33		menubar.AddCascade(Lbl("Edit"), Underline(0), Mnu(editMenu))
    34	
    35		helpMenu := menubar.Menu()
    36		helpMenu.AddCommand(Lbl("Help Index"))
    37		helpMenu.AddCommand(Lbl("About..."))
    38		menubar.AddCascade(Lbl("Help"), Underline(0), Mnu(helpMenu))
    39	
    40		App.WmTitle(fmt.Sprintf("%s on %s", App.WmTitle(""), runtime.GOOS))
    41		ActivateTheme("azure light")
    42		App.Configure(Mnu(menubar), Width("8c"), Height("6c")).Wait()
    43	}

Menus on darwin are now using the system-managed menu bar.    

![text](_examples/text.png "text")

Rich text using markup (_examples/text.go).

     1	package main
     2	
     3	import . "modernc.org/tk9.0"
     4	import _ "modernc.org/tk9.0/themes/azure"
     5	
     6	func main() {
     7		ActivateTheme("azure light")
     8		var scroll *TScrollbarWidget
     9		t := Text(Font("helvetica", 10), Yscrollcommand(func(e *Event) { e.ScrollSet(scroll) }), Setgrid(true), Wrap("word"), Padx("2m"), Pady("2m"))
    10		scroll = TScrollbar(Command(func(e *Event) { e.Yview(t) }))
    11		Grid(t, Sticky("news"), Pady("2m"))
    12		Grid(scroll, Row(0), Column(1), Sticky("nes"), Pady("2m"))
    13		GridRowConfigure(App, 0, Weight(1))
    14		GridColumnConfigure(App, 0, Weight(1))
    15		Grid(TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
    16		t.TagConfigure("bgstipple", Background(Black), Borderwidth(0), Bgstipple(Gray12))
    17		t.TagConfigure("big", Font("helvetica", 12, "bold"))
    18		t.TagConfigure("bold", Font("helvetica", 10, "bold", "italic"))
    19		t.TagConfigure("center", Justify("center"))
    20		t.TagConfigure("color1", Background("#a0b7ce"))
    21		t.TagConfigure("color2", Foreground(Red))
    22		t.TagConfigure("margins", Lmargin1("12m"), Lmargin2("6m"), Rmargin("10m"))
    23		t.TagConfigure("overstrike", Overstrike(1))
    24		t.TagConfigure("raised", Relief("raised"), Borderwidth(1))
    25		t.TagConfigure("right", Justify("right"))
    26		t.TagConfigure("spacing", Spacing1("10p"), Spacing2("2p"), Lmargin1("12m"), Lmargin2("6m"), Rmargin("10m"))
    27		t.TagConfigure("sub", Offset("-2p"), Font("helvetica", 8))
    28		t.TagConfigure("sunken", Relief("sunken"), Borderwidth(1))
    29		t.TagConfigure("super", Offset("4p"), Font("helvetica", 8))
    30		t.TagConfigure("tiny", Font("times", 8, "bold"))
    31		t.TagConfigure("underline", Underline(1))
    32		t.TagConfigure("verybig", Font(CourierFont(), 22, "bold"))
    33		t.InsertML(`Text widgets like this one allow you to display information in a variety of styles. Display styles are controlled
    34	using a mechanism called <bold>tags</bold>. Tags are just textual names that you can apply to one or more ranges of characters within a
    35	text widget. You can configure tags with various display styles. If you do this, then the tagged characters will be displayed with the
    36	styles you chose. The available display styles are:
    37	<br><br><big>1. Font.</big> You can choose any system font, <verybig>large</verybig> or <tiny>small</tiny>.
    38	<br><br><big>2. Color.</big> You can change either the <color1>background</color1> or <color2>foreground</color2> color, or
    39	<color1><color2>both</color2></color1>.
    40	<br><br><big>3. Stippling.</big> You can cause the <bgstipple>background</bgstipple> information to be drawn with a stipple fill instead
    41	of a solid fill.
    42	<br><br><big>4. Underlining.</big> You can <underline>underline</underline> ranges of text.
    43	<br><br><big>5. Overstrikes.</big> You can <overstrike>draw lines through</overstrike> ranges of text.
    44	<br><br><big>6. 3-D effects.</big> You can arrange for the background to be drawn with a border that makes characters appear either
    45	<raised>raised</raised> or <sunken>sunken</sunken>.
    46	<br><br><big>7. Justification.</big> You can arrange for lines to be displayed <br>left-justified <br><right>right-justified, or</right>
    47	<br><center>centered.</center>
    48	<br><br><big>8. Superscripts and subscripts.</big> You can control the vertical position of text to generate superscript effects like
    49	10<super>n</super> or subscript effects like X<sub>i</sub>.
    50	<br><br><big>9. Margins.</big> You can control the amount of extra space left on each side of the text
    51	<br><br><margins>This paragraph is an example of the use of margins. It consists of a single line of text that wraps around on the
    52	screen.  There are two separate left margin values, one for the first display line associated with the text line, and one for the
    53	subsequent display lines, which occur because of wrapping. There is also a separate specification for the right margin, which is used to
    54	choose wrap points for lines.</margins>
    55	<br><br><big>10. Spacing.</big> You can control the spacing of lines with three separate parameters. "Spacing1" tells how much extra
    56	space to leave above a line, "spacing3" tells how much space to leave below a line, and if a text line wraps, "spacing2" tells how much
    57	space to leave between the display lines that make up the text line.
    58	<br><spacing>These indented paragraphs illustrate how spacing can be used. Each paragraph is actually a single line in the text widget,
    59	which is word-wrapped by the widget.</spacing>
    60	<br><spacing>Spacing1 is set to 10 points for this text, which results in relatively large gaps between the paragraphs. Spacing2 is set
    61	to 2 points, which results in just a bit of extra space within a pararaph. Spacing3 isn't used in this example.</spacing>
    62	<br><spacing>To see where the space is, select ranges of text within these paragraphs. The selection highlight will cover the extra
    63	space.</spacing>`)
    64		App.Center().Wait()
    65	}

![svg](_examples/svg.png "svg")

Using svg (_examples/svg.go).

     1	package main
     2	
     3	import . "modernc.org/tk9.0"
     4	import _ "modernc.org/tk9.0/themes/azure"
     5	
     6	// https://en.wikipedia.org/wiki/SVG
     7	const svg = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
     8	<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
     9	<svg width="391" height="391" viewBox="-70.5 -70.5 391 391" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
    10	<rect fill="#fff" stroke="#000" x="-70" y="-70" width="390" height="390"/>
    11	<g opacity="0.8">
    12		<rect x="25" y="25" width="200" height="200" fill="lime" stroke-width="4" stroke="pink" />
    13		<circle cx="125" cy="125" r="75" fill="orange" />
    14		<polyline points="50,150 50,200 200,200 200,100" stroke="red" stroke-width="4" fill="none" />
    15		<line x1="50" y1="50" x2="200" y2="200" stroke="blue" stroke-width="4" />
    16	</g>
    17	</svg>`
    18	
    19	func main() {
    20		ActivateTheme("azure light")
    21		Pack(Label(Image(NewPhoto(Data(svg)))),
    22			TExit(),
    23			Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
    24		App.Center().Wait()
    25	}

![calc](_examples/calc.png "calc")

A simple calculator (_examples/calc.go).

     1	package main
     2	
     3	import "github.com/expr-lang/expr"
     4	import . "modernc.org/tk9.0"
     5	import _ "modernc.org/tk9.0/themes/azure"
     6	
     7	func main() {
     8		ActivateTheme("azure light")
     9		out := Label(Height(2), Anchor("e"), Txt("(123+232)/(123-10)"))
    10		Grid(out, Columnspan(4), Sticky("e"))
    11		var b *TButtonWidget
    12		for i, c := range "C()/789*456-123+0.=" {
    13			b = TButton(Txt(string(c)),
    14				Command(
    15					func() {
    16						switch c {
    17						case 'C':
    18							out.Configure(Txt(""))
    19						case '=':
    20							x, err := expr.Eval(out.Txt(), nil)
    21							if err != nil {
    22								MessageBox(Icon("error"), Msg(err.Error()), Title("Error"))
    23								x = ""
    24							}
    25							out.Configure(Txt(x))
    26						default:
    27							out.Configure(Txt(out.Txt() + string(c)))
    28						}
    29					},
    30				),
    31				Width(-4))
    32			Grid(b, Row(i/4+1), Column(i%4), Sticky("news"), Ipady("2.6m"), Padx("0.5m"), Pady("0.5m"))
    33		}
    34		Grid(b, Columnspan(2))
    35		b.Configure(Style("Accent.TButton"))
    36		App.Configure(Padx(0), Pady(0)).Wait()
    37	}

![font](_examples/font.png "font")

A font previewer (_examples/font.go).

     1	package main
     2	
     3	import "fmt"
     4	import "slices"
     5	import . "modernc.org/tk9.0"
     6	import _ "modernc.org/tk9.0/themes/azure"
     7	
     8	func main() {
     9		ActivateTheme("azure light")
    10		var scroll *TScrollbarWidget
    11		t := Text(Wrap("none"), Setgrid(true), Yscrollcommand(func(e *Event) { e.ScrollSet(scroll) }))
    12		scroll = TScrollbar(Command(func(e *Event) { e.Yview(t) }))
    13		fonts := FontFamilies()
    14		slices.Sort(fonts)
    15		Grid(t, Sticky("news"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
    16		Grid(scroll, Row(0), Column(1), Sticky("nes"), Pady("2m"))
    17		GridRowConfigure(App, 0, Weight(1))
    18		GridColumnConfigure(App, 0, Weight(1))
    19		Grid(TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
    20		m := map[string]bool{}
    21		for i, font := range fonts {
    22			if m[font] {
    23				continue
    24			}
    25			m[font] = true
    26			tag := fmt.Sprintf("t%v", i)
    27			t.TagConfigure(tag, Font(NewFont(Family(font))))
    28			t.Insert("end", font+": ", "", "Lorem ipsum dolor sit amet, consectetur adipiscing elit...\n", tag)
    29		}
    30		App.Center().Wait()
    31	}

![splot](_examples/splot.png "surface plot")

Surface plot (_examples/splot.go). This example requires Gnuplot 5.4+ installation.

     1	package main
     2	
     3	import . "modernc.org/tk9.0"
     4	import _ "modernc.org/tk9.0/themes/azure"
     5	
     6	var cm = int(TkScaling()*72/2.54 + 0.5)
     7	
     8	func main() {
     9		ActivateTheme("azure light")
    10		Pack(Label(Image(NewPhoto(Width(20*cm), Height(15*cm)).Graph("set grid; splot x**2+y**2, x**2-y**2"))),
    11			TExit(),
    12			Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
    13		App.Center().Wait()
    14	}

![tori](_examples/tori.png "interlocked tori")

Interlocked tori plot (_examples/tori.go). This example requires Gnuplot 5.4+ installation.

     1	package main
     2	
     3	import . "modernc.org/tk9.0"
     4	import _ "modernc.org/tk9.0/themes/azure"
     5	
     6	// https://gnuplot.sourceforge.net/demo_5.4/hidden2.html
     7	const script = `
     8	set multiplot title "Interlocking Tori"
     9	set title "PM3D surface\nno depth sorting"
    10	set parametric
    11	set urange [-pi:pi]
    12	set vrange [-pi:pi]
    13	set isosamples 50,20
    14	set origin -0.02,0.0
    15	set size 0.55, 0.9
    16	unset key
    17	unset xtics
    18	unset ytics
    19	unset ztics
    20	set border 0
    21	set view 60, 30, 1.5, 0.9
    22	unset colorbox
    23	set pm3d scansbackward
    24	splot cos(u)+.5*cos(u)*cos(v),sin(u)+.5*sin(u)*cos(v),.5*sin(v) with pm3d,1+cos(u)+.5*cos(u)*cos(v),.5*sin(v),sin(u)+.5*sin(u)*cos(v) with pm3d
    25	set title "PM3D surface\ndepth sorting"
    26	set origin 0.40,0.0
    27	set size 0.55, 0.9
    28	set colorbox vertical user origin 0.9, 0.15 size 0.02, 0.50
    29	set format cb "%.1f"
    30	set pm3d depthorder
    31	splot cos(u)+.5*cos(u)*cos(v),sin(u)+.5*sin(u)*cos(v),.5*sin(v) with pm3d,1+cos(u)+.5*cos(u)*cos(v),.5*sin(v),sin(u)+.5*sin(u)*cos(v) with pm3d
    32	unset multiplot`
    33	
    34	var cm = int(TkScaling()*72/2.54 + 0.5)
    35	
    36	func main() {
    37		ActivateTheme("azure light")
    38		Pack(Label(Image(NewPhoto(Width(20*cm), Height(15*cm)).Graph(script))),
    39			TExit(),
    40			Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
    41		App.Center().Wait()
    42	}

![tori-canvas](_examples/tori_canvas.png "interlocked tori on canvas")

Interlocked tori plot on canvas (_examples/tori_canvas.go). This example requires Gnuplot 5.4+ installation.

     1	package main
     2	
     3	import . "modernc.org/tk9.0"
     4	import _ "modernc.org/tk9.0/themes/azure"
     5	
     6	// https://gnuplot.sourceforge.net/demo_5.4/surface2.9.gnu
     7	const script = `
     8	set dummy u, v
     9	set key bmargin center horizontal Right noreverse enhanced autotitle nobox
    10	set parametric
    11	set view 50, 30, 1, 1
    12	set isosamples 50, 20
    13	set hidden3d back offset 1 trianglepattern 3 undefined 1 altdiagonal bentover
    14	set style data lines
    15	set xyplane relative 0
    16	set title "Interlocking Tori" 
    17	set grid
    18	set urange [ -3.14159 : 3.14159 ] noreverse nowriteback
    19	set vrange [ -3.14159 : 3.14159 ] noreverse nowriteback
    20	set xrange [ * : * ] noreverse writeback
    21	set x2range [ * : * ] noreverse writeback
    22	set yrange [ * : * ] noreverse writeback
    23	set y2range [ * : * ] noreverse writeback
    24	set zrange [ * : * ] noreverse writeback
    25	set cbrange [ * : * ] noreverse writeback
    26	set rrange [ * : * ] noreverse writeback
    27	set colorbox vertical origin screen 0.9, 0.2 size screen 0.05, 0.6 front  noinvert bdefault
    28	NO_ANIMATION = 1
    29	splot cos(u)+.5*cos(u)*cos(v),sin(u)+.5*sin(u)*cos(v),.5*sin(v) with lines,1+cos(u)+.5*cos(u)*cos(v),.5*sin(v),sin(u)+.5*sin(u)*cos(v) with lines`
    30	
    31	var cm = int(TkScaling()*72/2.54 + 0.5)
    32	
    33	func main() {
    34		Pack(Canvas(Width(20*cm), Height(15*cm), Background(White)).Graph(script),
    35			TExit(),
    36			Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
    37		ActivateTheme("azure light")
    38		App.Center().Wait()
    39	}

![tex](_examples/tex.png "TeX")

Rendering plain TeX (_examples/tex.go). No runtime dependencies required.

     1	package main
     2	
     3	import . "modernc.org/tk9.0"
     4	import _ "modernc.org/tk9.0/themes/azure"
     5	
     6	func main() {
     7		tex := `$$\int _0 ^\infty {{\sin ax \sin bx}\over{x^2}}\,dx = {\pi a\over 2}$$`
     8		Pack(Label(Relief("sunken"), Image(NewPhoto(Data(TeX(tex, 2*TkScaling()*72/600))))),
     9			TExit(),
    10			Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
    11		ActivateTheme("azure light")
    12		App.Center().Wait()
    13	}

![embed](_examples/embed.png "embed")

Embedding pictures, TeX and other widgets in Text (_examples/embed.go).

     1	package main
     2	
     3	import . "modernc.org/tk9.0"
     4	import _ "embed"
     5	import _ "modernc.org/tk9.0/themes/azure"
     6	
     7	//go:embed gotk.png
     8	var icon []byte
     9	
    10	func main() {
    11		fontSize := int(10*TkScaling()/NativeScaling + 0.5)
    12		font := Font("helvetica", fontSize)
    13		var scroll *TScrollbarWidget
    14		t := Text(font, Height(22), Yscrollcommand(func(e *Event) { e.ScrollSet(scroll) }), Setgrid(true), Wrap("word"),
    15			Padx("4p"), Pady("12p"))
    16		scroll = TScrollbar(Command(func(e *Event) { e.Yview(t) }))
    17		Grid(t, Sticky("news"), Pady("2m"))
    18		Grid(scroll, Row(0), Column(1), Sticky("nes"), Pady("2m"))
    19		GridRowConfigure(App, 0, Weight(1))
    20		GridColumnConfigure(App, 0, Weight(1))
    21		Grid(TExit())
    22		t.TagConfigure("c", Justify("center"))
    23		t.TagConfigure("e", Offset("-2p"))
    24		t.TagConfigure("t", Font("times", fontSize))
    25		sym := " <t>T<e>E</e>X</t> "
    26		tex := `$Q(\xi) = \lambda_1 y_1^2 \sum_{i=2}^n \sum_{j=2}^n y_i b_{ij} y_j$`
    27		t.InsertML(`<c>Hello Go + Tk`, NewPhoto(Data(icon)), Padx("4p"), `users!
    28	<br><br>Hello Go + Tk +`, sym, tex, ` users! (\$inline math\$)
    29	<br><br>Hello Go + Tk +`, sym, `$`+tex+`$`, ` users! (\$\$display math\$\$)</c>
    30	<br><br>The above exemplifies embeding pictures and`, sym, `scripts. A text widget can also embed other widgets. For example,
    31	when a`, TButton(Txt("<Tbutton>")), Padx("4p"), Pady("2p"), Align("center"), `and
    32	a`, TEntry(Textvariable("<TEntry>"), Background(White), Width(8)), Padx("4p"), Pady("2p"), Align("center"), `are part of
    33	the markup, they will reflow when their containing text widget is resized.`)
    34		ActivateTheme("azure light")
    35		App.Center().Wait()
    36	}

The above screen shot is from '$ TK9_SCALE=1.2 go run _examples/embed.go'.

![tbutton](_examples/tbutton.png "tbutton")

Styling a button (_examples/tbutton.go). See the discussion at [Tutorial: Modifying a ttk button's style]

     1	package main
     2	
     3	import _ "embed"
     4	import . "modernc.org/tk9.0"
     5	
     6	//go:embed red_corner.png
     7	var red []byte
     8	
     9	//go:embed green_corner.png
    10	var green []byte
    11	
    12	func main() {
    13		StyleElementCreate("Red.Corner.TButton.indicator", "image", NewPhoto(Data(red)))
    14		StyleElementCreate("Green.Corner.TButton.indicator", "image", NewPhoto(Data(green)))
    15		StyleLayout("Red.Corner.TButton",
    16			"Button.border", Sticky("nswe"), Border(1), Children(
    17				"Button.focus", Sticky("nswe"), Children(
    18					"Button.padding", Sticky("nswe"), Children(
    19						"Button.label", Sticky("nswe"),
    20						"Red.Corner.TButton.indicator", Side("right"), Sticky("ne")))))
    21		StyleLayout("Green.Corner.TButton",
    22			"Button.border", Sticky("nswe"), Border(1), Children(
    23				"Button.focus", Sticky("nswe"), Children(
    24					"Button.padding", Sticky("nswe"), Children(
    25						"Button.label", Sticky("nswe"),
    26						"Green.Corner.TButton.indicator", Side("right"), Sticky("ne")))))
    27		opts := Opts{Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m")}
    28		rb := TButton(Txt("Red"))
    29		gb := TButton(Txt("Green"))
    30		Grid(rb, gb, opts)
    31		Grid(TButton(Txt("Use style"), Command(func() {
    32			rb.Configure(Style("Red.Corner.TButton"))
    33			gb.Configure(Style("Green.Corner.TButton"))
    34		})), TExit(), opts)
    35		App.Wait()
    36	}


![azure](_examples/azure.png "azure")

Example usage of the theme register. (_examples/azure.go)

     1	package main
     2	
     3	import . "modernc.org/tk9.0"
     4	import _ "modernc.org/tk9.0/themes/azure"
     5	
     6	func main() {
     7		Pack(TButton(Txt("Light"), Command(func() { ActivateTheme("azure light") })),
     8			TButton(Txt("Dark"), Command(func() { ActivateTheme("azure dark") })),
     9			TExit(),
    10			Pady("2m"), Ipady("1m"))
    11		App.Wait()
    12	}

![b5](_examples/b5.png "b5")

Technology preview of a Bootstrap 5-like theme buttons (_examples/b5.go). Only
a partial prototype/problem study/work in progress at the moment. But it may
get there, eventually. (_examples/b5.go)

     1	package main
     2	
     3	import (
     4		. "modernc.org/tk9.0"
     5		"modernc.org/tk9.0/b5"
     6	)
     7	
     8	func main() {
     9		background := White
    10		primary := b5.Colors{b5.ButtonText: "#fff", b5.ButtonFace: "#0d6efd", b5.ButtonFocus: "#98c1fe"}
    11		secondary := b5.Colors{b5.ButtonText: "#fff", b5.ButtonFace: "#6c757d", b5.ButtonFocus: "#c0c4c8"}
    12		success := b5.Colors{b5.ButtonText: "#fff", b5.ButtonFace: "#198754", b5.ButtonFocus: "#9dccb6"}
    13		danger := b5.Colors{b5.ButtonText: "#fff", b5.ButtonFace: "#dc3545", b5.ButtonFocus: "#f0a9b0"}
    14		warning := b5.Colors{b5.ButtonText: "#000", b5.ButtonFace: "#ffc107", b5.ButtonFocus: "#ecd182"}
    15		info := b5.Colors{b5.ButtonText: "#000", b5.ButtonFace: "#0dcaf0", b5.ButtonFocus: "#85d5e5"}
    16		light := b5.Colors{b5.ButtonText: "#000", b5.ButtonFace: "#f8f9fa", b5.ButtonFocus: "#e9e9ea"}
    17		dark := b5.Colors{b5.ButtonText: "#fff", b5.ButtonFace: "#212529", b5.ButtonFocus: "#a0a2a4"}
    18		link := b5.Colors{b5.ButtonText: "#1774fd", b5.ButtonFace: "#fff", b5.ButtonFocus: "#c2dbfe"}
    19		StyleThemeUse("default")
    20		opts := Opts{Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m")}
    21		Grid(TButton(Txt("Primary"), Style(b5.ButtonStyle("primary.TButton", primary, background, false))),
    22			TButton(Txt("Secondary"), Style(b5.ButtonStyle("secondary.TButton", secondary, background, false))),
    23			TButton(Txt("Success"), Style(b5.ButtonStyle("success.TButton", success, background, false))),
    24			opts)
    25		Grid(TButton(Txt("Danger"), Style(b5.ButtonStyle("danger.TButton", danger, background, false))),
    26			TButton(Txt("Warning"), Style(b5.ButtonStyle("warning.TButton", warning, background, false))),
    27			TButton(Txt("Info"), Style(b5.ButtonStyle("info.TButton", info, background, false))),
    28			opts)
    29		Grid(TButton(Txt("Light"), Style(b5.ButtonStyle("light.TButton", light, background, false))),
    30			TButton(Txt("Dark"), Style(b5.ButtonStyle("dark.TButton", dark, background, false))),
    31			TButton(Txt("Link"), Style(b5.ButtonStyle("link.TButton", link, background, false))),
    32			opts)
    33		Grid(TButton(Txt("Primary"), Style(b5.ButtonStyle("focused.primary.TButton", primary, background, true))),
    34			TButton(Txt("Secondary"), Style(b5.ButtonStyle("focused.secondary.TButton", secondary, background, true))),
    35			TButton(Txt("Success"), Style(b5.ButtonStyle("focused.success.TButton", success, background, true))),
    36			opts)
    37		Grid(TButton(Txt("Danger"), Style(b5.ButtonStyle("focused.danger.TButton", danger, background, true))),
    38			TButton(Txt("Warning"), Style(b5.ButtonStyle("focused.warning.TButton", warning, background, true))),
    39			TButton(Txt("Info"), Style(b5.ButtonStyle("focused.info.TButton", info, background, true))),
    40			opts)
    41		Grid(TButton(Txt("Light"), Style(b5.ButtonStyle("focused.light.TButton", light, background, true))),
    42			TButton(Txt("Dark"), Style(b5.ButtonStyle("focused.dark.TButton", dark, background, true))),
    43			TButton(Txt("Link"), Style(b5.ButtonStyle("focused.link.TButton", link, background, true))),
    44			opts)
    45		Grid(TExit(), Columnspan(3), opts)
    46		App.Configure(Background(background)).Wait()
    47	}

----

Gallery (_examples/demo.go)

Darwin(macOS) Sequoia 15.0

![darwin](_examples/darwin.png "darwin")

FreeBSD Xfce4

![freebsd](_examples/freebsd.png "freebsd")

Linux Mate 1.26.0

![linux](_examples/linux.png "linux")

Windows 11

![windows11](_examples/windows11.png "windows11")

[![Go Reference](https://pkg.go.dev/badge/modernc.org/tk9.0.svg)](https://pkg.go.dev/modernc.org/tk9.0)

[Tutorial: Modifying a ttk button's style]: https://wiki.tcl-lang.org/page/Tutorial%3A+Modifying+a+ttk+button%27s+style
