package main

import (
	"fmt"
	"math"
	"time"
)

const deltaT = 40 * time.Millisecond

func main() {

	// 25 FPS
	ticker := time.NewTicker(deltaT)

	// Moving objects (circles)
	movingObjects := []Circle{
		{Point{0.5, 0.5}, 0.01, Point{0.2, 0.0}},
	}

	// Fixed objects (lines)
	fixedObjects := []Line{
		// Floor
		{Point{0.0, 0.0}, Point{1.0, 0.0}},
		// Right wall
		{Point{1.0, 0.0}, Point{1.0, 1.0}},
		// Left wall
		{Point{0.0, 0.0}, Point{0.0, 1.0}},
		// Middle diagonal wall
		{Point{0.2, 0.2}, Point{0.5, 0.8}},
	}

	count := 0
	for _ = range ticker.C {
		count++
		movingObjects = next(movingObjects, fixedObjects)
		// Draw every 10th frame
		if count%10 == 0 {
			draw(count, movingObjects, fixedObjects)
		}
	}
}

// Calculate the next step for the numerical simulation
func next(movingObjects []Circle, fixedObjects []Line) []Circle {
	previousMovingObjects := make([]Circle, len(movingObjects))
	copy(previousMovingObjects, movingObjects)
	for i := range movingObjects {
		movingObjects[i].Velocity.Y = movingObjects[i].Velocity.Y - 0.1*deltaT.Seconds()
		movingObjects[i].Position.Y = movingObjects[i].Position.Y + movingObjects[i].Velocity.Y*deltaT.Seconds()
		movingObjects[i].Position.X = movingObjects[i].Position.X + movingObjects[i].Velocity.X*deltaT.Seconds()

		movement := Line{movingObjects[i].Position, previousMovingObjects[i].Position}

		for _, line := range fixedObjects {
			intersected := intersect(
				movement.Position1.X,
				movement.Position1.Y,
				movement.Position2.X,
				movement.Position2.Y,
				line.Position1.X,
				line.Position1.Y,
				line.Position2.X,
				line.Position2.Y,
			)
			if intersected {
				_, _, x, y := normal(line.Position1.X, line.Position1.Y, line.Position2.X, line.Position2.Y)
				fmt.Println("x, y", x, y)
				velX := movingObjects[i].Velocity.X * x
				velY := movingObjects[i].Velocity.Y * y
				fmt.Println("1) velox, veloy", movingObjects[i].Velocity.X, movingObjects[i].Velocity.Y)
				fmt.Println("vx, vy", velX, velY)
				if movingObjects[i].Velocity.Y > 0 {
					movingObjects[i].Velocity.Y -= 2.0 * math.Abs(velY)
				} else {
					movingObjects[i].Velocity.Y += 2.0 * math.Abs(velY)
				}
				if movingObjects[i].Velocity.X > 0 {
					movingObjects[i].Velocity.X -= 2.0 * math.Abs(velX)
				} else {
					movingObjects[i].Velocity.X += 2.0 * math.Abs(velX)
				}
				fmt.Println("2) velox, veloy", movingObjects[i].Velocity.X, movingObjects[i].Velocity.Y)
				movingObjects[i].Position.Y = previousMovingObjects[i].Position.Y + movingObjects[i].Velocity.Y*deltaT.Seconds()
				movingObjects[i].Position.X = previousMovingObjects[i].Position.X + movingObjects[i].Velocity.X*deltaT.Seconds()
			}
		}

		fmt.Println(movement.Position2.X, movement.Position2.Y)
	}
	return movingObjects
}

type Radius float64

type Point struct {
	X float64
	Y float64
}

type Circle struct {
	Position Point
	Radius   float64
	Velocity Point
}

type Line struct {
	Position1 Point
	Position2 Point
}
