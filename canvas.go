// Copyright 2025 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

// Graph — use gnuplot to draw on a canvas. Graph returns 'w'.
//
// The 'script' argument is passed to a gnuplot executable, which must be
// installed on the machine.  See the [gnuplot site] for documentation about
// producing graphs. The script must not use the 'set term <device>' command.
//
// [gnuplot site]: http://www.gnuplot.info/
func (w *CanvasWidget) Graph(script string) *CanvasWidget {
	script = fmt.Sprintf("set terminal tkcanvas size %s, %s\n%s", w.Width(), w.Height(), script)
	out, err := gnuplot(script)
	if err != nil {
		fail(fmt.Errorf("plot: executing script: %s", err))
		return w
	}

	evalErr(fmt.Sprintf("%s\ngnuplot %s", out, w))
	return w
}

func gnuplot(script string) (out []byte, err error) {
	f, err := os.CreateTemp("", "tk9.0-")
	if err != nil {
		return nil, err
	}

	defer os.Remove(f.Name())

	if err := os.WriteFile(f.Name(), []byte(script), 0o660); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), gnuplotTimeout)

	defer cancel()

	return exec.CommandContext(ctx, "gnuplot", f.Name()).Output()
}

func (w *CanvasWidget) create(typ string, args ...any) string {
	return evalErr(fmt.Sprintf("%s create %s %s", w, typ, collectAny(args...)))
}

// Canvas — Create and manipulate 'canvas' hypergraphics drawing surface widgets
//
// # Description
//
// Create a new arc in pathName of type type. The exact format of the
// arguments after type depends on type, but usually they consist of the
// coordinates for one or more points, followed by specifications for zero or
// more item options. See the subsections on individual item types below for
// more on the syntax of this command. This command returns the id for the new
// item.
//
// Items of type arc appear on the display as arc-shaped regions. An arc is a
// section of an oval delimited by two angles (specified by either the -start
// and -extent options or the -height option) and displayed in one of several
// ways (specified by the -style option). Arcs are created with widget commands
// of the following form:
//
//	pathName create arc x1 y1 x2 y2 ?option value ...?
//
// The arguments x1, y1, x2, and y2 or coordList give the coordinates of two
// diagonally opposite corners of a rectangular region enclosing the oval that
// defines the arc (except when -height is specified - see below). After the
// coordinates there may be any number of option-value pairs, each of which
// sets one of the configuration options for the item. These same option-value
// pairs may be used in itemconfigure widget commands to change the item's
// configuration. An arc item becomes the current item when the mouse pointer
// is over any part that is painted or (when fully transparent) that would be
// painted if both the -fill and -outline options were non-empty.
//
// The following standard options are supported by arcs:
//
//   - [Dash]
//   - [Activedash]
//   - [Disableddash]
//   - [Dashoffset]
//   - [Fill]
//   - [Activefill]
//   - [Disabledfill]
//   - [Offset]
//   - [Outline]
//   - [Activeoutline]
//   - [Disabledoutline]
//   - [Outlineoffset]
//   - [Outlinestipple]
//   - [Activeoutlinestipple]
//   - [Disabledoutlinestipple]
//   - [Stipple]
//   - [Activestipple]
//   - [Disabledstipple]
//   - [State]
//   - [Tags]
//   - [Width]
//   - [Activewidth]
//   - [Disabledwidth]
//
// The following extra options are supported for arcs:
//
//   - [Extent] degrees
//
//   - [Start] degrees
//
//   - [Height] distance
//
// Provides a shortcut for creating a circular arc segment by defining the
// distance of the mid-point of the arc from its chord. When this option is
// used the coordinates are interpreted as the start and end coordinates of
// the chord, and the options -start and -extent are ignored. The value of
// distance has the following meaning:
//
//	distance > 0 creates a clockwise arc,
//	distance < 0 creates an counter-clockwise arc,
//	distance = 0 creates an arc as if this option had not been specified.
//
// If you want the arc to have a specific radius, r, use the formula:
//
// distance = r ± sqrt(r**2 - (chordLength / 2)**2)
//
// choosing the minus sign for the minor arc and the plus sign for the major
// arc.
//
// Note that itemcget -height always returns 0 so that introspection code can
// be kept simple.
//
//   - [Style] type
//
// Specifies how to draw the arc. If type is pieslice (the default) then the
// arc's region is defined by a section of the oval's perimeter plus two line
// segments, one between the center of the oval and each end of the perimeter
// section. If type is chord then the arc's region is defined by a section of
// the oval's perimeter plus a single line segment connecting the two end
// points of the perimeter section. If type is arc then the arc's region
// consists of a section of the perimeter alone. In this last case the -fill
// option is ignored.
//
// More information might be available at the [Tcl/Tk canvas] page.
//
// [Tcl/Tk canvas]: https://www.tcl.tk/man/tcl9.0/TkCmd/canvas.html
func (w *CanvasWidget) CreateArc(x1, y1, x2, y2 any, options ...any) (r string) {
	return w.create("arc", append([]any{x1, y1, x2, y2}, options...)...)
}

