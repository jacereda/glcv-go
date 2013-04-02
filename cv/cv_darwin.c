#include <objc/runtime.h>
#include <objc/message.h>
#include "cv.c"
#define evType evTypeC
#include <Carbon/Carbon.h> // For keycodes
#undef evType
#include <ApplicationServices/ApplicationServices.h>
#include <OpenGL/gl.h>
#include <OpenGL/OpenGL.h>
typedef CGPoint NSPoint;
typedef CGSize NSSize;
typedef CGRect NSRect;

#define WINDOWED_MASK NSTitledWindowMask | NSClosableWindowMask |       \
        NSMiniaturizableWindowMask | NSResizableWindowMask
#define FULLSCREEN_MASK NSBorderlessWindowMask


#define cls objc_getClass
#define sel sel_registerName
#define send objc_msgSend
#define fsend objc_msgSend_fpret
#define ssend objc_msgSendSuper

#define NIL (id)(0)

static inline NSPoint mkpoint(float x, float y) {
        NSPoint p;
        p.x = x; p.y = y; 
        return p;
}

static inline NSSize mksize(float w, float h) {
        NSSize s;
        s.width = w;
        s.height = h;
        return s;
}

static inline NSPoint psend(id i, SEL s) {
	NSPoint (*fn)() = (NSPoint(*)())objc_msgSend;
        return fn(i, s);
}

static inline NSRect rsend(id i, SEL s) {
	NSRect (*fn)() = (NSRect(*)())objc_msgSend_stret;
        return fn(i, s);
}

static inline NSRect rsend2(id i, SEL s, NSRect r, unsigned u) {
	NSRect (*fn)() = (NSRect(*)())objc_msgSend_stret;
        return fn(i, s, r, u);
}

typedef struct objc_object NSString;
typedef unsigned long NSUInteger;
typedef long NSInteger;
typedef unsigned unichar;
typedef  uint32_t NSOpenGLPixelFormatAttribute;
typedef NSUInteger NSBackingStoreType;

enum {
        NSAlphaShiftKeyMask         = 1 << 16,
        NSShiftKeyMask              = 1 << 17,
        NSControlKeyMask            = 1 << 18,
        NSAlternateKeyMask          = 1 << 19,
        NSCommandKeyMask            = 1 << 20,
        NSNumericPadKeyMask         = 1 << 21,
        NSHelpKeyMask               = 1 << 22,
        NSFunctionKeyMask           = 1 << 23,
#if MAC_OS_X_VERSION_MAX_ALLOWED >= MAC_OS_X_VERSION_10_4
        NSDeviceIndependentModifierFlagsMask    = 0xffff0000UL
#endif
};

enum {
        NSBorderlessWindowMask		= 0,
        NSTitledWindowMask			= 1 << 0,
        NSClosableWindowMask		= 1 << 1,
        NSMiniaturizableWindowMask		= 1 << 2,
        NSResizableWindowMask		= 1 << 3
};

