package utility

import "testing"

func TestNewNum(t *testing.T) {
	num := NewNum("1D5onNHcr04WgKi79XlFkRUZGymShTpjx2zJVsqLtOY3wvdIaMufQ6P8CEeAbB")
	var id = 0
	cid, err := num.ToInt(num.ToStr(id))
	if err != nil {
		panic(err)
	}
	if cid != id {
		t.Fatal("error")
	}

	id = 123
	cid, err = num.ToInt(num.ToStr(id))
	if err != nil {
		panic(err)
	}
	if cid != id {
		t.Fatal("error")
	}

	id = 7536493459
	cid, err = num.ToInt(num.ToStr(id))
	if err != nil {
		panic(err)
	}
	if cid != id {
		t.Fatal("error")
	}
}
