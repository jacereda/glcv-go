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
	"runtime"
	"unsafe"
)

const (
	KEY_NONE           = C.CVK_NONE
	KEY_MOUSEWHEELUP   = C.CVK_MOUSEWHEELUP
	KEY_MOUSEWHEELDOWN = C.CVK_MOUSEWHEELDOWN
	KEY_MOUSELEFT      = C.CVK_MOUSELEFT
	KEY_MOUSERIGHT     = C.CVK_MOUSERIGHT
	KEY_MOUSEMIDDLE    = C.CVK_MOUSEMIDDLE
	KEY_A              = C.CVK_A
	KEY_B              = C.CVK_B
	KEY_C              = C.CVK_C
	KEY_D              = C.CVK_D
	KEY_E              = C.CVK_E
	KEY_F              = C.CVK_F
	KEY_G              = C.CVK_G
	KEY_H              = C.CVK_H
	KEY_I              = C.CVK_I
	KEY_J              = C.CVK_J
	KEY_K              = C.CVK_K
	KEY_L              = C.CVK_L
	KEY_M              = C.CVK_M
	KEY_N              = C.CVK_N
	KEY_O              = C.CVK_O
	KEY_P              = C.CVK_P
	KEY_Q              = C.CVK_Q
	KEY_R              = C.CVK_R
	KEY_S              = C.CVK_S
	KEY_T              = C.CVK_T
	KEY_U              = C.CVK_U
	KEY_V              = C.CVK_V
	KEY_W              = C.CVK_W
	KEY_X              = C.CVK_X
	KEY_Y              = C.CVK_Y
	KEY_Z              = C.CVK_Z
	KEY_0              = C.CVK_0
	KEY_1              = C.CVK_1
	KEY_2              = C.CVK_2
	KEY_3              = C.CVK_3
	KEY_4              = C.CVK_4
	KEY_5              = C.CVK_5
	KEY_6              = C.CVK_6
	KEY_7              = C.CVK_7
	KEY_8              = C.CVK_8
	KEY_9              = C.CVK_9
	KEY_EQUAL          = C.CVK_EQUAL
	KEY_MINUS          = C.CVK_MINUS
	KEY_RIGHTBRACKET   = C.CVK_RIGHTBRACKET
	KEY_LEFTBRACKET    = C.CVK_LEFTBRACKET
	KEY_QUOTE          = C.CVK_QUOTE
	KEY_SEMICOLON      = C.CVK_SEMICOLON
	KEY_BACKSLASH      = C.CVK_BACKSLASH
	KEY_COMMA          = C.CVK_COMMA
	KEY_SLASH          = C.CVK_SLASH
	KEY_PERIOD         = C.CVK_PERIOD
	KEY_GRAVE          = C.CVK_GRAVE
	KEY_KEYPADDECIMAL  = C.CVK_KEYPADDECIMAL
	KEY_KEYPADMULTIPLY = C.CVK_KEYPADMULTIPLY
	KEY_KEYPADPLUS     = C.CVK_KEYPADPLUS
	KEY_KEYPADCLEAR    = C.CVK_KEYPADCLEAR
	KEY_KEYPADDIVIDE   = C.CVK_KEYPADDIVIDE
	KEY_KEYPADENTER    = C.CVK_KEYPADENTER
	KEY_KEYPADMINUS    = C.CVK_KEYPADMINUS
	KEY_KEYPADEQUALS   = C.CVK_KEYPADEQUALS
	KEY_KEYPAD0        = C.CVK_KEYPAD0
	KEY_KEYPAD1        = C.CVK_KEYPAD1
	KEY_KEYPAD2        = C.CVK_KEYPAD2
	KEY_KEYPAD3        = C.CVK_KEYPAD3
	KEY_KEYPAD4        = C.CVK_KEYPAD4
	KEY_KEYPAD5        = C.CVK_KEYPAD5
	KEY_KEYPAD6        = C.CVK_KEYPAD6
	KEY_KEYPAD7        = C.CVK_KEYPAD7
	KEY_KEYPAD8        = C.CVK_KEYPAD8
	KEY_KEYPAD9        = C.CVK_KEYPAD9
	KEY_RETURN         = C.CVK_RETURN
	KEY_TAB            = C.CVK_TAB
	KEY_SPACE          = C.CVK_SPACE
	KEY_DELETE         = C.CVK_DELETE
	KEY_ESCAPE         = C.CVK_ESCAPE
	KEY_COMMAND        = C.CVK_COMMAND
	KEY_SHIFT          = C.CVK_SHIFT
	KEY_CAPSLOCK       = C.CVK_CAPSLOCK
	KEY_OPTION         = C.CVK_OPTION
	KEY_CONTROL        = C.CVK_CONTROL
	KEY_RIGHTSHIFT     = C.CVK_RIGHTSHIFT
	KEY_RIGHTOPTION    = C.CVK_RIGHTOPTION
	KEY_RIGHTCONTROL   = C.CVK_RIGHTCONTROL
	KEY_FUNCTION       = C.CVK_FUNCTION
	KEY_VOLUMEUP       = C.CVK_VOLUMEUP
	KEY_VOLUMEDOWN     = C.CVK_VOLUMEDOWN
	KEY_MUTE           = C.CVK_MUTE
	KEY_F1             = C.CVK_F1
	KEY_F2             = C.CVK_F2
	KEY_F3             = C.CVK_F3
	KEY_F4             = C.CVK_F4
	KEY_F5             = C.CVK_F5
	KEY_F6             = C.CVK_F6
	KEY_F7             = C.CVK_F7
	KEY_F8             = C.CVK_F8
	KEY_F9             = C.CVK_F9
	KEY_F10            = C.CVK_F10
	KEY_F11            = C.CVK_F11
	KEY_F12            = C.CVK_F12
	KEY_F13            = C.CVK_F13
	KEY_F14            = C.CVK_F14
	KEY_F15            = C.CVK_F15
	KEY_F16            = C.CVK_F16
	KEY_F17            = C.CVK_F17
	KEY_F18            = C.CVK_F18
	KEY_F19            = C.CVK_F19
	KEY_F20            = C.CVK_F20
	KEY_HELP           = C.CVK_HELP
	KEY_HOME           = C.CVK_HOME
	KEY_PAGEUP         = C.CVK_PAGEUP
	KEY_FORWARDDELETE  = C.CVK_FORWARDDELETE
	KEY_END            = C.CVK_END
	KEY_PAGEDOWN       = C.CVK_PAGEDOWN
	KEY_LEFTARROW      = C.CVK_LEFTARROW
	KEY_RIGHTARROW     = C.CVK_RIGHTARROW
	KEY_DOWNARROW      = C.CVK_DOWNARROW
	KEY_UPARROW        = C.CVK_UPARROW
	KEY_SCROLL         = C.CVK_SCROLL
	KEY_NUMLOCK        = C.CVK_NUMLOCK
	KEY_CLEAR          = C.CVK_CLEAR
	KEY_SYSREQ         = C.CVK_SYSREQ
	KEY_PAUSE          = C.CVK_PAUSE
	KEY_CAMERA         = C.CVK_CAMERA
	KEY_CENTER         = C.CVK_CENTER
	KEY_AT             = C.CVK_AT
	KEY_SYM            = C.CVK_SYM
	KEY_MAX            = C.CVK_MAX

	ON_INIT   = C.CVE_INIT
	ON_TERM   = C.CVE_TERM
	ON_GLINIT = C.CVE_GLINIT
	//	ON_GLTERM  = C.CVE_GLTERM
	ON_DOWN    = C.CVE_DOWN
	ON_UP      = C.CVE_UP
	ON_UNICODE = C.CVE_UNICODE
	ON_MOTION  = C.CVE_MOTION
	ON_CLOSE   = C.CVE_CLOSE
	ON_RESIZE  = C.CVE_RESIZE
	ON_UPDATE  = C.CVE_UPDATE

	HINT_NAME    = C.CVQ_NAME
	HINT_BORDERS = C.CVQ_BORDERS
	HINT_XPOS    = C.CVQ_XPOS
	HINT_YPOS    = C.CVQ_YPOS
	HINT_WIDTH   = C.CVQ_WIDTH
	HINT_HEIGHT  = C.CVQ_HEIGHT
)