enum {
	NSOpenGLPFAAllRenderers       =   1,	/* choose from all available renderers          */
	NSOpenGLPFADoubleBuffer       =   5,	/* choose a double buffered pixel format        */
	NSOpenGLPFAStereo             =   6,	/* stereo buffering supported                   */
	NSOpenGLPFAAuxBuffers         =   7,	/* number of aux buffers                        */
	NSOpenGLPFAColorSize          =   8,	/* number of color buffer bits                  */
	NSOpenGLPFAAlphaSize          =  11,	/* number of alpha component bits               */
	NSOpenGLPFADepthSize          =  12,	/* number of depth buffer bits                  */
	NSOpenGLPFAStencilSize        =  13,	/* number of stencil buffer bits                */
	NSOpenGLPFAAccumSize          =  14,	/* number of accum buffer bits                  */
	NSOpenGLPFAMinimumPolicy      =  51,	/* never choose smaller buffers than requested  */
	NSOpenGLPFAMaximumPolicy      =  52,	/* choose largest buffers of type requested     */
	NSOpenGLPFAOffScreen          =  53,	/* choose an off-screen capable renderer        */
	NSOpenGLPFAFullScreen         =  54,	/* choose a full-screen capable renderer        */
	NSOpenGLPFASampleBuffers      =  55,	/* number of multi sample buffers               */
	NSOpenGLPFASamples            =  56,	/* number of samples per multi sample buffer    */
	NSOpenGLPFAAuxDepthStencil    =  57,	/* each aux buffer has its own depth stencil    */
	NSOpenGLPFAColorFloat         =  58,	/* color buffers store floating point pixels    */
#if MAC_OS_X_VERSION_MAX_ALLOWED >= MAC_OS_X_VERSION_10_4
	NSOpenGLPFAMultisample        =  59,    /* choose multisampling                         */
	NSOpenGLPFASupersample        =  60,    /* choose supersampling                         */
	NSOpenGLPFASampleAlpha        =  61,    /* request alpha filtering                      */
#endif
	NSOpenGLPFARendererID         =  70,	/* request renderer by ID                       */
	NSOpenGLPFASingleRenderer     =  71,	/* choose a single renderer for all screens     */
	NSOpenGLPFANoRecovery         =  72,	/* disable all failure recovery systems         */
	NSOpenGLPFAAccelerated        =  73,	/* choose a hardware accelerated renderer       */
	NSOpenGLPFAClosestPolicy      =  74,	/* choose the closest color buffer to request   */
	NSOpenGLPFARobust             =  75,	/* renderer does not need failure recovery      */
	NSOpenGLPFABackingStore       =  76,	/* back buffer contents are valid after swap    */
	NSOpenGLPFAMPSafe             =  78,	/* renderer is multi-processor safe             */
	NSOpenGLPFAWindow             =  80,	/* can be used to render to an onscreen window  */
	NSOpenGLPFAMultiScreen        =  81,	/* single window can span multiple screens      */
	NSOpenGLPFACompliant          =  83,	/* renderer is opengl compliant                 */
	NSOpenGLPFAScreenMask         =  84,	/* bit mask of supported physical screens       */
	NSOpenGLPFAPixelBuffer        =  90,	/* can be used to render to a pbuffer           */
	NSOpenGLPFARemotePixelBuffer  =  91,	/* can be used to render offline to a pbuffer   */
#if MAC_OS_X_VERSION_MAX_ALLOWED >= MAC_OS_X_VERSION_10_5
	NSOpenGLPFAAllowOfflineRenderers = 96,  /* allow use of offline renderers               */
#endif
	NSOpenGLPFAAcceleratedCompute =  97,	/* choose a hardware accelerated compute device */
	NSOpenGLPFAVirtualScreenCount = 128	/* number of virtual screens in this format     */
};

enum {
        NSBackingStoreRetained	 = 0,
        NSBackingStoreNonretained	 = 1,
        NSBackingStoreBuffered	 = 2
};


typedef enum {
	NSOpenGLCPSwapRectangle       = 200,	/* Set or get the swap rectangle {x, y, w, h}       */
	NSOpenGLCPSwapRectangleEnable = 201,	/* Enable or disable the swap rectangle             */
	NSOpenGLCPRasterizationEnable = 221,	/* Enable or disable all rasterization              */
	NSOpenGLCPSwapInterval        = 222,	/* 0 -> Don't sync, 1 -> Sync to vertical retrace   */
#if MAC_OS_X_VERSION_MAX_ALLOWED >= MAC_OS_X_VERSION_10_2
	NSOpenGLCPSurfaceOrder        = 235,	/* 1 -> Above Window (default), -1 -> Below Window  */
	NSOpenGLCPSurfaceOpacity      = 236,	/* 1-> Surface is opaque (default), 0 -> non-opaque */
#endif
	NSOpenGLCPStateValidation     = 301	/* Validate state for multi-screen functionality    */
} NSOpenGLContextParameter;

#define  NSAnyEventMask (~0)

extern NSString * const NSDefaultRunLoopMode;
extern NSString *NSDeviceRGBColorSpace;


static int g_done = 0;

#define dbg cvReport
#define err cvReport

typedef struct Window {
	char storage[8192];
} Window;

typedef struct View {
	char storage[8192];
	unsigned int _prevflags;
} View;


static id g_win;
static id g_view;
static NSRect g_rect;
static id g_cur = 0;

static void setWindowMode(int style, NSRect rect) {
        send(g_win, sel("setStyleMask:"), style);
        send(g_win, sel("setFrame:display:"), rect, TRUE);
        send(g_win, sel("setLevel:"), (NSInteger)CGWindowLevelForKey(
                     style == FULLSCREEN_MASK? 
                     kCGPopUpMenuWindowLevelKey : kCGNormalWindowLevelKey));
        send(g_win, sel("makeFirstResponder:"), g_view);
}

