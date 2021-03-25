package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersect(t *testing.T) {
	cases := []struct {
		ax0      float64
		ay0      float64
		ax1      float64
		ay1      float64
		bx0      float64
		by0      float64
		bx1      float64
		by1      float64
		expected bool
	}{
		{
			0, 0, 1, 0,
			0, 1, 1, 1,
			false,
		},
		{
			0.5, 0.5, 0.5, -0.5,
			0, 0, 1, 0,
			true,
		},
	}

	for _, c := range cases {
		got := intersect(c.ax0, c.ay0, c.ax1, c.ay1, c.bx0, c.by0, c.bx1, c.by1)
		assert.Equal(t, c.expected, got)
	}
}
