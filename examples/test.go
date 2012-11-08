package main

import (
	"code.google.com/p/glcv-go/cv"
	gl "github.com/chsc/gogl/gl21"
	"log"
)

type testCanvas struct {
	cv.Canvas
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

func (c *testCanvas) OnPress(k int) {
	log.Println("Press", cv.KeyName(k))
	if k == cv.KEY_ESCAPE {
		c.Quit()
	}
}

func (c *testCanvas) OnRelease(k int) {
	log.Println("Release", cv.KeyName(k))
}

func (c *testCanvas) OnUnicode(u uint) {
	log.Printf("Unicode %c", u)
}

func (c *testCanvas) OnMotion(x, y int) {
	c.NextResponder().OnMotion(x, y)
	log.Println("Motion", x, y)
}

func (c *testCanvas) OnResize(w, h int) {
	log.Println("Resize", w, h)
}

func (c *testCanvas) OnUpdate() {
	if c.Pressed(cv.KEY_A) {
		log.Println("A pressed")
	}
	if c.Released(cv.KEY_A) {
		log.Println("A released")
	}
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (c *testCanvas) Name() string {
	return "test"
}

func (c *testCanvas) Geometry() (x, y, w, h int) {
	return 20, 20, 512, 512
}

func (c *testCanvas) Borders() bool {
	return true
}

func main() {
	var c testCanvas
	c.InitCanvas(&c)
	c.Go()
}
