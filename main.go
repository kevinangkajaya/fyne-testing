package main

import "fyne-testing/greeting"

func main() {
	// only one can run between these apps
	// goroutine.TimeTicker()
	greeting.Greeting()
}
