package tk9_0 // import "modernc.org/tk9.0"

import (
	"fmt"

	"golang.org/x/sys/windows"
)

// Create a new Tcl command whose name is the widget's pathname, and
// whose action is to dispatch on the operation passed to the widget:
func (proxy *WidgetProxy) registerEventDispatcher() {
	runCmdProxy := windows.NewCallback(proxy.eventDispatcher)
	if proxy.commandName, Error = cString(proxy.window.String()); Error != nil {
		return
	}
	cmd, _, _ := createCommandProc.Call(interp, interp, proxy.commandName, runCmdProxy, 0, 0)
	if cmd == 0 {
		Error = fmt.Errorf("registering widget proxy event dispatcher proxy failed: %v", getObjResultProc)
		return
	}
}

func (proxy *WidgetProxy) unregisterEventDispatcher() {
	defer func() {
		allocator.UintptrFree(proxy.commandName)
		proxy.commandName = 0 // nil
	}()
	_, _, _ = deleteCommandProc.Call(interp, proxy.commandName)
}
