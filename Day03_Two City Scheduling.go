package leetcode

import "sort"

/*
There are 2N people a company is planning to interview. The cost of flying the i-th person to city A is costs[i][0], and the cost of flying the i-th person to city B is costs[i][1].

Return the minimum cost to fly every person to a city such that exactly N people arrive in each city.

Example 1:

Input: [[10,20],[30,200],[400,50],[30,20]]
Output: 110
Explanation:
The first person goes to city A for a cost of 10.
The second person goes to city A for a cost of 30.
The third person goes to city B for a cost of 50.
The fourth person goes to city B for a cost of 20.
*/
func twoCitySchedCost(costs [][]int) int {
	var total int
	sort.Slice(costs, func(a, b int) bool {
		return costs[a][0]-costs[a][1] < costs[b][0]-costs[b][1]
	})

	for i := 0; i < len(costs); i++ {
		if i < len(costs)/2 {
			total += costs[i][0]
		} else {
			total += costs[i][1]
		}
	}
	return total

}
