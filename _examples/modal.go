package main

import (
	_ "embed"
	"strconv"

	tk "modernc.org/tk9.0"
)

const APPNAME = "Modal"

type App struct {
	label        *tk.TLabelWidget
	entry        *tk.TEntryWidget
	buttonFrame  *tk.TFrameWidget
	configButton *tk.TButtonWidget
	quitButton   *tk.TButtonWidget
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
	me.label = tk.TLabel(tk.Txt("This is modal.go; see also modeless.go"),
		tk.Relief(tk.GROOVE), tk.Background(tk.LightYellow))
	me.entry = tk.TEntry(
		tk.Textvariable("Cannot get focus when Config visible"))
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
	tk.Grid(me.configButton, tk.Row(0), tk.Column(0), tk.Sticky(tk.W), opts)
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
	data := ConfigDialogData{Scale: tk.TkScaling()}
	dlg := NewConfigDialog(&data)
	dlg.ShowModal()
	if data.Ok {
		tk.TkScaling(data.Scale)
	}
}

func (me *App) onQuit() { tk.Destroy(tk.App) }

type ConfigDialogData struct {
	Ok    bool
	Scale float64
}

type ConfigDialog struct {
	data         *ConfigDialogData
	win          *tk.ToplevelWidget
	scaleLabel   *tk.TLabelWidget
	scaleSpinbox *tk.TSpinboxWidget
	okButton     *tk.TButtonWidget
	cancelButton *tk.TButtonWidget
}

func NewConfigDialog(data *ConfigDialogData) *ConfigDialog {
	dlg := &ConfigDialog{data: data}
	dlg.win = tk.App.Toplevel()
	dlg.win.WmTitle("Config")
	// NOTE below doesn't work - BUG?
	// tk.WmProtocol(dlg.win, tk.WM_DELETE_WINDOW, dlg.onCancel)
	dlg.scaleLabel = dlg.win.TLabel(tk.Txt("Application Scale"))
	// TODO set initial value to 1.0
	dlg.scaleSpinbox = dlg.win.TSpinbox(tk.Format("%.1f"),
		tk.Increment(0.1), tk.From(0.5), tk.To(5.0))
	dlg.okButton = dlg.win.TButton(tk.Txt("OK"), tk.Command(dlg.onOk))
	dlg.cancelButton = dlg.win.TButton(tk.Txt("Cancel"),
		tk.Command(dlg.onCancel))
	tk.Grid(dlg.scaleLabel, tk.Row(0), tk.Column(0), tk.Sticky(tk.W))
	tk.Grid(dlg.scaleSpinbox, tk.Row(0), tk.Column(1), tk.Sticky(tk.WE))
	tk.Grid(dlg.okButton, tk.Row(1), tk.Column(0), tk.Sticky(tk.E))
	tk.Grid(dlg.cancelButton, tk.Row(1), tk.Column(1), tk.Sticky(tk.W))
	tk.GridColumnConfigure(dlg.win, 1, tk.Weight(1))
	return dlg
}

func (me *ConfigDialog) onOk() {
	text := "1.2" // me.scaleSpinbox.Get() // FIXME how to get text?
	if scale, err := strconv.ParseFloat(text, 64); err == nil {
		me.data.Scale = scale
		me.data.Ok = true
	}
	tk.Destroy(me.win)
}

func (me *ConfigDialog) onCancel() { tk.Destroy(me.win) }

func (me *ConfigDialog) ShowModal() {
	me.win.Raise(tk.App)
	tk.Focus(me.win)
	tk.Focus(me.scaleSpinbox)
	tk.GrabSet(me.win)
	me.win.Wait()
}
