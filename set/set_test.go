package set

import "testing"

type hint int

func (h hint) Hash() int64 {
	return int64(int(h))
}

func TestSet(t *testing.T) {
	set := Set{}
	set.Init(10)

	set.Add(hint(4))
	set.Add(hint(7))
	set.Add(hint(14))

	set.Add(hint(9))
	set.Remove(hint(9))

	containsCases := []Hashable{hint(4), hint(7), hint(14), hint(17), hint(9)}
	containsExpected := []bool{true, true, true, false, false}

	for i := range containsCases {
		contains := set.Contains(containsCases[i])
		if contains != containsExpected[i] {
			t.Errorf("Expected contains %d to return %v, was %v", containsCases[i].Hash(), containsExpected[i], contains)
		}
	}

	size := set.Size()
	if size != 3 {
		t.Errorf("Expected size of set to be 3, was %d", size)
	}

	all := set.All()
	for _, d := range all {
		if d != hint(4) && d != hint(7) && d != hint(14) {
			t.Errorf("Got unexpected value %d in set.All()", d.Hash())
		}
	}
}
