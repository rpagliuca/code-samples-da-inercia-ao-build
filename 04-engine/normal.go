package main

import "math"

func normal(x0, y0, x1, y1 float64) (float64, float64, float64, float64) {
	//Rotation by 90º
	// | 0 -1 | (X) = (Xf)
	// | 1  0 | (Y) = (Yf)
	return normalize(-y0, x0, -y1, x1)
}

func normalize(x0, y0, x1, y1 float64) (float64, float64, float64, float64) {
	magnitude := math.Pow(math.Pow(x1-x0, 2)+math.Pow(y1-y0, 2), 0.5)
	return 0, 0, (x1 - x0) / magnitude, (y1 - y0) / magnitude
}
