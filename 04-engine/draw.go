package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

func draw(frameId int, circles []Circle, lines []Line) {
	scale := 1000.0
	dc := gg.NewContext(int(scale), int(scale))

	// Grey background
	dc.DrawRectangle(0, 0, 1000, 1000)
	dc.SetRGB(0.1, 0.1, 0.1)
	dc.Fill()

	// Draw pink lines
	for _, l := range lines {
		dc.DrawLine(
			l.Position1.X*scale,
			(1.0-l.Position1.Y)*scale,
			l.Position2.X*scale,
			(1.0-l.Position2.Y)*scale,
		)
		dc.SetRGB(1.0, 20./255., 147./255.)
		dc.SetLineWidth(20)
		dc.Stroke()
	}

	// Draw blue circles
	for _, c := range circles {
		dc.DrawCircle(c.Position.X*scale, (1.0-c.Position.Y)*scale, c.Radius*scale)
		dc.SetRGB(0, 1, 0)
		dc.Fill()
	}

	dc.SavePNG(fmt.Sprintf("out/%05d.png", frameId))
}
