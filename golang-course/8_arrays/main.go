package main

import "fmt"

func main() {
	var nums[4]int
	fmt.Println(len(nums))

	nums[0] = 1
	fmt.Println(nums[0])

	fmt.Println(nums)

	var vals[4]bool
	fmt.Println(vals)
	vals[2] = true
	fmt.Println(vals)

	var name[3]string
	name[0] = "golan"
	fmt.Println(name)

	numbs := [3]int{1, 2, 3}
	fmt.Println(numbs)

	numb2s := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(numb2s)
}