static id alloc(const char * class) {
        return send(cls(class), sel("alloc"));
}

static void rel(id o) {
        send(o, sel("release"));
}

static void relcursor() {
        rel(g_cur);
        g_cur = 0;
        send(g_win, sel("invalidateCursorRectsForView:"), g_view);
}

static id curcursor() {
        return g_cur? g_cur : send(cls("NSCursor"), sel("arrowCursor"));
}

static void setcursor(uint8_t * crgba, int16_t hotx, int16_t hoty) {
        static uint8_t srgba[32*32*4];
        uint8_t * rgba = srgba;
        id img, b;
        relcursor();
        memcpy(srgba, crgba, sizeof(srgba));
        b = alloc("NSBitmapImageRep");
        send(b, sel("initWithBitmapDataPlanes:"
                    "pixelsWide:"
                    "pixelsHigh:"
                    "bitsPerSample:"
                    "samplesPerPixel:"
                    "hasAlpha:"
                    "isPlanar:"
                    "colorSpaceName:"
                    "bytesPerRow:"
                    "bitsPerPixel:"),
             &rgba,
             (NSInteger)32,
             (NSInteger)32,
             (NSInteger)8,
             (NSInteger)4,
             YES,
             NO,
             NSDeviceRGBColorSpace,
             (NSInteger)32*4,
             (NSInteger)32);
        img = alloc("NSImage");
        send(img, sel("initWithSize:"), mksize(32, 32));
        send(img, sel("addRepresentation:"), b);
        rel(b);
        g_cur = alloc("NSCursor");
        send(g_cur, sel("initWithImage:hotSpot:"), img , mkpoint(hotx, hoty));
        rel(img);
}

static NSRect scrFrame() {
        return rsend(send(send(cls("NSScreen"), sel("screens")),
                          sel("objectAtIndex:"), (NSUInteger)0),
                     sel("frame"));
}

intptr_t osEvent(ev * e) {
        static int fullscreen = 0;
        intptr_t ret = 1;
        switch (evType(e)) {
        case CVE_QUIT: g_done = 1; break;
        case CVE_SETCURSOR: 
                setcursor((uint8_t*)evArg0(e), 
                          evArg1(e) >> 16, evArg1(e) & 0xffff); 
                break;
        case CVE_DEFAULTCURSOR: 
                relcursor();
                break;
        case CVE_FULLSCREEN:
                if (!fullscreen) {
                        g_rect = rsend(g_win, sel("frame"));
                        setWindowMode(FULLSCREEN_MASK, scrFrame());
                        fullscreen = 1;
                }
                break;
        case CVE_WINDOWED: 
                if (fullscreen) {
                        setWindowMode(WINDOWED_MASK, g_rect);
                        fullscreen = 0;
                }
                break;
        default: ret = 0;
        }
        return ret;
}