// Dash option.
//
// This option specifies the dash pattern for the normal state of an item. If
// the dash option is omitted then the default is a solid outline.
//
// Many items support the notion of a dash pattern for outlines.  The syntax is
// a list of integers. Each element represents the number of pixels of a line
// segment. Only the odd segments are drawn using the “outline” color. The
// other segments are drawn transparent.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Dash(pattern ...int) Opt {
	list := fmt.Sprint(pattern)
	return rawOption(fmt.Sprintf(`-dash {%s}`, list[1:len(list)-1]))
}

// Activedash option.
//
// This option specifies the dash pattern for the active state of an item. If
// the dash option is omitted then the default is a solid outline.
//
// Many items support the notion of a dash pattern for outlines.  The syntax is
// a list of integers. Each element represents the number of pixels of a line
// segment. Only the odd segments are drawn using the “outline” color. The
// other segments are drawn transparent.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Activedash(pattern ...int) Opt {
	list := fmt.Sprint(pattern)
	return rawOption(fmt.Sprintf(`-activedash {%s}`, list[1:len(list)-1]))
}

// Disableddash option.
//
// This option specifies the dash pattern for the disabled state of an item. If
// the dash option is omitted then the default is a solid outline.
//
// Many items support the notion of a dash pattern for outlines.  The syntax is
// a list of integers. Each element represents the number of pixels of a line
// segment. Only the odd segments are drawn using the “outline” color. The
// other segments are drawn transparent.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Disableddash(pattern ...int) Opt {
	list := fmt.Sprint(pattern)
	return rawOption(fmt.Sprintf(`-disableddash {%s}`, list[1:len(list)-1]))
}

// Dashoffset option.
//
// The starting offset in pixels into the pattern provided by the -dash option.
// -dashoffset is ignored if there is no -dash pattern.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Dashoffset(offset any) Opt {
	return rawOption(fmt.Sprintf(`-dashoffset %s`, tclSafeString(fmt.Sprint(offset))))
}

// Activefill option.
//
// This option specifies the color to be used to fill item's area in its active
// state. The even-odd fill rule is used. Color may have any of the forms
// accepted by Tk_GetColor. For the line item, it specifies the color of the
// line drawn. For the text item, it specifies the foreground color of the
// text. If color is an empty string (the default for all canvas items except
// line and text), then the item will not be filled.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Activefill(color any) Opt {
	return rawOption(fmt.Sprintf(`-activefill %s`, tclSafeString(fmt.Sprint(color))))
}

// Disabledfill option.
//
// This option specifies the color to be used to fill item's area in its disabled
// state. The even-odd fill rule is used. Color may have any of the forms
// accepted by Tk_GetColor. For the line item, it specifies the color of the
// line drawn. For the text item, it specifies the foreground color of the
// text. If color is an empty string (the default for all canvas items except
// line and text), then the item will not be filled.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Disabledfill(color any) Opt {
	return rawOption(fmt.Sprintf(`-disabledfill %s`, tclSafeString(fmt.Sprint(color))))
}

// Outline option.
//
// This option specifies the color that should be used to draw the outline of
// the item in its normal state. Color may have any of the forms accepted by
// Tk_GetColor. If color is specified as an empty string then no outline is
// drawn for the item.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Outline(color any) Opt {
	return rawOption(fmt.Sprintf(`-outline %s`, tclSafeString(fmt.Sprint(color))))
}

