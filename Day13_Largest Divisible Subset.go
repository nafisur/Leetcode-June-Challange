package leetcode

import "sort"

/*
Given a set of distinct positive integers, find the largest subset such that every pair (Si, Sj) of elements in this subset satisfies:

Si % Sj = 0 or Sj % Si = 0.

If there are multiple solutions, return any subset is fine.

Example 1:

Input: [1,2,3]
Output: [1,2] (of course, [1,3] will also be ok)
Example 2:

Input: [1,2,4,8]
Output: [1,2,4,8]
*/
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([]int, n)
	parent := make([]int, n)
	var maxN, maxI int
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if nums[j]%nums[i] == 0 && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				parent[i] = j
				if dp[i] > maxN {
					maxN = dp[i]
					maxI = i
				}
			}
		}
	}

	res := make([]int, 0, n)
	currIndex := maxI
	for i := 0; i < maxN; i++ {
		res = append(res, nums[currIndex])
		currIndex = parent[currIndex]
	}
	return res
}
