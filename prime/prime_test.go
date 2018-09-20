package prime

import "testing"

func TestPrimesTo(t *testing.T) {
	a := [...]int{2, 3, 5, 7, 11, 13, 17, 19}
	p := PrimesTo(20)
	for i, val := range a {
		f := p[i]
		if f != val {
			t.Errorf("Got PrimesTo(20)[%d] %d, expected %d", i, f, val)
		}
	}
}