// Activeoutline option.
//
// This option specifies the color that should be used to draw the outline of
// the item in its active state. Color may have any of the forms accepted by
// Tk_GetColor. If color is specified as an empty string then no outline is
// drawn for the item.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Activeoutline(color any) Opt {
	return rawOption(fmt.Sprintf(`-activeoutline %s`, tclSafeString(fmt.Sprint(color))))
}

// Disabledline option.
//
// This option specifies the color that should be used to draw the outline of
// the item in its disabled state. Color may have any of the forms accepted by
// Tk_GetColor. If color is specified as an empty string then no outline is
// drawn for the item.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Disabledoutline(color any) Opt {
	return rawOption(fmt.Sprintf(`-disabledoutline %s`, tclSafeString(fmt.Sprint(color))))
}

// Outlineoffset option.
//
// Specifies the offset of the stipple pattern used for outlines, in the same
// way that the -outline option controls fill stipples. (See the -outline
// option for a description of the syntax of offset.)
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Outlineoffset(offset any) Opt {
	return rawOption(fmt.Sprintf(`-outlineoffset %s`, tclSafeString(fmt.Sprint(offset))))
}

// Outlinestipple option.
//
// This option specifies the stipple pattern that should be used to draw the
// outline of the item in its normal state. Indicates
// that the outline for the item should be drawn with a stipple pattern; bitmap
// specifies the stipple pattern to use, in any of the forms accepted by
// Tk_GetBitmap. If the -outline option has not been specified then this option
// has no effect. If bitmap is an empty string (the default), then the outline
// is drawn in a solid fashion. Note that stipples are not well supported on
// platforms that do not use X11 as their drawing API.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Outlinestipple(bitmap any) Opt {
	return rawOption(fmt.Sprintf(`-outlinstipple %s`, tclSafeString(fmt.Sprint(bitmap))))
}

// Activeoutlinestipple option.
//
// This option specifies the stipple pattern that should be used to draw the
// outline of the item in its active state. Indicates
// that the outline for the item should be drawn with a stipple pattern; bitmap
// specifies the stipple pattern to use, in any of the forms accepted by
// Tk_GetBitmap. If the -outline option has not been specified then this option
// has no effect. If bitmap is an empty string (the default), then the outline
// is drawn in a solid fashion. Note that stipples are not well supported on
// platforms that do not use X11 as their drawing API.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Activeoutlinestipple(bitmap any) Opt {
	return rawOption(fmt.Sprintf(`-activeoutlinestipple %s`, tclSafeString(fmt.Sprint(bitmap))))
}

// Disabledoutlinestipple option.
//
// This option specifies the stipple pattern that should be used to draw the
// outline of the item in its disabled state. Indicates
// that the outline for the item should be drawn with a stipple pattern; bitmap
// specifies the stipple pattern to use, in any of the forms accepted by
// Tk_GetBitmap. If the -outline option has not been specified then this option
// has no effect. If bitmap is an empty string (the default), then the outline
// is drawn in a solid fashion. Note that stipples are not well supported on
// platforms that do not use X11 as their drawing API.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Disabledoutlinestipple(bitmap any) Opt {
	return rawOption(fmt.Sprintf(`-disabledoutlinestipple %s`, tclSafeString(fmt.Sprint(bitmap))))
}

// Stipple option.
//
// This option specifies the stipple patterns that should be used to fill the item in
// its normal state. bitmap specifies the stipple pattern
// to use, in any of the forms accepted by Tk_GetBitmap. If the -fill option
// has not been specified then this option has no effect. If bitmap is an empty
// string (the default), then filling is done in a solid fashion. For the text
// item, it affects the actual text. Note that stipples are not well supported
// on platforms that do not use X11 as their drawing API.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
//   - [CanvasWidget.CreateBitmap] (widget specific)
func Stipple(bitmap any) Opt {
	return rawOption(fmt.Sprintf(`-stipple %s`, tclSafeString(fmt.Sprint(bitmap))))
}

// Activestipple option.
//
// This option specifies the stipple patterns that should be used to fill the item in
// its active state. bitmap specifies the stipple pattern
// to use, in any of the forms accepted by Tk_GetBitmap. If the -fill option
// has not been specified then this option has no effect. If bitmap is an empty
// string (the default), then filling is done in a solid fashion. For the text
// item, it affects the actual text. Note that stipples are not well supported
// on platforms that do not use X11 as their drawing API.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Activestipple(bitmap any) Opt {
	return rawOption(fmt.Sprintf(`-activestipple %s`, tclSafeString(fmt.Sprint(bitmap))))
}

