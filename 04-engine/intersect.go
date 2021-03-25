package main

func intersect(x0, y0, x1, y1, x2, y2, x3, y3 float64) bool {
	p0 := (y3-y2)*(x3-x0) - (x3-x2)*(y3-y0)
	p1 := (y3-y2)*(x3-x1) - (x3-x2)*(y3-y1)
	p2 := (y1-y0)*(x1-x2) - (x1-x0)*(y1-y2)
	p3 := (y1-y0)*(x1-x3) - (x1-x0)*(y1-y3)
	return p0*p1 <= 0 && p2*p3 <= 0
}
