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

// Discretization of time
const deltaT = 40 * time.Millisecond

func main() {
	start := time.Now()
	vPrevious := v0
	sPrevious := s0
	// 25 FPS (frames per second) => 40 milliseconds per frame
	ticker := time.NewTicker(deltaT)
	for now := range ticker.C {
		// Forward Euler method
		v := vPrevious + a*deltaT.Seconds()
		// WOW! The world has stopped!
		s := sPrevious + vPrevious*deltaT.Seconds()
		// Print current position
		t := now.Sub(start).Seconds()
		fmt.Printf("t=%.2f s=%.2f\n", t, s)
		if s <= 0 {
			// End of simulation
			fmt.Printf("Hit the ground!")
			break
		}
		vPrevious = v
		sPrevious = s
	}
}
