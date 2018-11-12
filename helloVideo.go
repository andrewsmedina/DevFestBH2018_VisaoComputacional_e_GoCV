// This example is all based on the tutorial existing in: https://gocv.io/writing-code/,
// with some changes to compile in Windows.

package main

import (
	"gocv.io/x/gocv"
)

func main() {
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Hello, DevFest BH :)!")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		gocv.WaitKey(1)
	}
}

// How to run (Windows):
//
// 		go run helloVideo.go
//
