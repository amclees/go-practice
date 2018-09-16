package list

import "testing"

func TestList(t *testing.T) {
	list := List{}

	list.Append(1)
	list.Prepend(2)
	list.Prepend(3)

	val, _ := list.Get(1)
	if val != 2 {
		t.Errorf("Expected list.Get(1) to be 2, got %d", val)
	}

	if list.First() != 3 {
		t.Errorf("Expected list.First() to be 3, got %d", list.First())
	}

	if list.Last() != 1 {
		t.Errorf("Expected list.Last() to be 1, got %d", list.Last())
	}

	list.Prepend(4)
	removed := list.RemoveReverse(1)
	removed2 := list.Remove(0)

	if !removed || !removed2 {
		t.Errorf("Expected remove calls to succeed")
	}

	val2, found := list.Get(1)
	if !found {
		t.Errorf("Expected list.Get(1) to succeed")
	}
	if val2 != 1 {
		t.Errorf("Expected list.Get(1) to be 1 after deletion, got %d", val2)
	}

	list.Append(3)
	list.Append(4)

	i1 := list.IndexOf(3)
	if i1 != 2 {
		t.Errorf("Expected list.IndexOfReverse(3) to be 2, got %d", i1)
	}

	i2 := list.IndexOfReverse(4)
	if i2 != 3 {
		t.Errorf("Expected list.IndexOfReverse(4) to be 3, got %d", i2)
	}

	length := list.Length()
	if length != 4 {
		t.Errorf("Expected list.Length() to be 4, got %d", length)
	}
}
