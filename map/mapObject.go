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

//Prepend and append 0 element to array given
func addElements(arr []int) []int {
	arr = append(arr, 0)
	arr = append([]int{0}, arr...)
	return arr
}

//Prepend and append 0 element to all arrays in Map
func (m *Map) modifyMap() {
	for i := range m.topology {
		m.topology[i] = addElements(m.topology[i])
	}
}

//Prepend and append a new arrays fill with zeros
func (m *Map) addNewRows() {
	m.topology = append(m.topology, make([]int, len(m.topology[0])))
	m.topology = append(make([][]int, 1), m.topology...)
	for i := 0; i < 1; i++ {
		m.topology[i] = make([]int, len(m.topology[i+1]))
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
	//Next helper methods are used to get a border in map
	m.modifyMap()
	m.addNewRows()
}

//Draw map in the output
func (m *Map) Draw() {
	for line := range m.topology {
		for cell := range m.topology[line] {
			if cell < (len(m.topology[line]) - 1) {
				switch m.topology[line][cell] {
				case 0:
					fmt.Print(" ")
				case 1:
					fmt.Print("_")
				case 2:
					fmt.Print("^")
				}
			} else {
				switch m.topology[line][cell] {
				case 0:
					fmt.Println(" ")
				case 1:
					fmt.Println("_")
				case 2:
					fmt.Println("^")
				}
			}
		}
	}
}
