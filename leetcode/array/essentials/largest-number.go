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

	// "30" > "3" true
	// "303" > "30" true

	//Input: nums = [3,30,34,5,9]
	//Output: "9534330"

	// If sorted by
	// sort.Slice(strs, func(i, j int) bool {
	// 	return strs[i]+strs[j] > strs[j]+strs[i]
	// })
	// result will be: 9,5,34,30,3 as "30" > "3" true
	// Therefore, "30"+"3" < "3" + "30" is being used, "3" will be ordered in front of "30"
	
	// a+b > b+a && b+c > c+b == a+c > c+a
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	if strs[0] == "0" {
		return "0"
	}

	largestNum := strings.Join(strs, "")

	return largestNum
}
