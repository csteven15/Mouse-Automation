package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("Starting script")

	fmt.Println("Enter how many programs there are:")
	var numApps int
	fmt.Scanf("%d", &numApps)
	fmt.Println("Scanning for", numApps, "clicks")

	// apps : slice of slices
	apps := make([][]int, numApps)

	for i := range apps {
		apps[i] = getAppLocation()
	}

	for i := range apps {
		fmt.Println("moving to ", apps[i])
		robotgo.MoveMouse(apps[i][0], apps[i][1])
		robotgo.MouseClick("left", true)
		time.Sleep(4 * time.Second)
	}

}

func getAppLocation() []int {
	getPoints := robotgo.AddEvent("mleft")
	if getPoints {
		x, y := robotgo.GetMousePos()
		fmt.Println("you clicked on", x, y)
		return []int{x, y}
	}
	return []int{-1, -1}
}
