package precomputation

// https://leetcode.com/problems/product-of-array-except-self/

func ProductExceptSelf(nums []int) []int {

	// The length of the input array
	lenNums := len(nums)

	// The left and right arrays as described in the algorithm
	l := make([]int, lenNums)
	r := make([]int, lenNums)

	// Final answer array to be returned
	ans := make([]int, lenNums)

	// l[i] contains the product of all the elements to the left
	// Note: for the element at index '0', there are no elements to the left,
	// so l[0] would be 1
	l[0] = 1

	for i := 1; i < lenNums; i++ {

		// l[i - 1] already contains the product of elements to the left of 'i - 1'
		// Simply multiplying it with nums[i - 1] would give the product of all
		// elements to the left of index 'i'
		l[i] = nums[i-1] * l[i-1]
	}

	// r[i] contains the product of all the elements to the right
	// Note: for the element at index 'length - 1', there are no elements to the right,
	// so the r[length - 1] would be 1
	r[lenNums-1] = 1
	for i := lenNums - 2; i >= 0; i-- {

		// r[i + 1] already contains the product of elements to the right of 'i + 1'
		// Simply multiplying it with nums[i + 1] would give the product of all
		// elements to the right of index 'i'
		r[i] = nums[i+1] * r[i+1]
	}

	// Constructing the answer array
	for i := 0; i < lenNums; i++ {

		// For the first element, r[i] would be product except self
		// For the last element of the array, product except self would be l[i]
		// Else, multiple product of all elements to the left and to the right
		ans[i] = r[i] * l[i]
	}

	return ans
}

func ProductExceptSelf2(nums []int) []int {

	// The length of the input array
	lenNums := len(nums)

	// Final answer array to be returned
	ans := make([]int, lenNums)

	// ans[i] contains the product of all the elements to the left
	// Note: for the element at index '0', there are no elements to the left,
	// so the ans[0] would be 1
	ans[0] = 1
	for i := 1; i < lenNums; i++ {
		ans[i] = nums[i-1] * ans[i-1]
	}

	// r contains the product of all the elements to the right
	// Note: for the element at index 'length - 1', there are no elements to the right,
	// so the r would be 1
	r := 1
	for i := lenNums - 1; i >= 0; i-- {

		// For the index 'i', r would contain the
		// product of all elements to the right. We update r accordingly
		ans[i] = ans[i] * r
		r *= nums[i]
	}

	return ans
}
