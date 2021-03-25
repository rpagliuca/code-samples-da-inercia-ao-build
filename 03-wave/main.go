package main

import (
	"fmt"
	"time"
)

const LEN = 100
const D = 0.06

type Grid [3][LEN]float64

// Wave equation
//
// d2u/dt2 = c2 d2u/dx2
//
// du/dt (forward): (u(t+dt) - u(t)) / dt
//
// d2u/dt2 (x fixo): ((u(t+dt) - u(t)) / dt - (u(t) - u(t-dt)) / dt) / dt
//				= (u(t+dt) - 2u(t) + u(t-dt)) / dt^2
//
// d2u/dx2 (t fixo): ((u(x+dx) - u(x)) / dx - (u(x) - u(x-dx)) / dx) / dx
//				= (u(x+dx) - 2u(x) + u(x-dx)) / dx^2
//
// Equação combinada:
// (u(t+dt) - 2u(t) + u(t-dt)) / dt^2 = C * (u(x+dx) - 2u(x) + u(x-dx)) / dx^2
//				=> u(t+dt) = D * (u(x+dx) - 2u(x) + u(x-dx)) + 2u(t) - u(t-dt)

func NextStep(grid Grid) Grid {

	for i := 1; i < LEN-1; i++ {
		// Combined equation
		grid[2][i] = D*(grid[1][i+1]-2.0*grid[1][i]+grid[1][i-1]) + 2.0*grid[1][i] - grid[0][i]
	}

	// Sane boundaries
	grid[2][0] = grid[2][1]
	grid[2][LEN-1] = grid[2][LEN-2]

	// Cycle temporal values
	grid[0] = grid[1]
	grid[1] = grid[2]

	return grid
}

func main() {

	start := time.Now()

	// 25 FPS (frames per second) => 40 milliseconds per frame
	ticker := time.NewTicker(40 * time.Millisecond)

	g := Grid{}
	g[0][0] = 100
	g[0][1] = 95
	g[0][2] = 90
	g[0][3] = 85
	g[0][4] = 80
	g[1] = g[0]

	for now := range ticker.C {
		fmt.Printf("t=%.2f u=%.2f\n", now.Sub(start).Seconds(), g[0][0])
		g = NextStep(g)
	}
}
