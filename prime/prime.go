package prime

import "math"

func PrimesTo(n int) []int {
    all := make([]bool, n)

    all[0] = true
    all[1] = true

    c := len(all)
    sqrt := int(math.Ceil(math.Sqrt(float64(n))))
    for i := 2; i <= sqrt; i++ {
        j := 2
        for j, k := 2, i * j; k < n; j, k = j + 1, i * j {
	    all[k] = true
	    c -= 1
        }
    }

    s := make([]int, c + 10)
    k := 0
    for i, composite := range all {
        if !composite {
            s[k] = i
	    k += 1
        }
    }

    return s
}
