package leetcode

/*
The set [1,2,3,...,n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:

"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.

Note:

Given n will be between 1 and 9 inclusive.
Given k will be between 1 and n! inclusive.
Example 1:

Input: n = 3, k = 3
Output: "213"
Example 2:

Input: n = 4, k = 9
Output: "2314"
*/
func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	var fact, index int
	var result []byte
	for len(nums) > 0 {
		fact = factorial(n - 1)
		index = (k - 1) / fact
		k = k - index*fact
		result = append(result, byte(nums[index]+48))
		nums = append(nums[:index], nums[index+1:]...)
		n--
	}
	return string(result)
}

func factorial(n int) int {
	result := 1
	for n > 0 {
		result *= n
		n--
	}
	return result
}
