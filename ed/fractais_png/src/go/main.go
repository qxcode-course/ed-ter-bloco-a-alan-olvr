package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func circulos(pen *Pen, raio float64, nivel int) {
	//condição de parada
	if nivel == 0 {
		return
	}

	//desenhando o circulo central
	pen.DrawCircle(raio)

	for i := 0; i < 6; i++ {
		pen.Up()
		pen.Walk(raio)
		pen.Down()
		circulos(pen, raio/2.5, nivel-1)
		pen.Up()
		pen.Walk(-raio)
		pen.Down()
		pen.Left(60)
	}

}


func main() {
	pen := NewPen(800, 800)   // cria um canvas de 500 de largura por 500 de altura
	pen.SetRGB(255, 255, 255)     // muda a cor do pincel para vermelho
	pen.SetPosition(400, 400) // move o pincel para x 250, y 500
	circulos(pen, 200, 4)
	pen.SetHeading(90) // coloca o pincel apontando para cima
	

	pen.SavePNG("circulos.png")
	fmt.Println("PNG file created successfully.")
}
