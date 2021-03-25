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

const experimentalConstant = 0.0015

func getAcceleration(height, velocity float64) float64 {
	// Newton's law of universal gravitation
	// plus Newton's second law
	g := -G * M / math.Pow(r+height, 2)
	// Air drag: https://www.scielo.br/pdf/rbef/v37n2/0102-4744-rbef-37-02-2306.pdf
	drag := experimentalConstant * getAirDensity(height) * math.Pow(velocity, 2)
	a := g + drag
	return a
}

func getAirDensity(height float64) float64 {
	return 1.225 * math.Pow(math.E, -height/8882)
}

// Initial height of the object
const s0 = 38970.0

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
			return v0 + getAcceleration(sPrevious, vPrevious)*t
		}
		vk1 := getAcceleration(sPrevious, vPrevious)
		vk2 := getAcceleration(sPrevious, vPrevious+deltaT.Seconds()*vk1/2.0)
		vk3 := getAcceleration(sPrevious, vPrevious+deltaT.Seconds()*vk2/2.0)
		vk4 := getAcceleration(sPrevious, vPrevious+deltaT.Seconds()*vk3)
		v := vPrevious + (1.0/6.0)*deltaT.Seconds()*(vk1+2.0*vk2+2.0*vk3+vk4)
		sk1 := getV(vPrevious, 0)
		sk2 := getV(vPrevious, deltaT.Seconds()/2.0)
		sk3 := getV(vPrevious, deltaT.Seconds()/2.0)
		sk4 := getV(vPrevious, deltaT.Seconds())
		s := sPrevious + (1.0/6.0)*deltaT.Seconds()*(sk1+2.0*sk2+2.0*sk3+sk4)
		// Print current position
		t := now.Sub(start).Seconds()
		fmt.Printf("t=%.2f airDensity=%.2f v=%.2f s=%.2f\n", t, getAirDensity(sPrevious), v, s)
		if s <= 0 {
			// End of simulation
			fmt.Printf("Hit the ground!")
			break
		}
		vPrevious = v
		sPrevious = s
	}
}
