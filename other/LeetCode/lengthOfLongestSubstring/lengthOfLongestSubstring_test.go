package lengthOfLongestSubstring

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	var cases = []struct {
		seq  string
		sseq string
		val  int
	}{
		{
			// duplicate values
			"abcabcbb",
			"abc",
			3,
		},
		{
			// all the same
			"bbbbbbb",
			"b",
			1,
		},
		{
			// mid value
			"pwwkew",
			"wke",
			3,
		},
		{
			// no string
			"",
			"",
			0,
		},
		{
			// last value
			"aaaaaaaaaaaab",
			"ab",
			2,
		},
		{
			// string, len 2
			"aa",
			"a",
			1,
		},
		{
			// string, len 1
			"b",
			"b",
			1,
		},
		{
			// string, len 1
			"dvdf",
			"vdf",
			3,
		},
	}
	for _, c := range cases {
		sseq, val := lengthOfLongestSubstring(c.seq)
		if val != c.val {
			t.Errorf("Values do not match - got %v want: %v\n", val, c.val)
		}
		if sseq != c.sseq {
			t.Errorf("Strings do not match - got %v want: %v\n", sseq, c.sseq)
		}
	}
}
