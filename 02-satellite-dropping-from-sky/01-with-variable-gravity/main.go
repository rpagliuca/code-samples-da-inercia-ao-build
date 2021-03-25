package main

import (
	"fmt"
	"math"
	"time"
)

// Universal gravitational constant (S.I. units)
const G = 6.67408e-11

// Earth's mass (in kg)
const M = 5.9722e24

// Sea level radius (in meters)
const r = 6371e3

func getAcceleration(height float64) float64 {
	// Newton's law of universal gravitation
	// plus Newton's second law
	return -G * M / math.Pow(r+height, 2)
}

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
		// Runge-Kutta 4th order
		getV := func(v0, t float64) float64 {
			return v0 + getAcceleration(sPrevious)*t
		}
		v := getV(vPrevious, deltaT.Seconds())
		k1 := getV(vPrevious, 0)
		k2 := getV(vPrevious, deltaT.Seconds()/2.0)
		k3 := getV(vPrevious, deltaT.Seconds()/2.0)
		k4 := getV(vPrevious, deltaT.Seconds())
		s := sPrevious + (1.0/6.0)*deltaT.Seconds()*(k1+2.0*k2+2.0*k3+k4)
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
