package main

import (
	"code.google.com/p/glcv-go/cv"
	gl "github.com/chsc/gogl/gl21"
	"log"
)

func handler(e *cv.Event) uintptr {
	switch e.Type() {
	case cv.HINT_NAME:
		return cv.String("cvtest")
	case cv.HINT_BORDERS:
		return 1
	case cv.HINT_XPOS:
		return 20
	case cv.HINT_YPOS:
		return 20
	case cv.HINT_WIDTH:
		return 640
	case cv.HINT_HEIGHT:
		return 640
	case cv.ON_INIT:
		log.Println("INIT")
		return 1
	case cv.ON_TERM:
		log.Println("TERM")
		return 1
	case cv.ON_GLINIT:
		log.Println("GLINIT")
		gl.Init()
		return 1
	case cv.ON_DOWN:
		log.Println("DOWN", cv.KeyName(e.Which()))
		return 1
	case cv.ON_UP:
		log.Println("UP", cv.KeyName(e.Which()))
		return 1
	case cv.ON_UNICODE:
		log.Printf("UNICODE %c\n", e.Unicode())
		return 1
	case cv.ON_MOTION:
		log.Println("MOTION", e.X(), e.Y())
		return 1
	case cv.ON_CLOSE:
		log.Println("CLOSE")
		cv.Quit()
		return 1
	case cv.ON_RESIZE:
		log.Println("RESIZE", e.Width(), e.Height())
		return 1
	case cv.ON_UPDATE:
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		return 1
	}
	return 0
}

func main() {
	log.Println("running")
	cv.Run(handler)
}