static cvkey mapkeycode(unsigned k) {
        cvkey ret;
        switch (k) {
        case kVK_ANSI_A: ret = CVK_A; break;
        case kVK_ANSI_S: ret = CVK_S; break;
        case kVK_ANSI_D: ret = CVK_D; break;
        case kVK_ANSI_F: ret = CVK_F; break;
        case kVK_ANSI_H: ret = CVK_H; break;
        case kVK_ANSI_G: ret = CVK_G; break;
        case kVK_ANSI_Z: ret = CVK_Z; break;
        case kVK_ANSI_X: ret = CVK_X; break;
        case kVK_ANSI_C: ret = CVK_C; break;
        case kVK_ANSI_V: ret = CVK_V; break;
        case kVK_ANSI_B: ret = CVK_B; break;
        case kVK_ANSI_Q: ret = CVK_Q; break;
        case kVK_ANSI_W: ret = CVK_W; break;
        case kVK_ANSI_E: ret = CVK_E; break;
        case kVK_ANSI_R: ret = CVK_R; break;
        case kVK_ANSI_Y: ret = CVK_Y; break;
        case kVK_ANSI_T: ret = CVK_T; break;
        case kVK_ANSI_1: ret = CVK_1; break;
        case kVK_ANSI_2: ret = CVK_2; break;
        case kVK_ANSI_3: ret = CVK_3; break;
        case kVK_ANSI_4: ret = CVK_4; break;
        case kVK_ANSI_6: ret = CVK_6; break;
        case kVK_ANSI_5: ret = CVK_5; break;
        case kVK_ANSI_Equal: ret = CVK_EQUAL; break;
        case kVK_ANSI_9: ret = CVK_9; break;
        case kVK_ANSI_7: ret = CVK_7; break;
        case kVK_ANSI_Minus: ret = CVK_MINUS; break;
        case kVK_ANSI_8: ret = CVK_8; break;
        case kVK_ANSI_0: ret = CVK_0; break;
        case kVK_ANSI_RightBracket: ret = CVK_RIGHTBRACKET; break;
        case kVK_ANSI_O: ret = CVK_O; break;
        case kVK_ANSI_U: ret = CVK_U; break;
        case kVK_ANSI_LeftBracket: ret = CVK_LEFTBRACKET; break;
        case kVK_ANSI_I: ret = CVK_I; break;
        case kVK_ANSI_P: ret = CVK_P; break;
        case kVK_ANSI_L: ret = CVK_L; break;
        case kVK_ANSI_J: ret = CVK_J; break;
        case kVK_ANSI_Quote: ret = CVK_QUOTE; break;
        case kVK_ANSI_K: ret = CVK_K; break;
        case kVK_ANSI_Semicolon: ret = CVK_SEMICOLON; break;
        case kVK_ANSI_Backslash: ret = CVK_BACKSLASH; break;
        case kVK_ANSI_Comma: ret = CVK_COMMA; break;
        case kVK_ANSI_Slash: ret = CVK_SLASH; break;
        case kVK_ANSI_N: ret = CVK_N; break;
        case kVK_ANSI_M: ret = CVK_M; break;
        case kVK_ANSI_Period: ret = CVK_PERIOD; break;
        case kVK_ANSI_Grave: ret = CVK_GRAVE; break;
        case kVK_ANSI_KeypadDecimal: ret = CVK_KEYPADDECIMAL; break;
        case kVK_ANSI_KeypadMultiply: ret = CVK_KEYPADMULTIPLY; break;
        case kVK_ANSI_KeypadPlus: ret = CVK_KEYPADPLUS; break;
        case kVK_ANSI_KeypadClear: ret = CVK_KEYPADCLEAR; break;
        case kVK_ANSI_KeypadDivide: ret = CVK_KEYPADDIVIDE; break;
        case kVK_ANSI_KeypadEnter: ret = CVK_KEYPADENTER; break;
        case kVK_ANSI_KeypadMinus: ret = CVK_KEYPADMINUS; break;
        case kVK_ANSI_KeypadEquals: ret = CVK_KEYPADEQUALS; break;
        case kVK_ANSI_Keypad0: ret = CVK_KEYPAD0; break;
        case kVK_ANSI_Keypad1: ret = CVK_KEYPAD1; break;
        case kVK_ANSI_Keypad2: ret = CVK_KEYPAD2; break;
        case kVK_ANSI_Keypad3: ret = CVK_KEYPAD3; break;
        case kVK_ANSI_Keypad4: ret = CVK_KEYPAD4; break;
        case kVK_ANSI_Keypad5: ret = CVK_KEYPAD5; break;
        case kVK_ANSI_Keypad6: ret = CVK_KEYPAD6; break;
        case kVK_ANSI_Keypad7: ret = CVK_KEYPAD7; break;
        case kVK_ANSI_Keypad8: ret = CVK_KEYPAD8; break;
        case kVK_ANSI_Keypad9: ret = CVK_KEYPAD9; break;
        case kVK_Return: ret = CVK_RETURN; break;
        case kVK_Tab: ret = CVK_TAB; break;
        case kVK_Space: ret = CVK_SPACE; break;
        case kVK_Delete: ret = CVK_DELETE; break;
        case kVK_Escape: ret = CVK_ESCAPE; break;
        case kVK_Command: ret = CVK_COMMAND; break;
        case kVK_Shift: ret = CVK_SHIFT; break;
        case kVK_CapsLock: ret = CVK_CAPSLOCK; break;
        case kVK_Option: ret = CVK_OPTION; break;
        case kVK_Control: ret = CVK_CONTROL; break;
        case kVK_RightShift: ret = CVK_RIGHTSHIFT; break;
        case kVK_RightOption: ret = CVK_RIGHTOPTION; break;
        case kVK_RightControl: ret = CVK_RIGHTCONTROL; break;
        case kVK_Function: ret = CVK_FUNCTION; break;
        case kVK_F17: ret = CVK_F17; break;
        case kVK_VolumeUp: ret = CVK_VOLUMEUP; break;
        case kVK_VolumeDown: ret = CVK_VOLUMEDOWN; break;
        case kVK_Mute: ret = CVK_MUTE; break;
        case kVK_F18: ret = CVK_F18; break;
        case kVK_F19: ret = CVK_F19; break;
        case kVK_F20: ret = CVK_F20; break;
        case kVK_F5: ret = CVK_F5; break;
        case kVK_F6: ret = CVK_F6; break;
        case kVK_F7: ret = CVK_F7; break;
        case kVK_F3: ret = CVK_F3; break;
        case kVK_F8: ret = CVK_F8; break;
        case kVK_F9: ret = CVK_F9; break;
        case kVK_F11: ret = CVK_F11; break;
        case kVK_F13: ret = CVK_F13; break;
        case kVK_F16: ret = CVK_F16; break;
        case kVK_F14: ret = CVK_F14; break;
        case kVK_F10: ret = CVK_F10; break;
        case kVK_F12: ret = CVK_F12; break;
        case kVK_F15: ret = CVK_F15; break;
        case kVK_Help: ret = CVK_HELP; break;
        case kVK_Home: ret = CVK_HOME; break;
        case kVK_PageUp: ret = CVK_PAGEUP; break;
        case kVK_ForwardDelete: ret = CVK_FORWARDDELETE; break;
        case kVK_F4: ret = CVK_F4; break;
        case kVK_End: ret = CVK_END; break;
        case kVK_F2: ret = CVK_F2; break;
        case kVK_PageDown: ret = CVK_PAGEDOWN; break;
        case kVK_F1: ret = CVK_F1; break;
        case kVK_LeftArrow: ret = CVK_LEFTARROW; break;
        case kVK_RightArrow: ret = CVK_RIGHTARROW; break;
        case kVK_DownArrow: ret = CVK_DOWNARROW; break;
        case kVK_UpArrow: ret = CVK_UPARROW; break;
        default: ret = CVK_NONE;
        }
        return ret;
}

