package problems

// TwoSumSolution
//
// Problem link: https://leetcode.com/problems/two-sum/
func TwoSumSolution(nums []int, target int) []int {
	var solution []int

	// Compare each value with another value
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				solution = []int{i, j}
			}
		}
	}

	return solution
}

// TODO
func TwoSumOptimizedSolution(nums []int, target int) []int {
	var solution []int
	return solution
}
