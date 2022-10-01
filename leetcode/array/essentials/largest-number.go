package essentials

import (
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/largest-number/

func LargestNumber(nums []int) string {

	var strs []string

	for i := 0; i < len(nums); i++ {
		strs = append(strs, strconv.Itoa(nums[i]))
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	if strs[0] == "0" {
		return "0"
	}

	largestNum := strings.Join(strs, "")

	return largestNum
}