static void _I_View_dealloc(id self, SEL _cmd) {
	struct objc_super x = { self, class_getSuperclass((Class)cls("View")) };
        ssend(&x, sel("dealloc"));
}

static NSPoint _I_View_toAbs_(id self, SEL _cmd, NSPoint p) {
	id x = send(self, sel("window"), sel("convertBaseToScreen:"), (NSPoint)p);
	return *(NSPoint*)&x;
}

static void _I_View_setConstrainedFrameSize_(id self, SEL _cmd, NSSize desiredSize) {}

static BOOL _I_View_acceptsFirstResponder(id self, SEL _cmd) {
        return YES;
}

static void _I_View_windowDidResize_(id self, SEL _cmd, id n) {
	NSRect fr = rsend(self, sel("frame"));
        NSSize sz = fr.size;
        send(send(cls("NSOpenGLContext"), sel("currentContext")), sel("update"));
        cvInject(CVE_RESIZE, sz.width, sz.height);
}

static BOOL _I_View_windowShouldClose_(id self, SEL _cmd, id s) {
        cvInject(CVE_CLOSE, 0, 0);
        return NO;
}

static void _I_View_handleKeyEvent_mode_(id self, SEL _cmd, id ev, int mode) {
        cvInject(mode, mapkeycode((uintptr_t)send(ev, sel("keyCode"))), 0);
}

static void _I_View_keyUp_(id self, SEL _cmd, id ev) {
        send(self, sel("handleKeyEvent:mode:"), ev, CVE_UP);
}

static void _I_View_keyDown_(id self, SEL _cmd, id ev) {
        send(self, sel("handleKeyEvent:mode:"), ev, CVE_DOWN);
        send(self, sel("interpretKeyEvents:"), send(cls("NSArray"), sel("arrayWithObject:"), ev));
}

static void _I_View_deleteBackward_(id self, SEL _cmd, id sender) {
        cvInject(CVE_UNICODE, CVK_DELETE, 0);
}

static void _I_View_insertText_(id self, SEL _cmd, id s) {
        int sl = (intptr_t)send(s, sel("length"));
        int i;
        for (i = 0; i < sl; i++)
                cvInject(CVE_UNICODE, (uintptr_t)send(s, sel("characterAtIndex:"), (NSUInteger)i), 0);
}

