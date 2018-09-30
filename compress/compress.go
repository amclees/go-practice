package compress

import (
	"math"
	"strconv"
	"strings"
)

const (
	countDigits = iota
	parseDigits
	gatherContent
)

func decompress(s string) string {
	finalContent := strings.Builder{}
	contents := []strings.Builder{finalContent}
	reps := make([]int, 0)
	content := &finalContent
	st := countDigits
	depth := 0
	iseq := 0
	l := 0
	rep := 0
	for i := 0; i < len(s); i++ {
		switch st {
		case countDigits:
			if s[i] <= '9' && s[i] >= '0' {
				l += 1
			} else {
				i = iseq - 1
				st = parseDigits
				rep = 0
			}
		case parseDigits:
			if s[i] <= '9' && s[i] >= '0' {
				digit, _ := strconv.Atoi(string(s[i]))
				rep += digit * int(math.Pow(10.0, float64(iseq+l-i-1)))
			} else {
				l = 0
				depth += 1
				reps = append(reps, rep)
				contents = append(contents, strings.Builder{})
				content = &contents[len(contents) - 1]
				st = gatherContent
			}
		case gatherContent:
			if s[i] == ']' {
				depth -= 1
				if depth == 0 {
					finalContent.WriteString(strings.Repeat(content.String(), reps[depth]))
					c := strings.Builder{}
					content = &c
				} else {
					oldContent := contents[len(contents) - 1].String()
					contents = contents[:len(contents) - 1]
					content = &contents[len(contents) - 1]
					content.WriteString(strings.Repeat(oldContent, reps[depth]))
				}
				reps = reps[:len(reps) - 1]
			} else if s[i] <= '9' && s[i] >= '0' {
				st = countDigits
				iseq = i
				i -= 1
			} else {
				content.WriteByte(s[i])
			}
		}
	}

	ls := content.String()
	if len(ls) > 0 {
		finalContent.WriteString(content.String())
	}

	return finalContent.String()
}
