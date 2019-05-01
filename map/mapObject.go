package mapobject

import (
	"fmt"
	"math/rand"
	"time"
)

//Map struct with two dimension array to display 2D
type Map struct {
	topology [][]int
}

//New map struct
func New(width, height int) *Map {
	c := make([][]int, height)
	for i := range c {
		c[i] = make([]int, width)
	}
	return &Map{
		topology: c,
	}
}

//FillMap method to generate a map
func (m *Map) FillMap() {
	rand.Seed(time.Now().UTC().UnixNano())
	for line := range m.topology {
		for cell := range m.topology[line] {
			m.topology[line][cell] = rand.Intn(3)
		}
	}
}

//Draw map in the output
func (m *Map) Draw() {
	for line := range m.topology {
		for cell := range m.topology[line] {
			if cell < (len(m.topology[line]) - 1) {
				switch m.topology[line][cell] {
				case 0:
					fmt.Print("*")
				case 1:
					fmt.Print("_")
				case 2:
					fmt.Print("^")
				}
			} else {
				switch m.topology[line][cell] {
				case 0:
					fmt.Println("*")
				case 1:
					fmt.Println("_")
				case 2:
					fmt.Println("^")
				}
			}
		}
	}
}
