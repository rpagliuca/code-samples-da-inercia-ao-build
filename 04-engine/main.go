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
		{Point{0.5, 0.5}, 0.01, Point{0.0, 0.0}},
	}

	// Fixed objects (lines)
	fixedObjects := []Line{
		// Floor
		{Point{0.0, 0.0}, Point{1.0, 0.0}},
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

		movement := Line{movingObjects[i].Position, previousMovingObjects[i].Position}

		for _, line := range fixedObjects {
			intersected := doLinesIntersect(movement, line)
			fmt.Println(intersected)
			if intersected {
				movingObjects[i].Position.Y = previousMovingObjects[i].Position.Y
				movingObjects[i].Velocity.Y = -movingObjects[i].Velocity.Y
			}
		}

		fmt.Println(movement)

		if movingObjects[i].Position.Y < 0 {
			panic("já foi")
		}

	}
	return movingObjects
}

func rotate45Deg(line Line) Line {
	theta := math.Pi / 4
	x1 := line.Position1.X
	y1 := line.Position1.Y
	x2 := line.Position2.X
	y2 := line.Position2.Y
	return Line{
		Point{x1*math.Cos(theta) - y1*math.Sin(theta), x1*math.Sin(theta) + y1*math.Cos(theta)},
		Point{x2*math.Cos(theta) - y2*math.Sin(theta), x2*math.Sin(theta) + y2*math.Cos(theta)},
	}
}

func doLinesIntersect(line1, line2 Line) bool {
	for line1.Position2.X-line1.Position1.X == 0 || line2.Position2.X-line2.Position1.X == 0 {
		fmt.Println("rotating!")
		line1 = rotate45Deg(line1)
		fmt.Println("l1", line1)
		line2 = rotate45Deg(line2)
		fmt.Println("l2", line2)
	}

	a1 := math.Abs((line1.Position2.Y - line1.Position1.Y) / (line1.Position2.X - line1.Position1.X))
	a2 := math.Abs((line2.Position2.Y - line2.Position1.Y) / (line2.Position2.X - line2.Position1.X))
	fmt.Println("a1, a2", a1, a2)

	if a1 == a2 {
		// They are parallel, so do not intersect
		return false
	}

	// y = ax + b
	// b = y - ax
	b1 := line1.Position1.Y - a1*line1.Position1.X
	b2 := line2.Position1.Y - a2*line2.Position1.X
	fmt.Println("b1, b2", b1, b2)

	// Where do they intersect?
	// y1 == y2, x1 == x2
	// y = a1x + b1
	// y = a2x + b2
	// => a1x + b1 = a2x + b2
	// => x (a1 - a2) = (b2 - b1)
	// => x = (b2 - b1) / (a1 - a2)

	x := (b2 - b1) / (a1 - a2)

	if ((x >= line1.Position1.X && x <= line1.Position2.X) ||
		(x <= line1.Position1.X && x >= line1.Position2.X)) &&
		((x >= line2.Position1.X && x <= line2.Position2.X) ||
			(x <= line2.Position1.X && x >= line2.Position2.X)) {
		return true
	}

	return false
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
