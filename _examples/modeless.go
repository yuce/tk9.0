package main

import (
	_ "embed"
	"fmt"
	"strconv"

	tk "modernc.org/tk9.0"
)

const APPNAME = "Modeless"

type App struct {
	label        *tk.TLabelWidget
	entry        *tk.TEntryWidget
	buttonFrame  *tk.TFrameWidget
	configButton *tk.TButtonWidget
	quitButton   *tk.TButtonWidget
	configDialog *ConfigDialog
}

func main() {
	NewApp().Run()
}

func NewApp() *App {
	app := &App{}
	tk.StyleThemeUse("clam")
	tk.WmWithdraw(tk.App)
	tk.App.WmTitle(APPNAME)
	tk.WmProtocol(tk.App, tk.WM_DELETE_WINDOW, app.onQuit)
	app.makeWidgets()
	app.makeLayout()
	app.makeBindings()
	return app
}

func (me *App) makeWidgets() {
	me.label = tk.TLabel(tk.Txt("This is modeless.go; see also modal.go"),
		tk.Relief(tk.GROOVE), tk.Background(tk.LightYellow))
	me.entry = tk.TEntry(
		tk.Textvariable("Can get focus when Config visible"))
	me.buttonFrame = tk.TFrame()
	me.configButton = me.buttonFrame.TButton(tk.Txt("Config…"),
		tk.Command(me.onConfig))
	me.quitButton = me.buttonFrame.TButton(tk.Txt("Quit"),
		tk.Command(me.onQuit))
}

func (me *App) makeLayout() {
	opts := tk.Opts{tk.Padx(3), tk.Pady(3)}
	tk.Grid(me.label, tk.Row(0), tk.Column(0), tk.Sticky(tk.WE), opts)
	tk.Grid(me.entry, tk.Row(1), tk.Column(0), tk.Sticky(tk.WE), opts)
	tk.Grid(me.buttonFrame, tk.Row(2), tk.Column(0), tk.Columnspan(2),
		tk.Sticky(tk.WE), opts)
	tk.GridColumnConfigure(tk.App, 0, tk.Weight(1))
	tk.Grid(me.configButton, tk.Row(0), tk.Column(0), tk.Sticky(tk.W),
		opts)
	tk.Grid(me.quitButton, tk.Row(0), tk.Column(1), tk.Sticky(tk.E), opts)
	tk.GridColumnConfigure(me.buttonFrame, 1, tk.Weight(1))
}

func (me *App) makeBindings() {
	tk.Bind(tk.App, "<Escape>", tk.Command(me.onQuit))
}

func (me *App) Run() {
	tk.App.Center()
	tk.WmDeiconify(tk.App)
	tk.App.Wait()
}

func (me *App) onConfig() {
	if me.configDialog == nil {
		me.configDialog = NewConfigDialog()
	}
	me.configDialog.Show()
}

func (me *App) onQuit() { tk.Destroy(tk.App) }

type ConfigDialog struct {
	win          *tk.ToplevelWidget
	scaleLabel   *tk.TLabelWidget
	scaleSpinbox *tk.TSpinboxWidget
	closeButton  *tk.TButtonWidget
}

func NewConfigDialog() *ConfigDialog {
	dlg := &ConfigDialog{}
	dlg.win = tk.App.Toplevel()
	dlg.win.WmTitle("Config")
	// NOTE below doesn't work - BUG?
	// tk.WmProtocol(dlg.win, tk.WM_DELETE_WINDOW, dlg.onHide)
	dlg.scaleLabel = dlg.win.TLabel(tk.Txt("Application Scale"))
	dlg.scaleSpinbox = dlg.win.TSpinbox(tk.Format("%.1f"),
		tk.Increment(0.1), tk.From(0.5), tk.To(5.0),
		tk.Textvariable(fmt.Sprintf("%f", tk.TkScaling())),
		tk.Command(dlg.onScaleChange))
	dlg.closeButton = dlg.win.TButton(tk.Txt("Close"),
		tk.Command(dlg.onHide))
	tk.Grid(dlg.scaleLabel, tk.Row(0), tk.Column(0), tk.Sticky(tk.W))
	tk.Grid(dlg.scaleSpinbox, tk.Row(0), tk.Column(1), tk.Sticky(tk.WE))
	tk.Grid(dlg.closeButton, tk.Row(1), tk.Column(0), tk.Columnspan(2))
	return dlg
}

func (me *ConfigDialog) onScaleChange() {
	text := me.scaleSpinbox.Textvariable()
	if scale, err := strconv.ParseFloat(text, 64); err == nil {
		tk.TkScaling(scale) // Live update FIXME Doesn't work
		tk.Update()
	}
}

func (me *ConfigDialog) onHide() {
	// FIXME doesn't work
	// tk.WmWithdraw(me.win)
	tk.GrabRelease(me.win)
}

func (me *ConfigDialog) Show() {
	// FIXME doesn't work
	// tk.WmDeiconify(me.win)
	me.win.Raise(tk.App)
	tk.Focus(me.win)
	tk.Focus(me.scaleSpinbox)
}