type Event C.ev

func (e *Event) Type() int {
	return int(C.evType((*C.ev)(e)))
}

func (e *Event) Name() string {
	return C.GoString(C.evName((*C.ev)(e)))
}

func (e *Event) Width() int {
	return int(C.evWidth((*C.ev)(e)))
}

func (e *Event) Height() int {
	return int(C.evHeight((*C.ev)(e)))
}

func (e *Event) Which() int {
	return int(C.evWhich((*C.ev)(e)))
}

func (e *Event) Unicode() uint {
	return uint(C.evUnicode((*C.ev)(e)))
}

func (e *Event) X() int {
	return int(C.evX((*C.ev)(e)))
}

func (e *Event) Y() int {
	return int(C.evY((*C.ev)(e)))
}

func (e *Event) Pos() (x, y int) {
	x, y = e.X(), e.Y()
	return
}

func String(s string) uintptr {
	return uintptr(unsafe.Pointer(C.CString(s)))
}

var handler func(e *Event) uintptr

func Run(h func(e *Event) uintptr) {
	handler = h
	C.run()
}

func Quit() {
	C.cvQuit()
}

func KeyName(k int) string {
	return C.GoString(C.keyName(C.cvkey(k)))
}

func MouseX() int {
	return int(C.cvMouseX())
}

func MouseY() int {
	return int(C.cvMouseY())
}

func Width() int {
	return int(C.cvWidth())
}

func Height() int {
	return int(C.cvHeight())
}

func Pressed(k int) bool {
	return C.cvPressed(C.cvkey(k)) == 1
}

func Released(k int) bool {
	return C.cvReleased(C.cvkey(k)) == 1
}

//export gohandle
func gohandle(e *C.ev) C.intptr_t {
	return C.intptr_t(handler((*Event)(e)))
}

func init() {
	runtime.LockOSThread()
}
