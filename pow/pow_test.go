package pow

import "testing"

func TestPow(t *testing.T) {
	xs := [...]float64{5, 5, 3, 10, 10}
	ns := [...]int{4, 25, 3, 11, 60}
	as := [...]float64{625, 298023223876953125, 27, 100000000000, 1e60}
	for i, val := range as {
		f := Pow(xs[i], ns[i])
		if f-val > 0.01 {
			t.Errorf("Got %e**%d == %e, expected %e (diff %e)", xs[i], ns[i], f, val, f-val)
		}
	}
}