// Disabledstipple option.
//
// This option specifies the stipple patterns that should be used to fill the item in
// its disabled state. bitmap specifies the stipple pattern
// to use, in any of the forms accepted by Tk_GetBitmap. If the -fill option
// has not been specified then this option has no effect. If bitmap is an empty
// string (the default), then filling is done in a solid fashion. For the text
// item, it affects the actual text. Note that stipples are not well supported
// on platforms that do not use X11 as their drawing API.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Disabledstipple(bitmap any) Opt {
	return rawOption(fmt.Sprintf(`-disabledstipple %s`, tclSafeString(fmt.Sprint(bitmap))))
}

// Tags option.
//
// Specifies a set of tags to apply to the item. TagList consists of a list of
// tag names, which replace any existing tags for the item. TagList may be an
// empty list.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
//   - [CanvasWidget.CreateBitmap] (widget specific)
func Tags(tagList ...string) Opt {
	list := fmt.Sprint(tagList)
	return rawOption(fmt.Sprintf(`-tags {%s}`, list[1:len(list)-1]))
}

// Activewidth option.
//
// This option specifies the width of the outline to be drawn around the item's
// region, in its active state. outlineWidth may be in
// any of the forms described in the COORDINATES section above. If the -outline
// option has been specified as an empty string then this option has no effect.
// This option defaults to 1.0. For arcs, wide outlines will be drawn centered
// on the edges of the arc's region.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Activewidth(outlineWidth any) Opt {
	return rawOption(fmt.Sprintf(`-activewidth %s`, tclSafeString(fmt.Sprint(outlineWidth))))
}

// Disabledwidth option.
//
// This option specifies the width of the outline to be drawn around the item's
// region, in its disabled state. outlineWidth may be in
// any of the forms described in the COORDINATES section above. If the -outline
// option has been specified as an empty string then this option has no effect.
// This option defaults to 1.0. For arcs, wide outlines will be drawn centered
// on the edges of the arc's region.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Disabledwidth(outlineWidth any) Opt {
	return rawOption(fmt.Sprintf(`-disabledwidth %s`, tclSafeString(fmt.Sprint(outlineWidth))))
}

// Extent option.
//
// Specifies the size of the angular range occupied by the arc. The arc's
// range extends for degrees degrees counter-clockwise from the starting
// angle given by the -start option. Degrees may be negative. If it is
// greater than 360 or less than -360, then degrees modulo 360 is used as
// the extent.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Extent(degrees any) Opt {
	return rawOption(fmt.Sprintf(`-extent %s`, tclSafeString(fmt.Sprint(degrees))))
}

// Start option.
//
// Specifies the beginning of the angular range occupied by the arc. Degrees
// is given in units of degrees measured counter-clockwise from the
// 3-o'clock position; it may be either positive or negative.
//
// Known uses:
//   - [CanvasWidget.CreateArc] (widget specific)
func Start(degrees any) Opt {
	return rawOption(fmt.Sprintf(`-start %s`, tclSafeString(fmt.Sprint(degrees))))
}

