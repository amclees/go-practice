package stack

import "testing"

func TestStack(t *testing.T) {
	st := Stack{}
	st.Init(1000)

	st.PushManual(5)
	st.Push(3)
	st.Push(4)

	a := st.Pop()
	b := st.Pop()

	if a != 4 {
		t.Errorf("Expected first pop to be 4, was %d", a)
	}
	if b != 3 {
		t.Errorf("Expected second pop to be 3, was %d", b)
	}

	for i := 0; i < 10000; i++ {
		st.Push(int64(i))
	}

	c := st.Pop()
	if c != 9999 {
		t.Errorf("Expected third pop after resize to be 9999, was %d", c)
	}
}

func BenchmarkStackPregrow(b *testing.B) {
	st := Stack{}
	st.Init(1)

	st.Grow(b.N)

	for i := 0; i < b.N; i++ {
		st.PushManual(1)
	}
}

func BenchmarkStackAutogrowAppend(b *testing.B) {
	st := Stack{}
	st.Init(1)

	for i := 0; i < b.N; i++ {
		st.Push(1)
	}
}

func BenchmarkStackAutogrowManual(b *testing.B) {
	st := Stack{}
	st.Init(1)

	for i := 0; i < b.N; i++ {
		st.PushManual(1)
	}
}
