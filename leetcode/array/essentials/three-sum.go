package essentials

import (
	"sort"
)

// https://leetcode.com/problems/3sum/

func ThreeSum(nums []int) [][]int {

	var ans [][]int

	// Sort the input array nums
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {

		// If the current value is greater than zero, break from the loop
		// If the current value is same as the value before, skip it
		// To skip duplicate
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Set low pointer to i+1, and high pointer r to the last index
		l, r := i+1, len(nums)-1

		// While low pointer is smaller than high
		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum < 0 {
				// If sum is less than 0, increment l
				//
				l++
			} else if sum > 0 {
				// If sum is greater than 0, decrement r
				r--
			} else {

				// Add it to result ans
				ans = append(ans, []int{nums[i], nums[l], nums[r]})

				// Increment l or decrement r
				// while the next value is the same as before to avoid duplicates in the result.

				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}

				// Decrement r and increment l
				l++
				r--
			}
		}

	}

	return ans
}

func ThreeSumHashSet(nums []int) [][]int {

	// Sort the input array nums
	sort.Ints(nums)
	var ans [][]int

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			continue
		}

		if i == 0 || nums[i-1] != nums[i] {
			seen := map[int]bool{}
			for j := i + 1; j < len(nums); j++ {
				complement := -nums[i] - nums[j]
				if _, ok := seen[complement]; ok {
					ans = append(ans, []int{nums[i], nums[j], complement})
					for j+1 < len(nums) && nums[j] == nums[j+1] {
						j++
					}
				}
				seen[nums[j]] = true
			}
		}
	}

	return ans
}

func ThreeSumNoSort(nums []int) [][]int {

	// Use another hash set dups to skip duplicates in the outer loop
	dups, seen := map[int]bool{}, map[int]int{}
	res := map[[3]int]bool{}

	for i := 0; i < len(nums); i++ {
		if _, ok := dups[nums[i]]; !ok {
			dups[nums[i]] = true
			for j := i + 1; j < len(nums); j++ {
				complement := -nums[i] - nums[j]
				if c, ok := seen[complement]; ok && c == i {

					triplet := []int{nums[i], nums[j], complement}
					sort.Ints(triplet)

					var tripletKey [3]int
					copy(tripletKey[:], triplet)

					res[tripletKey] = true
				}
				seen[nums[j]] = i
			}
		}
	}
	triplets := make([][]int, 0, len(res))

	for k, _ := range res {
		k := k
		triplets = append(triplets, k[:])
	}

	return triplets
}
