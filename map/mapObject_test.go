package mapobject

import (
	"testing"
)

func TestAddElementToArray(t *testing.T) {
	a := New(2, 2)
	a.FillMap()
	lenBeforeChange := len(a.topology[0])
	a.topology[0] = addElements(a.topology[0])
	if len(a.topology[0]) == lenBeforeChange {
		t.Fail()
	}
}

func TestAddElementToMap(t *testing.T) {
	a := New(2, 2)
	a.FillMap()
	lenBeforeChange := len(a.topology[0])
	a.modifyMap()
	if len(a.topology[1]) == lenBeforeChange {
		t.Fail()
	}
}
func TestAddNewRows(t *testing.T) {
	a := New(2, 2)
	a.FillMap()
	lenBeforeChange := len(a.topology[0])
	a.modifyMap()
	a.addNewRows()
	if len(a.topology[1]) == lenBeforeChange {
		t.Fail()
	}
}
