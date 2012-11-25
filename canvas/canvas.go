// Copyright (c) 2012, Jorge Acereda Maciá. All rights reserved.  
//
// Use of this source code is governed by a BSD-style license that can
// be found in the LICENSE file.

// Package canvas provides an OpenGL canvas and IO handling.
package canvas

import (
	"code.google.com/p/glcv-go/cv"
	"code.google.com/p/glcv-go/key"
)

type Responder interface {
	OnInit()                    // The OpenGL context is about to be created.
	OnTerm()                    // The OpenGL context is has been destroyed.
	OnGLInit()                  // The OpenGL context has been created.
	OnGLTerm()                  // The OpenGL context is about to be destroyed.
	OnPress(k key.Id)           // The key (KEY_*) has been pressed.
	OnRelease(k key.Id)         // The key (KEY_*) has been released.
	OnUnicode(u uint)           // The unicode character has been entered.
	OnMotion(x, y int)          // The mouse has moved.
	OnClose()                   // The window close button has been pressed.
	OnResize(w, h int)          // The window has been resized.
	OnUpdate()                  // Called once per frame.
	Name() string               // Should return the desired name for the app.
	Geometry() (x, y, w, h int) // Should return a hint for the desired geometry.
	Borders() bool              // Should the window have borders?
}

// The canvas structure, you should embed it in your derived canvas.
type Canvas struct {
	r Responder
}

var g_canvas *Canvas

// InitCanvas initializes the canvas with the given event responder
// (usually the derived canvas).
func (c *Canvas) InitCanvas(r Responder) {
	c.r = r
	g_canvas = c
}

// NextResponder returns the next event responder in the chain. Useful
// if you want to call the parent canvas in one of your handlers.
func (c *Canvas) NextResponder() Responder {
	return c
}

// Go starts the event loop.
func (c *Canvas) Go() {
	cv.Run(event)
}

// Quit requests the application to leave.
func (c *Canvas) Quit() {
	cv.Quit()
}

// ShowCursor makes the mouse cursor visible.
func (c *Canvas) ShowCursor() {
	cv.ShowCursor()
}

// HideCursor makes the mouse cursor invisible.
func (c *Canvas) HideCursor() {
	cv.HideCursor()
}

func (c *Canvas) init() {
	c.r.OnInit()
}

func (c *Canvas) OnInit() {}

func (c *Canvas) term() {
	c.r.OnTerm()
}

func (c *Canvas) OnTerm() {}

func (c *Canvas) close() {
	c.r.OnClose()
}

func (c *Canvas) OnClose() {}

func (c *Canvas) resize(w, h int) {
	c.r.OnResize(int(w), int(h))
}

func (c *Canvas) OnResize(w, h int) {}

func (c *Canvas) glinit() {
	c.r.OnGLInit()
}

func (c *Canvas) OnGLInit() {}

func (c *Canvas) glterm() {
	c.r.OnGLTerm()
}

func (c *Canvas) OnGLTerm() {}

func (c *Canvas) press(k key.Id) {
	c.r.OnPress(k)
}

func (c *Canvas) OnPress(k key.Id) {}

func (c *Canvas) release(k key.Id) {
	c.r.OnRelease(k)
}

func (c *Canvas) OnRelease(k key.Id) {}

func (c *Canvas) unicode(u uint) {
	c.r.OnUnicode(u)
}

func (c *Canvas) OnUnicode(u uint) {}

func (c *Canvas) motion(x, y int) {
	c.r.OnMotion(x, y)
}

func (c *Canvas) OnMotion(x, y int) {}

func (c *Canvas) update() {
	c.r.OnUpdate()
}

func (c *Canvas) OnUpdate() {}

func (c *Canvas) name() uintptr {
	return cv.String(c.r.Name())
}

func (c *Canvas) Name() string {
	return "canvas"
}

func (c *Canvas) borders() uintptr {
	if c.Borders() {
		return 1
	}
	return 0
}

func (c *Canvas) Borders() bool {
	return true
}

func (c *Canvas) geometry() (x, y, w, h uintptr) {
	ix, iy, iw, ih := c.r.Geometry()
	x, y, w, h = uintptr(ix), uintptr(iy), uintptr(iw), uintptr(ih)
	return
}

func (c *Canvas) Geometry() (int, int, int, int) {
	return -1, -1, -1, -1
}

func (c *Canvas) Size() (int, int) {
	return c.Width(), c.Height()
}

func (c *Canvas) Width() int {
	return cv.Width()
}

func (c *Canvas) Height() int {
	return cv.Height()
}

func (c *Canvas) Mouse() (int, int) {
	return cv.MouseX(), cv.MouseY()
}

func (c *Canvas) Pressed(k key.Id) bool {
	return cv.Pressed(k)
}

func (c *Canvas) Released(k key.Id) bool {
	return cv.Released(k)
}

func (c *Canvas) KeyName(k key.Id) string {
	return cv.KeyName(k)
}

func event(e *cv.Event) uintptr {
	r := uintptr(1)
	switch e.Type() {
	case cv.ON_INIT:
		g_canvas.init()
	case cv.ON_TERM:
		g_canvas.term()
	case cv.ON_GLINIT:
		g_canvas.glinit()
	case cv.ON_DOWN:
		g_canvas.press(e.Which())
	case cv.ON_UP:
		g_canvas.release(e.Which())
	case cv.ON_UNICODE:
		g_canvas.unicode(e.Unicode())
	case cv.ON_MOTION:
		g_canvas.motion(e.X(), e.Y())
	case cv.ON_CLOSE:
		g_canvas.close()
	case cv.ON_RESIZE:
		g_canvas.resize(e.Width(), e.Height())
	case cv.ON_UPDATE:
		g_canvas.update()
	case cv.HINT_NAME:
		r = g_canvas.name()
	case cv.HINT_BORDERS:
		r = g_canvas.borders()
	case cv.HINT_XPOS:
		r, _, _, _ = g_canvas.geometry()
	case cv.HINT_YPOS:
		_, r, _, _ = g_canvas.geometry()
	case cv.HINT_WIDTH:
		_, _, r, _ = g_canvas.geometry()
	case cv.HINT_HEIGHT:
		_, _, _, r = g_canvas.geometry()
	default:
		r = 0
	}
	return r
}
