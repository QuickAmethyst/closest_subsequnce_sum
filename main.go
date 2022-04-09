package main

import (
	"fmt"
	"math"
	"sort"
)

func getAllSubsequenceSum(nums []int) []int {
	result := make([]int, 0)
	result = append(result, 0)

	for _, num := range nums {
		tmp := append(result)

		for _, x := range tmp {
			tmp = append(tmp, x+num)
		}

		result = tmp
	}

	return result
}

func closestSubsequenceSum(input []int, target int) (result int) {
	mid := len(input) / 2
	firstHalfSums := getAllSubsequenceSum(input[:mid])
	secondHalfSums := getAllSubsequenceSum(input[mid:])

	sort.Slice(firstHalfSums, func(i, j int) bool {
		return firstHalfSums[i] < firstHalfSums[j]
	})

	result = int(math.Abs(float64(target)))
	for _, secondHalfSum := range secondHalfSums {
		var x = target - secondHalfSum

		v := Ceiling(firstHalfSums, 0, len(firstHalfSums)-1, x)
		if v >= 0 {
			result = int(math.Min(float64(result), float64(firstHalfSums[v] - x)))
		}

		v = Floor(firstHalfSums, 0, len(firstHalfSums)-1, x)
		if v >= 0 {
			result = int(math.Min(float64(result), float64(x - firstHalfSums[v])))
		}
	}

	return
}

func bfs(nums []int, left int, right int, target int) (int, int, int) {
	if right >= left {
		mid := (left + right) / 2

		// If the element is present at the middle
		// itself
		if nums[mid] == target {
			return mid, left, right
		}

		// If element is smaller than mid, then
		// it can only be present in left subarray
		if nums[mid] > target {
			return bfs(nums, left, mid-1, target)
		}

		// Else the element can only be present
		// in right subarray
		return bfs(nums, mid+1, right, target)
	}

	return -1, right, left
}

// Ceiling return the least element greater than or equal to target, or -1 if there is no such element.
// The search process is using binary search
func Ceiling(nums []int, left int, right int, target int) int {
	result, l, r := bfs(nums, left, right, target)

	if result >= 0 {
		return result
	}

	if l >= 0 && l <= len(nums)-1 && nums[l] > target {
		return l
	}

	if r >= 0 &&r <= len(nums)-1 && nums[r] > target {
		return r
	}

	return -1
}

// Floor return the greatest element less than or equal to target, or -1 if there is no such element.
// The search process is using binary search
func Floor(nums []int, left int, right int, target int) int {
	result, l, r := bfs(nums, left, right, target)

	if result >= 0 {
		return result
	}

	if l >= 0 && l <= len(nums)-1 && nums[l] < target {
		return l
	}

	if r >= 0 && r <= len(nums)-1 && nums[r] < target {
		return r
	}

	return -1
}

func main() {
	var input = []int{7, -9, 15, -2}
	fmt.Println("Result: ", closestSubsequenceSum(input, -5))
}
