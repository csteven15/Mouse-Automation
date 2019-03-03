package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("Starting script")

	nextTime := time.Now().Truncate(time.Minute)

	fmt.Println("Enter how many clicks there are:")
	var numClicks int
	fmt.Scanf("%d", &numClicks)
	fmt.Println("Scanning for", numClicks, "clicks")

	// clicks : slice of slices
	clicks := make([][]int, numClicks)

	for i := range clicks {
		clicks[i] = getAppLocation()
	}

	for {
		for i := range clicks {
			fmt.Println("Moving to", clicks[i])
			// robotgo.MoveMouseSmooth(clicks[i][0], clicks[i][1], 1.0, 1.0)
			robotgo.MoveMouse(clicks[i][0], clicks[i][1])
			fmt.Println("Clicking on", clicks[i])
			robotgo.MouseClick("left", true)
			time.Sleep(10 * time.Second)
		}
		nextTime = nextTime.Add(time.Second * 10)
	}

}

func getAppLocation() []int {
	getPoints := robotgo.AddEvent("mleft")
	if getPoints {
		x, y := robotgo.GetMousePos()
		fmt.Println("You clicked on", x, y)
		return []int{x, y}
	}
	return []int{-1, -1}
}
