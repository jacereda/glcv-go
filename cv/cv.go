// Copyright (c) 2012, Jorge Acereda Maciá. All rights reserved.  
//
// Use of this source code is governed by a BSD-style license that can
// be found in the LICENSE file.

// Package cv provides an OpenGL canvas and IO handling. This is the
// low-level package, you probably want the canvas package instead.
package cv

/*
#cgo CFLAGS: -DCV_NO_MAIN -DCV_EXPLICIT_ENTRY 
#cgo darwin LDFLAGS: -framework AppKit -lobjc
#cgo linux CFLAGS: -I/usr/X11R6/include
#cgo linux LDFLAGS: -L/usr/X11R6/lib -lGL -lX11
#cgo windows LDFLAGS: -lgdi32 -lopengl32
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

// Run starts the event loop using the passed event handler.
func Run(h func(e *Event) uintptr) {
	handler = h
	C.run()
}

// MosueX returns the current X mouse position.
func MouseX() int {
	return int(C.cvMouseX())
}

// MouseY returns the current Y mouse position.
func MouseY() int {
	return int(C.cvMouseY())
}

// Width returns the current canvas width.
func Width() int {
	return int(C.cvWidth())
}

// Height returns the current canvas height.
func Height() int {
	return int(C.cvHeight())
}

// Pressed returns if the key is pressed.
func Pressed(k key.Id) bool {
	return C.cvPressed(C.cvkey(k)) == 1
}

// Released returs if the key just been released.
func Released(k key.Id) bool {
	return C.cvReleased(C.cvkey(k)) == 1
}

// Quits requests the application to leave.
func Quit() {
	C.cvQuit()
}

// ShowCursor makes the mouse cursor visible.
func ShowCursor() {
	C.cvShowCursor()
}

// HideCursor makes the mouse cursor invisible.
func HideCursor() {
	C.cvHideCursor()
}

//export gohandle
func gohandle(e *C.ev) C.intptr_t {
	return C.intptr_t(handler((*Event)(e)))
}

func init() {
	runtime.LockOSThread()
}

func KeyName(k key.Id) string {
	return C.GoString(C.keyName(C.cvkey(k)))
}
