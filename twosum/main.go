package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
    var arr = make([]int, 2)
	var number *int
    for i := 0; i < len(nums); i++ {
		arr = arr[:0] // empty element in slice
        number = &nums[i]
        arr = append(arr, i)
        for j := 0; j < len(nums); j++ {
			if(i == j) {continue;}
            if 	*number + nums[j] == target{
                arr = append(arr, j)
                return arr
            }
    	}
    }
    return arr
}  

func main (){
	fmt.Print(twoSum([]int{2,7,11,15}, 9))
}
