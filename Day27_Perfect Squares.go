package leetcode

/*
Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.

Example 1:

Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.
Example 2:

Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.
*/
func numSquares(n int) int {
	m := make([]int, n+1)
	m[1] = 1
	for i := 2; i <= n; i++ {
		min := i
		for j := 1; j*j <= i; j++ {
			current := 1 + m[i-j*j]
			if min > current {
				min = current
			}
		}
		m[i] = min
	}

	return m[n]
}
