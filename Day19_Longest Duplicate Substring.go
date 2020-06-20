package leetcode

/*
Given a string S, consider all duplicated substrings: (contiguous) substrings of S that occur 2 or more times.  (The occurrences may overlap.)

Return any duplicated substring that has the longest possible length.  (If S does not have a duplicated substring, the answer is "".)

Example 1:

Input: "banana"
Output: "ana"
Example 2:

Input: "abcd"
Output: ""


Note:

2 <= S.length <= 10^5
S consists of lowercase English letters.
*/
func longestDupSubstring(S string) string {
	var result, str string
	left, right := 0, len(S)-1
	for left <= right {
		mid := left + (right-left)/2
		str = ""
		m := make(map[string]bool, len(S)-mid)
		for i := 0; i+mid <= len(S); i++ {
			if _, ok := m[S[i:i+mid]]; !ok {
				m[S[i:i+mid]] = true
			} else {
				str = S[i : i+mid]
				break
			}
		}
		if str == "" {
			right = mid - 1
		} else {
			left = mid + 1
			result = str
		}
	}
	return result
}
