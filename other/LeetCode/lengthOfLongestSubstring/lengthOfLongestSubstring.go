/*
from leetcode:
	"Given a string, find the length of the longest substring without repeating
	characters."

Example solution:
	Given;
		- input: `"abcabcbb"`, solution: 3
			`"abc"`
		- input: `"bbbbb"`, solution: 1
			`"b"`
		- input: `"pwwkew"`, solution: 3
			`"wke"`
				- "pwke" is a subsequence and not a substring
	Solution;
		- 7 -> 0 -> 8

*/
package lengthOfLongestSubstring

// func allUnique(str string) bool {

// 	// allUnique iterates through a string and create a set of runes present in
// 	// the string and returns true if all values are unique, false if not.
// 	set := make(map[rune]bool)
// 	for _, r := range str {
// 		_, ok := set[r]
// 		switch ok {
// 		case true:
// 			return false
// 		default:
// 			set[r] = true
// 		}
// 	}
// 	return true
// }

func lengthOfLongestSubstring(str string) (string, int) {
	// TODO: revisit this problem and implement a sliding window
	ans := 0
	si := 0
	ei := 0

	// // Brute force solution. RT = O(n^3). We iterate through all possible
	// // substring combinations and store the indexes of the longest substring.
	// for i := range str {
	// 	for j := range str[i:] {
	// 		sstr := str[i : j+i+1]
	// 		if allUnique(sstr) {
	// 			if len(sstr) > llen {
	// 				ans = len(sstr)
	// 				si = i
	// 				ei = j + i + 1
	// 			}
	// 		}
	// 	}
	// }

	// Sliding window solution. RT = O(n).
	// solution explaination can be found here:
	// https://leetcode.com/articles/longest-substring-without-repeating-characters/
	set := make(map[rune]bool)
	n := len(str)
	i := 0
	j := 0
	for i < n && j < n {
		_, ok := set[rune(str[j])]
		if !ok {
			set[rune(str[j])] = true
			j++
			if j-i > ans {
				ans = j - i
				si = i
				ei = j
			}
		} else {
			delete(set, rune(str[i]))
			i++
		}
	}

	//fmt.Printf("String: %v, lss = %v, val = %v\n", str, str[si:ei], llen)
	return str[si:ei], ans
}
