package main

import (
	"fmt"
	"slices"
)

// Slice --> dynamic
// most used construct in go
// usefull methods
func main() {
	var nums[]int

    fmt.Println(nums)
    fmt.Println(nums == nil)
    fmt.Println(len(nums))

    foods := make([]int, 0, 5)
    fmt.Println(foods)
    fmt.Println(len(foods))
    fmt.Println(foods == nil)
    fmt.Println(cap(foods))

    foods = append(foods, 1)
    fmt.Println(foods)

    foods = append(foods, 2)
    foods = append(foods, 3)
    foods = append(foods, 4)
    foods = append(foods, 5)
    foods = append(foods, 6)
    foods = append(foods, 7)
    foods = append(foods, 8)
    foods = append(foods, 9)
    foods = append(foods, 10)
    fmt.Println(foods)
    fmt.Println(cap(foods))

    cuisine := []int{1,2,3}
    fmt.Println(cuisine)
    fmt.Println(cap(cuisine))

    // copy
    newCuisine := []int{}
    copy(newCuisine, cuisine)
    fmt.Println(newCuisine)

    // slice operator
    cups := []int{1,2,3}
    fmt.Println(cups[0:2])
    fmt.Println(cups[:1])
    fmt.Println(cups[1:])

    nums1 := []int{1, 2}
    nums2 := []int{1, 3}
    fmt.Println(slices.Equal(nums1, nums2))

    nums3 := [][]int{{ 1, 2 }, {2, 3}}
    fmt.Println(nums3)
}