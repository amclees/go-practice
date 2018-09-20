package pow

import "math"

func Pow(x float64, n int) float64 {
	for n > 1 {
		k := HighestFactor(n)
		x = SlowPow(x, n/k)
		n = k
	}
	return x
}

func HighestFactor(n int) int {
	sqrt := int(math.Floor(math.Sqrt(float64(n))))
	for i := 2; i < sqrt; i++ {
		if n%i == 0 {
			return n / i
		}
	}
	return 1
}

func SlowPow(x float64, n int) float64 {
	pow := 1.0
	for i := 0; i < n; i++ {
		pow *= x
	}
	return pow
}
