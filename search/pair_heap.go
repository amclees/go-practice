package search

type pair struct {
	n, w int
}

type dqueue struct {
	s []pair
}

func (q *dqueue) next() (pair, bool) {
	if len(q.s) == 0 {
		return pair{}, false
	}
	if len(q.s) == 1 {
		d := q.s[0]
		q.s = []pair{}
		return d, true
	}

	last := q.s[len(q.s)-1]
	q.s = q.s[0 : len(q.s)-1]
	d := q.s[0]
	q.s[0] = last

	q.settle(0)

	return d, true
}

const maxInt int = int(^uint(0) >> 1)
const minInt int = -maxInt - 1

func (q *dqueue) settle(i int) {
	l, r := i*2+1, i*2+2
	dl, dr := maxInt, maxInt
	if l < len(q.s) {
		dl = q.s[l].w
	}
	if r < len(q.s) {
		dr = q.s[r].w
	}

	d := q.s[i].w
	if d < dl && d < dr {
		return
	}
	if dl > dr {
		q.s[i].w, q.s[r].w = q.s[r].w, q.s[i].w
		q.settle(r)
	}
	if dr > dl {
		q.s[i].w, q.s[l].w = q.s[l].w, q.s[i].w
		q.settle(l)
	}
}

func (q *dqueue) set(node int, w int) bool {
	i := len(q.s)

	for fi := 0; fi < len(q.s); fi++ {
		if q.s[fi].n == node {
			if q.s[fi].w <= w {
				return false
			}
			q.s[fi].w = w
			i = fi
			break
		}
	}

	if i == len(q.s) {
		p := pair{node, w}
		q.s = append(q.s, p)
	}

	for ; i != 0; i /= 2 {
		if q.s[i].w < q.s[i/2].w {
			q.s[i], q.s[i/2] = q.s[i/2], q.s[i]
		}
	}

	return true
}
