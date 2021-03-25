package main

import (
	"fmt"
	"time"
)

// Gravity acceleration at sea level
const a = -9.8

// Initial height of the object
const s0 = 10.0

// Initial velocity of the object
const v0 = 0.0

func main() {
	start := time.Now()
	// 25 FPS (frames per second) => 40 milliseconds per frame
	ticker := time.NewTicker(40 * time.Millisecond)
	for now := range ticker.C {
		// Seconds since start of simulation
		t := now.Sub(start).Seconds()
		// Analytical solution for constant gravity
		s := s0 + v0*t + a*t*t/2.0
		// Print current position
		fmt.Printf("t=%.2f s=%.2f\n", t, s)
		if s <= 0 {
			// End of simulation
			fmt.Printf("Hit the ground!")
			break
		}
	}
}