static unsigned int _I_View_keyFor_(id self, SEL _cmd, unsigned int mask) {
        unsigned ret = 0;
        switch (mask) {
        case NSShiftKeyMask: ret = CVK_SHIFT; break;
        case NSControlKeyMask: ret = CVK_CONTROL; break;
        case NSAlternateKeyMask: ret = CVK_OPTION; break;
        case NSCommandKeyMask: ret = CVK_COMMAND; break;
        default: assert(0);
        }
        return ret;
}

static void _I_View_handleMod_flags_(id self, SEL _cmd, unsigned int mask, unsigned int flags) {
        View * v = (View*)self;
        unsigned delta = flags ^ (v->_prevflags);
        unsigned pressed = delta & flags;
        unsigned released = delta & ~flags;
        if (mask & pressed)
                cvInject(CVE_DOWN, (uintptr_t)send(self, sel("keyFor:"), (unsigned int)mask), 0);
        if (mask & released)
                cvInject(CVE_UP, (uintptr_t)send(self, sel("keyFor:"), (unsigned int)mask), 0);
}

static void _I_View_flagsChanged_(id self, SEL _cmd, id ev) {
        View * v = (View*)self;
        unsigned flags = (uintptr_t)send(ev, sel("modifierFlags"));
        send(self, sel("handleMod:flags:"), (unsigned int)NSShiftKeyMask, (unsigned int)flags);
        send(self, sel("handleMod:flags:"), (unsigned int)NSControlKeyMask, (unsigned int)flags);
        send(self, sel("handleMod:flags:"), (unsigned int)NSAlternateKeyMask, (unsigned int)flags);
        send(self, sel("handleMod:flags:"), (unsigned int)NSCommandKeyMask, (unsigned int)flags);
	v->_prevflags = flags;
	struct objc_super x  = { self, class_getSuperclass((Class)cls("View")) };
        ssend(&x, sel("flagsChanged:"), ev);
}

static void _I_View_handleMouseEvent_mode_(id self, SEL _cmd, id ev, int mode) {
        cvInject(mode, (intptr_t)send(ev, sel("buttonNumber")) + CVK_MOUSELEFT, 0);
}

static void _I_View_mouseDown_(id self, SEL _cmd, id ev) {
        send(self, sel("handleMouseEvent:mode:"), ev, CVE_DOWN);
}

static void _I_View_mouseUp_(id self, SEL _cmd, id ev) {
        send(self, sel("handleMouseEvent:mode:"), ev, CVE_UP);
}

static void _I_View_rightMouseDown_(id self, SEL _cmd, id ev) {
        send(self, sel("handleMouseEvent:mode:"), ev, CVE_DOWN);
}

static void _I_View_rightMouseUp_(id self, SEL _cmd, id ev) {
        send(self, sel("handleMouseEvent:mode:"), ev, CVE_UP);
}

static void _I_View_otherMouseDown_(id self, SEL _cmd, id ev) {
        send(self, sel("handleMouseEvent:mode:"), ev, CVE_DOWN);
}

static void _I_View_otherMouseUp_(id self, SEL _cmd, id ev) {
        send(self, sel("handleMouseEvent:mode:"), ev, CVE_UP);
}

static void _I_View_scrollWheel_(id self, SEL _cmd, id ev) {
        unsigned k = (CGFloat)fsend(ev, sel("deltaY")) > 0? 
                CVK_MOUSEWHEELUP : CVK_MOUSEWHEELDOWN;
        cvInject(CVE_DOWN, k, 0);
        cvInject(CVE_UP, k, 0);
}

static void _I_View_handleMotion_(id self, SEL _cmd, id ev) {
        NSPoint l = psend(ev, sel("locationInWindow"));
        cvInject(CVE_MOTION, l.x, rsend(self, sel("frame")).size.height - l.y);
}

static void _I_View_mouseMoved_(id self, SEL _cmd, id ev) {
        send(self, sel("handleMotion:"), ev);
}

static BOOL _I_View_isOpaque(id self, SEL _cmd) {
        return YES;
}

static void _I_View_resetCursorRects(id self) {
        id cur = curcursor();
        send(self, sel("discardCursorRects"));
        send(self, sel("addCursorRect:cursor:"), 
             rsend(self, sel("visibleRect")), cur);
        send(cur, sel("set"));
}

static void _I_Window_dealloc(id self, SEL _cmd) {
	struct objc_super x =  {self, class_getSuperclass((Class)cls("Window")) };
        ssend(&x, sel("dealloc"));
}

