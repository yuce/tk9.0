package tk9_0 // import "modernc.org/tk9.0"

import (
	"fmt"
	"unsafe"
)

// OperationCallback is the signature for a WidgetProxy operation callback.
type OperationCallback func(args []string)

// widgetProxy wraps a widget window, providing the means to intercept its
// internal operations, and modify their behaviour.
// For example, it can provide access to a TextWidget's internal 'insert' and
// 'delete' operations.
type widgetProxy struct {
	window       *Window
	originalPath string
	commandName  uintptr
	operations   map[string]OperationCallback
}

// newWidgetProxy creates a proxy that facilitates hooking in to
// a widget window's internal operations.
func newWidgetProxy(window *Window) widgetProxy {
	proxy := widgetProxy{
		window:       window,
		originalPath: window.String() + "_original",
		operations:   make(map[string]OperationCallback),
	}

	// Rename the Tcl command within Tcl:
	evalErr(fmt.Sprintf("rename %s %s", window, proxy.originalPath))
	// Register an event dispatcher for the window.
	proxy.registerEventDispatcher()

	return proxy
}

// Close undoes the wrapping of its Window.
// All registered operations are unregistered.
func (proxy *widgetProxy) Close() {
	// Unregister all registered operations.
	for operation := range proxy.operations {
		proxy.Unregister(operation)
	}

	// Restore the original widget Tcl command.
	proxy.unregisterEventDispatcher()
	// Restore the Tcl command within Tcl:
	evalErr(fmt.Sprintf("rename %s %s", proxy.originalPath, proxy.window))
}

// Register registers a callback for an operation supported by the wrapped Window.
//
// The operation name is widget-specific. It is not possible to validate whether
// the operation name is support by the wrapped widget. If an unsupported
// operation name is used it will be silently ignored.
//
// The operation's arguments are passed to the callback. The callback may perform
// the operation on the wrapped widget by calling the EvalWrapped method with the
// arguments. The arguments may be modified if desired.
func (proxy *widgetProxy) Register(operation string, callback OperationCallback) {
	proxy.operations[operation] = callback
}

// Unregister unregisters an operation callback.
// If there is no registered callback for the operation, it is silently ignored.
func (proxy *widgetProxy) Unregister(operation string) {
	delete(proxy.operations, operation)
}

// EvalWrapped evaluates the arguments as a raw Tcl string against the wrapped Window.
func (proxy *widgetProxy) EvalWrapped(args []string) {
	evalErr(fmt.Sprintf("%s %s", proxy.originalPath, tclSafeStrings(args...)))
}

func (proxy *widgetProxy) eventDispatcher(clientData, in uintptr, argc int32, argv uintptr) uintptr {
	// Expect at least arguments for the path and the operation.
	if argc < 2 {
		setResult(fmt.Sprintf("WidgetProxy eventDispatcher internal error: argc=%v", argc))
		return tcl_error
	}

	argsPointers := unsafe.Slice((*uintptr)(unsafe.Pointer(argv)), argc)
	argsPointers = argsPointers[1:] // skip path
	args := make([]string, len(argsPointers))
	for i := 0; i < len(argsPointers); i++ {
		args[i] = goString(argsPointers[i])
	}

	operation := args[0]
	if callback, ok := proxy.operations[operation]; ok {
		// Dispatch to the operation's registered callback.
		callback(args)
	} else {
		// Process as normal.
		proxy.EvalWrapped(args)
	}

	return tcl_ok
}

// TextWidgetProxy wraps a TextWidget.
// It provides the ability to intercept the widget's internal operations (such as
// 'insert' and 'delete'), and modify their behaviour.
type TextWidgetProxy struct {
	*TextWidget
	widgetProxy
}

// NewTextWidgetProxy creates a TextWidgetProxy, wrapping the
// provided TextWidget.
func NewTextWidgetProxy(widget *TextWidget) TextWidgetProxy {
	return TextWidgetProxy{
		TextWidget:  widget,
		widgetProxy: newWidgetProxy(widget.Window),
	}
}
