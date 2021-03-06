package main

import (
	"github.com/jacereda/glcv-go/canvas"
	"github.com/jacereda/glcv-go/key"
	gl "github.com/chsc/gogl/gl21"
	"log"
	"time"
)

type testCanvas struct {
	canvas.Canvas
}

func (c *testCanvas) OnInit() {
	log.Println("Init")
}

func (c *testCanvas) OnTerm() {
	log.Println("Term")
}

func (c *testCanvas) OnGLInit() {
	log.Println("GLInit")
	gl.Init()
}

func (c *testCanvas) OnGLTerm() {
	log.Println("GLTerm")
}

func (c *testCanvas) OnPress(k key.Id) {
	log.Println("Press", c.KeyName(k))
}

func (c *testCanvas) OnRelease(k key.Id) {
	log.Println("Release", c.KeyName(k))
	switch k {
	case key.ESCAPE:
		c.Quit()
	case key.S:
		c.DefaultCursor()
	case key.H:
		c.HideCursor()
	case key.F:
		c.Fullscreen()
	case key.W:
		c.Windowed()
	}
}

func (c *testCanvas) OnUnicode(u uint) {
	log.Printf("Unicode %c", u)
}

func (c *testCanvas) OnMotion(x, y int) {
	c.NextResponder().OnMotion(x, y)
	log.Println("Motion", x, y)
}

func (c *testCanvas) OnClose() {
	log.Println("Close")
	c.Quit()
}

func (c *testCanvas) OnResize(w, h int) {
	log.Println("Resize", w, h)
	gl.Viewport(0, 0, gl.Sizei(w), gl.Sizei(h))
}

func (c *testCanvas) OnUpdate() {
	if c.Pressed(key.A) {
		log.Println("A pressed")
	}
	if c.Released(key.A) {
		log.Println("A released")
	}
	n := gl.Float(time.Now().Nanosecond()) / 1000000000
	gl.ClearColor(0, n, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (c *testCanvas) Name() string {
	return "test"
}

func (c *testCanvas) Geometry() (x, y, w, h int) {
	return 20, 20, 512, 512
}

func main() {
	var c testCanvas
	c.InitCanvas(&c)
	c.Go()
}
