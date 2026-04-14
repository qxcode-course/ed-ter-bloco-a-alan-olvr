package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func carpet(pen *Pen, x, y, size float64, depth int) {

	if depth == 0 {
		drawLeaf(pen, x, y, size)
		return
	}

	cell := size / 4.0

	pen.SetRGB(255, 255, 255)
	pen.SetLineWidth(1.0)

	drawRect(pen, x, y, size, size)

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			if row >= 1 && row <= 2 && col >= 1 && col <= 2 {
				continue
			}

			cx := x + float64(col)*cell
			cy := y + float64(row)*cell
			carpet(pen, cx, cy, cell, depth-1)
		}
	}
}

func drawLeaf(pen *Pen, x, y, size float64) {
	pen.SetRGB(255, 255, 255)
	pen.SetLineWidth(1.0)

	half := size / 2.0
	quarter := size / 4.0

	for row := 0; row < 2; row++ {
		for col := 0; col < 2; col++ {
			mx := x + float64(col)*half + quarter*0.15
			my := y + float64(row)*half + quarter*0.15
			s := half * 0.7
			drawRect(pen, mx, my, s, s)
			inner := s * 0.45
			ox := mx + quarter*0.15
			oy := my + quarter*0.15
			drawRect(pen, ox, oy, inner, inner)
		}
	}
}

func drawRect(pen *Pen, x, y, w, h float64) {
	pen.Up()
	pen.SetPosition(x, y)
	pen.Down()
	pen.Goto(x+w, y)
	pen.Goto(x+w, y+h)
	pen.Goto(x, y+h)
	pen.Goto(x, y)
	pen.Up()
}
func main() {
	pen := NewPen(1000.0, 1000.0)

	pen.SetPosition(0, 0)
	pen.SetRGB(0, 0, 0)
	pen.FillSquare(1000.0, 1000.0)

	const margin = 20.0
	carpet(pen, margin, margin, 1000.0-2*margin, 4)

	pen.SavePNG("carpet.png")
	fmt.Println("PNG file created successfully.")
}
