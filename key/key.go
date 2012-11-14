package key

/*
#cgo CFLAGS: -DCV_NO_MAIN -DCV_EXPLICIT_ENTRY -I../cv/glcv/src
#include "cv.h"
*/
import "C"

type Id int

const (
	NONE           = C.CVK_NONE
	MOUSEWHEELUP   = C.CVK_MOUSEWHEELUP
	MOUSEWHEELDOWN = C.CVK_MOUSEWHEELDOWN
	MOUSELEFT      = C.CVK_MOUSELEFT
	MOUSERIGHT     = C.CVK_MOUSERIGHT
	MOUSEMIDDLE    = C.CVK_MOUSEMIDDLE
	A              = C.CVK_A
	B              = C.CVK_B
	C              = C.CVK_C
	D              = C.CVK_D
	E              = C.CVK_E
	F              = C.CVK_F
	G              = C.CVK_G
	H              = C.CVK_H
	I              = C.CVK_I
	J              = C.CVK_J
	K              = C.CVK_K
	L              = C.CVK_L
	M              = C.CVK_M
	N              = C.CVK_N
	O              = C.CVK_O
	P              = C.CVK_P
	Q              = C.CVK_Q
	R              = C.CVK_R
	S              = C.CVK_S
	T              = C.CVK_T
	U              = C.CVK_U
	V              = C.CVK_V
	W              = C.CVK_W
	X              = C.CVK_X
	Y              = C.CVK_Y
	Z              = C.CVK_Z
	N0             = C.CVK_0
	N1             = C.CVK_1
	N2             = C.CVK_2
	N3             = C.CVK_3
	N4             = C.CVK_4
	N5             = C.CVK_5
	N6             = C.CVK_6
	N7             = C.CVK_7
	N8             = C.CVK_8
	N9             = C.CVK_9
	EQUAL          = C.CVK_EQUAL
	MINUS          = C.CVK_MINUS
	RIGHTBRACKET   = C.CVK_RIGHTBRACKET
	LEFTBRACKET    = C.CVK_LEFTBRACKET
	QUOTE          = C.CVK_QUOTE
	SEMICOLON      = C.CVK_SEMICOLON
	BACKSLASH      = C.CVK_BACKSLASH
	COMMA          = C.CVK_COMMA
	SLASH          = C.CVK_SLASH
	PERIOD         = C.CVK_PERIOD
	GRAVE          = C.CVK_GRAVE
	KEYPADDECIMAL  = C.CVK_KEYPADDECIMAL
	KEYPADMULTIPLY = C.CVK_KEYPADMULTIPLY
	KEYPADPLUS     = C.CVK_KEYPADPLUS
	KEYPADCLEAR    = C.CVK_KEYPADCLEAR
	KEYPADDIVIDE   = C.CVK_KEYPADDIVIDE
	KEYPADENTER    = C.CVK_KEYPADENTER
	KEYPADMINUS    = C.CVK_KEYPADMINUS
	KEYPADEQUALS   = C.CVK_KEYPADEQUALS
	KEYPAD0        = C.CVK_KEYPAD0
	KEYPAD1        = C.CVK_KEYPAD1
	KEYPAD2        = C.CVK_KEYPAD2
	KEYPAD3        = C.CVK_KEYPAD3
	KEYPAD4        = C.CVK_KEYPAD4
	KEYPAD5        = C.CVK_KEYPAD5
	KEYPAD6        = C.CVK_KEYPAD6
	KEYPAD7        = C.CVK_KEYPAD7
	KEYPAD8        = C.CVK_KEYPAD8
	KEYPAD9        = C.CVK_KEYPAD9
	RETURN         = C.CVK_RETURN
	TAB            = C.CVK_TAB
	SPACE          = C.CVK_SPACE
	DELETE         = C.CVK_DELETE
	ESCAPE         = C.CVK_ESCAPE
	COMMAND        = C.CVK_COMMAND
	SHIFT          = C.CVK_SHIFT
	CAPSLOCK       = C.CVK_CAPSLOCK
	OPTION         = C.CVK_OPTION
	CONTROL        = C.CVK_CONTROL
	RIGHTSHIFT     = C.CVK_RIGHTSHIFT
	RIGHTOPTION    = C.CVK_RIGHTOPTION
	RIGHTCONTROL   = C.CVK_RIGHTCONTROL
	FUNCTION       = C.CVK_FUNCTION
	VOLUMEUP       = C.CVK_VOLUMEUP
	VOLUMEDOWN     = C.CVK_VOLUMEDOWN
	MUTE           = C.CVK_MUTE
	F1             = C.CVK_F1
	F2             = C.CVK_F2
	F3             = C.CVK_F3
	F4             = C.CVK_F4
	F5             = C.CVK_F5
	F6             = C.CVK_F6
	F7             = C.CVK_F7
	F8             = C.CVK_F8
	F9             = C.CVK_F9
	F10            = C.CVK_F10
	F11            = C.CVK_F11
	F12            = C.CVK_F12
	F13            = C.CVK_F13
	F14            = C.CVK_F14
	F15            = C.CVK_F15
	F16            = C.CVK_F16
	F17            = C.CVK_F17
	F18            = C.CVK_F18
	F19            = C.CVK_F19
	F20            = C.CVK_F20
	HELP           = C.CVK_HELP
	HOME           = C.CVK_HOME
	PAGEUP         = C.CVK_PAGEUP
	FORWARDDELETE  = C.CVK_FORWARDDELETE
	END            = C.CVK_END
	PAGEDOWN       = C.CVK_PAGEDOWN
	LEFTARROW      = C.CVK_LEFTARROW
	RIGHTARROW     = C.CVK_RIGHTARROW
	DOWNARROW      = C.CVK_DOWNARROW
	UPARROW        = C.CVK_UPARROW
	SCROLL         = C.CVK_SCROLL
	NUMLOCK        = C.CVK_NUMLOCK
	CLEAR          = C.CVK_CLEAR
	SYSREQ         = C.CVK_SYSREQ
	PAUSE          = C.CVK_PAUSE
	CAMERA         = C.CVK_CAMERA
	CENTER         = C.CVK_CENTER
	AT             = C.CVK_AT
	SYM            = C.CVK_SYM
	MAX            = C.CVK_MAX
)

// Key name
//func (k Id) Name() string {
//	return C.GoString(cv.KeyName(C.cvkey(k)))
//}
