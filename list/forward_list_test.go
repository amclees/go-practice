package list

import "testing"

func TestForwardList(t *testing.T) {
	list := ForwardList{}

	list.Append(5)
	// 5

	if list.Length() != 1 {
		t.Errorf("Expected list length %d, was %d", 1, list.Length())
	}

	for i := 0; i < 5; i++ {
		list.Append(int64(i))
	}
	// 5 0 1 2 3 4

	list.Prepend(1)
	// 1 5 0 1 2 3 4

	first, _ := list.Get(0)
	if first != 1 {
		t.Errorf("Expected first value of list after prepend to be 1, was %d", first)
	}

	removed := list.Remove(1)
	// 1 0 1 2 3 4
	if !removed {
		t.Errorf("Expected removal to succeed")
	}

	second, _ := list.Get(1)
	if second != 0 {
		t.Errorf("Expected second value of list after deletion to be 0, was %d", second)
	}

	index := list.IndexOf(2)
	if index != 3 {
		t.Errorf("Expected IndexOf(2) to be 3, was %d", index)
	}

	index = list.IndexOf(5)
	if index != -1 {
		t.Errorf("Expected IndexOf(5) to be -1, was %d", index)
	}
}
