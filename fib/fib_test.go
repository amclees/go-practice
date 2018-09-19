package fib

import "testing"

func TestFib(t *testing.T) {
	a := [...]int{1, 1, 2, 3, 5, 8, 13, 21}
	for i, val := range a {
		f := Fib(i)
		if f != val {
			t.Errorf("Got Fib(%d) %d, expected %d", i, f, val)
		}
	}
}
