package leetcode

import "sort"

/*
Given a list of airline tickets represented by pairs of departure and arrival airports [from, to], reconstruct the itinerary in order. All of the tickets belong to a man who departs from JFK. Thus, the itinerary must begin with JFK.

Note:

If there are multiple valid itineraries, you should return the itinerary that has the smallest lexical order when read as a single string. For example, the itinerary ["JFK", "LGA"] has a smaller lexical order than ["JFK", "LGB"].
All airports are represented by three capital letters (IATA code).
You may assume all tickets form at least one valid itinerary.
One must use all the tickets once and only once.
Example 1:

Input: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
Output: ["JFK", "MUC", "LHR", "SFO", "SJC"]
Example 2:

Input: [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]
Explanation: Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"].
             But it is larger in lexical order.
*/

func findItinerary(tickets [][]string) []string {
	m := make(map[string][]string, len(tickets)+1)
	var ans []string

	for _, t := range tickets {
		m[t[0]] = append(m[t[0]], t[1])
	}

	for k := range m {
		sort.Strings(m[k])
	}

	DFS("JFK", m, &ans)

	i, j := 0, len(ans)-1
	for i < j {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return ans
}

func DFS(start string, m map[string][]string, ans *[]string) {
	for len(m[start]) > 0 {
		cur := m[start][0]
		m[start] = m[start][1:]
		DFS(cur, m, ans)
	}

	*ans = append(*ans, start)
}
