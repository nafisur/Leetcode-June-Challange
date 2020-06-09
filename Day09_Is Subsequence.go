package leetcode

/*
Given a string s and a string t, check if s is subsequence of t.

A subsequence of a string is a new string which is formed from the original string by deleting
some (can be none) of the characters without disturbing the relative positions of the remaining
characters.
(ie, "ace" is a subsequence of "abcde" while "aec" is not).

Example 1:

Input: s = "abc", t = "ahbgdc"
Output: true
Example 2:

Input: s = "axc", t = "ahbgdc"
Output: false
*/

func isSubsequence(s string, t string) bool {
	if len(t) < len(s) {
		return false
	}
	var j int
	for i := 0; i < len(t) && j < len(s); i++ {
		if s[j] == t[i] {
			j++
		}
	}
	return j == len(s)
}
