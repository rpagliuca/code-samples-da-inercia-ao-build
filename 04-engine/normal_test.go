package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestNormal(t *testing.T) {
	cases := []struct {
		x0  float64
		y0  float64
		x1  float64
		y1  float64
		nx0 float64
		ny0 float64
		nx1 float64
		ny1 float64
	}{
		{
			0, 0, 2, 0,
			0, 0, 0, 1,
		},
	}

	for _, c := range cases {
		nx0, ny0, nx1, ny1 := normal(c.x0, c.y0, c.x1, c.y1)
		assert.Equal(t, nx0, c.nx0)
		assert.Equal(t, ny0, c.ny0)
		assert.Equal(t, nx1, c.nx1)
		assert.Equal(t, ny1, c.ny1)
	}
}
