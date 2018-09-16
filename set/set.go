package set

type Hashable interface {
	Hash() int64
}

type Set struct {
	s    [][]Hashable
	size int
}

func (set *Set) Init(cap int) {
	set.s = make([][]Hashable, cap)
	for i := range set.s {
		set.s[i] = make([]Hashable, 1)
	}
}

func (set *Set) Add(d Hashable) bool {
	if set.Contains(d) {
		return false
	}

	si := compress(len(set.s), d.Hash())
	set.s[si] = append(set.s[si], d)
	set.size += 1
	return true
}

func (set *Set) Remove(d Hashable) bool {
	dh := d.Hash()
	ss := set.s[compress(len(set.s), dh)]

	for i, sd := range ss {
		if sd != nil && sd.Hash() == dh {
			ss[i] = nil
			set.size -= 1
			return true
		}
	}

	return false
}

func (set *Set) Contains(d Hashable) bool {
	dh := d.Hash()
	ss := set.s[compress(len(set.s), dh)]

	for _, sd := range ss {
		if sd != nil && sd.Hash() == dh {
			return true
		}
	}

	return false
}

func (set *Set) All() []Hashable {
	all := make([]Hashable, 0, len(set.s))
	for _, ss := range set.s {
		for _, d := range ss {
			if d != nil {
				all = append(all, d)
			}
		}
	}
	return all
}

func (set *Set) Size() int {
	return set.size
}

func compress(span int, hash int64) int {
	return int(hash % int64(span))
}
