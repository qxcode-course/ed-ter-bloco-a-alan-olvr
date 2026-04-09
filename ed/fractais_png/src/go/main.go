package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func circulo(pen *Pen, radius float64, steps int) {
	//ponto de parada
	if steps == 0 {
		return
	}

	if steps == 4 {
    	pen.SetRGB(0, 0, 0) 
	} else {
    	intensidade := float64(5-steps) * 60 
    	pen.SetRGB(intensidade, intensidade/2, intensidade)
	}
	
	pen.FillCircle(radius)
	pen.SetRGB(255, 255, 255)
	pen.SetLineWidth(1.5)
	pen.DrawCircle(radius)

	for i := 0; i < 6; i++ {
		pen.Up()
		pen.Walk(radius * 1.3)
		pen.Down()
		circulo(pen, radius/3.0, steps-1)
		pen.Up()
		pen.Walk(-radius * 1.1)
		pen.Left(60)
	}

}

func main() {
	pen := NewPen(800, 800)   

	pen.SetPosition(0, 0) // Centraliza o ponto de partida
	pen.SetRGB(0, 0, 0)
	pen.FillSquare(800, 800)

	pen.SetPosition(400, 400)
	pen.SetRGB(255, 0, 0)
	circulo(pen, 200, 4)

	pen.SavePNG("circulo.png")
	fmt.Println("PNG file created successfully.")
}
