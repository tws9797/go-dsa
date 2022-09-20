package array

import "fmt"

func Arrays() {

	// Initialize array
	var a = [5]int{3, 5, 7, 4, 9}
	fmt.Println(a)

	// Index out of bound
	// a[6] = 0

	// Access element
	var b = a[0]
	fmt.Println(b)

	// Will initialize array of 5 with {5,5,5,5,5}
	var c [5]int
	fmt.Println(c)

	// Unlike C, C++, Java, Golang makes a copy of array
	var d [5]int = a
	fmt.Println(d)

	d[0] = 1
	fmt.Println(d)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for i, val := range a {
		fmt.Printf("Index: %d Value: %d\n", i, val)
	}

	e := [2][2]int{
		{3, 5},
		{7, 9},
	}
	fmt.Println(e)
}
