package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	texto   *List[*List[rune]]
	it_line *Node[*List[rune]]
	it_char *Node[rune]
	screen  tcell.Screen
	style   tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.it_char = e.it_line.Value.Insert(e.it_char, r)
	e.it_char = e.it_char.Next()
}

func (e *Editor) KeyLeft() {
	if e.it_char.Prev() != e.it_line.Value.End() { // Se o cursor não está no início da linha
		e.it_char = e.it_char.Prev() // Move o cursor para a esquerda
		return
	}
	// Estamos no início da linha
	if e.it_line != e.texto.Front() { // Se não está na primeira linha
		e.it_line = e.it_line.Prev()      // Move para a linha anterior
		e.it_char = e.it_line.Value.End() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyEnter() {
	new := NewList[rune]()

	e.texto.Insert(e.it_line.Next(), new)

	insertedLine := e.it_line.Next()

	for n := e.it_char; n != e.it_line.Value.End(); {
		next := n.Next()
		insertedLine.Value.Insert(insertedLine.Value.End(), n.Value)
		e.it_line.Value.Erase(n)
		n = next
	}

	e.it_line = insertedLine
	e.it_char = e.it_line.Value.Front()
}

func (e *Editor) KeyRight() {
	if e.it_char.Next() != e.it_line.Value.End() {
		e.it_char = e.it_char.Next()
		return
	}

	if e.it_line != e.texto.End() { // Se não está na primeira linha
		e.it_line = e.it_line.Next()      // Move para a linha anterior
		e.it_char = e.it_line.Value.Front() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyUp() {
	prev := e.it_line.Prev()
	if prev == nil {
		return
	}
	e.it_line = prev
	e.it_char = e.it_line.Value.Front()
}

func (e *Editor) KeyDown() {
	next := e.it_line.Next()
	if next == nil {
		return 
	}
	e.it_line = next
	e.it_char = e.it_line.Value.Front()
}

func (e *Editor) KeyBackspace() {
	if e.it_char.Prev() != e.it_line.Value.End() {
		e.it_char = e.it_line.Value.Erase(e.it_char.Prev())
		return
	} 

	if e.it_line.Prev() != e.texto.End() {
		prevLine := e.it_line.Prev()
		oldLine := e.it_line

		e.it_char = prevLine.Value.End()

		for n := oldLine.Value.Front(); n != oldLine.Value.End(); {
			next := n.Next()
			prevLine.Value.Insert(prevLine.Value.End(), n.Value)
			n = next
		}

		e.texto.Erase(oldLine)
		e.it_line = prevLine
	}
}

func (e *Editor) KeyDelete() {
	if e.it_char != e.it_line.Value.End() {
		e.it_char = e.it_line.Value.Erase(e.it_char)
		return
	}

	if e.it_line.Next() != e.texto.End() {
		nextLine := e.it_line.Next()

		for n := nextLine.Value.Front(); n != nextLine.Value.End(); {
			next := n.Next()
			e.it_line.Value.Insert(e.it_line.Value.End(), n.Value)
			n = next
		}

		e.texto.Erase(nextLine)
	}
}

func main() {
	// Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.texto = NewList[*List[rune]]()
	e.texto.PushBack(NewList[rune]())
	e.it_line = e.texto.Front()
	e.it_char = e.it_line.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.texto.Front(); line != e.texto.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '⤶'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.it_char {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}
