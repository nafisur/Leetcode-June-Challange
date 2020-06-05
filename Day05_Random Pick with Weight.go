package leetcode

/*
Given an array w of positive integers, where w[i] describes the weight of index i, write a function pickIndex which randomly picks an index in proportion to its weight.

Note:

1 <= w.length <= 10000
1 <= w[i] <= 10^5
pickIndex will be called at most 10000 times.
Example 1:

Input:
["Solution","pickIndex"]
[[[1]],[]]
Output: [null,0]
Example 2:

Input:
["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
[[[1,3]],[],[],[],[],[]]
Output: [null,0,1,1,1,0]
*/
import "math/rand"

type Solution struct {
	distance [][]int
	total    int
	count    int
}

func Constructor(w []int) Solution {
	d := [][]int{}
	total := 0
	for _, weight := range w {
		total += weight
		interval := []int{total - weight, total}
		d = append(d, interval)
	}
	return Solution{
		distance: d,
		total:    total,
		count:    0,
	}
}

func (this *Solution) PickIndex() int {
	this.count++
	random := rand.Intn(this.total)
	low, mid, high := 0, 0, len(this.distance)-1

	for low <= high {
		mid = (low + high) / 2
		if this.distance[mid][1] > random && this.distance[mid][0] <= random {
			return mid
		}
		if this.distance[mid][1] <= random {
			low = mid + 1
		} else if this.distance[mid][0] > random {
			high = mid - 1
		}
	}
	return mid
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
