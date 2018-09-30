package compress

import "testing"

func TestCompress(t *testing.T) {
	table := [][2]string{{"5[a]", "aaaaa"},
		{"3[abb]", "abbabbabb"},
		{"11[ab]", "ababababababababababab"},
		{"3[abc]4[ab]c", "abcabcabcababababc"},
		{"2[3[a]b]", "aaabaaab"},
		{"2[3[5[abb]c]d]e", "abbabbabbabbabbcabbabbabbabbabbcabbabbabbabbabbcdabbabbabbabbabbcabbabbabbabbabbcabbabbabbabbabbcde"},
		{"2[3[a]b]2[3[a]b]", "aaabaaabaaabaaab"}}

	for _, row := range table {
		d := decompress(row[0])
		if d != row[1] {
			t.Errorf("Expected decompress(%v) = %v, was %v", row[0], row[1], d)
		}
	}
}
