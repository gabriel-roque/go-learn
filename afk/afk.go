package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	screenWidth, screenHeight := robotgo.GetScreenSize()
	fmt.Println(screenWidth, screenHeight)

	centerX, centerY := screenWidth/2, screenHeight/2

	offset := 100

	for {
		robotgo.MoveSmooth(centerX, centerY, 0.2, 0.2)
		time.Sleep(5 * time.Second)

		robotgo.MoveSmooth(centerX+offset, centerY, 0.2, 0.2)
		time.Sleep(5 * time.Second)

		robotgo.MoveSmooth(centerX, centerY, 0.2, 0.2)
		time.Sleep(5 * time.Second)
	}
}