static BOOL _I_Window_canBecomeKeyWindow(id self, SEL _cmd) {
        return YES;
}


static BOOL _I_Window_isReleasedWhenClosed(id self, SEL _cmd) {
        return NO;
}

static void init() {
	Class win = objc_allocateClassPair((Class)cls("NSWindow"), "Window", 0);
	objc_registerClassPair(win);
#define am class_addMethod

	am(win, sel("dealloc"), (IMP)_I_Window_dealloc, "v16@0:8");
	am(win, sel("canBecomeKeyWindow"), (IMP)_I_Window_canBecomeKeyWindow, "c16@0:8");
	am(win, sel("isReleasedWhenClosed"), (IMP)_I_Window_isReleasedWhenClosed, "c16@0:8");

	
	Class view = objc_allocateClassPair((Class)cls("NSTextView"), "View", 0);
	objc_registerClassPair(view);

	am(view, sel("dealloc"), (IMP)_I_View_dealloc, "v16@0:8");
	am(view, sel("toAbs:"), (IMP)_I_View_toAbs_, "{CGPoint=dd}32@0:8{CGPoint=dd}16");
	am(view, sel("setConstrainedFrameSize:"), (IMP)_I_View_setConstrainedFrameSize_, "v32@0:8{CGSize=dd}16");
	am(view, sel("acceptsFirstResponder"), (IMP)_I_View_acceptsFirstResponder, "c16@0:8");
	am(view, sel("windowDidResize:"), (IMP)_I_View_windowDidResize_, "v24@0:8@16");
	am(view, sel("windowShouldClose:"), (IMP)_I_View_windowShouldClose_, "c24@0:8@16");
	am(view, sel("handleKeyEvent:mode:"), (IMP)_I_View_handleKeyEvent_mode_, "v28@0:8@16i24");
	am(view, sel("keyUp:"), (IMP)_I_View_keyUp_, "v24@0:8@16");
	am(view, sel("keyDown:"), (IMP)_I_View_keyDown_, "v24@0:8@16");
	am(view, sel("deleteBackward:"), (IMP)_I_View_deleteBackward_, "v24@0:8@16");
	am(view, sel("insertText:"), (IMP)_I_View_insertText_, "v24@0:8@16");
	am(view, sel("keyFor:"), (IMP)_I_View_keyFor_, "I20@0:8I16");
	am(view, sel("handleMod:flags:"), (IMP)_I_View_handleMod_flags_, "v24@0:8I16I20");
	am(view, sel("flagsChanged:"), (IMP)_I_View_flagsChanged_, "v24@0:8@16");
	am(view, sel("handleMouseEvent:mode:"), (IMP)_I_View_handleMouseEvent_mode_, "v28@0:8@16i24");
	am(view, sel("mouseDown:"), (IMP)_I_View_mouseDown_, "v24@0:8@16");
	am(view, sel("mouseUp:"), (IMP)_I_View_mouseUp_, "v24@0:8@16");
	am(view, sel("rightMouseDown:"), (IMP)_I_View_rightMouseDown_, "v24@0:8@16");
	am(view, sel("rightMouseUp:"), (IMP)_I_View_rightMouseUp_, "v24@0:8@16");
	am(view, sel("otherMouseDown:"), (IMP)_I_View_otherMouseDown_, "v24@0:8@16");
	am(view, sel("otherMouseUp:"), (IMP)_I_View_otherMouseUp_, "v24@0:8@16");
	am(view, sel("scrollWheel:"), (IMP)_I_View_scrollWheel_, "v24@0:8@16");
	am(view, sel("handleMotion:"), (IMP)_I_View_handleMotion_, "v24@0:8@16");
	am(view, sel("mouseMoved:"), (IMP)_I_View_mouseMoved_, "v24@0:8@16");
	am(view, sel("isOpaque"), (IMP)_I_View_isOpaque, "c16@0:8");
        am(view, sel("resetCursorRects"), (IMP)_I_View_resetCursorRects, "v16@0:8");
#undef am

}

