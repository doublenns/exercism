package main

import "fmt"

func main() {
	slice_to_map := func(nums []int) map[int][]int {
		result := make(map[int][]int)
		for i, v := range nums {
			_, found := result[v]
			if found {
				result[v] = append(result[v], i)
			} else {
				result[v] = []int{i}
			}
		}
		return result
	}

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(nums)
	fmt.Println(slice_to_map(nums))

}
