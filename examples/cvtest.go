package main

import (
	"github.com/jacereda/glcv-go/cv"
	"github.com/jacereda/glcv-go/key"
	gl "github.com/chsc/gogl/gl21"
	"log"
	"time"
)

func handler(e *cv.Event) uintptr {
	switch e.Type() {
	default:
		return 0
	case cv.HINT_NAME:
		return cv.String("cvtest")
	case cv.HINT_XPOS:
		return 20
	case cv.HINT_YPOS:
		return 20
	case cv.HINT_WIDTH:
		return 640
	case cv.HINT_HEIGHT:
		return 480
	case cv.ON_INIT:
		log.Println("INIT")
	case cv.ON_TERM:
		log.Println("TERM")
	case cv.ON_GLINIT:
		log.Println("GLINIT")
		gl.Init()
	case cv.ON_GLTERM:
		log.Println("GLTERM")
	case cv.ON_DOWN:
		log.Println("DOWN", cv.KeyName(e.Which()))
		switch e.Which() {
		case key.ESCAPE:
			cv.Quit()
		case key.S:
			cv.DefaultCursor()
		case key.H:
			cv.HideCursor()
		case key.F:
			cv.Fullscreen()
		case key.W:
			cv.Windowed()
		default:
			return 0
		}
	case cv.ON_UP:
		log.Println("UP", cv.KeyName(e.Which()))
	case cv.ON_UNICODE:
		log.Printf("UNICODE %c\n", e.Unicode())
	case cv.ON_MOTION:
		log.Println("MOTION", e.X(), e.Y())
	case cv.ON_CLOSE:
		log.Println("CLOSE")
		cv.Quit()
	case cv.ON_RESIZE:
		log.Println("RESIZE", e.Width(), e.Height())
	case cv.ON_UPDATE:
		n := gl.Float(time.Now().Nanosecond()) / 1000000000
		gl.ClearColor(0, n, 0, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		if cv.Pressed(key.A) {
			log.Println("A pressed")
		}
		if cv.Released(key.A) {
			log.Println("A released")
		}
	}
	return 1
}

func main() {
	cv.Run(handler)
}
