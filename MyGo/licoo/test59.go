package main

import "fmt"

func main()  {
	fmt.Println("yes")
}
func generateMatrix(n int) [][]int {
	left,right,top,bottom:=0,n-1,0,n-1;
	a := make([][]int, n)
	for i := range a {
		a[i]=make([]int,n)
	}
	num:=1
	for left<=right&&top<=bottom {
		for i := left; i <=right; i++ {
			a[left][i]=num
			num++
		}

		for i := top+1; i <=bottom ; i++ {
			a[i][right]=num
			num++
		}

		for i := right-1; i>left ; i-- {
			a[bottom][i]=num
			num++
		}

		for i := bottom; i >top ; i-- {
			a[i][left]=num
			num++
		}
		left++
		right--
		bottom--
		top++
	}
	return a
}