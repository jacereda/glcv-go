package cv

/*
#cgo CFLAGS: -DCV_EXPLICIT_ENTRY -Iglcv/src 
#cgo LDFLAGS: -Lglcv/b -lglcv
#include "cv.h"
static void run() {
  intptr_t gohandle(ev *);
  cvRun((intptr_t(*)(const ev *))gohandle);
}
*/
import "C"

import (
	"code.google.com/p/glcv-go/key"
	"runtime"
	"unsafe"
)

// Helper function for HINT_NAME event.
func String(s string) uintptr {
	return uintptr(unsafe.Pointer(C.CString(s)))
}

var handler func(e *Event) uintptr

// Start the event loop using the passed event handler.
func Run(h func(e *Event) uintptr) {
	handler = h
	C.run()
}

// Request the application to leave.
func Quit() {
	C.cvQuit()
}

// Mouse current X position.
func MouseX() int {
	return int(C.cvMouseX())
}

// Mouse current Y position.
func MouseY() int {
	return int(C.cvMouseY())
}

// Current canvas width.
func Width() int {
	return int(C.cvWidth())
}

// Current canvas height.
func Height() int {
	return int(C.cvHeight())
}

// Is the key pressed?
func Pressed(k key.Id) bool {
	return C.cvPressed(C.cvkey(k)) == 1
}

// Has the key just been released?
func Released(k key.Id) bool {
	return C.cvReleased(C.cvkey(k)) == 1
}

//export gohandle
func gohandle(e *C.ev) C.intptr_t {
	return C.intptr_t(handler((*Event)(e)))
}

func init() {
	runtime.LockOSThread()
}