int cvrun(int argc, char ** argv) {
        id arp = send(alloc("NSAutoreleasePool"), sel("init"));
        ProcessSerialNumber psn = { 0, kCurrentProcess };
        NSRect frm;
        id win;
        id view;
        int style = WINDOWED_MASK;
        GLint param = 1;
        id ctx = 0;
        CGDirectDisplayID dpy = kCGDirectMainDisplay;
        NSOpenGLPixelFormatAttribute attr[] = {
                NSOpenGLPFAFullScreen,
                NSOpenGLPFAScreenMask, CGDisplayIDToOpenGLDisplayMask(dpy),
                NSOpenGLPFAColorSize, 24,
                NSOpenGLPFADepthSize, 32,
                NSOpenGLPFAAlphaSize, 8,
                NSOpenGLPFADoubleBuffer,
                NSOpenGLPFAAccelerated,
                NSOpenGLPFANoRecovery,
                0, 0, 0, 0, 0, 0, 0, 0,
        };
        id fmt;
        NSRect rect;
        id app;
        puts("");
        TransformProcessType(&psn, kProcessTransformToForegroundApplication);
        SetFrontProcess(&psn);
	init();

        app = send(cls("NSApplication"), sel("sharedApplication"));
        send(app, sel("activateIgnoringOtherApps:"), YES);
        send(app, sel("finishLaunching"));

        cvInject(CVE_INIT, argc, (intptr_t)argv);
        rect.origin.x = cvInject(CVQ_XPOS, 0, 0);
        rect.origin.y = cvInject(CVQ_YPOS, 0, 0);
        rect.size.width = cvInject(CVQ_WIDTH, 0, 0);
        rect.size.height = cvInject(CVQ_HEIGHT, 0, 0);
        rect.origin.y = scrFrame().size.height - 1 - rect.origin.y - rect.size.height;
        win = send(alloc("Window"), sel("initWithContentRect:styleMask:backing:defer:"), (NSRect)rect, (NSUInteger)style, (NSBackingStoreType)NSBackingStoreBuffered, NO);
        frm = rsend2(cls("Window"), sel("contentRectForFrameRect:styleMask:"), rsend(win, sel("frame")), (NSUInteger)style);
        view = alloc("View");
	view = send(view, sel("initWithFrame:"), (NSRect)frm);
        send(view, sel("setSelectable:"), NO);
        g_win = win;
        g_view = view;
        send(win, sel("setReleasedWhenClosed:"), NO);
        send(win, sel("setDelegate:"), view);
        send(win, sel("setContentView:"), view);
        send(win, sel("setAcceptsMouseMovedEvents:"), YES);
        send(win, sel("makeKeyAndOrderFront:"), view);
        send(win, sel("makeFirstResponder:"), view);

        fmt = send(alloc("NSOpenGLPixelFormat"), sel("initWithAttributes:"), (const NSOpenGLPixelFormatAttribute *)attr);
        ctx = send(alloc("NSOpenGLContext"), sel("initWithFormat:shareContext:"), fmt, NIL);
        rel(fmt);
        send(ctx, sel("setView:"), view);
        send(ctx, sel("setValues:forParameter:"), (const GLint *)&param, (NSOpenGLContextParameter)NSOpenGLCPSwapInterval);
        send(ctx, sel("makeCurrentContext"));
        cvInject(CVE_GLINIT, 0, 0);
        rect = rsend(view, sel("frame"));
        cvInject(CVE_RESIZE, rect.size.width, rect.size.height);
        send(arp, sel("drain"));
        do {
                arp = send(alloc("NSAutoreleasePool"), sel("init"));
                id ev = send(app, sel("nextEventMatchingMask:untilDate:inMode:dequeue:"), NSAnyEventMask, send(cls("NSDate"), sel("distantPast")), NSDefaultRunLoopMode, YES);
                send(app, sel("sendEvent:"), ev);
                cvInject(CVE_UPDATE, 0, 0);
                send(arp, sel("drain"));
                send(ctx, sel("flushBuffer"));
        } while (!g_done);
        arp = send(alloc("NSAutoreleasePool"), sel("init"));
        cvInject(CVE_GLTERM, 0, 0);
        send(ctx, sel("clearDrawable"));
        rel(ctx);
        send(win, sel("makeKeyAndOrderFront:"), NIL);
        send(win, sel("makeFirstResponder:"), NIL);
        send(win, sel("setDelegate:"), NIL);
        send(win, sel("setContentView:"), NIL);
        send(win, sel("close"));
        rel(win);
        rel(view);
        send(arp, sel("drain"));
        return cvInject(CVE_TERM, 0, 0);
}
/* 
   Local variables: **
   c-file-style: "bsd" **
   c-basic-offset: 8 **
   indent-tabs-mode: nil **
   End: **
*/
