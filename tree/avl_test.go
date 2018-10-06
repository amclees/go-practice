package tree

import (
	"math"
	"testing"
)

func TestAVLBalance(t *testing.T) {
	avl := &AVL{h: -1}
	for i := 0; i < 262144; i++ {
		avl.Insert(i, i)
	}

	c := 262144
	h := avl.Height()
	lowerBound := int(math.Floor(math.Log2(float64(c + 1))))
	upperBound := int(math.Ceil(1.44 * math.Log2(float64(c + 2)) - 0.328))
	if lowerBound >= h || upperBound <= h {
		t.Errorf("Expected %d < h < %d; c was %d", lowerBound, upperBound, h)
	}
}
