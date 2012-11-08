package cv

type Responder interface {
	OnInit()                    // The OpenGL context is about to be created.
	OnTerm()                    // The OpenGL context is has been destroyed.
	OnGLInit()                  // The OpenGL context has been created.
	OnGLTerm()                  // The OpenGL context is about to be destroyed.
	OnPress(k int)              // The key (KEY_*) has been pressed.
	OnRelease(k int)            // The key (KEY_*) has been released.
	OnUnicode(u uint)           // The unicode character has been entered.
	OnMotion(x, y int)          // The mouse has moved.
	OnClose()                   // The window close button has been pressed.
	OnResize(w, h int)          // The window has been resized.
	OnPreUpdate()               // Called once per frame before OnUpdate().
	OnUpdate()                  // Called once per frame.
	OnPostUpdate()              // Called once per frame after OnUpdate().
	Name() string               // Should return the desired name for the app.
	Geometry() (x, y, w, h int) // Should return a hint for the desired geometry.
	Borders() bool              // Should the window have borders?
}

// The canvas structure, you should embed it in your derived canvas.
type Canvas struct {
	r Responder
}

var g_canvas *Canvas

// Initialize the canvas with the given event responder (usually the
// derived canvas).
func (c *Canvas) InitCanvas(r Responder) {
	c.r = r
	g_canvas = c
}

// Returns the next event responder in the chain. Useful if you want
// to call the parent canvas in one of your handlers.
func (c *Canvas) NextResponder() Responder {
	return c
}

// Start the event loop.
func (c *Canvas) Go() {
	Run(event)
}

// Request the application to leave.
func (c *Canvas) Quit() {
	Quit()
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

func (c *Canvas) press(k int) {
	c.r.OnPress(k)
}

func (c *Canvas) OnPress(k int) {}

func (c *Canvas) release(k int) {
	c.r.OnRelease(k)
}

func (c *Canvas) OnRelease(k int) {}

func (c *Canvas) unicode(u uint) {
	c.r.OnUnicode(u)
}

func (c *Canvas) OnUnicode(u uint) {}

func (c *Canvas) motion(x, y int) {
	c.r.OnMotion(x, y)
}

func (c *Canvas) OnMotion(x, y int) {}

func (c *Canvas) update() {
	c.r.OnPreUpdate()
	c.r.OnUpdate()
	c.r.OnPostUpdate()
}

func (c *Canvas) OnUpdate() {}

func (c *Canvas) OnPreUpdate() {}

func (c *Canvas) OnPostUpdate() {}

func (c *Canvas) name() uintptr {
	return ctring(c.r.Name())
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
	return Width(), Height()
}

func (c *Canvas) Width() int {
	return Width()
}

func (c *Canvas) Height() int {
	return Height()
}

func (c *Canvas) Mouse() (int, int) {
	return MouseX(), MouseY()
}

func (c *Canvas) Pressed(k int) bool {
	return Pressed(k)
}

func (c *Canvas) Released(k int) bool {
	return Released(k)
}

func event(e *Event) uintptr {
	r := uintptr(1)
	switch e.Type() {
	case ON_INIT:
		g_canvas.init()
	case ON_TERM:
		g_canvas.term()
	case ON_GLINIT:
		g_canvas.glinit()
	case ON_DOWN:
		g_canvas.press(e.Which())
	case ON_UP:
		g_canvas.release(e.Which())
	case ON_UNICODE:
		g_canvas.unicode(e.Unicode())
	case ON_MOTION:
		g_canvas.motion(e.X(), e.Y())
	case ON_CLOSE:
		g_canvas.close()
	case ON_RESIZE:
		g_canvas.resize(e.Width(), e.Height())
	case ON_UPDATE:
		g_canvas.update()
	case HINT_NAME:
		r = g_canvas.name()
	case HINT_BORDERS:
		r = g_canvas.borders()
	case HINT_XPOS:
		r, _, _, _ = g_canvas.geometry()
	case HINT_YPOS:
		_, r, _, _ = g_canvas.geometry()
	case HINT_WIDTH:
		_, _, r, _ = g_canvas.geometry()
	case HINT_HEIGHT:
		_, _, _, r = g_canvas.geometry()
	default:
		r = 0
	}
	return r
}
