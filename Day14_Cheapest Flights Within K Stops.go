package leetcode

/*
There are n cities connected by m flights. Each flight starts from city u and arrives at v with a price w.

Now given all the cities and flights, together with starting city src and the destination dst,
your task is to find the cheapest price from src to dst with up to k stops.
If there is no such route, output -1.

Example 1:
Input:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 1
Output: 200
The cheapest price from city 0 to city 2 with at most 1 stop costs 200


Example 2:
Input:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 0
Output: 500
The cheapest price from city 0 to city 2 with at most 0 stop costs 500
*/
import "math"

type node struct {
	key   int
	price int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	graph := map[int][]node{}
	for _, flight := range flights {
		u, v, p := flight[0], flight[1], flight[2]
		graph[u] = append(graph[u], node{v, p})
	}
	queue := []node{{src, 0}}
	res := math.MaxInt32
	for len(queue) > 0 {
		for size := len(queue); size > 0; size-- {
			u, c := queue[0].key, queue[0].price
			if u == dst {
				res = min(res, c)
			}
			for _, neigh := range graph[u] {
				v, p := neigh.key, neigh.price
				if c+p > res {
					continue
				}
				queue = append(queue, node{v, c + p})
			}
			queue = queue[1:]
		}
		if K < 0 {
			break
		}
		K--
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
