package main

import "fmt"

func main()  {
	nums:= []int{3,2,2,3}
	removeElement(nums,3)
}
//方法一
func removeElement(nums []int, val int) int {
	i := len(nums)
	for index, num := range nums {
		if num==val {
			nums[index]=0
			i--
		}
	}
	fmt.Println(i,",nums=",nums)
	return i
}
//方法二
func removeElement2(nums []int, val int) int {
	left,right:=0, len(nums)-1
	for left<right {
		if nums[left]==val {
			nums[left]=nums[right]
			right--
		}else {
			left++
		}
	}
	fmt.Println(left+1,",",nums)
	return left
}
//方法3
func removeElement3(nums []int, val int) int {
	left := 0
	for _, v := range nums { // v 即 nums[right]
		if v != val {
			nums[left] = v
			left++
		}
	}
	fmt.Println(left+1,",",nums)
	return left
}