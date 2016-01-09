package cv

/*
#cgo CFLAGS: -DCV_EXPLICIT_ENTRY -Iglcv/src 
#include "cv.h"
*/
import "C"

import (
	"github.com/jacereda/glcv-go/key"
)

type Event C.ev

const (
	ON_INIT    = C.CVE_INIT    // The OpenGL context is about to be created.
	ON_TERM    = C.CVE_TERM    // The OpenGL context is has been destroyed.
	ON_GLINIT  = C.CVE_GLINIT  // The OpenGL context has been created.
	ON_GLTERM  = C.CVE_GLTERM  // The OpenGL context is about to be destroyed.
	ON_DOWN    = C.CVE_DOWN    // The key (KEY_*) has been pressed.
	ON_UP      = C.CVE_UP      // The key (KEY_*) has been released.
	ON_UNICODE = C.CVE_UNICODE // The unicode character has been entered.
	ON_MOTION  = C.CVE_MOTION  // The mouse has moved.
	ON_CLOSE   = C.CVE_CLOSE   // The window close button has been pressed.
	ON_RESIZE  = C.CVE_RESIZE  // The window has been resized.
	ON_UPDATE  = C.CVE_UPDATE  // Called once per frame.

	HINT_NAME    = C.CVQ_NAME    // Should return the desired name for the app.
	HINT_XPOS    = C.CVQ_XPOS    // Desired X position (hint).
	HINT_YPOS    = C.CVQ_YPOS    // Desired Y position (hint).
	HINT_WIDTH   = C.CVQ_WIDTH   // Desired width (hint). -1 == fullscreen.
	HINT_HEIGHT  = C.CVQ_HEIGHT  // Desired height.
)

// Event type (ON_* / HINT_*).
func (e *Event) Type() int {
	return int(C.evType((*C.ev)(e)))
}

// Event name.
func (e *Event) Name() string {
	return C.GoString(C.evName((*C.ev)(e)))
}

// Width for ON_RESIZE events.
func (e *Event) Width() int {
	return int(C.evWidth((*C.ev)(e)))
}

// Height for ON_RESIZE events.
func (e *Event) Height() int {
	return int(C.evHeight((*C.ev)(e)))
}

// Key identifier for ON_DOWN/ON_UP events.
func (e *Event) Which() key.Id {
	return key.Id(C.evWhich((*C.ev)(e)))
}

// Unicode character for ON_UNICODE events.
func (e *Event) Unicode() uint {
	return uint(C.evUnicode((*C.ev)(e)))
}

// X mouse position for ON_MOTION events.
func (e *Event) X() int {
	return int(C.evX((*C.ev)(e)))
}

// Y mouse position for ON_MOTION events.
func (e *Event) Y() int {
	return int(C.evY((*C.ev)(e)))
}

// Position for ON_MOTION events.
func (e *Event) Pos() (x, y int) {
	x, y = e.X(), e.Y()
	return
}
