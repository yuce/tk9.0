//go:build (linux && (amd64 || arm64)) || (darwin && (amd64 || arm64)) || (freebsd && (amd64 || arm64))

package tk9_0 // import "modernc.org/tk9.0"

import (
	"fmt"

	"github.com/ebitengine/purego"
)

// Create a new Tcl command whose name is the widget's pathname, and
// whose action is to dispatch on the operation passed to the widget:
func (proxy *widgetProxy) registerEventDispatcher() {
	runCmdProxy := purego.NewCallback(proxy.eventDispatcher)
	if proxy.commandName, Error = cString(proxy.window.String()); Error != nil {
		return
	}
	cmd, _, _ := purego.SyscallN(createCommandProc, interp, proxy.commandName, runCmdProxy, 0, 0)
	if cmd == 0 {
		Error = fmt.Errorf("registering widget proxy event dispatcher proxy failed: %v", getObjResultProc)
		return
	}
}

func (proxy *widgetProxy) unregisterEventDispatcher() {
	_, _, _ = purego.SyscallN(deleteCommandProc, interp, proxy.commandName)
}