// Canvas — Create and manipulate 'canvas' hypergraphics drawing surface widgets
//
// # Description
//
// Items of type bitmap appear on the display as images with two colors,
// foreground and background. Bitmaps are created with widget commands of the
// following form:
//
//	pathName create bitmap x y ?option value ...?
//
// The arguments x and y or coordList (which must have two elements) specify
// the coordinates of a point used to position the bitmap on the display, as
// controlled by the -anchor option. After the coordinates there may be any
// number of option-value pairs, each of which sets one of the configuration
// options for the item. These same option-value pairs may be used in
// itemconfigure widget commands to change the item's configuration. A bitmap
// item becomes the current item when the mouse pointer is over any part of its
// bounding box.
//
// The following standard options are supported by bitmaps:
//
//   - [Anchor]
//   - [State]
//   - [Tags]
//
// The following extra options are supported for bitmaps:
//
//   - [Background] color
//
//   - [Activebackground] color
//
//   - [Disabledbackground] color
//
//     Specifies the color to use for each of the bitmap's “0” valued pixels in
//     its normal, active and disabled states. Color may have any of the forms
//     accepted by Tk_GetColor. If this option is not specified, or if it is
//     specified as an empty string, then nothing is displayed where the bitmap
//     pixels are 0; this produces a transparent effect.
//
//   - [Bitmap] bitmap
//
//   - [Activebitmap] bitmap
//
//   - [Disabledbitmap] bitmap:
//
//     These options specify the bitmaps to display in the item in its normal,
//     active and disabled states. Bitmap may have any of the forms accepted by
//     Tk_GetBitmap.
//
//   - [Foreground] color
//
//   - [Activeforeground] color
//
//   - [Disabledforeground] color:
//
//     These options specify the color to use for each of the bitmap's “1” valued
//     pixels in its normal, active and disabled states. Color may have any of the
//     forms accepted by Tk_GetColor.
//
// More information might be available at the [Tcl/Tk canvas] page.
//
// [Tcl/Tk canvas]: https://www.tcl.tk/man/tcl9.0/TkCmd/canvas.html
func (w *CanvasWidget) CreateBitmap(x, y any, options ...any) (r string) {
	return w.create("bitmap", append([]any{x, y}, options...)...)
}

// Activebitmap option.
//
// This option specify the bitmap to display in the item in its active state.
// Bitmap may have any of the forms accepted by Tk_GetBitmap.
//
// Known uses:
//   - [CanvasWidget.CreateBitmap] (widget specific)
func Activebitmap(bitmap any) Opt {
	return rawOption(fmt.Sprintf(`-activebitmap %s`, tclSafeString(fmt.Sprint(bitmap))))
}

// Disabledbitmap option.
//
// This option specify the bitmap to display in the item in its disabled state.
// Bitmap may have any of the forms accepted by Tk_GetBitmap.
//
// Known uses:
//   - [CanvasWidget.CreateBitmap] (widget specific)
func Disabledbitmap(bitmap any) Opt {
	return rawOption(fmt.Sprintf(`-disabledbitmap %s`, tclSafeString(fmt.Sprint(bitmap))))
}
<<<<<<< HEAD
=======

// Canvas — Create and manipulate 'canvas' hypergraphics drawing surface widgets
//
// # Description
//
// Items of type image are used to display images on a canvas. Images are
// created with widget commands of the following form:
//
//	pathName create image x y ?option value ...?
//
// The arguments x and y or coordList specify the coordinates of a point used
// to position the image on the display, as controlled by the -anchor option.
// After the coordinates there may be any number of option-value pairs, each of
// which sets one of the configuration options for the item. These same
// option-value pairs may be used in itemconfigure widget commands to change
// the item's configuration. An image item becomes the current item when the
// mouse pointer is over any part of its bounding box.
//
// The following standard options are supported by bitmaps:
//
//   - [Anchor]
//   - [State]
//   - [Tags]
//
// The following extra options are supported for bitmaps:
//
//   - [Image] name
//
//   - [Activeimage] name
//
//   - [Disabledimage] name
//
//     Specifies the name of the images to display in the item in is normal,
//     active and disabled states. This image must have been created previously
//     with the image create command.
//
// More information might be available at the [Tcl/Tk canvas] page.
//
// [Tcl/Tk canvas]: https://www.tcl.tk/man/tcl9.0/TkCmd/canvas.html
func (w *CanvasWidget) CreateImage(x, y any, options ...any) (r string) {
	return w.create("image", append([]any{x, y}, options...)...)
}

// Activeimage option.
//
// Specifies the name of the image to display in the item in its active state.
// This image must have been created previously with the image create command.
//
// Known uses:
//   - [CanvasWidget.CreateImage] (widget specific)
func Activeimage(val any) Opt {
	return rawOption(fmt.Sprintf(`-activeimage %s`, optionString(val)))
}

// Disabledimage option.
//
// Specifies the name of the image to display in the item in its disabled state.
// This image must have been created previously with the image create command.
//
// Known uses:
//   - [CanvasWidget.CreateImage] (widget specific)
func Disabledimage(val any) Opt {
	return rawOption(fmt.Sprintf(`-disabledimage %s`, optionString(val)))
}
>>>>>>> ok
